package main

//

 // https://mp.weixin.qq.com/s/jgQYQx7YDH_XQ9ZkBQDmkA
//


// TODO [下载地址](https://studygolang.com/dl)
// TODO 文档 https://studygolang.com/pkgdoc
// TODO 免费在线解密 pdf文件 https://smallpdf.com/cn/unlock-pdf
/* TODO Windows 环境配置
1. 网站下载： https://studygolang.com/dl  语言安装包
2. 直接下一步安装， 选择安装路径即可
3. 配置环境变量：系统变量中添加
  变量： GOROOT
  	值： C:\Go\  （就是安装go的路径）， 好像会自动配置，有就不用管了

需要配置： GOPATH 路径。 ---》 就是写代码的路径配置，  使用Goland软件可以不用在系统中配置，
我一般是在Goland 中配置的， (软件设置中Go下面就能看到)
系统中配置 ： 先创建创建一个项目文件目录，然后再里面创建三个文件夹
比如：
	new_golang/scr/
	new_golang/bing/
	new_golang/pak/
系统变量中添加 :  变量： GOPATH  值： D:\new_golang\

TODO 目录简介
$ GOPATH目录下预定三个子目录
>> src 存放源代码(比如：.go .c .h .s等)   按照golang默认约定，go run，go install等命令的当前工作路径（即在此路径下执行上述命令）。
>> pkg 编译时生成的中间文件（比如：.a）　　golang编译包时
>> bin 编译后生成的可执行文件（为了方便，可以把此目录加入到 $PATH 变量中，如果有多个gopath，那么使用${GOPATH//://bin:}/bin添加所有的bin目录）

go命令依赖一个重要的环境变量：$GOPATH
GOPATH允许多个目录，当有多个目录时，请注意分隔符，多个目录的时候Windows是分号;
当有多个GOPATH时默认将go get获取的包存放在第一个目录下


go get

go get会做两件事：
1. 从远程下载需要用到的包
2. 执行go install

go install

go install 会生成可执行文件直接放到bin目录下，当然这是有前提的
你编译的是可执行文件，如果是一个普通的包，会被编译生成到pkg目录下该文件是.a结尾



*/


// TODO Goland 软件中运行代码
/*
刚开始学习的时候 新建.go文件总是要新建文件夹，这样很麻烦，
使用Goland， 一个package 多个文件 都想直接运行的话， 运行时 把
run kind  改成 file 运行单个文件即可， 如果是 package就是运行不了的

*/

/*
TODO 编译
使用go build
1. 在项目目录下执行 go build
2. 在其他路径下执行go build ,需要再后面加上项目的路径，编译之后的可执行文件就保存再当前目录下
3. go build -o hello.exe
*/

