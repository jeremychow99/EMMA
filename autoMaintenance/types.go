package main

type MaintenanceResp struct {
	Code int             `json:"code"`
	Data MaintenanceData `json:"data"`
}

type MaintenanceData struct {
	Maintenance []Maintenance `json:"maintenance"`
}

type Equipment1 struct {
	EquipmentID       string `json:"equipment_id"`
	EquipmentLocation string `json:"equipment_location"`
	EquipmentName     string `json:"equipment_name"`
}
type Partlist struct {
	PartName string `json:"PartName"`
	Qty      int    `json:"Qty"`
	ID       string `json:"_id"`
}
type Technician struct {
	Name         string `json:"name"`
	Phone        int    `json:"phone"`
	TechnicianID string `json:"technician_id"`
}

type Maintenance struct {
	ID           string     `json:"_id"`
	Equipment    Equipment1 `json:"equipment"`
	Partlist     []Partlist `json:"partlist"`
	ScheduleDate string     `json:"schedule_date"`
	Status       string     `json:"status"`
	Technician   Technician `json:"technician"`
}

type SubmitTechnician struct {
	ID    string `json:"technician_id"`
	Name  string `json:"name"`
	Phone int    `json:"phone"`
}

type EqpResp struct {
	Code int `json:"code"`
	Data struct {
		EquipmentData []Equipment `json:"equipment"`
	} `json:"data"`
}

type Equipment struct {
	ID                string `json:"_id"`
	EquipmentLocation string `json:"equipment_location"`
	EquipmentName     string `json:"equipment_name"`
	LastMaintained    string `json:"last_maintained"`
	EquipmentStatus   string `json:"equipment_status"`
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

// type eqpResp struct {
// 	Code int `json:"code"`
// 	Data struct {
// 		ID                string `json:"_id"`
// 		EquipmentLocation string `json:"equipment_location"`
// 		EquipmentName     string `json:"equipment_name"`
// 		EquipmentStatus   string `json:"equipment_status"`
// 		LastMaintained    string `json:"last_maintained"`
// 	} `json:"data"`
// }

type eqpResp struct {
	Code int  `json:"code"`
	Data Eqp `json:"data"`
}
type Eqp struct {
	ID                string `json:"_id"`
	EquipmentLocation string `json:"equipment_location"`
	EquipmentName     string `json:"equipment_name"`
	LastMaintained    string `json:"last_maintained"`
	EquipmentStatus   string `json:"equipment_status"`
}