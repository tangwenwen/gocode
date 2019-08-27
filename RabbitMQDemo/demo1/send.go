package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

func main() {
	//建立socket的MQ连接
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672")
	failOnError(err, "failed to connect to RabbitMQ")
	defer conn.Close()
	//建立管道
	ch, err := conn.Channel()
	failOnError(err, "failed to open a channel")
	defer ch.Close()
	//建立队列来存储/转发信息
	q, err := ch.QueueDeclare(
		"hello",
		false,
		false,
		false,
		false,
		nil)
	failOnError(err, "failed to declare q queue")
	body := "Hell"
	err = ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	failOnError(err, "failed to publish a message")
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s:%s", msg, err)
		panic(fmt.Sprintf("%s:%s", msg, err))
	}
}
