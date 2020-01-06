####  阿里云仓库地址
```java
公网： registry.cn-hangzhou.aliyuncs.com/d_storage/fruit_dish
专网： registry-vpc.cn-hangzhou.aliyuncs.com/d_storage/fruit_dish
经典： registry-internal.cn-hangzhou.aliyuncs.com/d_storage/fruit_dish

登入： docker login --username=15607195236 registry.cn-hangzhou.aliyuncs.com 
密码：
上传
docker tag [ImageId] registry.cn-hangzhou.aliyuncs.com/d_storage/fruit_dish:[镜像版本号]
docker push registry.cn-hangzhou.aliyuncs.com/d_storage/fruit_dish:[镜像版本号]

下载pull registry.cn-hangzhou.aliyuncs.com/d_storage/fruit_dish:[镜像版本号]

```


k8s + docker 维护


https://wiki.openvz.org/Download/template/precreated
git clone https://git.oschina.net/dockerf/docker-training.git

yum install kde-l10n-Chinese
yum reinstall glibc-common


####Docker 安装

- 1.windows 旧版安装 Docker Toolbox工具箱
- 2.windows Docker for Windows 安装

- 修改源，
```python 
# 终端连接 docker-machine ssh default（虚拟机名）
# 连接上后输入 sudo sed -i "s|EXTRA_ARGS='|EXTRA_ARGS='--registry-mirror="https://d46cq005.mirror.aliyuncs.com"|g" /var/lib/boot2docker/profile
# 阿里docker源加速 https://d46cq005.mirror.aliyuncs.com
# 重启default 虚拟机
# docker-machine restart default
#远程
#进入default虚拟机 到/var/lib/boot2docker/etc/docker 路径下  进去daemon.json 文件 没有就创建 添加远程连接的地址
{"insecure-registries":["192.168.10.104:5000"]}

# liunx系统中 /etc/docker/daemon.json， 添加加速源
"registry-mirrors": ["https://d46cq005.mirror.aliyuncs.com"]
```

#### Docker Linux版本查看，命令

```python
# liunx系统安装Dockaer 最新版本
curl -sSL https://get.daocloud.io/docker | sh
# 如果没得curl 就先安装 install curl
#开启server
service docker start 
#查看版本
docker version
#守护进程启动
systemctl daemon-reload
#查看docker进程服务
systemctl status docker.service 
# 重启命令
service docker restart
systemctl restart docker 
#关闭docker
service docker stop
systemctl stop docker




```

### Dcoekr镜像导入

```java
// 本地镜像导入
cat centos-7-x86_64.tar.gz | docker import - name:0.1
// 保存镜像到本滴
docker save -o 存为本地镜像名字.tar  需要保存的镜像名称版本号
// 导入用save导出的静下心
docker load -i 镜像.tar包


//设置时间为亚洲上海
ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
//查看ip时bash: ifconfig: command not found 解决办法
yum install net-tools      ##一般是没有安装这个

没有就创建，没有语言就下 yum install kde-l10n-Chinese
解决中问乱码： vim /etc/locale.conf 文件中 添加 LANG="en_US.UTF-8"



----------python 环境配置------------

---- pip 安装 ----
####解决 bash: pip: command not found 办法 pip 安装 
wget https://bootstrap.pypa.io/get-pip.py
python get-pip.py # 在查看版本
##### -------------- 下载安装包  ------------------- ###
## 没有make就安装
yum -y install  make 

#安装python前 先安装python依赖包
yum install -y openssl-devel openssl-static zlib-devel lzma tk-devel xz-devel bzip2-devel ncurses-devel gdbm-devel readline-devel sqlite-devel gcc libffi-devel
wget https://www.python.org/ftp/python/3.7.1/Python-3.7.1.tgz (Gzipped source tarball 连接)

#移动到要安装的路径下
mv Python-3.7.1.tgz   /use/local/
#解压
tar -xvf Python-3.7.1.tgz 
#进入路径，编译安装
cd /usr/local/Python-3.7.1/
./configure
make
make install
#建立软连接
ln -s /usr/local/Python-3.7.1/python /usr/bin/python3

```





####Docker 镜像命令
```python

#docker中 删除所有的镜像
docker rmi -f <image ID>
## docker rmi $(docker images | awk '{print $3}' |tail -n +2)
** tail -n +2 表示从第二行开始读取

## 删除所有正在运行的容器
docker rm -f `docker ps -aq`
docker rm -f  容器ID
# 创建运行容器
sudo docker run -i -t ubuntu /bin/bash


#修改镜像标签
 docker tag 镜像ID  name:
```

