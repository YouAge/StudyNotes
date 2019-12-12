

### Python liunx下安装

```java
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
mv Python-3.7.1.tgz   /usr/local/  
#解压
tar -xvf Python-3.7.1.tgz  // tar -xvf /usr/local/Python-3.7.1.tgz -C /usr/local/ 指定目录下
# 进入路径，编译安装
cd /usr/local/Python-3.7.1/
./configure
make
make install
#建立软连接
2529343637

```






#### Python 面向对象
-->python 语言  -- 鸭子类型

1. 多态 -->  python 语言本事就支持多继承，多态


// 接口类和抽象类  再 python当中的应用并不是非常重要




#### python 封装

--- 广义上面向对象的封装：代码的保护， 面向对象的思想本事就是一种封装，只能让自己的对象调用自己类方法

--- 侠义上的封装 ---  面向对象上的三大特性之一 ，  属性 和方法都藏起来 不让你看距

```python
class ClassName(object):
	"""docstring for ClassName"""
	def __init__(self, name, passed):
		super(ClassName, self).__init__()
		self.name = name
		self.__passwd = passed  #私有属性

	def get_pwd(self):
		print(self.__dict__)
		return self.__passwd

alex = ClassName('alex','alex665')
print(alex.__passwd) #调用不到
print(alex._ClassName__passwd) # 可以调用到 —类名__属性名

print(alex.get_pwd())

```

--- > property  内置装饰器函数， 直再面向对象中使用



### Python测试函数时间装饰器
```python
def get_time(func):
    def wraper(*args, **kwargs):
        start_time = time.time()
        result = func(*args, **kwargs)
        end_time = time.time()
        print("Spend:", end_time - start_time)
        return result
    return wraper


```



#### Python loggin 





#### hashillb







#### python时间格式互转，时间获取

```python 

import datetime
import time


# 字符串转 datetime, 但是不能自动补齐，不能有缺少
dataestr = '2018-12-25 11:35:56'
datat = datetime.datetime.strptime(datestr, "%Y-%m-%d %H:%M:%S")

## 获取当前时间
now = datetime.datetime.now()
## 获取昨天的时间
yes_now = now + datetime.timedelta(days=-1)
## 获取几个小时以前的时间
yes_now = now + datetime.timedelta(hours=-1)
## 获取星期以前的时间
yes_now = now + datetime.timedelta(weeks=-1)
## 获取去年时间
datetime.date.today() - relativedelta(months=12)

# 当前时间戳
sp = time.time()

## 把datetime装成字符串
now.strftime("%Y-%m-%d %H:%M:%S")
## 把字符串装成datetime
datetime.datetime.strptime(dataestr, "%Y-%m-%d %H:%M:%S")
## 把字符串转成时间戳形式
time.mktime(time.strptime(dataestr, "%Y-%m-%d %H:%M:%S"))
## 把时间戳转成字符串形式
time.strftime("%Y-%m-%d %H:%M:%S", time.localtime(sp))
## 把datetime类型转外时间戳形式
time.mktime(now.timetuple())



## 时间加减， 把数据都装成时间类型，进行加减，再获取天数，秒数(seconds)

```


###  python list操作
```python

# 列表分割成多个列表

a = [ i for i in range(20)]
lsits = [ a[i:i+4] for i in range(0,len(a),4)     ]



```




   <input type="hidden" name="csrf_token" value="{{ csrf_token() }}"/>






####   python pandas

```python
import pandas as pd

// 读写excel文件

#使用pandas读取excel文件
xls_file=pd.ExcelFile('./data/workbook.xls')
xls_file.sheet_names#显示出读入excel文件中的表名字
table1=xls_file.parse('first_sheet')
table2=xls_file.parse('second_sheet')

xlsx_file=pd.ExcelFile("./demo.xlsx")
x1=xlsx_file.parse(0)
x2=xlsx_file.parse(1)

#excel文件的写出
#data.to_excel("abc.xlsx",sheet_name="abc",index=False,header=True)  #该条语句会运行失败，原因在于写入的对象是np数组而不是DataFrame对象,只有DataFrame对象才能使用to_excel方法。

DataFrame(data).to_excel("abc.xlsx",sheet_name="123",index=False,header=True)
#以数组形式
#excel文件和pandas的交互读写，主要使用到pandas中的两个函数,一个是pd.ExcelFile函数,一个是to_excel函数



// 一起写入excel文件 

for xin in collection.find():
	data = dict(
	            时间=xin['publish_date'],
	            标题=xin['title'],
	            正文=xin['content'],
	            来源=xin['source_name'],
	            区县=xin['district'],
	            链接 =xin['url']
	        )
	ali.append(data)
data = pd.DataFrame(ali)
data.to_excel('./新闻.xlsx',sheet_name='Data1',index=False)




import pandas as pd
# 存cvs
data = pd.DataFrame({'类': [xin['_id']]})
data.to_csv('wenti.csv', mode='a', encoding='utf-8', header=False, index=True)



//2种转换
df = pd.read_csv('xxx.csv')
df.to_excel('xxxx.xlsx',sheet_name='Data1',index=False)


```




