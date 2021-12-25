FROM ubuntu:latest

RUN sed -i 's/deb.debian.org/mirrors.ustc.edu.cn/g' /etc/apt/sources.list
RUN sed -i 's/security.debian.org/mirrors.ustc.edu.cn/g' /etc/apt/sources.list
RUN apt-get update 
RUN apt-get install -y libaio1 numactl libncurses5 net-tools

# 配置文件
COPY --chown=root:root conf/my.cnf  /etc/my.cnf
COPY --chown=root:root conf/config.ini /etc/config.ini
# 可执行文件
WORKDIR /usr/local/mysql/mysql-cluster
COPY ndb ndb/
RUN ln -s  /usr/local/mysql/mysql-cluster/ndb/bin/* /usr/local/bin/
# 数据目录
RUN mkdir data
RUN mysqld --initialize --user=root

# 开机启动
# RUN cp ndb/support-files/mysql.server /etc/init.d/
# RUN chmod +x /etc/init.d/mysql.server
# RUN update-rc.d mysql.server defaults