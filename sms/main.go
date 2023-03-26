package main

import (
	"log"
	amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	//RABBITMQ CONSUMERS
	maintenanceMsgs, err := ch.Consume(
		"Maintenance", // queue
		"",                // consumer
		true,              // auto ack
		false,             // exclusive
		false,             // no local
		false,             // no wait
		nil,               // args
	)
	failOnError(err, "Failed to register a consumer")
	
	// execMsgs, err := ch.Consume(
	// 	"Execute_Maintenance", // queue
	// 	"",                // consumer
	// 	true,              // auto ack
	// 	false,             // exclusive
	// 	false,             // no local
	// 	false,             // no wait
	// 	nil,               // args
	// )
	// failOnError(err, "Failed to register a consumer")

	// orderMsgs, err := ch.Consume(
	// 	"Order_Parts", // queue
	// 	"",                // consumer
	// 	true,              // auto ack
	// 	false,             // exclusive
	// 	false,             // no local
	// 	false,             // no wait
	// 	nil,               // args
	// )
	// failOnError(err, "Failed to register a consumer")

	// returnMsgs, err := ch.Consume(
	// 	"Return_Parts", // queue
	// 	"",                // consumer
	// 	true,              // auto ack
	// 	false,             // exclusive
	// 	false,             // no local
	// 	false,             // no wait
	// 	nil,               // args
	// )
	// failOnError(err, "Failed to register a consumer")

	
	var forever chan struct{}

	go func() {
		for d := range maintenanceMsgs {
			log.Printf(" [x] %s", d.Body)
			SendMessage()
		}
	}()

	log.Printf(" [*] Waiting for logs. To exit press CTRL+C")
	<-forever
}
