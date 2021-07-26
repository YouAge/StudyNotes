# postgresql查询

## window安装

1. 下载地址 <https://www.enterprisedb.com/downloads/postgres-postgresql-downloads>
2. 下一步，全勾上，设置密码， 去掉勾选，直接点 Finish，完成

## 设置pgAdmin4 为中文

1. 点击 file-1-选择语言

## 创建数据库

## python操作postgresql

```python
import psycopg2
conn = psycopg2.connect(database="postgres", user="postgres", password="123456789", host="10.128.212.164", port="5432")
cur = conn.cursor()  # 创建指针对象
cur.execute("select * from statu")
results = cur.fetchall()
conn.commit()
cur.close()
conn.close()
```

## python Pandas 读取和写入数据库


### 读取数据库

```python

import psycopg2
import pandas as pd

# 配置文件
localhost=dict(
    database="demo1",
    user="postgres",
    password="123456789",
    host="localhost",
    port="5432"
)

def redPostGresql(config,sql):
    """读取数据 config配置文件 sql语句"""
    conn = psycopg2.connect(**config)  # 连接数据库
    try:
        data = pd.read_sql(sql, con=conn) # 读取数据库数据
        return data
    except Exception as e:
        print(e)
        return False
    finally:
        conn.close()
# 直接获取数据库数据， 以DataFrame格式返回，（）
IfsstepID = redPostGresql(cnctum0pgsql02,'SELECT * FROM "Report"."IFSStepID"')
```


### 写入数据库
```python 
from sqlalchemy import create_engine
import pandas as pd
# 创建连接
engine = create_engine("postgresql://postgres:123456789@localhost:5432/demo1") 
# 获取要写入的数据，
posData = pd.DataFrame(snWIPlist)
try:
 	"""if_exists：
    * fail: Raise a ValueError.
    * replace: Drop the table before inserting new values.
    * append: Insert new values to the existing table.
    """
    posData.to_sql('wip',engine,index=False,if_exists='replace') # 写入数据库
except Exception as e:
    print(e)
```

### LIKE 模糊查询

#### %

> 任意0个和多个字符

#### _

> 任意单个字符

#### []  

> [正则表达式]

#### [^]



```sql
SELECT *  FROM "public".ifactoryreportdata  WHERE  daytime >='2020-07-09' and daytime <='2020-07-23'  and reportname LIKE 'ic(ng)%'
```


### 时间段查询

```sql
方法一：
select * from user_info where create_date>= '2015-07-01' and create_date < '2015-08-15';
 
方法二：
select * from user_info where create_date between '2015-07-01' and '2015-08-15';
 
方法三：
select * from user_info where create_date >= '2015-07-01'::timestamp and create_date < '2015-08-15'::timestamp;
 
方法四：
 
select * from user_info where create_date between to_date('2015-07-01','YYYY-MM-DD') and to_date('2015-08-15','YYYY-MM-DD');

select * from user_info where create_date between to_date('2015-07-01','yyyy-mm-dd hh24:mi:ss') and to_date('2015-08-15','yyyy-mm-dd hh24:mi:ss');
```


### 表连接

### union不合并重复数据

```sql
select * from T1 union all select * from T2 
```

### union合并重复数据

```sql
 select * from T1 union select * from T2 
```

### except 左查询中返回右查询没有找到的所有非重复值


### intersect 两个结果集的交集（即两个查询都返回的所有非重复值）



### 关联查询

1. `inner join` 内连接查询 (2表 合并，显示相同的数据)

```sql
SELECT a.*,b.* FROM table_a a INNER JOIN table_b b ON a.id=b.id
```

2. `left jion` 左关联查询

```sql
SELECT a.*,b.* FROM table_a a LEFT JOIN table_b b ON a.id=b.id
```

3. right join 右关联查询 

4. 左连接-内连接 （取左表的部分集合，但又不存在右表中）

