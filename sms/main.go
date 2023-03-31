package main

import (
	"encoding/json"
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

// type MaintenanceMsg struct {
// 	EquipmentID      string `json:"equipment_id"`
// 	ScheduleDatetime string `json:"schedule_datetime"`
// 	Partlist         []struct {
// 		PartName string `json:"PartName"`
// 		Qty      int    `json:"Qty"`
// 		ID       string `json:"_id"`
// 	} `json:"partlist"`
// }

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
		"",            // consumer
		true,          // auto ack
		false,         // exclusive
		false,         // no local
		false,         // no wait
		nil,           // args
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
			var maintenanceSMS MaintenanceSMS

			err := json.Unmarshal(d.Body, &maintenanceSMS)
			if err != nil {
				fmt.Println(err)
			}

			fmt.Println(maintenanceSMS.Equipment.EquipmentID)
			fmt.Println(maintenanceSMS.Technician.Phone)
			// SendMessage(&maintenanceSMS)
		}
	}()

	log.Printf(" [*] Waiting for logs. To exit press CTRL+C")
	<-forever
}
