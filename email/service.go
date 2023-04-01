package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)


func sendMaintenanceEmail(email MaintenanceMessage) {
	from := mail.NewEmail("Jeremy Chow", os.Getenv("SENDER_EMAIL"))
	subject := "Maintenance Scheduled"
	to := mail.NewEmail("Example User", os.Getenv("RECEIVER_EMAIL"))

	plainTextContent := "A Maintenance has been scheduled for " + email.ScheduleDate + ".\n" +
		"Technician: " + email.Technician.Name + "(" + email.Technician.TechnicianID + ")\n" +
		"Equipment Name: " + email.Equipment.EquipmentName + "\n\n" +
		"Location: " + email.Equipment.EquipmentLocation

	htmlContent := "A Maintenance has been scheduled for " + email.ScheduleDate + ".\n" +
		"Technician: " + email.Technician.Name + "(" + email.Technician.TechnicianID + ")\n\n" +
		"Equipment Name: " + email.Equipment.EquipmentName + "\n" +
		"Location: " + email.Equipment.EquipmentLocation

	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	response, err := client.Send(message)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}

func sendProcurementEmail(email ProcurementMessage) {
	from := mail.NewEmail("Jeremy Chow", os.Getenv("SENDER_EMAIL"))
	subject := "Maintenance Scheduled"
	to := mail.NewEmail("Procurement Manager", "sclim.2021@scis.smu.edu.sg")
	plainTextContent := "Dear Procurement Manager, we would like to request for the following parts: \n"
	htmlContent := "Dear Procurement Manager, we would like to request for the following parts: " + "<br>"
	for i := range email {
		fmt.Println(email[i].PartName)
		plainTextContent += email[i].PartName + ": " + strconv.Itoa(email[i].Qty) + "<br>"
		htmlContent += email[i].PartName + ": " + strconv.Itoa(email[i].Qty) + "<br>"
	}

	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	response, err := client.Send(message)

	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}

func sendStartEmail(msg StartMessage) {
	from := mail.NewEmail("Jeremy Chow", os.Getenv("SENDER_EMAIL"))
	subject := "Maintenance Started"
	to := mail.NewEmail("Procurement Manager", "xelloxzxd@gmail.com")
	plainTextContent := "Dear Admin,\n\n Maintenance for the following equipment has been started.\n " + 
	"Equipment ID: " + msg.Equipment.EquipmentID + "\n" +
	"Equipment Name: "  + msg.Equipment.EquipmentName + "\n" +
	"Start Time: "  + msg.StartDatetime + "\n"

	htmlContent := "Dear Admin,<br><br>Maintenance for the equipment has started<br> " + 
	"Equipment ID: " + msg.Equipment.EquipmentID + "<br>" +
	"Equipment Name: "  + msg.Equipment.EquipmentName + "<br>" +
	"Start Time: "  + msg.StartDatetime + "<br>"

	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	response, err := client.Send(message)

	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}

func sendEndEmail(msg EndMessage) {
	from := mail.NewEmail("Jeremy Chow", os.Getenv("SENDER_EMAIL"))
	subject := "Maintenance Started"
	to := mail.NewEmail("Procurement Manager", "xelloxzxd@gmail.com")
	
	plainTextContent := ""
	htmlContent := ""

	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	response, err := client.Send(message)

	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}


