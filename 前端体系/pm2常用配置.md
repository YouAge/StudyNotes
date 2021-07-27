# pm2 使用配置说明



## pm2命令使用（项目更目录下使用）：

```sh
$ pm2 start app.js  启动app项目

$ pm2 list 列出由pm2管理的所有进程信息，还会显示一个进程会被启动多少次，因为没处理的异常。

$ pm2 monit 监视每个node进程的CPU和内存的使用情况

$ pm2 logs 显示所有进程日志

$ pm2 stop all 停止所有进程

$ pm2 restart all 重启所有进程

$ pm2 reload all 0秒停机重载进程 (用于 NETWORKED 进程)

$ pm2 stop 0 停止指定的进程

$ pm2 restart 0 重启指定的进程

$ pm2 startup 产生 init 脚本 保持进程活着

$ pm2 web 运行健壮的 computer API endpoint (http://localhost:9615)

$ pm2 delete 0 杀死指定的进程

$ pm2 delete all 杀死全部进程
```
 

##　启动进程的方式详细：

```sh
$ pm2 start app.js -i max 根据有效CPU数目启动最大进程数目

$ pm2 start app.js -i 3 启动3个进程

$ pm2 start app.js -x 用fork模式启动 app.js 而不是使用 cluster

$ pm2 start app.js -x -- -a 23 用fork模式启动 app.js 并且传递参数 (-a 23)

$ pm2 start app.js --name serverone 启动一个进程并把它命名为 serverone

$ pm2 stop serverone 停止 serverone 进程

$ pm2 start app.json 启动进程, 在 app.json里设置选项

$ pm2 start app.js -i max -- -a 23 在--之后给 app.js 传递参数

```

## json配置文件

```json
{
 "apps": [
 {
  "name": "mywork",
  "cwd": "/srv/node-app/current",
  "script": "bin/xxx",
  "log_date_format": "YYYY-MM-DD HH:mm:ss",
  "error_file": "/var/log/node-app/node-app.stderr.log",
  "out_file": "log/node-app.stdout.log",
  "pid_file": "pids/node-geo-api.pid",
  "instances": 6,
  "min_uptime": "200s",
  "max_restarts": 10,
  "max_memory_restart": "1M",
  "cron_restart": "1 0 * * *",
  "watch": false,
  "merge_logs": true,
  "exec_interpreter": "node",
  "exec_mode": "fork",
  "autorestart": false,
  "vizion": false
 }
 ]
}
```
```
apps:json结构，apps是一个数组，每一个数组成员就是对应一个pm2中运行的应用
name:应用程序名称
cwd:应用程序所在的目录
script:应用程序的脚本路径
log_date_format:
error_file:自定义应用程序的错误日志文件
out_file:自定义应用程序日志文件
pid_file:自定义应用程序的pid文件
instances:
min_uptime:最小运行时间，这里设置的是60s即如果应用程序在60s内退出，pm2会认为程序异常退出，此时触发重启max_restarts设置数量
max_restarts:设置应用程序异常退出重启的次数，默认15次（从0开始计数）
cron_restart:定时启动，解决重启能解决的问题
watch:是否启用监控模式，默认是false。如果设置成true，当应用程序变动时，pm2会自动重载。这里也可以设置你要监控的文件。
merge_logs:
exec_interpreter:应用程序的脚本类型，这里使用的shell，默认是nodejs
exec_mode:应用程序启动模式，这里设置的是cluster_mode（集群），默认是fork
autorestart:启用/禁用应用程序崩溃或退出时自动重启
vizion:启用/禁用vizion特性(版本控制)
```


##　windows配置自起

### 1. 安装

```sh

npm install pm2-windows-startup -g
pm2-startup install
pm2 save  



npm i -g pm2-windows-service
```

### 2. 添加环境变量


> 在系统变量中，添加， 
> PM2_HOME
> D:\.pm2（这路径根据自己需要定） // D:\Program Files\nodejs\node_global\node_modules\pm2  到文件夹即可

### 3. 新开cmd执行 

```sh
pm2-service-install    # 
```
> n， 取消安装就可以了，然后就可以在服务中查看pm2

> 完成后，通过WINDOWS+R  services.msc 查看服务 pm2


### 启动程序

```
cd到当前文件夹，输入pm2 start app.js --name web_examine_system启动，其中--name web_examine_system是给当前项目

```

```sh
pm2 save # 保存到服务列表，开机自起  
```



## pm2 log配置


### 安装

```
pm2 install pm2-logrotate
```

### 查看和配置

```sh
pm2 conf  # 查看

# -----默认配置
pm2 set pm2-logrotate:max_size 10M
pm2 set pm2-logrotate:retain 30
pm2 set pm2-logrotate:compress false
pm2 set pm2-logrotate:dateFormat YYYY-MM-DD_HH-mm-ss
pm2 set pm2-logrotate:workerInterval 30
pm2 set pm2-logrotate:rotateInterval 0 0 * * *
pm2 set pm2-logrotate:rotateModule true
# ---- 


>`max_size` (Defaults to 10M): 配置项默认是 10MB，并不意味着切割出来的日志文件大小一定就是 10MB，而是检查时发现日志文件大小达到 max_size，则触发日志切割。
>`retain` (Defaults to 30 file logs): 这个数字是在任何一个时间保留已分割的日志的数量，这意味着如果您保留7个，那么您将最多有7个已分割日志和您当前的一个
>`compress` (Defaults to false): 对所有已分割的日志启用 gzip 压缩
>`dateFormat` (Defaults to YYYY-MM-DD_HH-mm-ss) : 文件名格式化的规则
>`rotateModule` (Defaults to true) : 像其他应用程序一样分割 pm2模块的日志
>`workerInterval` (Defaults to 30 in secs) : 您可以控制工作线程检查日志大小的间隔(最小值为1）单位为秒（控制模块检查log日志大小的循环时间，默认30s检查一次）
>>`rotateInterval` (Defaults to 0 0 * * * everyday at midnight): 多久备份一次，默认值是0 0 * * *，意思是每天晚上0点分割


pm2 restart all  # 配置生效

```




# 执行Python脚本

```
pm2 start main.py -x --interpreter python

pm2 start -i max python main.py --name xx 

pm2 start  python main.py --name kappa_v1 -i 5
```


"c:\python39\python.exe|c:\python39\lib\site-packages\wfastcgi.py"