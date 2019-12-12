
##  MySQL数据库常用的命令


### MySQL数据库安装
```python
----windows下安装，下一步

-----liunx下安装--

## 网上资料 
yum install mysql
yum install mysql-server
yum install mysql-devel
### 在安装mysql-server是出现报错 
"""
Loaded plugins: fastestmirror
Loading mirror speeds from cached hostfile
 * base: mirrors.sina.cn
 * extras: mirrors.sina.cn
 * updates: mirrors.sina.cn
No package mysql-server available.
Error: Nothing to do"""
方法一 安装 yum install mariadb-server mariadb 
"""
mariadb数据库的相关命令是：
systemctl start mariadb  #启动MariaDB
systemctl stop mariadb  #停止MariaDB
systemctl restart mariadb  #重启MariaDB
systemctl enable mariadb  #设置开机启动
所以先启动数据库
"""
## 启动数据库，进入后感觉怪怪 的 mysql -u root -p

放法二 官网下载mysql-server 
## 下载地址  https://dev.mysql.com/downloads/repo/yum/ 找到适合版本的liunx
rpm -ivh（安装软件包并显示安装进度） 版本 .rpn
yum -y install mysql-community-server

"""
service mysqld start (5.0版本是mysqld) 启动
service mysql start (5.5.7版本是mysql)
----停止----
service mysqld stop
----重启----
service mysqld restart 
service mysql restart (5.5.7版本命令)
----查看MySQL运行状态--
systemctl status mysqld.service

2、使用 mysqld 脚本启动：
/etc/inint.d/mysqld stop
"""

------------------初始登入------->
## 如果首次登入不上，需要查看一下零时密码
## 查看临时密码 grep "password" /var/log/mysqld.log 
"""
使用临死密码登入， 成功后
ALTER USER 'root'@'localhost' IDENTIFIED BY 'Root_12root';
只能先修改密码 不能进行任何操作，设置密码也有规范不能太简单
先修改一个符合条件的密码 Root_12root
查看密码设置条件 SHOW VARIABLES LIKE 'validate_password%';
通过设置 
set global validate_password.policy=0;
set global validate_password.length=1;
这样就可以修改简单密码了，在重置一下密码

ALTER USER 'root'@'localhost' IDENTIFIED BY 'mysql';
"""

----------------注-----
如果使用方法二去覆盖方法一的安装， 此时会出现MySQL 服务器开不起来的解决方法
出现错误 Failed to start MySQL Server
"""
查看 cat /etc/my.cnf
datadir=/var/lib/mysql
socket=/var/lib/mysql/mysql.sock

删除--   rm -rf /var/lib/mysql/*
在启动mysql服务器， 安装方法二的操作

"""


```



### 解决Navicat 连接MySQL 8.0.11 出现2059错误	
```python
mysql -uroot -ppassword #登录

use mysql; #选择数据库

## --->  select user,plugin from user where user='root';   查看加密方式
ALTER USER 'root'@'localhost' IDENTIFIED BY 'password' PASSWORD EXPIRE NEVER; #更改加密方式

ALTER USER 'root'@'localhost' IDENTIFIED WITH mysql_native_password BY 'password'; #更新用户密码

FLUSH PRIVILEGES; #刷新权限

-----------注-------------
"""
修改时出现 ERROR 1819 (HY000): Your password does not satisfy the current policy requirements 错误
需要修改数据库的策略

查看mysql 初始的密码策略  SHOW VARIABLES LIKE 'validate_password%';
需要设置密码的验证强度等级，设置 validate_password_policy 的全局参数为 LOW 即可
set global validate_password.policy=LOW;
当前密码长度为 8 ，如果不介意的话就不用修改了
 set global validate_password.length=6; 

update user set password=('123.com') where user='root';

关于 mysql 密码策略相关参数；
1）、validate_password_length  固定密码的总长度；
2）、validate_password_dictionary_file 指定密码验证的文件路径；
3）、validate_password_mixed_case_count  整个密码中至少要包含大/小写字母的总个数；
4）、validate_password_number_count  整个密码中至少要包含阿拉伯数字的个数；
5）、validate_password_policy 指定密码的强度验证等级，默认为 MEDIUM；
关于 validate_password_policy 的取值：
0/LOW：只验证长度；
1/MEDIUM：验证长度、数字、大小写、特殊字符；
2/STRONG：验证长度、数字、大小写、特殊字符、字典文件；
6）、validate_password_special_char_count 整个密码中至少要包含特殊字符的个数；
--------------------- 
"""

-------------（liunx）如果mysql登入密码忘记，--------------------
首先 vim /etc/my.cnf 
在[mysql] 下面增加一行  skip-grant-tables
重启mysql重新登入不用输入密码
更具网上的重置更新密码的语法 update user set authentication_string=password('123.com') where user='root'; 如果报语法错误就使用下面一条，密码已经是加密过的 为mysql  （先查看一样，密码的加密方式是否是 mysql_native_password）是的就是可行的，
update mysql.user set authentication_string='*E74858DB86EBA20BC33D0AECAE8A8108C56B17FA where user'='root';



-----查看密码---
select host,user,authentication_string from user;


------------更改mysql远程连接---------------
方式一 直接修改修改  update user set host = '%' where user = 'root';
 


------查看修改后的状态---
select host, user from user;

```