```sql
SELECT a.*,b.* FROM table_a a LEFT JOIN table_b b ON a.id=b.id WHERE b.id IS NULL

	-- SELECT a.*,b.* FROM table_a a LEFT JOIN table_b b ON a.id=b.id WHERE b.id IS NULL  只需要几秒钟中完成
	-- SELECT a.*,b.* FROM table_a a LEFT JOIN table_b b ON a.id !=b.id   谨慎测试，单机十几万的数据，直接把库跑挂了
	-- 根据测试，，你会发现， 后者比前者的查询速度相差太大了
```
5. 右连接-内连接 (取有表的部分数据，但又不存在左表中)

```sql
SELECT a.*,b.* FROM table_a a RIGHT JOIN table_b b ON a.id=b.id WHERE a.id IS NULL
```


## 创建临时表查询 whit

1. ` WITH abcd AS (select * from table)` 语法

```sql
-- 例如
 WITH abcd AS (
         SELECT a."SN",
            a."Process",
            a."ProcessID",
            to_char(now(), 'yyyy-mm-dd'::text) AS lasttime,
            a."Project"
           FROM "Report"."mainSnWithFirstTimeStationAndResult" a
          WHERE (to_char(now(), 'yyyy-mm-dd'::text) = to_char(a."FirstTime", 'yyyy-mm-dd'::text))
        )
 SELECT a."SN",
    upper((a."Process")::text) AS process,
    a.lasttime,
    a."ProcessID",
    a."Project"
   FROM abcd a
  WHERE ("RawData".isnumeric(a."Process") <> true)
UNION ALL
 SELECT a."SN",
    upper(b."StepName") AS process,
    a.lasttime,
    a."ProcessID",
    a."Project"
   FROM (abcd a
     JOIN "Report"."IFSStepID" b ON (((a."Process")::text = b."StepID")));

   -- 创建 临时表 abcd， 
-- 创建多个临时表， 只需要用 , 分割就好， 最后一定要有sql语句
-- 多个语句语法语法
WITH so1 AS (),so2 AS(),so3 AS ()

```

## 排序ORDER BY 

> 查询后，按什么字段排序 升序：`ASC`；降序：`DESC`；


```sql
-- 例句：

```


## 聚合查询,分组

>GROUP BY我们可以先从字面上来理解，GROUP表示分组，BY后面写字段名，就表示根据哪个字段进行分组，如果有用Excel比较多的话，GROUP BY比较类似Excel里面的透视表。
GROUP BY必须得配合聚合函数来用，分组之后你可以计数（COUNT），求和（SUM），求平均数（AVG）等


### 常用的 聚合查询

>1. `count()` 计数
>2. `sum()` 求和
>3. `avg()` 平均数
>4. `max()` 最大值
>5. `min()` 最小值



```sql
-- 以a.customer, a.factory, a.lineid, a.reportprocess, a.wip分组， 计算sum(a.sn) 的值

SELECT a.customer,
    a.factory,
    a.lineid,
    a.reportprocess,
    a.wip,
    to_char(now(), 'yyyy-mm-dd'::text) AS lasttime,
    sum(a.sn) AS wipsn
   FROM stationamecount a
  GROUP BY a.customer, a.factory, a.lineid, a.reportprocess, a.wip;
```

### HAVING 筛选

> `HAVING`是对于`GROUP BY`对象进行筛选， 可以使用聚合函数筛选

```sql
SELECT a.customer,
    a.factory,
    a.lineid,
    a.reportprocess,
    a.wip,
    to_char(now(), 'yyyy-mm-dd'::text) AS lasttime,
    sum(a.sn) AS wipsn
   FROM stationamecount a
  GROUP BY a.customer, a.factory, a.lineid, a.reportprocess, a.wip
  HAVING a.wip='yes'  -- 筛选除 wip=yes的 组， 
```


## Postgresql 子查询

> 子查询或称为内部查询、嵌套查询，指的是在 PostgreSQL 查询中的 WHERE 子句中嵌入查询语句

