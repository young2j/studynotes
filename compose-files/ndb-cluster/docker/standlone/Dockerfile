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
RUN ln -s  $PWD/ndb/bin/* /usr/local/bin/
# 数据目录
RUN mkdir -p data/mgmd_data data/ndbd_2_data data/ndbd_3_data data/mysqld_data

# 启动脚本
COPY start.sh start.sh
RUN chmod 755 start.sh

## 启动管理节点
# ndb_mgmd --initial -f /etc/config.ini
# ndb_mgm -e "SHOW"
# ndb_mgm -e "ALL STATUS"

# # 启动数据节点
# ndbd --initial 
# ndbd --initial

# # 启动sql节点
# mysqld --initialize --user=root
# mysqld_safe --user=root &

# mysql -u -p
# ALTER USER user() IDENTIFIED BY 'rootroot';
# use mysql;
# update user set host = '%' where user = 'root';
# flush privileges;

# ctrl+p ctrl+q
# docker run -it --name=ndb-cluster-standlone -v "$PWD/data/mgmd_data:/usr/local/mysql/mysql-cluster/mgmd_data" -v "$PWD/data/ndbd_2_data:/usr/local/mysql/mysql-cluster/ndbd_2_data" -v "$PWD/data/ndbd_3_data:/usr/local/mysql/mysql-cluster/ndbd_3_data" -v "$PWD/data/mysqld_data:/usr/local/mysql/mysql-cluster/mysqld_data" registry.cn-chengdu.aliyuncs.com/yangsj/ndb-cluster:standlone bash
# docker commit -m 'ndb-cluster-standlone' ndb-cluster-standlone registry.cn-chengdu.aliyuncs.com/yangsj/ndb-cluster:standlone
# docker run -d -it --name=ndb-standlone -p 3306:3306 -v data:/usr/local/mysql/mysql-cluster/data registry.cn-chengdu.aliyuncs.com/yangsj/ndb-cluster:standlone bash start.sh
# docker run exec ndb-standlone ndb_mgm -e "SHOW"