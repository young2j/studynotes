import os


def runserver():
    os.system("python ./demosite/manage.py runserver")


def makemigrations():
    os.system("python ./demosite/manage.py makemigrations polls")


def migrate():
    os.system("python ./demosite/manage.py migrate")


def flakeformat():
    os.system("python -m flake8 .")


def shell():
    os.system("python ./demosite/manage.py shell")
