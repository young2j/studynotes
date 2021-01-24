# gitlab-cicd 全流程示例

# 背景

现在很多公司的项目代码都托管在gitlab上，可以利用gitlab的cicd配置实现项目的持续集成与部署。通过在项目根目录中创建`.gitlab-ci.yml`文件，当代码提交到远程仓库中时，gitlab可以自动识别该文件，并运行文件中配置的流程。对于已经熟悉ci的配置及运行的人来说，这其实并不是一件难事，但是对于刚接触的人就不一定了。我在刚接触ci时，尽管搜过不少资料，第一次实际操作起来仍然一脸抓瞎。于是在gitlab上创建了一个小项目，进行了各种尝试。

贴图为证，我足足花了周末两天时间，提交了63次，才彻底跑成功了整个流程。本文以此项目为示例，讲解gitlab ci如何配置、运行以及自动部署你的项目。

![image-20210123215702980](D:\programs\studynotes\gitlab-cicd\docs\image-20210123215702980.png)

> 当然，本文默认你已经掌握了docker的基本使用。(没有接触过docker的人建议去尝试一下下，我从接触到到目前也才使用两个月，挺香的)

# 准备工作

## 运行的项目

首先，需要准备一个用来运行的项目，这个项目可以是任何你喜欢的语言、你喜欢的框架项目，这不是关键。作为示例，本文使用的是`go`语言写的程序，程序主要是从`mysql`中查询了一行记录然后返回。我们需要先了解下项目的目录结构，这有助于稍后理解ci文件中的部分内容：

![image-20210123220640822](D:\programs\studynotes\gitlab-cicd\docs\image-20210123220640822.png)

* `.gitlab-ci.yml`：这是本文的主角，gitlab会自动识别该文件名，然后运行你定义的工作流程；

* `db.sql`：里面是sql语句，主要用于在ci中创建测试所需的数据库和测试数据。

* `deploy.sh`：这个是拿来自动部署时运行的命令行，主要工作是 拉取最新镜像=>运行最新镜像=>清除tag为none的无用镜像；

  ```bash
  #!/bin/bash
  
  docker-compose down
  docker-compose pull
  docker-compose up -d prod
  docker image prune -f
  ```

* `docker-compose.yml,Dockerfile,Dockerfile.for_go_build`：docker相关的三个文件，`deploy.sh`中的命令根据`docker-compose.yml`中的内容运行应用镜像。`Dockerfile`将项目构建为应用部署的镜像，`Dockerfile.for_go_build`文件专门用于构建一个额外的镜像，这个镜像通常会安装许多额外的依赖包，主要为编译go程序、测试等工作提供基础依赖环境。这样可以保证在生产环境运行的应用镜像体积小、干净纯粹。

  ```dockerfile
  # Dockerfile.for_go_build 构建为 registry.cn-chengdu.aliyuncs.com/yangsj/go-build-alpine:latest
  FROM golang:1.15-alpine
  RUN sed -i "s/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g" /etc/apk/repositories
  RUN apk --update --no-cache add mariadb-client g++ rsync openssh
  RUN go env -w GOPROXY=https://goproxy.cn,direct
  ```

  ```dockerfile
  # Dockerfile 构建为应用程序运行的镜像 registry.cn-chengdu.aliyuncs.com/yangsj/gitlab-cicd:latest
  FROM registry.cn-chengdu.aliyuncs.com/yangsj/go-build-alpine:latest AS builder
  WORKDIR /code
  COPY . /code
  RUN go fmt && go mod tidy
  RUN go build main.go

  FROM alpine:latest
  WORKDIR /code
  COPY --from=builder /code/main /code/main
  CMD [ "/code/main" ]
  ```

* 其他文件： 不重要，忽略。

## 运行的工具

现在准备好了可以运行的项目，还需要准备可以运行ci的工具人——`gitlab-runner`。但是，注意如果你的gitlab中已经存在可以运行的runner了，那么下面的步骤可以省略了。现在假定不存在任何runner，需要我们自己为项目注册一个runner：

首先安装这个工具人：

  ```bash
# linux[ubuntu]
sudo apt install  gitlab-runner
# Mac
brew install gitlab-runner
  ```

安装好了之后，可以进行两种操作：

