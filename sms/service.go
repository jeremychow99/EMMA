package main

import (
	"encoding/json"
	"fmt"
	"os"
	// "strconv"

	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
)

func SendMessage(sms *MaintenanceSMS) {
	// toNum := "+65" + strconv.Itoa(sms.Technician.Phone)
	from := os.Getenv("TWILIO_FROM_PHONE_NUMBER")
	to := os.Getenv("TWILIO_TO_PHONE_NUMBER")

	fmt.Println(sms.Technician.Name)
	client := twilio.NewRestClient()
	params := &twilioApi.CreateMessageParams{}
	params.SetTo(to)
	params.SetFrom(from)
	params.SetBody(
		"Hello " + sms.Technician.Name + ", you have a new maintenance scheduled.\n" +
		"Date: " + sms.ScheduleDate + "\n" +
		"Equipment ID: " + sms.Equipment.EquipmentID  + "\n" +
		"Equipment Name: " + sms.Equipment.EquipmentName + "\n" +
		"Location: " + sms.Equipment.EquipmentLocation)

	resp, err := client.Api.CreateMessage(params)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		response, _ := json.Marshal(*resp)
		fmt.Println("Response: " + string(response))
	}
}
