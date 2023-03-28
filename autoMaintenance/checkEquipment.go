package main

import (
	// "bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func getJson(url string, target interface{}) error {
	r, err := http.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

func getBusyTechs(dateStr string) []string {
	url := "http://localhost:5000/maintenance/busy_technicians/" + dateStr
	var resp BusyTechsResp
	err := getJson(url, &resp)
	if err != nil {
		fmt.Println(err)
	}
	return resp.Data
}

func getTechnicians() []string {
	technicians := []string{}
	url := "http://localhost:3001/all"
	var resp UserResp
	err := getJson(url, &resp)
	if err != nil {
		fmt.Println(err)
	}
	// for each technician, get name and append to slice
	for _, u := range resp.Users {
		if u.Role == "TECHNICIAN" {
			technicians = append(technicians, u.ID)
		}
	}
	return technicians
}

// func getMaintenance() []Maintenance {
// 	url := "http://localhost:5000/maintenance"
// 	var resp MaintenanceResp
// 	err := getJson(url, &resp)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	return resp.Data.Maintenance
// }

func checkEquipment() {
	url := "http://localhost:4999/equipment"
	var resp EqpResp
	err := getJson(url, &resp)
	if err != nil {
		fmt.Println(err)
	}
	equipmentList := resp.Data.Equipment
	technicians := getTechnicians()
	for _, e := range equipmentList {
		status := false

		if e.EquipmentStatus == "Down" {
			// add a day to current date.
			date := time.Now().AddDate(0, 0, 1)
			// while status == false, meaning havent schedule maintenance
			for !status {
				dateStr := date.Format("2006-01-02")
				availList := technicians
				busyTechs := getBusyTechs(dateStr)
				// remove technician from available list
				for i := range availList {
					// HARDCODE VALUE for now, supposed to be mtn.assignedTechnician
					if contains(busyTechs, availList[i]){
						// remove value
						fmt.Println("hello")
						availList = append(availList[:i], availList[i+1:]...)
					}
				}

				if len(availList) != 0 {
					// invoke maintenance controller ot schedule maintenance
					status = true
					fmt.Println("Scheduled for " + dateStr)
				} else {

					date = date.AddDate(0, 0, 1)
					fmt.Println("next day")
				}
			}

		}

	}
}

// TO-DO

// var jsonStr = []byte(`{"title":"Buy cheese and bread for breakfast."}`)
// req, err := http.NewRequest("POST", "http://localhost:8080/schedule_maintenance", bytes.NewBuffer(jsonStr))
// if err != nil {
// 	panic(err)
// }
// req.Header.Set("Content-Type", "application/json")

// client := &http.Client{}
// resp, err := client.Do(req)
// if err != nil {
// 	panic(err)
// }
// defer resp.Body.Close()
