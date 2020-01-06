package main

import (
	"github.com/streadway/amqp"
	"log"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

// 只能在安装 rabbitmq 的服务器上操作
func main() {
	conn, err := amqp.Dial("amqp://guest:guest@10.253.100.29:5672/yunce",)
	failOnError(err, "连接到RabbitMQ失败")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "未能打开通道")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"test_cloud_queue1", // 队列name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "未能声明队列")
	err =ch.QueueBind(
		q.Name,
		"",    // Routing key
		"",  // 交换机
		false,
		nil,
		)

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "注册用户失败")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("收到一个消息: %s", d.Body)
		}
	}()

	log.Printf(" [*] 等待消息。退出按CTRL+C")
	<-forever
}