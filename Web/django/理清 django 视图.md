### django 视图
#### FBV

> 需要根据请求方法不同做条件分支返回不同数据

```python
def func(request,**kwargs):
    if request.method.lower() == 'get':
        pass
    elif request.method.lower() == 'post': 
        pass
```
#### CBV

> 视图类初始化时自动根据请求方法dispatch到不同的`handler`处理

```python
class demoView(View):
    def get(self,request, *args, **kwargs):
        pass
    def post(self,request, *args, **kwargs):
        pass
```
### rest_framework 视图
#### APIView

> 实现了多种请求认证、解析、控制等机制

> super().__init__()
>
> =>dispatch
>
> =>initial_request封装原生request、在initial中实现根据不同请求method进行handler处理

```python
class demoView(APIView):
    renderer_classes = api_settings.DEFAULT_RENDERER_CLASSES
    parser_classes = api_settings.DEFAULT_PARSER_CLASSES
    authentication_classes = api_settings.DEFAULT_AUTHENTICATION_CLASSES
    throttle_classes = api_settings.DEFAULT_THROTTLE_CLASSES
    permission_classes = api_settings.DEFAULT_PERMISSION_CLASSES
    content_negotiation_class = api_settings.DEFAULT_CONTENT_NEGOTIATION_CLASS
    metadata_class = api_settings.DEFAULT_METADATA_CLASS
    versioning_class = api_settings.DEFAULT_VERSIONING_CLASS

    def get(self,request, *args, **kwargs):
        pass
    def post(self,request, *args, **kwargs):
        pass
```
#### GenericAPIView

> 通用api视图实现了自定义查询集、序列化类、筛选、过滤、分页处理。

```python
queryset = None # 查询集
serializer_class = None # 序列化类
lookup_field = 'pk' # 查找字段
lookup_url_kwarg = None # 
filter_backends = api_settings.DEFAULT_FILTER_BACKENDS # 筛选过滤类
pagination_class = api_settings.DEFAULT_PAGINATION_CLASS # 分页类

def get_queryset():
    #...
    return queryset
 
def get_object():
    # queryset执行filter后返回obj
    return obj

# 返回序列化类，可根据条件自定义serializer_class
def get_serializer_class():
    #....
    return self.serializer_class

# 获取序列化类实例
def get_serializer():
    #...调用self.get_serializer_class()
    return serializer_class(*args, **kwargs)

# 获取序列化类的上下文
def get_serializer_context(self):
    """
    Extra context provided to the serializer class.
    """
    return {
            'request': self.request,
            'format': self.format_kwarg,
            'view': self
        }

# 分页查询集
def paginate_queryset():
    #....
    return self.paginator.paginate_queryset(queryset, self.request, view=self)

# 通过分页器获得分页响应
def get_paginated_response(self, data):
    """
    Return a paginated style `Response` object for the given output data.
    """
    assert self.paginator is not None
    return self.paginator.get_paginated_response(data)
```

#### Mixin

> 	mixins.CreateModelMixin
> 	mixins.RetrieveModelMixin,
> 	mixins.UpdateModelMixin,
> 	mixins.DestroyModelMixin,
> 	分别重写了GenericAPIView的方法，必须和GenericAPIView同时继承，否则无意义

#### ViewSetMixin

> 主要重写了视图as_view()方法，可以将http method 绑定到对应的action，如：
>
> MyViewSet.as_view({'get': 'list', 'post': 'create'})
>
> 实现 类方法，通过装饰器@action获得对应的视图处理
>
> ```python
> @classmethod
> def get_extra_actions(cls):
>     """
>     Get the methods that are marked as an extra ViewSet `@action`.
>     """
>     return [_check_attr_name(method, name)
> 			for name, method in getmembers(*cls*, _is_extra_action)]
> ```

#### ViewSet

> 主要继承了ViewSetMixin和APIView，没做任何处理
>
> 只能自动生成@action装饰的路由api

####  GenericViewSet

> 主要继承了ViewSetMixin和GenericAPIView，没做任何处理
>
> 只能自动生成@action装饰的路由api

#### ModelViewSet

> 继承了		mixins.CreateModelMixin,
>         			mixins.RetrieveModelMixin,
>                     mixins.UpdateModelMixin,
>                     mixins.DestroyModelMixin,
>                     mixins.ListModelMixin,GenericViewSet
>
> 没做任何处理

### 视图继承关系

```python
class View:
    pass

class APIView(View):
    pass

class GenericAPIView(APIView):
    pass

# 以下类无继承，对应重写了GenericAPIView的方法
mixins.RetrieveModelMixin
mixins.UpdateModelMixin
mixins.DestroyModelMixin
mixins.CreateModelMixin
mixins.ListModelMixin

class ViewSetMixin:
    pass

class ViewSet(ViewSetMixin,APIView):
    pass

class GenericViewSet(ViewSetMixin,GenericAPIView):
    pass

class ReadOnlyModelViewSet(mixins.RetrieveModelMixin,
                           mixins.ListModelMixin,
                           GenericViewSet):
    pass


class ModelViewSet(mixins.CreateModelMixin,
                   mixins.RetrieveModelMixin,
                   mixins.UpdateModelMixin,
                   mixins.DestroyModelMixin,
                   mixins.ListModelMixin,
                   GenericViewSet):
    pass
```

