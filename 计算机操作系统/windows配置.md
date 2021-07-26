
## 修改cmd携带管理权限

```
运行中 搜索：  `gpedit.msc`
 计算机配置 -> Windows 设置 -> 安全设置 -> 本地策略 -> 安全选项

  "用户帐户控制：以管理员批准模式运行所有管理员", 禁用掉
```




# iis 下部署vue 静态网页

## 1. 打包好vue编译后的文件 找一个部署的目录存放

```
```

## 2.启动IIS，点击网址，右键添加网站，

1. 随便输入项目名称， 
2. 选择网址文件路径， 就是刚刚vue项目的路径，到有 index.html 下
3. 设置端口即可， 点击运行
4. 访问网址，如果出现权限问题， 需要右键配置权限， 添加账号密码
5. 如果刷新页面，出现400，需要配置url重写功能， （可以直接再 vue项目下，添加 web.config 文件）

```xml
<?xml version="1.0" encoding="UTF-8"?>
<configuration>
<system.webServer>
<rewrite>
<rules>
<rule name="Handle History Mode and custom 404/500" stopProcessing="true">
<match url="(.*)" />
<conditions logicalGrouping="MatchAll">
<add input="{REQUEST_FILENAME}" matchType="IsFile" negate="true" />
<add input="{REQUEST_FILENAME}" matchType="IsDirectory" negate="true" />
</conditions>
<action type="Rewrite" url="/" />
</rule>
</rules>
</rewrite>
</system.webServer>
</configuration>
```


## 3. 如果还是出现权限文件， "/"

1. 出现 "/"应用程序中的服务器错误 访问被拒绝...
2. 错误消息 401.3: 您无权使用您提供的凭据查看此目录或页(由于访问控制列表而导致访问被拒绝)

**解决办法**
```
iis里找到有问题的网站，右键编辑权限，安全，高级；
再权限组中，添加  Everyone  账号进去给权限即可
再把最下面的， （使用可以此对象继承的权限项目替代所有子对象权限...）打勾
```



# 如果windows服务器系统 桌面被删除 

```js
1. 使用下面命令，等待安装，后重启即可, （试过可行）

c:\user\absb> Dism /online /enable-feature /featurename:Server-Gui-Mgmt /featurename:Server-Gui-Shell /featurename:ServerCore-FullServer /all
2. 或者使用(未试用)

c:\user\absb> dism /online /enable-feature /all /featurename:servercore-fullserver /featurename:server-gui-shell /featurename:server-gui-mgmt
```