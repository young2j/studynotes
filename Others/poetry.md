<p style="text-align:center;font-size:28px;font-weight:bold;">poetry</p>

> **《相比 Pipenv，Poetry 是一个更好的选择》**https://zhuanlan.zhihu.com/p/81025311

### 安装、卸载与更新

```shell
# 安装
# poetry 会自动找到当前环境的python版本，并据此创建环境变量
python get-poetry.py # 下载脚本本地安装
pip install poetry #官方不推荐，可能导致依赖冲突
# 默认安装在用户home目录中，可以通过环境变量改变安装目录
POETRY_HOME=/etc/poetry python get-poetry.py

# 卸载
python get-poetry.py --uninstall
POETRY_UNINSTALL=1 python get-poetry.py

# 更新
poetry self update
```

### 全局选项

```shell
poetry 
--verbose (-v|vv|vvv): Increase the verbosity of messages: "-v" for normal output, "-vv" for more verbose output and "-vvv" for debug.
--help (-h) : Display help information.
--quiet (-q) : Do not output any message.
--ansi: Force ANSI output.
--no-ansi: Disable ANSI output.
--version (-V): Display this application version.
poetry -V
poetry [command] -h
```

### 创建项目

```shell
# 交互式在当前目录生成pyproject.toml,可以是空目录，可以是已有项目
poetry init 
--name: Name of the package.
--description: Description of the package.
--author: Author of the package.
--python Compatible Python versions.
--dependency: Package to require with a version constraint. Should be in format foo:1.0.0.
--dev-dependency: Development requirements, see --require.

# 新建项目 
poetry new polls
polls
├── README.rst
├── polls
│   └── __init__.py
├── pyproject.toml
└── tests
    ├── __init__.py
    └── test_polls.py

# 改变源代码目录名为src,不不改变包名
poetry new polls --src    
polls
├── README.rst
├── pyproject.toml
├── src
│   └── polls
│       └── __init__.py
└── tests
    ├── __init__.py
    └── test_polls.py
# 改变包名同时改名源码目录名
poetry new polls --name demo

polls
├── README.rst
├── demo
│   └── __init__.py
├── pyproject.toml
└── tests
    ├── __init__.py
    └── test_demo.py
    
# pyproject.toml
[tool.poetry]
name = "demo"
version = "0.1.0"
description = ""    
```

### 配置

配置文件的保存路径：

- macOS: `~/Library/Application Support/pypoetry`
- Windows: `C:\Users\<username>\AppData\Roaming\pypoetry`
- unix: `~/.config/pypoetry`

```shell
# 展示当前项目配置 
poetry config --list
# 展示单条配置
poetry config [key]
# 新增、改变配置
poetry config key value1 value2 ... [--local] # 添加--local只会改变当前项目的配置
poetry config cache-dir . # 依赖包缓存目录
poetry config virtualenvs.in-project true # 虚拟环境存放路径
#If set to true, the virtualenv wil be created and expected in a folder named .venv within the root directory of the project.
#If not set explicitly (default), poetry will use the virtualenv from the .venv directory when one is available. If set to false, poetry will ignore any existing .venv directory.

# 删除配置
poetry config key --unset

# 以环境变量作配置，必须以POETRY_开头
export POETRY_VIRTUALENVS_PATH=/path/to/virtualenvs/directory
```

#### cache-dir

- macOS: `~/Library/Caches/pypoetry`
- Windows: `C:\Users\<username>\AppData\Local\pypoetry\Cache`
- Unix: `~/.cache/pypoetry`

#### virtualenvs.create

如果设置为`false`，则必须保证有python环境，并且pip模块可用。

#### virtualenvs.in-project

* 设置为`true`，会在当前项目目录下创建`.env`文件夹，虚拟环境相关会保存在该文件夹中。
* 设置为`flase`，即使当前目录下有`.env`，也会忽视它。

* 默认为None。如果当前目录下存在`.env`，会采用它为虚拟环境目录。

#### virtualenvs.path

虚拟环境创建的路径。默认为`{cache-dir}/virtualenvs`

#### repository.\<name>

仓库路径，默认使用pypi作为包的下载和发布仓库。

```shell
# 配置私有仓库
poetry config repositories.foo https://foo.bar/simple/
# 配置仓库用户名和密码
poetry config http-basic.foo username password

poetry config http-basic.pypi username password
poetry config pypi-token.pypi my-token # 官方推荐使用token，而不是用户名和密码

# 也可以使用环境变量作为配置
export POETRY_PYPI_TOKEN_PYPI=my-token
export POETRY_HTTP_BASIC_PYPI_USERNAME=username
export POETRY_HTTP_BASIC_PYPI_PASSWORD=password
```

