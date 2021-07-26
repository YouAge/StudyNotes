

# gitlab搭建

> <https://about.gitlab.com/install/#centos-7> 安装地址

##　安装发送邮件功能


```
sudo yum install postfix
sudo systemctl enable postfix
sudo systemctl start postfix
```

```
curl https://packages.gitlab.com/install/repositories/gitlab/gitlab-ee/script.rpm.sh | sudo bash

yum install -y gitlab-ee
```


## gitlab命令

gitlab-ctl reconfigure 更新配置

1. 关闭自动启动： systemctl disable gitlab-runsvdir.service
2. 开启自动启动: systemctl enable gitlab-runsvdir.service
3.  启动systemctl start gitlab-runsvdir.service

  查看服务状态：gitlab-ctl status
  开启服务： gitlab-ctl start
  关闭服务：gitlab-ctl stop 



## 备份gitlab


1. 使用备份命令

```sh
gitlab-rake gitlab:backup:create
```

2. 备份的文件默认使用存错在 `/var/opt/gitlab/backups/`


##删除旧的备份文件

随着时间的推移gitlab备份文件越来越多，服务器的磁盘空间也不够大。
此时我们就要删除部分旧的备份文件，gitlab也提供了删除旧的备份文件功能。该功能在gitlab的配置文件中，进行配置即可。
在此以保留7天之前的备份文件为例


１.　`vi /etc/gitlab/gitlab.rb`


```sh
gitlab_rails[‘backup_keep_time’] = 604800
```

其中backup_keep_time是以秒为单位进行计算的，然后执行命令gitlab-ctl reconfigure即可



## gitlab创库恢复


> 新服务器上的gitlab的版本号必须与创建备份时的gitlab版本号相同。

>要验证gitlab备份的有效性，我们可以把该备份文件复制到已经安装好gitlab服务器的/var/opt/gitlab/backups/目录下。然后进行数据恢复，最后访问并查看其数据完整性即可。
通过gitlab备份文件可以恢复gitlab所有的信息，包括仓库、数据库、用户、用户组、用户密钥、权限等信息。


### 停止相关数据连接服务器

```
gitlab-ctl stop unicorn
gitlab-ctl stop sidekiq
```

### 恢复gitlab仓库

```
gitlab-rake gitlab:backup:restore BACKUP='备份文件'
```

### 重启服务器，

```
gitlab-ctl start
```

>`https://www.ilanni.com/?p=13890` 参考地址















# Git 命令


## Git global setup

```git
git config --global user.name "Administrator"
git config --global user.email "admin@example.com"
```

## Create a new repository
```
git clone http://localhost:8888/gitlab-instance-516e7dd4/monitoring.git
cd monitoring
touch README.md
git add README.md
git commit -m "add README"
git push -u origin master
```
## Push an existing folder
```
cd existing_folder
git init
git remote add origin http://localhost:8888/gitlab-instance-516e7dd4/monitoring.git
git add .
git commit -m "Initial commit"
git push -u origin master
```
##　Push an existing Git repository
```
cd existing_repo
git remote rename origin old-origin
git remote add origin http://localhost:8888/gitlab-instance-516e7dd4/monitoring.git
git push -u origin --all
git push -u origin --tags
```


