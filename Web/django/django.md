<p style="text-align:center;font-size:28px;font-weight:bold;">django</p>

```shell
python -m django --version
```

```bash
mkdir macaws
django-admin startproject main macaws

django-admin startproject demosite
demosite
├── manage.py
└── demosite -->main
    ├── __init__.py
    ├── asgi.py
    ├── settings.py
    ├── urls.py
    └── wsgi.py

python manage.py runserver [host:port]
```
```
python manage.py startapp polls
demosite/polls
├── __init__.py
├── admin.py
├── apps.py
├── migrations
│   └── __init__.py
├── models.py
├── tests.py
└── views.py

python manage.py makemigrations polls

python manage.py sqlmigrate polls 0001 # sqlmigrate 命令接收一个迁移的名称，然后返回对应的 SQL

python manage.py check # 这个命令帮助你检查项目中的问题

#检查 INSTALLED_APPS 设置，为其中的每个应用创建需要的数据表,如果不需要某个或某些应用，可以在运行 migrate 前删除
python manage.py migrate [--database=default]
python manage.py migrate <app> 0002 # 撤销0002版本之后的所有迁移
python manage.py migrate <app> zero # 撤销所有迁移使用zero

python manage.py squashmigrations <app> 0004 # 压缩迁移文件
```

```shell
# 创建超级用户
python manage.py createsuperuser

# 创建普通用户
>>> from django.contrib.auth.models import User
>>> user = User.objects.create_user('john', 'lennon@thebeatles.com', 'johnpassword')
# At this point, user is a User object that has already been saved
# to the database. You can continue to change its attributes
# if you want to change other fields.
>>> user.last_name = 'Lennon'
>>> user.save()

# 更改密码
>>> from django.contrib.auth.models import User
>>> u = User.objects.get(username='john')
>>> u.set_password('new password')
>>> u.save()

# 用户验证
from django.contrib.auth import authenticate

user = authenticate(username='john', password='secret')
if user is not None:
    # A backend authenticated the credentials
else:
    # No backend authenticated the credentials
```

```shell
# 在tests.py中编写测试用例
python manage.py test polls
```

以下是自动化测试的运行过程：

- `python manage.py test polls` 将会寻找 `polls` 应用里的测试代码
- 它找到了 [`django.test.TestCase`](https://docs.djangoproject.com/zh-hans/3.1/topics/testing/tools/#django.test.TestCase) 的一个子类
- 它创建一个特殊的数据库供测试使用
- 它在类中寻找测试方法——以 `test` 开头的方法。
- 在 `test_was_published_recently_with_future_question` 方法中，它创建了一个 `pub_date` 值为 30 天后的 `Question` 实例。
- 接着使用 `assertls()` 方法，发现 `was_published_recently()` 返回了 `True`，而我们期望它返回 `False`。

```shell
# 使用虚拟client交互测试
python manage.py shell
>>> from django.test.utils import setup_test_environment
>>> setup_test_environment()
from django.test import Client
# create an instance of the client for our use
>>> client = Client()
>>> from django.urls import reverse
>>> response = client.get(reverse('polls:index'))
>>> response.status_code
```

打包示例

https://docs.djangoproject.com/zh-hans/3.1/intro/reusable-apps/

```shell
python manage.py diffsettings # 显示当前配置文件与 Django 默认配置的差异。
```

