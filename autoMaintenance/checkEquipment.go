package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func getJson(url string, target interface{}) error {
	r, err := http.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

// func doCheck(c pb.SchedulerClient, equipmentID string) {
// 	log.Println("doCheck invoked")
// 	res, err := c.ScheduleMaintenance(context.Background(), &pb.Equipment{
// 		Name: equipmentID,
// 	})
// 	if err != nil {
// 		log.Fatalf("could not check: %v\n", err)
// 	}

// 	log.Printf(res.Status)
// }

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

// var addr string = "0.0.0.0:50051"

func checkEquipment() {
	url := "http://localhost:4999/equipment"
	var response Response
	err := getJson(url, &response)
	if err != nil {
		fmt.Println(err)
	}

	equipmentList := response.Data.Equipment
	for _, e := range equipmentList {
		// if equipment status == require maintenance, then invoke gRPC
		// conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
		// if err != nil {
		// 	log.Fatalf("Failed to connect %v\n", err)
		// }
		// defer conn.Close()
		// c := pb.NewSchedulerClient(conn)
		// doCheck(c, e.EquipmentName)

		fmt.Println(e.EquipmentName)
		var jsonStr = []byte(`{"title":"Buy cheese and bread for breakfast."}`)
		req, err := http.NewRequest("POST", "http://localhost:8080/schedule_maintenance", bytes.NewBuffer(jsonStr))
		if err != nil {
			panic(err)
		}
		req.Header.Set("Content-Type", "application/json")
		
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
	}
	// TO-DO
}