### python 字符串，列表，元組，字典，流程控制 

```python
## 切割 



# 压缩和解压
import zlib
data='hello world'
data = zlib.compress(data）#压缩
#data此时的值为：'xx9cxcbHxcdxc9xc9W(xcf/xcaIx01x00x1ax0bx04]'
data=zlib.decompress(data) #解压
#data此时的值为：'hello world'




```

### Python 内置函数
```python

hasattr()  函数用于判断对象是否包含对应的属性。   
------>hasattr(object, name)







```


### Python异步IO asyncio

```python






```


## python scrapy 去重
```python

bloomfiter文件最后一行中添加去重key的过期时间
self.server.expire(self.key, 100)



```


### lxml删除节点
```

ele = tree.xpath('//script')
for e in ele:
    e.getparent().remove(e)


```




### Git命令使用
```java

// 本地git 代码管理

管理文件，初始化
$ git init

查看当前文件夹的状态
$ git status  
对指定文件进行版本控制
git add 文件
git add .  对指定文件下的所有文件及子目录进行版本控制
创建提交记录（版本）
git commit -m ‘详细的描述信息’

推送到远程git
git push -u origin master

查看log
git log
查看所有的log
git reflog 

回归到某一个版本
git reset --hard 版本号

删除本地缓存
git rm -r --cached.
重新添加，上传


帮助我们暂时存储已经开发了一些新功能的代码，继续开发
缓存
零时新功能，移走修改过没有提交的
git stash 
git stash pop 合并第一个缓存
git stash list 缓存过的所有的记录
git stash clear  清除
git stash apply  合并哪个一个
git stash drop 清除哪一个

123.com

git branch
创建分支 dev
git branch dev

删除
git branch -d 分支 
切换分支
git checkout 分支

git mast dev 合并2
```



### python asyncio异步编程

```python

简介：asyncio模块提供了使用协程构建并发应用的工具。它使用一种单线程单进程的的方式实现并发，应用的各个部分彼此合作, 可以显示的切换任务，一般会在程序阻塞I/O操作的时候发生上下文切换如等待读写文件,或者请求网络。同时asyncio也支持调度代码在将来的某个特定事件运行，从而支持一个协程等待另一个协程完成，以处理系统信号和识别其他一些事件。




import asyncio
 
 
async def main():
    print("主协程")
    print("等待result1协程运行")
    res1 = await result1()
    print("等待result2协程运行")
    res2 = await result2(res1)
    return (res1,res2)
 
 
async def result1():
    print("这是result1协程")
    return "result1"
 
 
async def result2(arg):
    print("这是result2协程")
    return f"result2接收了一个参数,{arg}"
 
 
if __name__ == '__main__':
    loop = asyncio.get_event_loop() 
    try:
        result = loop.run_until_complete(main())
        print(f"获取返回值:{result}")
    finally:
        print("关闭事件循环")
        loop.close()



------call_soon---call_later-----call_at--用法------------------
clal_soon  优先
调用不同函数, loop.clal_soon(callback,*arge,context=None) #即刻执行
call_later 指定时间后执行
call_at 

import asyncio
import functools
 
 
def callback(args, *, kwargs="defalut"):
    print(f"普通函数做为回调函数,获取参数:{args},{kwargs}")
 
 
async def main(loop):
    print("注册callback")
    loop.call_soon(callback, 1) # 
    wrapped = functools.partial(callback, kwargs="not defalut")
    loop.call_soon(wrapped, 2)
    await asyncio.sleep(0.2)
 
 
if __name__ == '__main__':
    loop = asyncio.get_event_loop()
try:
    loop.run_until_complete(main(loop))
finally:
    loop.close()




```





#### python虚拟环境使用
```python


$ pip install virtualenv
$ virtualenv venv  #创建一个venv名字的虚拟环境
$ source venv/bin/activate  #激活虚拟环境
$ virtualenv -p /usr/bin/python3 venv  # -p参数指定Python解释器程序路径
$ deactivate  # 退出虚拟环境
$ rm -rf venv  # 删除虚拟环境




pycahrm，密钥
http://jetbrains-license-server
```












