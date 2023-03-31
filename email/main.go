package main

import (
	"encoding/json"
	"log"
	"fmt"
	// "github.com/joho/godotenv"
	amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

// func init() {
// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Fatal("Error loading .env file")
// 	}
// }

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()


	maintenanceEmails, err := ch.Consume(
		"Maintenance_Email", // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")
	
	procurementEmails, err := ch.Consume(
		"Order_Parts", // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	var forever chan struct{}

	go func() {
		for d := range maintenanceEmails {
			var maintenanceEmail MaintenanceEmail
			log.Printf("Received a message: %s", d.Body)
			err := json.Unmarshal(d.Body, &maintenanceEmail)
			if err != nil {
				fmt.Println(err)
			}
			SendMaintenanceEmail(&maintenanceEmail)
		}
	}()
	go func() {
		for d := range procurementEmails {
			var procurementEmail ProcurementEmail
			log.Printf("Received a message: %s", d.Body)
			err := json.Unmarshal(d.Body, &procurementEmail)
			if err != nil {
				fmt.Println(err)
			}
			for i := range procurementEmail {
				fmt.Println(procurementEmail[i].PartName)
			}
			SendProcurementEmail(procurementEmail)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
