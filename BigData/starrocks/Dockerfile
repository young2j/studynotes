FROM centos:centos7

# RUN sed -i "s/archive.ubuntu.com/mirrors.aliyun.com/g" /etc/apt/sources.list
# RUN sed -i "s/deb.debian.org/mirrors.aliyun.com/g" /etc/apt/sources.list

# Prepare StarRocks Installer.
RUN mkdir -p /data/deploy/
ADD StarRocks-2.4.0.tar.gz /data/deploy/

# Install Java JDK.
RUN yum -y install java-1.8.0-openjdk-devel.x86_64
RUN rpm -ql java-1.8.0-openjdk-devel.x86_64 | grep bin$
RUN /usr/lib/jvm/java-1.8.0-openjdk-1.8.0.352.b08-2.el7_9.x86_64/bin/java -version

# Create directory for FE meta and BE storage in StarRocks.
RUN mkdir -p /data/deploy/StarRocks-2.4.0/fe/meta
RUN jps
RUN mkdir -p /data/deploy/StarRocks-2.4.0/be/storage

# Install relevant tools.
RUN yum -y install mysql net-tools telnet

# Run Setup script.
COPY run_script.sh /data/deploy/run_script.sh
RUN chmod +x /data/deploy/run_script.sh
CMD /data/deploy/run_script.sh