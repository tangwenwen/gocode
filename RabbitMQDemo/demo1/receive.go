package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"sync"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672")
	failOnError(err, "Failed to connect to server")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to connect to channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello", //name
		false,   //durable
		false,   //delete when usused
		false,   // exclusive
		false,   //no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")
	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil)
	failOnError(err, "Failed to register a consumer")

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		for d := range msgs {
			log.Printf("received a message:%s", d.Body)
		}
	}()
	wg.Wait()
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s:%s", msg, err)
		panic(fmt.Sprintf("%s:%s", msg, err))
	}
}
