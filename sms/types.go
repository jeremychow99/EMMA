package main

// type MaintenanceSMS struct {
// 	Equipment    Equipment  `json:"equipment"`
// 	Partlist     []any      `json:"partlist"`
// 	ScheduleDate string     `json:"schedule_date"`
// 	Technician   Technician `json:"technician"`
// }
// type Equipment struct {
// 	EquipmentID       string `json:"equipment_id"`
// 	EquipmentLocation string `json:"equipment_location"`
// 	EquipmentName     string `json:"equipment_name"`
// 	LastMaintained    string `json:"last_maintained"`
// 	EquipmentStatus   string `json:"equipment_status"`
// }
// type Technician struct {
// 	TechnicianID string `json:"technician_id"`
// 	Name         string `json:"name"`
// 	Phone        int    `json:"phone"`
// }

type MaintenanceSMS struct {
	Equipment struct {
		EquipmentID       string `json:"equipment_id"`
		EquipmentLocation string `json:"equipment_location"`
		EquipmentName     string `json:"equipment_name"`
		LastMaintained    string `json:"last_maintained"`
		EquipmentStatus   string `json:"equipment_status"`
	} `json:"equipment"`
	Partlist     []any  `json:"partlist"`
	ScheduleDate string `json:"schedule_date"`
	Technician   struct {
		TechnicianID string `json:"technician_id"`
		Name         string `json:"name"`
		Phone        int    `json:"phone"`
	} `json:"technician"`
}