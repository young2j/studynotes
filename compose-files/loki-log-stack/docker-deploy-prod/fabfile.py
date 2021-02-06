import os
from fabric.api import env, task, cd, run, execute
from fabric.contrib.project import rsync_project
from fabric.contrib.console import confirm


env.user = "root"
env.hosts = ["139.9.240.168"]
env.port = 22


@task
def deploy_loki():
    rsync_project(remote_dir="/data/opt", local_dir="./loki")
    with cd("/data/opt/loki"):
        run("docker-compose down")
        run("docker-compose up -d")
        run("docker image prune -f")


@task
def deploy_nginx():
    rsync_project("/etc/nginx/sites-available", "./nginx")
    with cd("/etc/nginx/sites-available"):
        run("nginx -s reload")


def deploy_the_promtail(proj):
    proj_dir = os.path.join("./promtail", proj)
    remote_dir = "/data/opt/loki-promtail"
    # rsync_project(remote_dir, proj_dir)
    with cd(remote_dir):
        # run("docker-compose down")
        # run("docker-compose up -d")
        # run("docker image prune -f")
        run("ls")


promtail_projs = os.listdir("./promtail")


@task
def deploy_promtail(proj="all"):
    if proj == "all":
        yes = confirm("需要部署全部项目吗？", default=False)
        if not yes:
            return

        for p in promtail_projs:
            execute(deploy_the_promtail, p)
    else:
        if proj not in promtail_projs:
            print("请输入正确的项目文件名!!!")
            return
        execute(deploy_the_promtail, proj)


if __name__ == "__main__":
    deploy_the_promtail()
