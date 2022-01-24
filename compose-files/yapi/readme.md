```shell
# 1. 运行mongo
docker-compose up -d mongo
# 2. 初始化yapi的数据库
docker run -it --rm \
  --link mongo:mongo \
  --entrypoint npm \
  --workdir /yapi/vendors \
  -v $PWD/config.json:/yapi/config.json \
  registry.cn-hangzhou.aliyuncs.com/anoyi/yapi \
  run install-server
# 3. 运行yapi
docker run -d \
    --name yapi \
    --link mongo:mongo \
    --workdir /yapi/vendors \
    -p 3000:3000 \
    -v $PWD/config.json:/yapi/config.json \
    registry.cn-hangzhou.aliyuncs.com/anoyi/yapi \
    server/app.js
```
