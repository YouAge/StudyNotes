#### Mongodb命令

-- 删除
 1. 删除指定的集合，表 `db.user.drop()`

```python
// 连接mongodb数据库
	from pymongo import MongoClient
	cont = MongoClient('ip',27017)
	db_name = 'work'
	db = cont[db_name]
	collection = db['task']
```

```python

-- 更新某字段 
collection.update({"_id":nam["_id"]},{"$unset":{"spider_id":''}}) #删除spider_id 字段 
db.task.update({'date':"2019-05-28",'engineer':'dongzhu'},{$unset:{'match_id':''}},false,true)
collection.update({"_id":nam["_id"]},{"$set":{"spider_id":1}}) #添加spider_id 字段
db.stu.update({},{$set:{gender:0}},{multi:true}) #multi:可选，默认是false，表示只更新找到的第⼀条记录，值为true表示把满足条件全部更新
db.task.updateMany({'finder["district"]':{$regex:'重庆'}},{$set:{"spider_id":1}})

-- 删除当前指向的数据库，如果数据库不存在，则什么也不做
db.dropDatanase()
-- 删除数据 
db.stu.remove({gender:0},{justOne:true})  #参数justOne:可选，如果设为true或1，则只删除⼀条，默认false，表示删除多条


-- 创建集合，指定大小
	name是要创建的集合的名称
	options是⼀个⽂档，⽤于指定集合的配置，选项​​参数是可选的，所以只需要到指定的集合名称
db.createCollection(name, options)
	参数capped：默认值为false表示不设置上限，值为true表示设置上限
	参数size：当capped值为true时，需要指定此参数，表示上限⼤⼩，当⽂档达到上限时，会将之前的数据覆盖，单位为字节
db.createCollection("sub", { capped : true, size : 10 } )


-- 插入数据
db.集合名称.insert(data)



```










### Window 安装MongoDB
```
1. 下载安装
2. cmd /bin/下 输入命令
设置对应的路径和创建文件
mongod --logpath "C:\Program Files\MongoDB\log\mongo.log" --logappend --dbpath "C:\Program Files\MongoDB\data\db" --serviceName "MongoDB" --install





```




