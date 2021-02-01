最近在工作中遇到一个需求，就是要开一个接口来接收供应商推送的数据。项目采用的`python`的`django`框架，我是想也没想，就直接一梭哈，写出了如下代码：

```python
class XXDataPushView(APIView):
    """
    接收xx数据推送
    """
		# ...
    @white_list_required
    def post(self, request, **kwargs):
        req_data = request.data or {}
				# ...
```

但随后，发现每日数据并没有任何变化，质问供应商是否没有做推送，在忽悠我们。然后对方给的答复是，他们推送的是`gzip`压缩的数据流，接收端需要主动进行解压。此前从没有处理过这种压缩的数据，对方具体如何做的推送对我来说也是一个黑盒。

因此，我要求对方给一个推送的简单示例，没想到对方**不讲武德**，仍过来一段没法单独运行的`java`代码：

```java
private byte[] compress(JSONObject body) {
    try {
        ByteArrayOutputStream out = new ByteArrayOutputStream();
        GZIPOutputStream gzip = new GZIPOutputStream(out);
        gzip.write(body.toString().getBytes());
        gzip.close();
        return out.toByteArray();
    } catch (Exception e) {
        logger.error("Compress data failed with error: " + e.getMessage()).commit();
    }
    return JSON.toJSONString(body).getBytes();
}

public void post(JSONObject body, String url, FutureCallback<HttpResponse> callback) {
    RequestBuilder requestBuilder = RequestBuilder.post(url);
    requestBuilder.addHeader("Content-Type", "application/json; charset=UTF-8");
    requestBuilder.addHeader("Content-Encoding", "gzip");

    byte[] compressData = compress(body);

    int timeout = (int) Math.max(((float)compressData.length) / 5000000, 5000);

    RequestConfig.Builder requestConfigBuilder = RequestConfig.custom();
    requestConfigBuilder.setSocketTimeout(timeout).setConnectTimeout(timeout);

    requestBuilder.setEntity(new ByteArrayEntity(compressData));

    requestBuilder.setConfig(requestConfigBuilder.build());

    excuteRequest(requestBuilder, callback);
}

private void excuteRequest(RequestBuilder requestBuilder, FutureCallback<HttpResponse> callback) {
    HttpUriRequest request = requestBuilder.build();
    httpClient.execute(request, new FutureCallback<HttpResponse>() {
        @Override
        public void completed(HttpResponse httpResponse) {
            try {
                int responseCode = httpResponse.getStatusLine().getStatusCode();
                if (callback != null) {
                    if (responseCode == 200) {
                        callback.completed(httpResponse);
                    } else {
                        callback.failed(new Exception("Status code is not 200"));
                    }
                }
            } catch (Exception e) {
                logger.error("Get error on " + requestBuilder.getMethod() + " " + requestBuilder.getUri() + ": " + e.getMessage()).commit();
                if (callback != null) {
                    callback.failed(e);
                }
            }

            EntityUtils.consumeQuietly(httpResponse.getEntity());
        }

        @Override
        public void failed(Exception e) {
            logger.error("Get error on " + requestBuilder.getMethod() + " " + requestBuilder.getUri() + ": " + e.getMessage()).commit();
            if (callback != null) {
                callback.failed(e);
            }
        }

        @Override
        public void cancelled() {
            logger.error("Request cancelled on  " + requestBuilder.getMethod() + " " + requestBuilder.getUri()).commit();
            if (callback != null) {
                callback.cancelled();
            }
        }
    });
}
```

从上述代码可以看出，对方将`json`数据压缩为了`gzip`数据流`stream`。于是搜索`django`的文档，只有这段关于`gzip`处理的装饰器描述:

> [`django.views.decorators.gzip`](https://docs.djangoproject.com/zh-hans/3.1/topics/http/decorators/#module-django.views.decorators.gzip) 里的装饰器控制基于每个视图的内容压缩。
>
> - `gzip_page`()[¶](https://docs.djangoproject.com/zh-hans/3.1/topics/http/decorators/#django.views.decorators.gzip.gzip_page)
>
>   如果浏览器允许 gzip 压缩，那么这个装饰器将压缩内容。它相应的设置了 `Vary` 头部，这样缓存将基于 `Accept-Encoding` 头进行存储。

但是，这个装饰器只是压缩发往浏览器的内容，我们目前的需求是解压缩接收的数据。这不是我们想要的。

幸运的是，在`flask`中有一个扩展叫`flask-inflate`，安装了此扩展会自动对请求来的数据做解压操作。查看该扩展的具体代码处理：

```python
# flask_inflate.py
import gzip
from flask import request

GZIP_CONTENT_ENCODING = 'gzip'


class Inflate(object):
    def __init__(self, app=None):
        if app is not None:
            self.init_app(app)

    @staticmethod
    def init_app(app):
        app.before_request(_inflate_gzipped_content)


def inflate(func):
    """
    A decorator to inflate content of a single view function
    """
    def wrapper(*args, **kwargs):
        _inflate_gzipped_content()
        return func(*args, **kwargs)

    return wrapper


def _inflate_gzipped_content():
    content_encoding = getattr(request, 'content_encoding', None)

    if content_encoding != GZIP_CONTENT_ENCODING:
        return

    # We don't want to read the whole stream at this point.
    # Setting request.environ['wsgi.input'] to the gzipped stream is also not an option because
    # when the request is not chunked, flask's get_data will return a limited stream containing the gzip stream
    # and will limit the gzip stream to the compressed length. This is not good, as we want to read the
    # uncompressed stream, which is obviously longer.
    request.stream = gzip.GzipFile(fileobj=request.stream)
```

上述代码的核心是:

```python
 request.stream = gzip.GzipFile(fileobj=request.stream)
```

于是，在`django`中可以如下处理：

```python
class XXDataPushView(APIView):
    """
    接收xx数据推送
    """
		# ...
    @white_list_required
    def post(self, request, **kwargs):
        content_encoding = request.META.get("HTTP_CONTENT_ENCODING", "")
        if content_encoding != "gzip":
            req_data = request.data or {}
        else:
            gzip_f = gzip.GzipFile(fileobj=request.stream)
            data = gzip_f.read().decode(encoding="utf-8")
            req_data = json.loads(data)
        # ... handle req_data
```

ok, 问题完美解决。还可以用如下方式测试请求：

```python
import gzip
import requests
import json

data = {}

data = json.dumps(data).encode("utf-8")
data = gzip.compress(data)

resp = requests.post("http://localhost:8760/push_data/",data=data,headers={"Content-Encoding": "gzip", "Content-Type":"application/json;charset=utf-8"})

print(resp.json())
```

