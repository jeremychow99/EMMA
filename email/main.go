package main

import (
	"encoding/json"
	"fmt"
	"log"

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

	maintenanceMsgs, err := ch.Consume(
		"Maintenance_Email", // queue
		"",                  // consumer
		true,                // auto-ack
		false,               // exclusive
		false,               // no-local
		false,               // no-wait
		nil,                 // args
	)
	failOnError(err, "Failed to register a consumer")

	procurementMsgs, err := ch.Consume(
		"Order_Parts", // queue
		"",            // consumer
		true,          // auto-ack
		false,         // exclusive
		false,         // no-local
		false,         // no-wait
		nil,           // args
	)
	failOnError(err, "Failed to register a consumer")

	startMaintenanceMsgs, err := ch.Consume(
		"Execute_Maintenance", // queue
		"",                    // consumer
		true,                  // auto-ack
		false,                 // exclusive
		false,                 // no-local
		false,                 // no-wait
		nil,                   // args
	)
	failOnError(err, "Failed to register a consumer")

	var forever chan struct{}

	go func() {
		for d := range maintenanceMsgs {
			var maintenanceMessage MaintenanceMessage
			log.Printf("Received a message: %s", d.Body)
			err := json.Unmarshal(d.Body, &maintenanceMessage)
			if err != nil {
				fmt.Println(err)
			}
			sendMaintenanceEmail(maintenanceMessage)
		}
	}()
	go func() {
		for d := range procurementMsgs {
			var procurementMessage ProcurementMessage
			log.Printf("Received a message: %s", d.Body)
			err := json.Unmarshal(d.Body, &procurementMessage)
			if err != nil {
				fmt.Println(err)
			}
			sendProcurementEmail(procurementMessage)
		}
	}()

	go func() {
		for d := range startMaintenanceMsgs {
			fmt.Println(d.RoutingKey)
			log.Printf("Received a message: %s", d.Body)

			if d.RoutingKey == "maintenance.start" {
				var msg StartMessage
				err := json.Unmarshal(d.Body, &msg)
				if err != nil {
					fmt.Println(err)
				}
				sendStartEmail(msg)
			} else {
				fmt.Println("recevied an end maintenance AMQP message")
				var msg EndMessage
				err := json.Unmarshal(d.Body, &msg)
				if err != nil {
					fmt.Println(err)
				}
				sendEndEmail(msg)
			}
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
