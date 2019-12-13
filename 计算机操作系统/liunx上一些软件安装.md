
在命令模式下，输入:.,$d 一回车就全没了。

whereis python3 

### 查看系统相关命令
```java
查看系统版本
cat /etc/redhat-release

查看时间  date 
查看磁盘信息 df  du
df -h 	 显示磁盘剩余空间 disk free
du -h [目录名] 	 显示目录下的文件大小 disk usage

ps -ef |grep python  # 查看python运行程序
ps -ef|grep python|grep -u dz| grep -v grep # 查看dz用户下的python进程 过滤掉grep进程
# 提取pid(awk以空格分割，显示第二个变量即为pid)
ps -ef|grep elasticsearch|grep -v grep|awk '{print $2}'

查看进程信息 ps
ps aux	process status 查看进程的详细状况
top	动态显示运行中的进程并且排序

kill [-9] 进程代号	终止指定代号的进程，-9 表示强行终止 

-----参数含义------>>>
a	显示终端上的所有进程，包括其他用户的进程
u	显示进程的详细状态
x	显示没有控制终端的进程
--------注意点----->>>
ps 默认只会显示当前用户通过终端启动的应用程序
使用 kill 命令时，最好只终止由当前用户开启的进程，而不要终止 root 身份开启的进程，否则可能导致系统崩溃
要退出 top 可以直接输入 q


```


### yum安装软件卸载
```
安装：yum -y install ...
卸载：yum -y remove  ...
```
### 通配付的使用
```

*	代表任意个数个字符
?	代表任意一个字符，至少 1 个
[]	表示可以匹配字符组中的任一一个
[abc]	匹配 a、b、c 中的任意一个
[a-f]	匹配从 a 到 f 范围内的的任意一个字符

```


### liunx文件和目录常用命令
```Python
查看目录 ls 
ls -a  (ll)显示指定目录下所有子目录与文件，包括隐藏文件
ls -l 以列表方式显示文件的详细信息
ls -h 配合 -l 以人性化的方式显示文件大小

创建文件 touch  ..
创建文件夹 mkdir ..
-p 可以递归创建目录

删除 rm   
-f:强制删除
-i:交互模式，在删除前询问用户是否操作
-r:递归删除，常用在目录的删除

复制 cp
-i	覆盖文件前提示
-r	若给出的源文件是目录文件，则 cp 将递归复制该目录下的所有子目录和文件，目标文件必须为一个目录名

移动 mv
-f:force，强制直接移动而不询问
-i:若目标文件(destination)已经存在，就会询问是否覆盖
-u:若目标文件已经存在，且源文件比较新，才会更新

-------------------------------
Linux 系统中 grep 命令是一种强大的文本搜索工具
grep允许对文本文件进行 模式查找，所谓模式查找，又被称为正则表达式，在就业班会详细讲解
-n	显示匹配行及行号
-v	显示不包含匹配文本的所有行（相当于求反）
-i	忽略大小写
常用的两种模式查找
^a	行首，搜寻以 a 开头的行
ke$	行尾，搜寻以 ke 结束的行


```

### 远程管理常用命令
```java

重启操作系统，now表现在
$ shutdown -r now
立刻关机
$ shutdown mow
系统在今天 20:25关机
$ shutdown 20:25
系统再过十分钟后关机
$ shutdown +10

取消之前制定的关机计划
shutdown -c

重启系统  reboot now

查看网卡 Ip
ifconfig
ping Ip地址

linux系统间文件传输：

# 把本地当前目录下的 01.py 文件 复制到 远程 家目录下的 Desktop/01.py
# 注意：`:` 后面的路径如果不是绝对路径，则以用户的家目录作为参照路径
scp -P port 01.py user@remote:Desktop/01.py

# 把远程 家目录下的 Desktop/01.py 文件 复制到 本地当前目录下的 01.py
scp -P port user@remote:Desktop/01.py 01.py

# 加上 -r 选项可以传送文件夹
# 把当前目录下的 demo 文件夹 复制到 远程 家目录下的 Desktop
scp -r demo user@remote:Desktop

# 把远程 家目录下的 Desktop 复制到 当前目录下的 demo 文件夹
scp -r user@remote:Desktop demo

--------参数含义
-r	若给出的源文件是目录文件，则 scp 将递归复制该目录下的所有子目录和文件，目标文件必须为一个目录名
-P	若远程 SSH 服务器的端口不是 22，需要使用大写字母 -P 选项指定端口
--------------------------------

```


### 文件权限管理 chmod
```java
chmod 可以修改 用户／组 对 文件／目录 的权限
chmod +/-rwx 文件名|目录名    -R 递归修改


# 修改文件|目录的拥有者
chown 用户名 文件名|目录名

# 递归修改文件|目录的组
chgrp -R 组名 文件名|目录名

# 递归修改文件权限
chmod -R 755 文件名|目录名

权限数字对应
r-->4
w-->2
x-->1
拥有者(rwx),组(rwx),其他(rwx)
rwx--7
rw- -->6


---查看log最后10行
tail  -n  10  test.log

tail -n +10 test.log    查询10行之后的所有日志;

head -n 10  test.log   查询日志文件中的头10行日志;

head -n -10  test.log   查询日志文件除了最后10行的其他所有日志;



grep “要查找的字符串” 文件名
```


