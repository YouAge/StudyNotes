package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

// 定义一个全局的pool
var pool *redis.Pool

func IsError(err error) {
	if err != nil {
		fmt.Println(err)
		return
	}

}

func selectDb(conn redis.Conn,db int) (reply interface{}, err error) {
	// 选择db
	return conn.Do("SELECT",db)
}

func redisClient()  {
	//初始发 redis连接

	pool = &redis.Pool{
		MaxIdle:8, // 最大的空闲连接数，表示即使没有redis连接时依然可以保持N个空闲的连接，而不被清除，随时处于待命状态。
		MaxActive:0, //最大的激活连接数，表示同时最多有N个连接
		IdleTimeout:200, // 最大的空闲连接等待时间，超过此时间后，空闲连接将被关闭
		Dial: func() (conn redis.Conn, e error) {
			return redis.Dial("tcp","localhost:6379")

		},
	}


}

func setRedis()  {

	conn := pool.Get() // 获取一个连接
	defer conn.Close()
	selectDb(conn,7)
	_,err :=conn.Do("Set","go_set","dop")
	IsError(err)
	selectDb(conn,9)
	r,err :=conn.Do("Set","go_set","dop8")
	IsError(err)
	fmt.Println(r)

}




func main()  {
	redisClient()
	setRedis()
}