1. 在本地用runner运行你ci中的job，例如需要运行测试流程，我们指定runner使用docker通过本地镜像来运行test_job：

   ```bash
   gitlab-runner exec docker test_job --docker-pull-policy=never
   ```

2. 为你gitlab中的项目注册runner，这样当你在进行代码提交、打tag等操作时，runner就会运行你的ci配置。可以输入如下命令进行注册：

   ```bash
   gitlab-runner register
   ```

   具体如何注册runner，这里就没有必要细说了，网上一搜一大把。这里需要说明的是，企业开发注册runner需要通过参数提供安全证书，个人普通用户可以直接进行注册。

工具人注册好了，在你的项目`settings->CI/CD`中就可以看到该runner：

![image-20210124001218072](D:\programs\studynotes\gitlab-cicd\docs\image-20210124001218072.png)

# CI文件说明

此部分详细说明.gitlab-ci.yml中的全部内容。为了便于后文对ci配置的理解，这里截个图（如下）。这是为`dev`分支定义的两个阶段`stages`(`test、build`)和对应的`jobs`(`test_job、build_job`)， `build`阶段依赖于`test`阶段，图中`test stage`正在运行：

![image-20210123212633857](D:\programs\studynotes\gitlab-cicd\docs\image-20210123212633857.png)

## CI阶段

ci中的各个阶段`stages`具体视项目实际需要而定，你可以定义多个阶段，也可以只定义一个。各阶段可以前后依赖、也可以各自独立，组成了整个工作流程。 示例项目中我定义了三个，分别为测试阶段`test`、编译阶段`build`和部署阶段`deploy`：

```yaml
stages:
  - test
  - build
  - deploy
```

这三个阶段意味着我定义的工作流程是：测试阶段运行测试程序=>测试成功后编译打包我的应用=>部署编译打包后的应用程序。这个流程你可以自由发挥。

## 测试阶段

定义好了`stages`后，我们需要为各个`stage`定义具体的工作内容，我定义的测试阶段需要做如下工作：

```yaml
test_job:
  stage: test
  image: registry.cn-chengdu.aliyuncs.com/yangsj/go-build-alpine:latest
  environment:
    name: test_ci
  variables:
    MYSQL_ROOT_PASSWORD: rootroot
    MYSQL_DATABASE: hello
    TEST_CICD_ENV: dev
  services:
    - name: mysql:5.7.17
      alias: mysql
  before_script:
    - mysql --version
    - mysql -uroot -p$MYSQL_ROOT_PASSWORD -hmysql < ./db.sql
   
  script:
    - go fmt && go mod tidy
    - CGO_ENABLED=0 go build main.go
    - go test -v
  
  # tags:
  #   - gitlab-cicd
  only:
    - dev
```

* `stage`: 我的`test_job`属于`test` 阶段；

* `image`: 运行`test_job`的基础环境由`go-build-alpine:latest`镜像提供，runner在运行`test_job`时会去拉取这个镜像作为基础运行环境;

* `variables`: 指定运行这个`job`所需要的环境变量为 mysql数据库的密码、mysql使用的数据库以及`TEST_CICD_ENV`这个变量;`TEST_CICD_ENV`是`main.go`运行需要的变量，根据其不同的值读取数据库不同的配置，所以你可以在`variables`中自由定义你需要的环境变量，然后在程序中读取它;

* `services`: 定义`test_job`运行的依赖服务为mysql5.7.17，并取别名为mysql，这其实是runner帮你运行了一个mysql容器;

* `before_script`：在运行正式脚本前，进行一些预处理，比如显示mysql版本、登录mysql并通过`db.sql`创建测试数据库及测试数据。这个不是必须的;

* `script`: 主要的测试命令，示例为编译go程序、并运行程序中所有测试文件`*_test.go`。`script`内容是必须的，否则runner会运行失败，并提示你需要有此部分内容;

* `tags`: 指定运行`test_job`的runner的tags为`gitlab-cicd`。runner的`tags`在注册时输入的。（如果你发现你的runner一直处于pending状态， 那么或许可以注掉这个`tags`试试。我尝试了好久才成功运行的，百度一大堆都是叫你编辑runner，勾选下面这个选项，然而对我不管用：)

  ![image-20210124011211942](D:\programs\studynotes\gitlab-cicd\docs\image-20210124011211942.png)

* `only`: 指定这个`test_job`只在`dev`分支起作用。

## 编译阶段

