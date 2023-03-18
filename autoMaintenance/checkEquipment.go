package main

import (
	pb "autoMaintenance/proto"
	"encoding/json"
	"fmt"
	"log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net/http"
)

type Response struct {
	Code int `json:"code"`
	Data struct {
		Equipment []struct {
			ID                string `json:"_id"`
			EquipmentID       string `json:"equipment_id"`
			EquipmentLocation string `json:"equipment_location"`
			EquipmentName     string `json:"equipment_name"`
			LastMaintained    string `json:"last_maintained"`
		} `json:"equipment"`
	} `json:"data"`
}

func getJson(url string, target interface{}) error {
	r, err := http.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

var addr string = "0.0.0.0:50051"

func checkEquipment() {
	url := "http://localhost:4999/equipment"
	var response Response
	err := getJson(url, &response)
	if err != nil {
		fmt.Println(err)
	}
	equipmentList := response.Data.Equipment
	for _, e := range equipmentList {
		fmt.Println(e.EquipmentName)
	}

	
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect %v\n", err)
	}
	defer conn.Close()

	c := pb.NewSchedulerClient(conn)
	doCheck(c)

	// TO-DO
}
