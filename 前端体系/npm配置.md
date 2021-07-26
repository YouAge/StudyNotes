## node.js Windons安装配置

```
官网下载直接安装（选择路径安装）
找到安装路径，下创建2个文件
node_cache, node_gloal
然后在cmd中输入
npm config set prefix "D:\Program Files\nodejs\node_global"
npm config set cache "D:\Program Files\nodejs\node_cache"

在在环境配置中修改
系统变量中 添加 
NODE_PATH   D:\Program Files\nodejs\node_cache
个人变量中path 中把原来的npm路径修改 成 D:\Program Files\nodejs\node_global
这样下载的全局安装就会在node_global 文件中 
```


## 杀掉yum进行，
rm -f /var/run/yum.pid   

ps aux | grep yum


npm 修改源。。。
```java

// 设置 淘宝镜像源
npm config set registry https://registry.npm.taobao.org

// 查看 使用的 镜像源
npm config get registry

// 安装 淘宝镜像源
npm install -g cnpm --registry=https://registry.npm.taobao.org

//重置
npm config set registry https://registry.npmjs.org/ 
```


npm install -global --production windows-build-tools


## centos安装node版本管理

```


```



### centos 安装

```

稳定版：

1. 软件预存：
$ yum clean all && yum makecache fast
$ yum install -y gcc-c++ make
$ curl -sL https://rpm.nodesource.com/setup_10.x | sudo -E bash -

2. sudo yum install nodejs -y

3. node -v， npm -v



最新版安装：

$ yum clean all && yum makecache fast
$ yum install -y gcc-c++ make
$ curl -sL https://rpm.nodesource.com/setup_12.x | sudo -E bash -


```







### 相关包安装
```

一. nrm 安装

	1. npm i nrm -g全局安装`nrm`包；
	2. nrm ls查看当前所有可用的镜像源地址以及当前所使用的镜像源地址；
	3. nrm use npm 或nrm use taobao`切换不同的镜像源地址；



```


### npm 错误

```
// unable to get local issuer certificate 错误时
// https://blog.darkthread.net/blog/npm-unable-get-local-issuer-cert-issue/

npm config set strict-ssl false 暫時關閉 SSL 檢查
yarn config set strict-ssl false

# 关闭GIt ssl 检查
git config --global http.sslVerify false
git config --system http.sslverify false
```