### mysql数据库基础命令操作
```SQL

----------------数据库-------
--查看所有数据库
>> show databases;
--使用数据库
>> use 数据库名；
--查看当前使用的数据库
>> select database();
--创建数据库
>> create database 库名 charset=utf8;
--删除数据库
>> drop database 库名;

--------------数据表------------
--查看当前数据库中所以的表
>> show tables;
-- 查看表结构
>> desc 表名;


----- 创建班级表 -----
>> create table classes (
    id int unsigned auto_increment primary key not null,
    name varchar(30) not null
);
-- students表
create table students(
    id int unsigned primary key auto_increment not null,
    name varchar(20) default '',
    age tinyint unsigned default 0,
    height decimal(5,2),
    gender enum('男','女','中性','保密') default '保密',
    cls_id int unsigned default 0,
    is_delete bit default 0
);

-- 导入数据 
>> insert into 表名 values (数据)
例如：
-- 向students表中插入数据
insert into students values
(0,'小明',18,180.00,2,1,0),
-- 向classes表中插入数据
insert into classes values (0, "python_01期"), (0, "python_02期");

```


### mysql表的增删改查
```sql
<<<<<<<<<<<<<<<<<<<<基础查询>>>>>>>>>>>>>>>>>>>>>>>>>>>
--查询表格所以字段
>> select * from 表名;
-- 指定字段查询
>> select 字1,字2,.. from 表名;
-- 对于表名 跟字段名长的可以通过 as 别名
>> select id as 标号, name as 名字, gender as 性别 from students;
-- 跟表名起别名
>> select s.id,s.name,s.gender from students as s;
-- 在select后面列前使用distinct可以消除重复的行
>> select distinct 列1,... from 表名;
	例：
	>> select distinct gender from students;

>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>条件查询<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
-- 使用where子句对表中的数据筛选，结果为true的行会出现在结果集中
>> select * from 表名 where 条件;
	例：
	>> select * from students where id=1;
----where后面支持多重运算符，进行比较
	>> 比较运算符 | 逻辑运算符| 模糊查询| 范围查询 | 空判断
---比较运行符 
	-- 等于: = | 大于: > |大于等于: >= |小于: < | 小于等于: <= | 不等于: != 或 <>
>> 例如：
	-- 例1：查询编号大于3的学生
	>> select * from students where id > 3;
	--例2：查询编号不大于4的学生
	>> select * from students where id <= 4;
	--例3：查询姓名不是“黄蓉”的学生
	>> select * from students where name != '黄蓉';
	--例4：查询没被删除的学生
	>> select * from students where is_delete=0;

--- 逻辑运算符、
	>> and | or | not
>> 例如：
	--例5：查询编号大于3的女同学
	>> select * from students where id > 3 and gender=0;
	--例6：查询编号小于4或没被删除的学生
	>> select * from students where id < 4 or is_delete=0;

-- 模糊查询
	like | %表示任意多个任意字符 | _表示一个任意字符
>> 例如：
	--例7：查询姓黄的学生
	select * from students where name like '黄%';
	--例8：查询姓黄并且名字是一个字的学生
	select * from students where name like '黄_';
	--例9：查询姓黄或叫靖的学生
	select * from students where name like '黄%' or name like '%靖';

-- 范围查询
	in 表示在一个非连续的范围内 | between ... and ...表示在一个连续的范围内
>> 例如：
	--例10：查询编号是1或3或8的学生
	select * from students where id in(1,3,8);

	--例11：查询编号为3至8的学生
	select * from students where id between 3 and 8;
	
	--例12：查询学生是3至8的男生
	select * from students where id between 3 and 8 and gender=1;

-- 空判断
	判空 is null  | 判非空is not null
>> 例如：
	--例13：查询没有填写身高的学生
	select * from students where height is null;

	--例14：查询填写了身高的学生
	select * from students where height is not null;

	--例15：查询填写了身高的男生
	select * from students where height is not null and gender=1;

-- 优先级
	--优先级由高到低的顺序为：小括号，not，比较运算符，逻辑运算符
	--and比or先运算，如果同时出现并希望先算or，需要结合()使用

>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>排序<<<<<<<<< order by <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<

>> select * from 表名 order by 列1 asc|desc [,列2 asc|desc,...]
	--将行数据按照列1进行排序，如果某些行列1的值相同时，则按照列2排序，以此类推
	--默认按照列值从小到大排列（asc）
	--asc从小到大排列，即升序
	--desc从大到小排序，即降序

	--例：查询未删除男生信息，按学号降序
	select * from students where gender=1 and is_delete=0 order by id desc;
	--例：查询未删除学生信息，按名称升序
	select * from students where is_delete=0 order by name;

>>>>>>>>>>>>>>>>>聚合函数<<<<<<<<<总数 count(*)<<最大值max()<<最小值min()<<求和sun()<<平均值avg()<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
-- count聚合函数不能在 where 中使用
>> 例子：
	--查询学生总数
	select count(*) from students;
	--查询女生的编号最大值
	select max(id) from students where gender=2;
	--查询未删除的学生最小编号
	select min(id) from students where is_delete=0;
	--查询男生的总年龄
	select sum(age) from students where gender=1;
	-- 平均年龄
	select sum(age)/count(*) from students where gender=1;
	--查询未删除女生的编号平均值
	select avg(id) from students where is_delete=0 and gender=2;

```


