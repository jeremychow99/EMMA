package main

import (
)

type MaintenanceResp struct {
	Code int  `json:"code"`
	Data MaintenanceData `json:"data"`
}

type Partlist struct {
	PartName    string `json:"PartName"`
	ReservedQty int    `json:"ReservedQty"`
	ID          string `json:"_id"`
}

type Maintenance struct {
	ID               string     `json:"_id"`
	EquipmentID      string     `json:"equipment_id"`
	Partlist         []Partlist `json:"partlist"`
	ScheduleDatetime string   `json:"schedule_datetime"`
	Description      string     `json:"description,omitempty"`
	EndDatetime      string     `json:"end_datetime,omitempty"`
	StartDatetime    string     `json:"start_datetime,omitempty"`
	Status           string     `json:"status,omitempty"`
}
type MaintenanceData struct {
	Maintenance []Maintenance `json:"maintenance"`
}


type EqpResp struct {
	Code int `json:"code"`
	Data struct {
		Equipment []struct {
			ID                string `json:"_id"`
			EquipmentLocation string `json:"equipment_location"`
			EquipmentName     string `json:"equipment_name"`
			LastMaintained    string `json:"last_maintained"`
			EquipmentStatus   string `json:"equipment_status"`
		} `json:"equipment"`
	} `json:"data"`
}

type UserResp struct {
	Users []User `json:"users"`
}
type User struct {
	ID       string `json:"_id"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Role     string `json:"role"`
	Phone    int    `json:"phone"`
	Email    string `json:"email"`
	V        int    `json:"__v"`
}

type BusyTechsResp struct {
	Code int      `json:"code"`
	Data []string `json:"data"`
}