```sql
with aa as
    (
        SELECT a.userid,a.username,max(a.jiontime) as jiontime,a.project FROM "public".userhistory as a  WHERE a.department='{0}' {1}
             GROUP BY  a.userid,a.username,a.project
    ),
    tpc as (
		    SELECT b.userid,b.username,b.jiontime,max(a.grade) as grade,b.project FROM "public".userhistory a ,aa b WHERE  a.userid=b.userid AND a.jiontime=b.jiontime GROUP BY 
    b.userid,b.username,b.jiontime,b.project
		)
	SELECT tpc.*,b."Status" FROM tpc, "OparationInfo"."OprationInformation" AS b WHERE b."ID" = tpc.userid {2}

-- {} 主要是Python中字符串格式化的一种操作，条件
```



## 函数集

### 常用函数

#### initcap(首字母大写)
#### lower(字母全部小写)
#### upper(字母全部大写)

### 日期函数


#### now() 获取当前日期
#### localtime 当前时间 `SELECT localtime`
#### localtimestamp 当前日期和时间 `SELECT localtimestamp`
#### current_date 当前日期 `select current_date;`


```sql
to_char(now(), 'yyyy-mm-dd'::text) -- 2020-07-14
```
#### timeofday()当前日期和时间

```sql
SELECT timeofday()
-- Tue Jul 14 09:32:24.894371 2020 HKT
```

```sql
-- 时间计算
select now() + interval '1 years' -- 一年后 
select now() + interval '1 month' -- 一月后
select now() - interval '3 week'  --三周后
select now() + '10 min' -- 10分钟后
select now() + '10 D'      -- 一天后

/*
Abbreviation	Meaning
	Y			Years
	M			Months (in the date part)
	W			Weeks
	D			Days
	H			Hours
	M			Minutes (in the time part)
	S			Seconds
*/
```

#### age() 计算2个时间差

```sql

select age(timestamp '2019-09-15'); -- 9 mons 29 days
select age(now(), timestamp '2022-02-05'); -- -1 years -6 mons -21 days -14:12:19.288452
select age(now(), timestamp '2019-02-05'); -- 1 year 5 mons 9 days 09:48:08.012659
```


#### EXTRACT(field FROM source)时间字段截取

```sql
-- 获取年
select extract(year from now()); 

-- 获取月份
select extract(month from now());    

-- 获取 天数
select extract(day from timestamp '2013-04-13');
SELECT EXTRACT(DAY FROM INTERVAL '40 days 1 minute');

-- 查看今天是一年中的多少天
select extract(doy from now());

```


### 删除

```sql
DELETE FROM 表名称 WHERE 列名称 = 值
```


### 自定义初始化自增长数据

```sql
select setval('op3_topic_id_seq',6000,false);

-- 设置冲当前最大值开始增长 tablename表面
select setval('tablename_id_seq',(select max(id) from tablename))
select setval('op3_users_id_seq',(select max(id) from op3_users))
```


### 转换函数

#### to_number(数值类型的字符):将字符转换为数值
#### to_char(数值或者是日期):将数值或者日期转换为字符
#### to_date(日期格式的字符)：将字符转换为日期



### 自定义函数

#### 函数基本语法

> 实践中，官方函数没法满足需求，就需要自己写函数， 那么了解 函数的语法构建

>自定义函数像内置函数一样返回标量值，也可以将结果集用表格变量返回。sql函数必须有返回值


```sql



```

#### 自定义函数整理

##### 判断为存数字函数

```sql
CREATE
OR REPLACE FUNCTION isnumeric (txtStr VARCHAR) RETURNS BOOLEAN AS $$
BEGIN
	RETURN txtStr ~ '^([0-9]+[.]?[0-9]*|[.][0-9]+)$' ;
END ; $$ LANGUAGE 'plpgsql';
```



### ALTER TABLE 曾删表的列

> alter table命令用于添加、修改、删除一张已经存在表的列, 也可以添加，删除约束

**语法**

```sql
-- table_name ：表名 ; column_name：新列名; datatype：数据类型
ALTER TABLE table_name ADD column_name datatype;
```



### 视图

> 视图： 就是查询语句的别名，生成的新的表

```sql
CREATE VIEW viewName AS select name,id,age FROM student WITH id>5;
-- 这样 viewName 就是一个视图， 可以理解为一张新的表在数据库了
```


## 事务

> 一组操作，要么成功，要么失败


