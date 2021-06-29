<p style='text-align:center;font-size:30px;font-weight:bold'>nsqadmin之nginx代理配置</p>

使用nsq消息队列后，在本地可以直接通过访问 http://localhost:4171 看到管理界面。但是生产环境中，不能直接对外暴露服务器4171端口，这时候需要配置nginx代理。然而nsqadmin配置nginx比较坑，简单的通过`proxy_pass http://127.0.0.1:4171/`依然是无法正常显示管理界面的。网上搜索到的各类文章，没一个能行。本文记录下nsqadmin正确的nginx代理配置过程。

1. 生成静态文件
   如果直接进行如下配置，浏览器访问地址 nsqadmin.example.com/nsq将显示css、js等各类静态资源找不到(404)，这时候需要同时代理静态资源目录。

   ```nginx
   server {
       listen       80;
       server_name  nsqadmin.example.com;
       auth_basic "NSQAdmin Auth";
       auth_basic_user_file auth-passwd/nsqadmin-htpasswd;
       location /nsq/ {
           proxy_pass http://127.0.0.1:4171/;
     }
   ```

   为了生成静态资源文件，根据官方仓库[nsq/nsqadmin at master · nsqio/nsq (github.com)](https://github.com/nsqio/nsq/tree/master/nsqadmin)的说明，需要进行如下步骤：

   ```shell
   # 1. clone repo
   git clone https://github.com/nsqio/nsq.git
   cd nsqadmin
   # 2. install go-bindata
   go get -u github.com/shuLhan/go-bindata
   # 3. install NodeJS 14.x (includes npm)
   # 4. install dependencies
   npm install
   # 5. 执行
   ./gulp --series clean watch # 或者执行 ./gulp --series clean build
   # 6. 执行
   go-bindata --debug --pkg=nsqadmin --prefix=static/build/ static/build/...
   ```

   > 如果嫌弃上述过程太麻烦，可以直接拷贝仓库里编译好的static文件夹，即nsqadmin目录下的static文件夹。

   上述过程把所有相关的静态文件都build到了`static/build/`目录下，我们需要将该目录拷贝至服务器的某个目录下，例如`/data/static`。然后nginx配置如下：

   ```nginx
   server {
       listen       80;
       server_name  nsqadmin.example.com;
       auth_basic "NSQAdmin Auth";
       auth_basic_user_file auth-passwd/nsqadmin-htpasswd;
       location /nsq/ {
           proxy_pass http://127.0.0.1:4171/;
       }
       location /static/ {
           alias  /data/static/ 
       }
   }
   ```

2.  代理`/api`
   上述nginx配置完成后，可以显示出nsqadmin的首页内容，但是当我们点击交互时依然会出错，查看页面请求，你会发现nsqadmin在请求`/api`路由，而这些路由显示为404。我们需要将`/api`代理到nsq内部。因此需要追加如下配置:

   ```nginx
   server {
       listen       80;
       server_name  nsqadmin.example.com;
       auth_basic "NSQAdmin Auth"; 
       auth_basic_user_file auth-passwd/nsqadmin-htpasswd;
       location /nsq/ {
           proxy_pass http://127.0.0.1:4171/;
       }
       location /static/ {
           alias  /data/static/ 
       }
       location /api {
           proxy_pass http://127.0.0.1:4171/nsq$request_uri;
     }
   }
   ```

最后重启nginx, 访问 nsqadmin.example.com/nsq，输入用户名、密码即可正常访问nsqadmin的管理界面。