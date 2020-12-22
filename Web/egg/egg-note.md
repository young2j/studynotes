# egg cli
```shell
mkdir egg-example && cd egg-example
npm init egg --type=simple # --type=ts 生成ts项目
npm i
npm run dev
```

# egg 目录约定
> app/
* view: 模板视图
* controller： 处理业务逻辑
* model：orm，定义数据模型
* service： 与数据库打交道，查询、请求数据
* extend: 扩展
* schedule：定时任务
> config/
* plugin: 插件

> app.js
* app启动时的一些自定义操作，比如监听app事件

# egg内置对象
## Application

### event

```js
//app.js
module.exports = app =>{
    app.once('server',server=>{
        //websocket
    })
    app.on('error',(err,ctx)=>{
        //report err
    })
    app.on('request',ctx=>{
        //log received request
    })
    app.on('response',ctx=>{
        const used = Date.now()-ctx.starttime //ctx.starttime 框架实现了的
        console.log('ellapse time:',used)
    })
}
```

### 获取app 对象

> 几乎所有被框架 [Loader](https://eggjs.org/zh-cn/advanced/loader.html) 加载的文件（Controller，Service，Schedule 等），都可以 export 一个函数，这个函数会被 Loader 调用，并使用 app 作为参数。
>
> 继承于 Controller, Service 基类的实例，可以直接通过 `this.app` 访问到 Application 对象。

```js
//app.js
module.exports = app =>{
    app.cache = new Cache()
}

//app/controller/user.js
class UserController extends Controller{
    async fetch(){
        //this.ctx.body = this.ctx.app.cache.get(this.ctx.query.id)
        this.ctx.body = this.app.cache.get(this.ctx.query.id)
    }
}
```

## Context

### 获取ctx对象

> 在中间件中，同koa一样，作为处理函数的参数

```js
//app/middleware/log.js
async function logMiddleware(ctx,next){
    console.log(ctx.query)
}
```

> 在非用户请求的场景下，如app.js中，可以通过`Application.createAnonymousContext()`创建一个匿名Contex实例。

```js
//app.js
module.exports = app =>{
    app.beforeStart(async ()=>{
        const ctx = app.createAnonymousContext()
        await ctx.service.posts.load() //preload
    })
}
```

> 在shedule中，每一个task都接受一个Context实例作为参数。

```js
//app/schedule/refresh.js
exports.task = async ctx =>{
    await ctx.service.posts.refresh()
}
```

### Request & Response

> - [Koa](http://koajs.com/) 会在 Context 上代理一部分 Request 和 Response 上的方法和属性;
> - 如 `ctx.request.query.id` = `ctx.query.id` ，`ctx.response.body`=`ctx.body`;
> - 需要注意的是，获取 POST 的 body 应该使用 `ctx.request.body`，而不是 `ctx.body`。

```js
// app/controller/user.js
class UserController extends Controller {
  async fetch() {
    const { app, ctx } = this;
    const id = ctx.request.query.id;
    ctx.response.body = app.cache.get(id);
  }
}
```

## Controller

> 继承自Controller的类有如下属性：

```js
const Controller = require('egg').Controller
class UserController extends Controller{
    const {
        ctx,
        app,
        config,
        service,
        logger
    } = this
}
module.exports = UserController
```

## Service

> 属性与Controller一致

```js
//app/service/user.js
const Service = require('egg').Service
class UserService extends Service{
        const {
        ctx,
        app,
        config,
        service,
        logger
    } = this
}
module.exports = UserService
```

## Helper

> Helper 用来提供一些实用的 utility 函数；
>
> 具有和Controller一样的属性；
>
> 可以通过ctx.helper获取。

```js
// app/controller/user.js
class UserController extends Controller {
  async fetch() {
    const { app, ctx } = this;
    const id = ctx.query.id;
    const user = app.cache.get(id);
    ctx.body = ctx.helper.formatUser(user);
  }
}
```

> 在app/extend中自定义helper方法

```js
//app/extend/helper.js
module.exports = {
    formatUser(user){
        return only(user, [ 'name', 'phone' ]);
    }
}
```

## Config

> 可以通过 `app.config` 从 Application 实例上获取到 config 对象；
>
> 也可以在 Controller, Service, Helper 的实例上通过 `this.config` 获取到 config 对象。

> ```
> config
> |- config.default.js
> |- config.prod.js
> |- config.unittest.js
> `- config.local.js
> ```

```js
//配置的三种写法
//1
module.exports = {
    logger:{
        dir:'/home/admin/logs/apploger'
    }
}
//2
exports.logger = {
    dir:'path/to/logfile',
    level:'DEBUG'
}

//3
module.exports = appInfo=>{
    //可以获得app的相关属性，如appInfo.baseDir
    return {
        logger:{
            dir:'path/to/logfile'
        }
    }
}
```

## Logger

### App Logger

可以通过 `app.logger` 来获取到它，应用级别的日志记录。

### App CoreLogger

可以通过 `app.coreLogger` 来获取到它，框架和插件通过它来打印应用级别的日志,一般我们在开发时都不应该用它。

### Context Logger

可以通过 `ctx.logger` 从 Context 实例上获取到它，是与请求相关的的日志。

### Context CoreLogger

可以通过 `ctx.coreLogger` 获取到它，一般只有插件和框架会通过它来记录日志。

### Controller Logger & Service Logger

可以在 Controller 和 Service 实例上通过 `this.logger` 获取到它们，本质上就是一个 Context Logger，不过在打印日志的时候还会额外的加上文件路径，方便定位日志的打印位置。



# Middleware

> 约定一个中间件是一个放置在 `app/middleware` 目录下的单独文件

## 有/无配置的中间件

```js
//app/middleware/gzip.js
//无配置的中间件
const isJSON = require('koa-is-json');
const zlib = require('zlib');

async function gzip(ctx, next) {
  await next();

  // 后续中间件执行完成后将响应体转换成 gzip
  let body = ctx.body;
  if (!body) return;
  if (isJSON(body)) body = JSON.stringify(body);

  // 设置 gzip body，修正响应头
  const stream = zlib.createGzip();
  stream.end(body);
  ctx.body = stream;
  ctx.set('Content-Encoding', 'gzip');
}

//有配置的中间件,接受options和app两个参数
const isJSON = require('koa-is-json');
const zlib = require('zlib');

module.exports = (options,app) => {
  return async function gzip(ctx, next) {
    await next();

    // 后续中间件执行完成后将响应体转换成 gzip
    let body = ctx.body;
    if (!body) return;

    // 支持 options.threshold
    if (options.threshold && ctx.length < options.threshold) return;

    if (isJSON(body)) body = JSON.stringify(body);

    // 设置 gzip body，修正响应头
    const stream = zlib.createGzip();
    stream.end(body);
    ctx.body = stream;
    ctx.set('Content-Encoding', 'gzip');
  };
};
```

## 使用中间件

### 在app中使用中间件

```js
//config/config.default.js
module.exports = {
    middleware:['gzip'], //注册gzip中间件
    gzip:{ //配置gzip的选项,就是options
        threshold:1024
    }
}
```

### 在框架和插件中使用中间件

> 框架和插件不支持在 `config.default.js` 中匹配 `middleware`，需要通过以下方式

```js
//app.js
module.exports = app =>{
    app.config.coreMiddleware.unshift('report')
}
//app/middleware/report.js
module.exports = ()=>{
    return async function (ctx,next){
        const startTime = Date.now()
        await next()
        //...
    }
}
```

### 在router中使用中间件

> 其他方式的中间件都是全局的，会处理每一个请求；
>
> 针对单个路由生效，只需在 `app/router.js` 中实例化和挂载。

```js
//app/router.js
module.exports = app =>{
    const gzip = app.middleware.gzip({threshold:1024})
    app.router.get('/gzip',gzip,app.controller.handler)
}
```

### 框架默认中间件

> 直接在config/config.default.js中修改配置。

```js
module.exports = {
    bodyParser:{
        jsonLimit:'10mb',
        formLimit:'1mb'
    }
}
```

### 使用koa中间件

```js
//app/middleware/compress.js
module.exports = require('koa-compress')

//config/config.default.js
module.exports = {
    middleware:['compress'],
    compress:{
        threshold:2048
    }
}
```

## 中间件的通用配置

### enable

```js
module.exports = {
    bodyParser:{
        enable:false //控制中间件是否开启
    }
}
```

### match/ignore

> 支持三种配置方式：
>
> 1. 字符串：当参数为字符串类型时，配置的是一个 url 的路径前缀，所有以配置的字符串作为前缀的 url 都会匹配上。 当然，你也可以直接使用字符串数组。
> 2. 正则：当参数为正则时，直接匹配满足正则验证的 url 的路径。
> 3. 函数：当参数为一个函数时，会将请求上下文传递给这个函数，最终取函数返回的结果（true/false）来判断是否匹配。

```js
//字符串
module.exports = {
  gzip: {
    match: '/static',
  },
};
//正则
module.exports = {
  gzip: {
    match:/iphone|ipad|ipod/i
  },
};
//函数
module.exports = {
  gzip: {
    match(ctx) {
      // 只有 ios 设备才开启
      const reg = /iphone|ipad|ipod/i;
      return reg.test(ctx.get('user-agent'));
    },
  },
};
```

# Router

## 路由定义

> 路由完整签名:
>
> router.verb('router-name', 'path-match', middleware1, ..., middlewareN, app.controller.action);

```js
// app/router.js
module.exports = app => {
  const { router, controller } = app;
  router.get('/home', controller.home);
  router.get('/user/:id', controller.user.page);
  router.post('/admin', isAdmin, controller.admin);
  router.post('/user', isLoginUser, hasAdminPermission, controller.user.create);
  router.post('/api/v1/comments', controller.v1.comments.create); // app/controller/v1/comments.js
};
```

> restful风格URL

```js
// app/router.js
module.exports = app => {
  const { router, controller } = app;
  router.resources('posts', '/api/posts', controller.posts);
  router.resources('users', '/api/v1/users', controller.v1.users); // app/controller/v1/users.js
};
```

## 路由参数获取

```js
// app/router.js
module.exports = app =>{
    const {router,controller} = app
    router.get('/search',controller.index.search) //query string
    router.get('user/:id/:name',controller.index.info) //动态路由
    router.post('/form',controller.index.form) //form data
}

//app/controller/index.js

// query string
// curl http://127.0.0.1:7001/search?name=egg
// curl http://127.0.0.1:7001/search?name=egg&name=koa
exports.search = async ctx=>{
    ctx.body = `search:${ctx.query.name}`
    // ctx.body = `search:${ctx.queries.name}` 
    // {
    //     name:['egg','koa']
    // }
}

// 动态路由
// curl http://127.0.0.1:7001/user/123/xiaoming
exports.info = async ctx =>{
    ctx.body =`user:${ctx.params.id}, ${ctx.params.name}`
}

// form data
// curl -X POST http://127.0.0.1:7001/form --data '{"name":"controller"}' --header 'Content-Type:application/json'
exports.form = async ctx=>{
    ctx.body = `body:${JSON.stringify(ctx.request.body)}`
}
```

## 表单校验

```js
// app/router.js
module.exports = app =>{
    //app.router.post('/user',app.controller.user) ?
    app.router.post('/user',app.controller.user.create)
}

// app/controller/user.js
const createRule = {
  username: {
    type: 'email',
  },
  password: {
    type: 'password',
    compare: 're-password',
  },
}
exports.create = async ctx => {
    ctx.validate(createRule)
    ctx.body = ctx.request.body
}
// curl -X POST http://127.0.0.1:7001/user --data 'username=abc@abc.com&password=111111&re-password=111111'
```

## 重定向

> 内部重定向 router.redirect(from,to)

```js
// app/router.js
module.exports = app => {
  app.router.get('index', '/home/index', app.controller.home.index);
  app.router.redirect('/', '/home/index', 302);
};

// app/controller/home.js
exports.index = async ctx => {
  ctx.body = 'hello controller';
};

// curl -L http://localhost:7001
```

> 外部重定向 ctx.redirect(to)

```js
// app/router.js
module.exports = app => {
  app.router.get('/search', app.controller.search.index);
};

// app/controller/search.js
exports.index = async ctx => {
  const type = ctx.query.type;
  const q = ctx.query.q || 'nodejs';

  if (type === 'bing') {
    ctx.redirect(`http://cn.bing.com/search?q=${q}`);
  } else {
    ctx.redirect(`https://www.google.co.kr/search?q=${q}`);
  }
};

// curl http://localhost:7001/search?type=bing&q=node.js
// curl http://localhost:7001/search?q=node.js
```

## 路由拆分

```js
// app/router.js
module.exports = app => {
  require('./router/news')(app);
  require('./router/admin')(app);
};

// app/router/news.js
module.exports = app => {
  app.router.get('/news/list', app.controller.news.list);
  app.router.get('/news/detail', app.controller.news.detail);
};

// app/router/admin.js
module.exports = app => {
  app.router.get('/admin/user', app.controller.admin.user);
  app.router.get('/admin/log', app.controller.admin.log);
};
```

# Controller

## 获取上传文件

> 浏览器上都是通过 <form method="POST" action="" enctype="multipart/form-data">格式发送文件的。框架通过内置 [Multipart](https://github.com/eggjs/egg-multipart) 插件来支持获取用户上传的文件。

### File模式

```js
//config/config.default.js
module.exports = {
    multipart:{
        mode:'file'
    }
}
```

> 上传接收单个文件

```js
// app/controller/upload.js
const Controller = require('egg').Controller;
const fs = require('mz/fs');

module.exports = class extends Controller {
  async upload() {
    const { ctx } = this;
    const file = ctx.request.files[0];
    const name = 'egg-multipart-test/' + path.basename(file.filename);
    let result;
    try {
      // 处理文件，比如上传到云端
      result = await ctx.oss.put(name, file.filepath);
    } finally {
      // 需要删除临时文件
      await fs.unlink(file.filepath);
    }

    ctx.body = {
      url: result.url,
      // 获取所有的字段值
      requestBody: ctx.request.body,
    };
  }
};
```

> 上传接收多个文件

```js
// app/controller/upload.js
const Controller = require('egg').Controller;
const fs = require('mz/fs');

module.exports = class extends Controller {
  async upload() {
    const { ctx } = this;
    console.log(ctx.request.body);
    console.log('got %d files', ctx.request.files.length);
    for (const file of ctx.request.files) {
      console.log('field: ' + file.fieldname);
      console.log('filename: ' + file.filename);
      console.log('encoding: ' + file.encoding);
      console.log('mime: ' + file.mime);
      console.log('tmp filepath: ' + file.filepath);
      let result;
      try {
        // 处理文件，比如上传到云端
        result = await ctx.oss.put('egg-multipart-test/' + file.filename, file.filepath);
      } finally {
        // 需要删除临时文件
        await fs.unlink(file.filepath);
      }
      console.log(result);
    }
  }
};
```

### stream模式

> 上传接收单个文件
>
> * 只支持上传一个文件。
> * 上传文件必须在所有其他的 fields 后面，否则在拿到文件流时可能还获取不到 fields。

```js
const path = require('path');
const sendToWormhole = require('stream-wormhole');
const Controller = require('egg').Controller;

class UploaderController extends Controller {
  async upload() {
    const ctx = this.ctx;
    const stream = await ctx.getFileStream();
    const name = 'egg-multipart-test/' + path.basename(stream.filename);
    // 文件处理，上传到云存储等等
    let result;
    try {
      result = await ctx.oss.put(name, stream);
    } catch (err) {
      // 必须将上传的文件流消费掉，要不然浏览器响应会卡死
      await sendToWormhole(stream);
      throw err;
    }

    ctx.body = {
      url: result.url,
      // 所有表单字段都能通过 `stream.fields` 获取到
      fields: stream.fields,
    };
  }
}

module.exports = UploaderController;
```

> 上传接收多个文件

```js
// config/config.default.js
module.exports = {
  multipart: {
    fileExtensions: [ '.apk' ] // 增加对 apk 扩展名的文件支持
  },
};
```



```js
const sendToWormhole = require('stream-wormhole');
const Controller = require('egg').Controller;

class UploaderController extends Controller {
  async upload() {
    const ctx = this.ctx;
    const parts = ctx.multipart();
    let part;
    // parts() 返回 promise 对象
    while ((part = await parts()) != null) {
      if (part.length) {
        // 这是 busboy 的字段
        console.log('field: ' + part[0]);
        console.log('value: ' + part[1]);
        console.log('valueTruncated: ' + part[2]);
        console.log('fieldnameTruncated: ' + part[3]);
      } else {
        if (!part.filename) {
          // 这时是用户没有选择文件就点击了上传(part 是 file stream，但是 part.filename 为空)
          // 需要做出处理，例如给出错误提示消息
          return;
        }
        // part 是上传的文件流
        console.log('field: ' + part.fieldname);
        console.log('filename: ' + part.filename);
        console.log('encoding: ' + part.encoding);
        console.log('mime: ' + part.mime);
        // 文件处理，上传到云存储等等
        let result;
        try {
          result = await ctx.oss.put('egg-multipart-test/' + part.filename, part);
        } catch (err) {
          // 必须将上传的文件流消费掉，要不然浏览器响应会卡死
          await sendToWormhole(part);
          throw err;
        }
        console.log(result);
      }
    }
    console.log('and we are done parsing the form!');
  }
}

module.exports = UploaderController;
```

##  Header

- `ctx.headers`，`ctx.header`，`ctx.request.headers`，`ctx.request.header`：这几个方法是等价的，都是获取整个 header 对象。
- `ctx.get(name)`，`ctx.request.get(name)`：获取请求 header 中的一个字段的值，如果这个字段不存在，会返回空字符串。
- 我们建议用 `ctx.get(name)` 而不是 `ctx.headers['name']`，因为前者会自动处理大小写。

## Cookie

> 配置

```js
//config/config.default.js
module.exports = {
  cookies: {
    // httpOnly: true | false,
    // sameSite: 'none|lax|strict',
  },
};
```

> 使用

```js
class CookieController extends Controller {
  async add() {
    const ctx = this.ctx;
    let count = ctx.cookies.get('count');
    count = count ? Number(count) : 0;
    ctx.cookies.set('count', ++count);
    ctx.body = count;
  }

  async remove() {
    const ctx = this.ctx;
    const count = ctx.cookies.set('count', null);
    ctx.status = 204;
  }
}
```

## Session

用来存储用户身份相关的信息，这份信息会加密后存储在 Cookie 中，实现跨请求的用户身份保持。

> 配置

```js
module.exports = {
    key:'EGG_SESS',// 承载 Session 的 Cookie 键值对名字
  	maxAge: 86400000, // Session 的最大有效时间
}
```

> 使用

```js
class PostController extends Controller {
  async fetchPosts() {
    const ctx = this.ctx;
    // 获取 Session 上的内容
    const userId = ctx.session.userId;
    const posts = await ctx.service.post.fetch(userId);
    // 修改 Session 的值
    ctx.session.visited = ctx.session.visited ? ++ctx.session.visited : 1;
    ctx.body = {
      success: true,
      posts,
    };
    // 删除cookie直接赋值为null即可
    // this.ctx.session = null;
  }
}
```

## Validate

> 参数校验通过 [Parameter](https://github.com/node-modules/parameter#rule) 完成，支持的校验规则可以在该模块的文档中查阅到。

```js
// config/plugin.js
exports.validate = {
    enable:true,
    package:'egg-validate'
}
```

```js
class PostController extends Controller {
  async create() {
    // 校验参数
    // 如果不传第二个参数会自动校验 `ctx.request.body`
    this.ctx.validate({
      title: { type: 'string' },
      content: { type: 'string' },
    });
  }
}
```

> 自定义校验规则 app.validator.addRule(type,check)

```js
//app.js
module.exports = app =>{
    app.validator.addRule('json',(rule,value)=>{
        try {
            JSON.parse(value)
        } catch(err){
            return 'must be json string'
        }
    })
}

//app/controller/post.js
class PostController extends Controller {
  async handler() {
    const ctx = this.ctx;
    // query.info 字段必须是 json 字符串
    const rule = { info: 'json' };
    ctx.validate(rule, ctx.query);
  }
};
```

# Service

```js
// app/router.js
module.exports = app => {
  app.router.get('/user/:id', app.controller.user.info);
};

// app/controller/user.js
const Controller = require('egg').Controller;
class UserController extends Controller {
  async info() {
    const { ctx } = this;
    const userId = ctx.params.id;
    const userInfo = await ctx.service.user.find(userId);
    ctx.body = userInfo;
  }
}
module.exports = UserController;

// app/service/user.js
const Service = require('egg').Service;
class UserService extends Service {
  // 默认不需要提供构造函数。
  // constructor(ctx) {
  //   super(ctx); 如果需要在构造函数做一些处理，一定要有这句话，才能保证后面 `this.ctx`的使用。
  //   // 就可以直接通过 this.ctx 获取 ctx 了
  //   // 还可以直接通过 this.app 获取 app 了
  // }
  async find(uid) {
    // 假如 我们拿到用户 id 从数据库获取用户详细信息
    const user = await this.ctx.db.query('select * from user where uid = ?', uid);

    // 假定这里还有一些复杂的计算，然后返回需要的信息。
    const picture = await this.getPicture(uid);

    return {
      name: user.user_name,
      age: user.age,
      picture,
    };
  }

  async getPicture(uid) {
    const result = await this.ctx.curl(`http://photoserver/uid=${uid}`, { dataType: 'json' });
    return result.data;
  }
}
module.exports = UserService;

// curl http://127.0.0.1:7001/user/1234
```

# Plugin

> 中间件的定位是拦截用户请求，并在它前后做一些事情，例如：鉴权、安全检查、访问日志等等。但实际情况是，有些功能是和请求无关的，例如：定时任务、消息订阅、后台逻辑等等。

## install

```shell
npm i egg-mysql --save
```

## register

> `plugin.js` 中的每个配置项支持：
>
> - `{Boolean} enable` - 是否开启此插件，默认为 true
> - `{String} package` - `npm` 模块名称，通过 `npm` 模块形式引入插件
> - `{String} path` - 插件绝对路径，跟 package 配置互斥
> - `{Array} env` - 只有在指定运行环境才能开启，会覆盖插件自身 `package.json` 中的配置

```js
// config/plugin.js
// 使用 mysql 插件
exports.mysql = {
  enable: true,
  package: 'egg-mysql',
};
```

## config

> 社区的插件可以 GitHub 搜索 [egg-plugin](https://github.com/topics/egg-plugin)。

```js
// config/config.default.js
exports.mysql = {
  client: {
    host: 'mysql.com',
    port: '3306',
    user: 'test_user',
    password: 'test_password',
    database: 'test',
  },
};
```

# [Schedule](https://eggjs.org/zh-cn/basics/schedule.html)

> 所有的定时任务都统一存放在 `app/schedule` 目录下，每一个文件都是一个独立的定时任务，可以配置定时任务的属性和要执行的方法。
>
> schedule支持如下参数
>
> -  `interval`: 任务执行间隔。
> -  `cron`: 任务执行的时间，cron表达式。
> -  `cronOptions`: 配置 cron 的时区等，参见 [cron-parser](https://github.com/harrisiirak/cron-parser#options) 文档
> -  `immediate`：配置了该参数为 true 时，这个定时任务会在应用启动并 ready 后立刻执行一次这个定时任务。
> -  `disable`：配置该参数为 true 时，这个定时任务不会被启动。
> -  `env`：数组，仅在指定的环境下才启动该定时任务。

```js
// 一个更新缓存的定时任务
// app/schedule/update_cache.js
const Subscription = require('egg').Subscription;

class UpdateCache extends Subscription {
  // 通过 schedule 属性来设置定时任务的执行间隔等配置
  static get schedule() {
    return {
      interval: '1m', // 1 分钟间隔
      type: 'all', // 指定所有的 worker 都需要执行
    };
  }

  // subscribe 是真正定时任务执行时被运行的函数
  async subscribe() {
    const res = await this.ctx.curl('http://www.api.com/cache', {
      dataType: 'json',
    });
    this.ctx.app.cache = res.data;
  }
}

module.exports = UpdateCache;

//可以更优雅的简写--------------
module.exports = {
  schedule: {
    interval: '1m', // 1 分钟间隔
    type: 'all', // 指定所有的 worker 都需要执行
  },
  async task(ctx) {
    const res = await ctx.curl('http://www.api.com/cache', {
      dataType: 'json',
    });
    ctx.app.cache = res.data;
  },
};
```

# app.js

> 框架提供了统一的入口文件（`app.js`）进行启动过程自定义，这个文件返回一个 Boot 类，我们可以通过定义 Boot 类中的**生命周期方法**来执行启动应用过程中的初始化工作 :
>
> - 配置文件即将加载，这是最后动态修改配置的时机（`configWillLoad`）
> - 配置文件加载完成（`configDidLoad`）
> - 文件加载完成（`didLoad`）
> - 插件启动完毕（`willReady`）
> - worker 准备就绪（`didReady`）
> - 应用启动完成（`serverDidReady`）
> - 应用即将关闭（`beforeClose`）