当测试通过后，我们需要编译程序，然后自动构建并上传镜像。`build_job`主要定义了这些内容：

```yaml
build_job:
  stage: build
  dependencies:
    - test_job
  environment:
    name: test_ci
  variables:
    MAIN_IMAGE: registry.cn-chengdu.aliyuncs.com/yangsj/gitlab-cicd:latest
    DOCKER_HOST: tcp://docker:2375
  image: docker:latest
  services:
    - docker:dind
  before_script:
    - docker login --username=<填你的> --password=<填你的> registry.cn-chengdu.aliyuncs.com
  script:
    - set -ex
    - docker build . -t $MAIN_IMAGE
    - docker push $MAIN_IMAGE
    - docker image prune -f
  only:
    - dev
```

* `stage`: 指定`build_job`属于`build`阶段;
* `dependencies`: 指定`build_job`依赖于`test_job`，只有依赖的`test_job`运行成功后才会运行`build_job`;
* `environment`: 指定`build_job`与`test_job`属于同一个环境，即ci测试阶段，这是自由定义的名称；
* `variables`: 指定环境变量，这里`MAIN_IMAGE`代表了要构建的镜像名，因为太长了，指定为变量后可以通过`$MAIN_IMAGE`来代替，`DOCKER_HOST`指定`docker daemon`监听的服务地址。
* `image、services`: 在ci中运行docker，是一个"docker in docker"问题，这两个是必须的，你可以认为这是一个固定写法组合；
* `before_script`: 运行docker前，进行登录，这个视具体需要相应更改；
* `script`: 构建镜像=>推送镜像=>删除无用镜像
* `only`: runner只在`dev`分支运行。

> `build_job`是一个`docker in docker` 问题， 也就是在docker中运行docker，这是一个并不推荐的做法，网上也有分析各种理由，这里只是作为演示这么做了。实际`build_job`在这个小示例中确实能够运行成功，但是长时间会处于运行中，有时会达到20-30分钟才会成功。总之，比较迷，因为我在公司的项目中也这么干了，但是却跑不起来。所以，我在实际使用中，采用的是手动构建镜像并推送。

## 部署阶段

构建镜像并推送至仓库后，我们需要自动在服务器上拉取最新的镜像并运行，示例中采用的是`deploy.sh`命令行根据`docker-compose.yml`文件来启动应用。所以，我们需要在部署阶段，将项目中的`docker-compose.yml`以及`deploy.sh`两个文件上传至服务器指定目录，然后运行`deploy.sh`。而上传文件，本文使用`rsync`通过`ssh`协议上传至服务器:

```yaml
deploy_job:
  stage: deploy
  dependencies:
    - build_job
  image: registry.cn-chengdu.aliyuncs.com/yangsj/go-build-alpine:latest
  environment:
    name: deploy_prod
  before_script:
    - mkdir -p ~/.ssh
    - eval $(ssh-agent -s)
    - echo "$TEST_KNOWN_HOST" > ~/.ssh/known_hosts
    - echo "$TEST_CONFIG" > ~/.ssh/config
    - echo "$SSH_PRIVATE_KEY" > ~/.ssh/id_rsa
    - chmod 0600 ~/.ssh/id_rsa
    - cat ~/.ssh/id_rsa
    - ssh-add
  script:
    - rsync -av docker-compose.yml deploy.sh root@huaweiyun:/dockers
    - ssh huaweiyun "cd /dockers;bash deploy.sh;"
  only:
    - master
```

* `stage`: `deploy_job`属于`deploy`阶段;
* `dependencies`:  `deploy_job`依赖于 `build_job`；
* `image`：指定`go-build-alpine:latest`为基础运行环境；
* `environment`: 指定 `deploy_job`的环境名为`deploy_prod`；
* `before_script`:  这部分主要在做ssh配置，以让你的runner能够通过ssh来上传部署需要的文件；
* `script`: 通过`rsync`将`docker-compose.yml、deploy.sh`上传至华为云服务器的`/dockers`目录中=>然后再ssh到华为云服务器并执行命令行`cd /dockers;bash deploy.sh;`
* `only`: 指定 `deploy_job`只在`master`分支运行。

# ssh话题

