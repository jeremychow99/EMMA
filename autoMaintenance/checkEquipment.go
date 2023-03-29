package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type alias struct{
	ID                string `json:"equipment_id"`
	EquipmentLocation string `json:"equipment_location"`
	EquipmentName     string `json:"equipment_name"`
	LastMaintained    string `json:"last_maintained"`
	EquipmentStatus   string `json:"equipment_status"`
}
func (e *Equipment) Convert() (alias) {
    var a alias = alias(*e)
	return a
}


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
	url := "http://host.docker.internal:5000/maintenance/busy_technicians/" + dateStr
	var resp BusyTechsResp
	err := getJson(url, &resp)
	if err != nil {
		fmt.Println(err)
	}
	return resp.Data
}

func getTechnicians() []User {
	technicians := []User{}
	url := "http://host.docker.internal:3001/all"
	var resp UserResp
	err := getJson(url, &resp)
	if err != nil {
		fmt.Println(err)
	}
	// for each technician, get name and append to slice
	for _, u := range resp.Users {
		if u.Role == "TECHNICIAN" {
			technicians = append(technicians, u)
		}
	}
	return technicians
}

func getScheduledEqp() []string {
	scheduledEqp := []string{}
	url := "http://host.docker.internal:5000/maintenance"
	var resp MaintenanceResp
	err := getJson(url, &resp)
	if err != nil {
		fmt.Println(err)
	}
	maintenances := resp.Data.Maintenance
	for i := range maintenances {
		scheduledEqp = append(scheduledEqp, maintenances[i].Equipment.EquipmentID)
	}
	return scheduledEqp
}

func checkEquipment() {
	fmt.Println(time.Now())
	url := "http://host.docker.internal:4999/equipment"
	var resp EqpResp
	err := getJson(url, &resp)
	if err != nil {
		fmt.Println(err)
	}
	equipmentList := resp.Data.EquipmentData
	technicians := getTechnicians()

	for _, e := range equipmentList {
		status := false
		scheduledEqp := getScheduledEqp()

		if e.EquipmentStatus == "Down" && !contains(scheduledEqp, e.ID) {
			fmt.Println(scheduledEqp)
			fmt.Println(e.ID)
			// add a day to current date.
			date := time.Now().AddDate(0, 0, 1)
			// while status == false, meaning havent schedule maintenance
			for !status {
				dateStr := date.Format("2006-01-02")
				availList := technicians
				busyTechs := getBusyTechs(dateStr)
				// remove technician from available list
				for i := range availList {
					if contains(busyTechs, availList[i].ID) {
						availList = append(availList[:i], availList[i+1:]...)
					}
				}

				if len(availList) > 0 {
					// invoke maintenance controller to schedule maintenance
					testarr := [] string{}
					e := e.Convert()
					var st SubmitTechnician
					st.ID = availList[0].ID
					st.Name = availList[0].Name
					st.Phone = availList[0].Phone
					details := map[string]interface{}{"equipment": e, "schedule_date": dateStr, "partlist": testarr, "technician": st}
					jsonData, err := json.Marshal(details)
					fmt.Println(details)
					if err != nil {
						fmt.Println(err)
					}

					resp, err := http.Post("http://host.docker.internal:8080/schedule_maintenance",
						"application/json",
						bytes.NewBuffer(jsonData))
					if err != nil {
						log.Fatal(err)
					}

					var res map[string]interface{}
					json.NewDecoder(resp.Body).Decode(&res)
					fmt.Println(res["json"])

					status = true
				} else {

					date = date.AddDate(0, 0, 1)

				}
			}

		}

	}
}