### mysql删除
```sql

>> delete from 表名 where 条件

例如：
	>> delete from students where id=5;

--逻辑删除，本质就是修改操作,表中有逻辑删除的字段is_delete
>> update students set is_delete=1 where id=1;


>>>>>>>备>>份>>> mysqldump>>>>>>>>>>>
--- 在未登入mysql下进行
mysqldump -uroot -p 库名 > python.sql; 
-- 提示输入mysql密码， 在
>>>>>导入备份数据，>>>>>>>>>>
-- 库名
mysql -uroot -p 库名 < python.sql


```


### mysql数据库创建表格关系
```sql

一对一关系表格创建






```







### 禅道的liunx上搭建
```java
-------------->>直接搭建

wget 下载版本

// 解压到/opt下
tar -zxvf ZenTaoPMS.11.4.1.zbox_old.64.tar.gz -C /opt/
// 进入录下
cd /opt/zbox/ 
// 开启禅道项目
./zbox start 
/* 会出现  
Start Apache success
Start Mysql success   
*/
// 开启数据库管理，需要设置帐号密码
./auth/adduser.sh
/*
Account: root
Password: Adding password for user root
*?


```












### jQuery 3.x

![](https://i.imgur.com/8UtgQNt.png)

```javaScript
1. 改变标签
2. 改变属性
3. 

查找标签 
$('')  -->$('#d1') 
找a标签class是c2的$('a.c2')
用逗号分隔，组合查找 

层级选择器
1.子子孙孙 --> $('x y')
2.儿子选择器 -->$('x>y')
3,弟弟选择器 -->$('x~y')
4.毗邻选择器 -->$('laber+input')

二 筛选器

1. 基本选择器
	1、 ：first
	2、 ：last
	3、 ：eq() -->重0开始




```








