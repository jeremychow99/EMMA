package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type MaintenanceEmail struct {
	Equipment struct {
		EquipmentID       string `json:"equipment_id"`
		EquipmentLocation string `json:"equipment_location"`
		EquipmentName     string `json:"equipment_name"`
		LastMaintained    string `json:"last_maintained"`
		EquipmentStatus   string `json:"equipment_status"`
	} `json:"equipment"`
	Partlist     []Part  `json:"partlist"`
	ScheduleDate string `json:"schedule_date"`
	Technician   struct {
		TechnicianID string `json:"technician_id"`
		Name         string `json:"name"`
		Phone        int    `json:"phone"`
	} `json:"technician"`
}

type ProcurementEmail []Part 

type Part struct {
	PartName string `json:"PartName"`
	Qty int `json:"Qty"`
	ID string `json:"_id"`
}


func SendMaintenanceEmail(email *MaintenanceEmail) {
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

func SendProcurementEmail (email ProcurementEmail) {

	for i := range email{
		fmt.Println(email[i].PartName)
	}
	from := mail.NewEmail("Jeremy Chow", os.Getenv("SENDER_EMAIL"))
	subject := "Maintenance Scheduled"
	to := mail.NewEmail("Procurement Manager", "sclim.2021@scis.smu.edu.sg")
	plainTextContent := "Dear Procurement Manager, we would like to for the following parts: \n"
	htmlContent := "Dear Procurement Manager, we would like to for the following parts: " + "<br>"
	for i := range email{
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
