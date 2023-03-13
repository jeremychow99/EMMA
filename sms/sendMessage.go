package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
)

func SendMessage() {
	from := os.Getenv("TWILIO_FROM_PHONE_NUMBER")
	to := os.Getenv("TWILIO_TO_PHONE_NUMBER")
	client := twilio.NewRestClient()

	params := &twilioApi.CreateMessageParams{}
	params.SetTo(to)
	params.SetFrom(from)
	params.SetBody("Hello there")

	resp, err := client.Api.CreateMessage(params)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		response, _ := json.Marshal(*resp)
		fmt.Println("Response: " + string(response))
	}
}
