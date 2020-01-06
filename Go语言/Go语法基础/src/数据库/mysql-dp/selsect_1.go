package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"math/rand"
	"time"
)

// mysql 配置常量
const (
	user = "root"
	password = "mysql"
	network = "tcp"
	server = "localhost"
	port = 3306
	database = "test"
	genre = "charset=utf8"
)


func IsError(err error) {
	if err != nil {
		fmt.Println(err)
		return
	}

}

// 插入，查询
func showDb() {
	// 连接mysql
	db, err := sql.Open("mysql", "root:mysql@tcp(127.0.0.1:3306)/test?charset=utf8")
	IsError(err)
	defer db.Close()
	//
	var result sql.Result
	// 向数据库中插入一条数据
	result, err = db.Exec("insert into person(name,age,IsBoy) values (?,?,?)",
		"张三", 18, true)
	result, err = db.Exec("insert into person(name,age,IsBoy) values (?,?,?)",
		"张三2", 18, true)
	IsError(err)
	lastId, _ := result.LastInsertId()
	fmt.Println("新插入的数据id为：", lastId)

	var row *sql.Row
	// 返回一行数据,
	row = db.QueryRow("select * from person")
	row = db.QueryRow("select * from person")

	var name string
	var id, age int
	var isBoy bool
	err = row.Scan(&id, &name, &age, &isBoy)
	IsError(err)
	fmt.Println(id, name, age, isBoy)

	// 再插入一条数据
	result, err = db.Exec("INSERT INTO person(name,age,IsBoy) values (?,?,?)", "老王", 20, false)
	result, err = db.Exec("INSERT INTO person(name,age,IsBoy) values (?,?,?)", "老王", 20, false)
	result, err = db.Exec("INSERT INTO person(name,age,IsBoy) values (?,?,?)", "老王", 20, false)
	print("----------------------\n")
	// 返回所有数据
	var rows *sql.Rows
	rows, err = db.Query("SELECT * FROM person")
	IsError(err)
	defer db.Exec("truncate table person")
	defer rows.Close() // 注意，Query 返回的 rows，取完数据后需要调用 Close 来释放资源
	for rows.Next() {
		var name string
		var id, age int
		var isBoy bool
		rows.Scan(&id, &name, &age, &isBoy)
		fmt.Println(id, name, age, isBoy)
	}

}

func main() {
	insertionDB()
}

func insertionDB()  {
	// 插入一万条数据 ，所用时间
	dns := fmt.Sprintf("%s:%s@%s(%s:%d)/%s?%s",user,password,network,server,port,database,genre)
	db,err := sql.Open("mysql",dns)
	defer db.Close()
	IsError(err)
	db.SetMaxOpenConns(100) // 设置最大连接数
	db.SetConnMaxLifetime(100*time.Second) // 设置最大连接周期，超过时常就差close
	db.SetMaxIdleConns(15) // 设置闲置连接数

	var smt *sql.Stmt
	smt,err = db.Prepare("insert into person(name,age,IsBoy) values (?,?,?)")
	IsError(err)
	fmt.Println("开始插入数据...",time.Now())
	r:= rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0;i<10000;i++{
		_,err =smt.Exec(fmt.Sprintf("张%d",r.Int()),r.Intn(50),r.Intn(100)%2)
		IsError(err)
	}
	fmt.Println("数据插入完成：",time.Now())
}