## Postgresql 触发器

> 触发器是数据库的回调函数，它会在指定的数据库事件发生时自动执行/调用

> - PostgreSQL 触发器 触发情况！
>   - 在执行操作之前（在检查约束并尝试插入、更新或删除之前）。
>   - 在执行操作之后（在检查约束并插入、更新或删除完成之后）。
>   - 更新操作（在对一个视图进行插入、更新、删除时）



## Postgresql索引

> 索引是加速搜索引擎检索数据的一种特殊表查询。简单地说，索引是一个指向表中数据的指针。 方便快速检索数据， 
> 创建的索引多了， 也会影响查询效率

>语法：  使用 CREATE INDEX 语句创建索引，它允许命名索引，指定表及要索引的一列或多列，并指示索引是升序排列还是降序排列。
>索引也可以是唯一的，与 UNIQUE 约束类似，在列上或列组合上防止重复条目。
 
### 语法
```sql
创建索引： CREATE INDEX index_name ON table_name;
删除所有： drop index index_name
```

### 索引类型

**类型**|**语法**
---|:--
单列索引|CREATE INDEX index_name ON table_name(column_name);
组合索引|CREATE INDEX index_name ON table_name(column1_name,column2_name);
唯一索引|CREATE UNIQUE INDEX index_name on table_name (column_name);
局部索引|CREATE INDEX index_name on table_name (conditional)
隐式索引|建表时由数据库服务器自动创建，一般为主键约束和唯一约束

> 不管是单列索引还是组合索引，该索引必须是在WHERE子句的过滤条件中使用非常频繁的列。
>唯一索引的使用，一方面提高了查询性能，同时也保护了数据的完整性，不允许任何重复的值插入到表中。

>- 什么情况下避免使用索引？
>   - 表数据较小时不需要使用索引
>   - 进行频繁插入或更新操作的表不需要使用
>   - 不应该使用在含有大量的NULL值的列上
>   - 不使用在频繁操作的列上




```sql
(to_char(now(), 'yyyy-mm-dd'::text) = to_char(a."LastTime", 'yyyy-mm-dd'::text)

 SELECT a.*,b.wipsn,b.wip FROM pvder as a LEFT JOIN "RawData"."wipCounts" as b ON upper(a.factory) = b.factory 
 and upper(a.customer)=b.customer 
 and upper(a.lineid) =  b.lineid 
 and upper(a.reportprocess) = b.reportprocess
 and a.date = b.lasttime WHERE a.date =b.lasttime
----------------------

SELECT a.*,b.wip FROM pvder as a LEFT JOIN "RawData"."wipCount" as b ON upper(a.factory) = b.factory 
and upper(a.customer)=b.customer 
and upper(a.lineid) =  b.lineid 
and upper(a.reportprocess) = b.reportprocess
and a.date =to_char(now(), 'yyyy-mm-dd') WHERE a.date = to_char(now(), 'yyyy-mm-dd')

```

```SQL
 SELECT testtable."Customer",
    testtable."StepName",
    testtable."StepID",
    testtable."DRI",
    testtable.remark,
    testtable.reportname,
    testtable.type,
    testtable.reportid,
    testtable.typeid
   FROM dblink('dbname=BG host=10.128.19.168 port=5432 user=postgres password=123456 '::text, 'SELECT "Customer", "StepName", "StepID", "DRI", remark, reportname, 
       type, reportid,typeid
  FROM public."IFSStepID"'::text) testtable("Customer" text, "StepName" text, "StepID" text, "DRI" text, remark text, reportname text, type text, reportid text, typeid text);
```



## not in 改 NOT EXISTS

```sql   
SELECT spk.sn,spk.stationid FROM spk WHERE spk.sn not IN (SELECT assy2.sn FROM assy2)
语句优化
SELECT * FROM spk WHERE NOT EXISTS(SELECT * FROM assy2 WHERE spk.sn = assy2.sn)
```



## 更新
```sql
UPDATE Person SET FirstName = 'Fred' WHERE LastName = 'Wilson' 
```







SELECT "public"."GetInspectorInfor"('2974241', '', '')