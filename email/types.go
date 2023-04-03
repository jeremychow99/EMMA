package main 

type Message map[string]interface{}

type MaintenanceMessage struct {
	Equipment struct {
		EquipmentID       string `json:"equipment_id"`
		EquipmentLocation string `json:"equipment_location"`
		EquipmentName     string `json:"equipment_name"`
		LastMaintained    string `json:"last_maintained"`
		EquipmentStatus   string `json:"equipment_status"`
	} `json:"equipment"`
	Partlist     []Part `json:"partlist"`
	ScheduleDate string `json:"schedule_date"`
	Technician   struct {
		TechnicianID string `json:"technician_id"`
		Name         string `json:"name"`
		Phone        int    `json:"phone"`
	} `json:"technician"`
}

type ProcurementMessage []Part

type Part struct {
	PartName string `json:"PartName"`
	Qty      int    `json:"Qty"`
	ID       string `json:"_id"`
}

type StartMessage struct {
	Equipment     Equipment `json:"equipment"`
	StartDatetime string    `json:"start_datetime"`
	Start         bool      `json:"start"`
}
type EndMessage struct {
	Equipment         Equipment `json:"equipment"`
	ScheduleDate      string    `json:"schedule_date"`
	Partlist          []Part     `json:"partlist"`
	ReturnPartlist    []any     `json:"return_partlist"`
	EndDatetime       string    `json:"end_datetime"`
	Description       string    `json:"description"`
	MaintenanceStatus string    `json:"maintenance_status"`
	Start             bool      `json:"start"`
}
type Equipment struct {
	EquipmentID       string `json:"equipment_id"`
	EquipmentLocation string `json:"equipment_location"`
	EquipmentName     string `json:"equipment_name"`
	EquipmentStatus   string `json:"equipment_status"`
	LastMaintained    string `json:"last_maintained"`
}