#### Docker 容器相关命令
```java
// 进入怎么在运行的容器中 
$:	docker exec -it  xxxx /bin/bash
$: sudo docker attach 容器ID 
// 启动容器
： docker start 容器ID
// 停止容器
$: docker stop 容器ID
// 重启容器
$: docker restart 容器ID
// 文件上传到容器中， 下载文件(路径对调)
$: docker cp /root/test.txt  容器ID:/root/

// 外部查看容器log  不加参数查看全部
$: docker logs -f -t --since="2018-02-08" --tail=100 CONTAINER_ID //查看指定时间后的日志，只显示最后100行：
$: docker logs --since 30m CONTAINER_ID //查看最近30分钟的日志:
$: docker logs -t --since="2018-02-08T13:23:37" CONTAINER_ID  //查看某时间之后的日志：
$: docker logs -t --since="2018-02-08T13:23:37" --until "2018-02-09T12:23:37" CONTAINER_ID  //查看某时间段日志：
$: docker logs -t e15adbcdfc81 | grep error >> logs_error.txt  // 把docker错误日志写入一个文件
// 本机上使用容器内部命令查看
$: docker exec  容器ID 命令(ls /root/) 

// docker中启动所有容器命令
$: docker start $(docker ps -a | awk '{ print $1}' | tail -n +2)
// docker中 关闭所有的容器命令
$: docker stop $(docker ps -a | awk '{ print $1}' | tail -n +2)
// docker中 删除所有的容器命令
$: docker rm $(docker ps -a | awk '{ print $1}' | tail -n +2)
	docker rm $(docker ps -a -q )


// 容器保存镜像
$: docker commit 容器ID  datap3:py3.6

```

docker-entrypoint.sh


### Dockerfile 文件，命令
```java
FROM scratch
ADD centos-7-docker.tar.xz /

LABEL org.label-schema.schema-version="1.0" \
    org.label-schema.name="CentOS Base Image" \
    org.label-schema.vendor="CentOS" \
    org.label-schema.license="GPLv2" \
    org.label-schema.build-date="20181204"

CMD ["/bin/bash"]



// 设置变量
EVNV MYPATH 
// 进入容器的落脚点
WORKDIR /root






```

### Docker运行相关命令
```java
// docker 容器内获取IP
1.通过环境变量传入docker run --env HOST_IP=192.168.0.160，通过环境变量$HOST_IP获取
2.运行docker时绑定hostdocker run --network host，通过ip route获取


```





#### Docker创库
```java







```




####centos 7 python3 安装
```python
配置好Python3.6和pip3
#安装EPEL和IUS软件源
yum install epel-release -y
yum install https://centos7.iuscommunity.org/ius-release.rpm -y

#安装Python3.6
yum install python36u -y
ln -s /bin/python3.6 /bin/python3
-----------
#安装pip3
yum install python36u-pip -y
ln -s /bin/pip3.6 /bin/pip3

```



###MQ安装
```python

——————————————————安装流程——————————————————>>>>

安装依赖包，有可以不用装
yum -y install make gcc gcc-c++ kernel-devel m4 ncurses-devel openssl-devel
yum -y install ncurses-devel

###Erlang 地址： https://github.com/rabbitmq/erlang-rpm
#vim /etc/yum.repos.d/rabbitmq-erlang.repo
[rabbitmq-erlang]
name=rabbitmq-erlang
baseurl=https://dl.bintray.com/rabbitmq/rpm/erlang/20/el/7
gpgcheck=1
gpgkey=https://dl.bintray.com/rabbitmq/Keys/rabbitmq-release-signing-key.asc
repo_gpgcheck=0
enabled=1


#上步骤后，安装MQ时会自动去安装erlang的版本
#MQ下载安装，MQ版本地址https://www.rabbitmq.com/news.html
wget https://github.com/rabbitmq/rabbitmq-server/releases/download/v3.7.11/rabbitmq-server-3.7.11-1.el7.noarch.rpm
yum -y install rabbitmq-server-3.7.11-1.el7.noarch.rpm




教程； https://blog.csdn.net/qq_36194413/article/details/85165382?tdsourcetag=s_pcqq_aiomsg
```

####MQ的命令
```python
#3启动RabbitMQ服务
service rabbitmq-server start
#状态查看
rabbitmqctl status
#启用插件,web管理页面
rabbitmq-plugins enable rabbitmq_management
#重启服务
service rabbitmq-server restart
#添加帐号:name 密码:passwd
rabbitmqctl add_user name passwd
#赋予其administrator角色
rabbitmqctl set_user_tags name administrator
#设置权限
rabbitmqctl set_permissions -p / name ".*" ".*" ".*"

```
###MQ配置

```java

// 修改 登录遇到问题：User can only log in via localhost问题

找到这个文件rabbit.app
/usr/lib/rabbitmq/lib/rabbitmq_server-3.7.7/ebin/rabbit.app

将：{loopback_users, [<<”guest”>>]}，
改为：{loopback_users, []}，
原因：rabbitmq从3.3.0开始禁止使用guest/guest权限通过除localhost外的访问

重启一下： systemctl restart rabbitmq-server.service
--------------------- 

```

### MQ集群配置 docker版本
```












```










### Anaconda，python虚拟环境命令
```python
#Anaconda命令
#创建环境
conda creat --name snowflake biopython
# 进去虚拟环境
activate  flowers 
#列出所以环境
conda info -e
#删除环境
remove -n flowers 
#退出环境
deactivate    ==》 conda  deactivate
#升级当前版本的conda
conda update conda
##

```





```


文件传输 加-r 是文件夹
scp /root/install.* root@192.168.1.12:/usr/local/src

```



```
删除root文件后

先创建 mkdir /root
cp -a /etc/skel/.[!.]* /root
主要是把 /etc/skel/里面的文件拷贝回去就行了

```




