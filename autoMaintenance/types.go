package main

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