### tar解压打包命令
```java

# 打包文件
tar -cvf 打包文件.tar 被打包的文件／路径...

# 解包文件
tar -xvf 打包文件.tar
------>>参数含义--------->>
c	生成档案文件，创建打包文件
x	解开档案文件
v	列出归档解档的详细过程，显示进度
f	指定档案文件名称，f 后面一定是 .tar 文件，所以必须放选项最后
--------->>>

gzip 压缩/解压
# 压缩文件
tar -zcvf 打包文件.tar.gz 被压缩的文件／路径...

# 解压缩文件
tar -zxvf 打包文件.tar.gz

# 解压缩到指定路径
tar -zxvf 打包文件.tar.gz -C 目标路径

-C	解压缩到指定目录，注意：要解压缩的目录必须存在
<<<<--------------------------


```





### **chromedriver的安装**
```
    - yum -y update
    - 安装Xvfb('yum -y install Xvfb')
    - self.order('yum -y install libXfont')
    - self.order('yum -y install xorg-x11-fonts*')
    - 命令安装最新的 Google Chrome('yum -y install https://dl.google.com/linux/direct/google-chrome-stable_current_x86_64.rpm')
    - 安装必要的库('yum -y install mesa-libOSMesa-devel gnu-free-sans-fonts wqy-zenhei-fonts')
    - chromedriver的下载地址 https://npm.taobao.org/mirrors/chromedriver
    - chromedriver的存放路径 /usr/bin/

  # 每次运行后都会产生缓存的垃圾文件， rm -rf  /tmp/.com.google.Chrome.* 需要清理掉

```






### nginx 部署web项目
```java

---安装 nginx  yum install nginx

--- 启动 systemctl start nginx.service

--- 查看版本  nginx -v
---查看端看  netstat -tnulp | grep nginx 








-- php 的环境包安装
yum install zlib-devel libxml2-devel libjpeg-devel libjpeg-turbo-devel freetype-devel libpng-devel gd-devel libcurl-devel libxslt-devel libxslt-devel -y




```

###　Supervisor　进程守护安装及其配置
```java

【安装】 pip install supervisor

【生成supervisor默认配置文件】
echo_supervisord_conf > /etc/supervisord.conf

【启动 】

supervisord -c /etc/supervisord.conf 





【centos 配置开启自动启动】

#supervisord.service
-------------------------------------------------
[Unit] 
Description=Supervisor daemon

[Service] 
Type=forking 
ExecStart=/usr/bin/supervisord -c /etc/supervisord.conf 
ExecStop=/usr/bin/supervisorctl shutdown 
ExecReload=/usr/bin/supervisorctl reload 
KillMode=process 
Restart=on-failure 
RestartSec=42s

[Install] 
WantedBy=multi-user.target
---------------------------------------------------


// 把文件深拷贝到 /usr/lib/systemd/system/
cp supervisord.service /usr/lib/systemd/system/

// 启动服务
systemctl enable supervisord


reboot new  重启电脑



-------------------相关命令-----------
supervisorctl status  --#查看所以进程
supervisorctl stop all  --#关闭所有
supervisorctl stop 加要关闭的程序名

restart重启
supervisorctl update #新增任务




systemctl stop firewalld.service  停止
systemctl disable firewalld.service 关闭开启

```







### Liunx定时任务crontab
```Golang

// TODO  查看是否安装：
rpm -qa | grep crontab
 
// TODO 任务存在路径， 对应用户名文件

/var/spool/cron

新增调度任务可用两种方法：
	1)、在命令行输入: crontab -e 然后添加相应的任务，wq存盘退出。
	2)、直接编辑/etc/crontab 文件，即vi /etc/crontab，添加相应的任务。

查看调度任务
	crontab -l //列出当前的所有调度任务
	crontab -l -u jp   //列出用户jp的所有调度任务

删除任务调度工作
	crontab -r   //删除所有任务调度工作


任务编写命令:
 	Minute Hour  Day Month Dayofweek  command
 	分钟    小时  天  月    天每星期     命令

--在这些字段里，除了“Command”是每次都必须指定的字段以外，其它字段皆为可选字段，
可视需要决定。对于不指定的字段，要用“*”来填补其位置


5      *       *       *     *     ll  指定每小时的第5分钟执行一次ll命令
30     5       *       *     *     ll  指定每天的 5:30 执行ll命令
30     7       8       *     *     ll  指定每月8号的7：30分执行ll命令
30     5       8       6     *     ll  指定每年的6月8日5：30执行ll命令
30     6       *       *     0     ll  指定每星期日的6:30执行ll命令[注：0表示星期天，1表示星期1，
以此类推，也可以用英文来表示，sun表示星期天，mon表示星期一等。]
30     3     10,20     *     *     ll  每月10号及20号的3：30执行ll命令[注：“，”用来连接多个不连续的时段]
25     8-11    *       *     *     ll  每天8-11点的第25分钟执行ll命令[注：“-”用来连接连续的时段]
*/15   *       *       *     *     ll  每15分钟执行一次ll命令 [即每个小时的第0 15 30 45 60分钟执行ls命令 ]
在shell脚本中
> ，这个表示向文件重新添加内容，该文件首先被清空
而 >> ，这个表示想文件追加内容，原来的内容不会被清理掉。

```