```shell
# 配置依赖包下载路径
[[tool.poetry.source]]
name = "tsinghua"
url = "https://pypi.tuna.tsinghua.edu.cn/simple/"
secondary=true # 指定为true，pypi会优先于它，否则反之
default=true # 指定为true，会禁用pypi源
```



### 依赖包管理

```shell
#这个命令会读取 pyproject.toml 中的所有依赖（包括开发依赖）并安装，如果不想安装开发依赖，可以附加 --no-dev 选项。如果项目根目录有 poetry.lock 文件，会安装这个文件中列出的锁定版本的依赖。
poetry install [--no-dev]
poetry install --no-root #install the dependencies only./Do not install the root package (your project)
poetry install --remove-untracked #remove old dependencies no longer present in the lock file


#  specify the extras you want installed
poetry install --extras "mysql pgsql"
poetry install -E mysql -E pgsql


# 安装包
poetry add [pkgname] [-D/--dev] [--path] [--optional] [--dry-run] [--lock]

poetry add django@^2.0 #指定版本
poetry add "django>=3.0"
poetry add django@latest

poetry add git+https://github.com/sdispater/pendulum.git#develop #git分支依赖
poetry add git+https://github.com/sdispater/pendulum.git#2.0.5 # git 版本依赖

# 相当于在pyproject.toml中添加如下
[tool.poetry.dependencies]
requests = { git = "https://github.com/kennethreitz/requests.git", branch = "next" } # Get the latest revision on the branch named "next"
flask = { git = "https://github.com/pallets/flask.git", rev = "38eb5d3b" } # Get a revision by its commit hash
numpy = { git = "https://github.com/numpy/numpy.git", tag = "v0.13.2" } # Get a revision by its tag


poetry add ./my-package/ # 本地依赖
poetry add ../my-package/dist/my-package-0.1.0.tar.gz
poetry add ../my-package/dist/my_package-0.1.0.whl
# 相当于在pyproject.toml中添加如下
[tool.poetry.dependencies]
my-package = { url = "https://example.com/my-package-0.1.0.tar.gz" }

# specify extra pkgs using []
poetry add requests[security,socks]
poetry add "requests[security,socks]~=2.22.0"
poetry add "git+https://github.com/pallets/flask.git@1.1.1[dotenv,dev]"

# 删除包
poetry remove [pkgname] [--dry-run] [--dev/-D]

# 查看包信息
poetry show [--no-dev] [--latest/-l] [--outdated/-o]# 所有依赖包信息
poetry show --tree # 所有包及其依赖关系
poetry show [pkgname] # 某一个包信息
poetry show [pkgname] --tree # 某一个包及其依赖关系
poetry show --outdated # 查看所有可以更新的依赖
poetry show django 

name         : django
version      : 3.1.2
description  : A high-level Python Web framework that encourages rapid development and clean, pragmatic design.

dependencies
 - asgiref >=3.2.10,<3.3.0
 - pytz *
 - sqlparse >=0.2.2

# 更新包
poetry update
poetry update [pkgname1] [pkgname2]
poetry update [--dry-run] # --verbose
poetry update --no-dev # 不更新dev依赖
poetry update --lock # 只更新lock file中的包

# 包缓存管理
poetry cache list

# 搜索包
poetry search requests pendulum

# 锁定依赖
poetry lock
```

### 虚拟环境管理

```shell
# 切换虚拟环境
poetry env use /full/path/to/python # 完整路径
poetry env use python3.7 # PATH中存在
poetry env use 3.7 # 简写
poetry env use system # 使用系统默认版本

# 显示虚拟环境信息
poetry env info

Virtualenv
Python:         3.7.3
Implementation: CPython
Path:           NA

System
Platform: win32
OS:       nt
Python:   C:\Users\jge\home\devs\Miniconda3

# 列出虚拟环境
poetry env list

# 删除虚拟环境
poetry env remove /full/path/to/python
poetry env remove python3.7
poetry env remove 3.7
poetry env remove test-O3eWbxRl-py3.7

# 自动寻找当前环境并执行命令
poetry run python app.py
# 主动进入虚拟环境
poetry shell

# [tool.poetry.scripts]
# my-script = "my_module:main"
poetry run my-script

# 检查pyproject.toml结构和错误
poetry check
```

### 打包

```shell
poetry build [--format/-f] [wheel/sdist] #将以两种格式打包：sdist源码格式+wheel编译格式
```

### 发布

```shell
poetry publish
poetry publish --build # 打包并发布
poetry publish [-repository/-r] my-repository [--username/-u] [--password/-p] [--dry-run] # 发布到仓库,默认pypi，可以通过config更改
```



