import os
from fabric.api import env, task, cd, run, execute, roles, settings
from fabric.contrib.project import rsync_project
from fabric.contrib.console import confirm


env.user = "opt"
env.roledefs = {
    "loki": ["139.155.5.58"],
    "test2": ["119.27.191.102"]
}

# =========================Loki===============================


@roles("loki")
@task
def deploy_loki():
    """部署loki
    """
    remote_dir = "/data/opt/loki-cluster"
    rsync_project(remote_dir, local_dir="./loki/")
    with cd(remote_dir):
        run("ls")
        run("docker-compose down")
        run("docker-compose up -d")
        run("docker image prune -f")


@roles("loki")
@task
def deploy_nginx():
    """部署loki的nginx配置
    """
    remote_dir = "/etc/nginx"
    with settings(user="root"):
        rsync_project(remote_dir, "./nginx/nginx.conf")
        rsync_project(remote_dir+'/sites-enabled', "./nginx/loki-cluster.conf")
        with cd(remote_dir):
            run("ls")
            run("ls sites-enabled")
            run("sudo nginx -s reload")


# =========================Promtail===============================

HOSTS = os.listdir("./promtail")


def deploy_one_promtail(host):
    """部署单个host的promtail
    """
    if host not in HOSTS:
        print("请输入正确的host文件名!!!")
        return
    host_dir = os.path.join("./promtail", host) + "/"
    remote_dir = "/data/opt/loki-promtail"
    rsync_project(remote_dir, host_dir)
    with cd(remote_dir):
        run("ls")
        run("docker-compose down")
        run("docker-compose up -d")
        run("docker image prune -f")


@task
def deploy_promtail(host="all"):
    """promtail部署的入口
    """
    if host == "all":
        all_yes = confirm("需要部署全部项目吗(默认否)？", default=False)
        if not all_yes:
            input_host = confirm("请输入部署的主机文件名(如test1)：")
            if not input_host:
                return
            execute(deploy_one_promtail, input_host, roles=[input_host])
            return

        for host in HOSTS:
            execute(deploy_one_promtail, host, roles=[host])
    else:
        execute(deploy_one_promtail, host, roles=[host])