如果你对ssh的流程不是很了解，那么上述部署阶段中`before_script`的内容，你可能会很懵圈，不知道在做什么。如果是，推荐你去看下阮一峰大佬的[ssh教程](http://wangdoc.com/ssh)。这里针对部署阶段的ssh内容做个简要的解释：

假定用户名为root，当我们要通过ssh连接到服务器时，需要输入命令：

```bash
ssh [-p 22] root@139.9.240.168
```

因为每次都要输入`用户名@服务器地址`，所以可以在`~/.ssh/config`文件中配置：

```
# ~/.ssh/config
Host huaweiyun
  Port 22
  HostName 139.9.240.168
  User root
  IdentityFile ~/.ssh/id_rsa
```

然后就可以直接输入如下简短的命令进行登录：

```bash
ssh huaweiyun
```

通过密钥登录的过程有如下四个步骤：

![image-20210124022728719](D:\programs\studynotes\gitlab-cicd\docs\image-20210124022728719.png)

1. 首先需要将你的公钥`id_rsa.pub`中的内容存入服务器中`authorized_keys`文件中，这一步你可以手动拷贝，也可以直接运行如下命令自动上传至服务器：

   ```bash
   ssh-copy-id -i ~/.ssh/id_rsa.pub root@139.9.240.168 # 注意服务器地址换成你自己的。
   ```


2. 当你通过命令ssh至服务器时，会询问你陌生服务器是否是你需要连接的目标服务器，选择yes会将服务器公钥的指纹存入本地的`known_hosts`文件中，下次连接时就不会再次询问；
3. 服务器会发送随机数据至本地，本地通过私钥`id_rsa`进行加密后发送至服务器；
4. 服务器通过公钥进行解密，通过则连接成功。(图中4画错了，公钥肯定是与本地私钥成对的公钥，不是服务器的公钥，这里不重新画了。)

所以，为了在ci中，能够自动通过ssh进行连接，我们需要准备4件事情：

1. 将你的公钥存至服务器；

2. 将你电脑中服务器的公钥指纹保持至ci的`~/.ssh/known_hosts`文件中。查看你电脑中目标服务器指纹的命令是：

   ```bash
   ssh-keyscan -p 22 139.9.240.168 # 请换成你的服务器端口和ip
   ```

3. 将你的ssh配置保存至ci的` ~/.ssh/config`文件中。其实你也可以不用config的形式，直接通过`username@ip`的形式连接，但是这样后续命令中如果需要再次ssh，就会重复编写冗长的代码。推荐config的形式，保密性更好，更简洁。

4. 将你的私钥保存至ci的 `~/.ssh/id_rsa`文件中，同时修改权限为0600, 否则ci会提示你文件权限过大而失败。

> 关于ssh会话ssh-agent不了解的可以自行去学习，其实就是避免重复输入密码。

ci中整个ssh的操作，其实就是在模拟你本地的ssh环境，因为你本地配置好后可以直接ssh到服务器，所以只需要拷贝相应的值到ci相应的文件中，ci就能够通过这些配置直接ssh到服务器了，然后就可以根据脚本自动上传文件。在实际的开发中，一般不会是你自己工作的电脑的ssh配置，往往是一个公共的配置，大家在ci中都使用一样的配置。

最后，再来看`deploy_job`中这段代码， 应该是很简单的了：

```yaml
  before_script:
    - mkdir -p ~/.ssh
    - eval $(ssh-agent -s)
    - echo "$TEST_KNOWN_HOST" > ~/.ssh/known_hosts
    - echo "$TEST_CONFIG" > ~/.ssh/config
    - echo "$SSH_PRIVATE_KEY" > ~/.ssh/id_rsa
    - chmod 0600 ~/.ssh/id_rsa
    - cat ~/.ssh/id_rsa
    - ssh-add
```

其中，为了安全性，具体的ssh配置值，如`$TEST_CONFIG`配置为runner中的变量，因为不可能将ssh的相关配置值直接暴露出来。关于runner变量的配置，可以在gitlab的`settings->CI/CD->Variables`中配置，如图：

![image-20210124031349109](D:\programs\studynotes\gitlab-cicd\docs\image-20210124031349109.png)

# 最后

gitlab的cicd体验下来，发现功能其实并不是很强大，`docker in docker`问题是最失败的，不知道专业的CICD工具如何，例如目前使用最多的`jenkens`，有时间也捣鼓捣鼓下。水平有限，如有不当之处，还请指正。


