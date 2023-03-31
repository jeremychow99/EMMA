package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type alias struct {
	ID                string `json:"equipment_id"`
	EquipmentLocation string `json:"equipment_location"`
	EquipmentName     string `json:"equipment_name"`
	EquipmentStatus   string `json:"equipment_status"`
	LastMaintained    string `json:"last_maintained"`
}

func (e *Eqp1) Convert() alias {
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

func getMaintenancesForEqp(eqpID string) []Maintenance {
	url := "http://host.docker.internal:5000/maintenance/equipment/" + eqpID
	var resp AutoGenerated
	err := getJson(url, &resp)
	if err != nil {
		fmt.Println(err)
	}
	return resp.Data
}

func getEqp(eqpID string) Eqp1 {
	url := "http://host.docker.internal:4999/equipment/" + eqpID
	var resp eqpResp
	err := getJson(url, &resp)
	if err != nil {
		log.Println(err)
	}
	return resp.Eqp1
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

// func getScheduledEqp() (int, []Maintenance, []string) {
// 	scheduledEqp := []string{}
// 	count := 0
// 	url := "http://host.docker.internal:5000/maintenance"
// 	var resp MaintenanceResp
// 	err := getJson(url, &resp)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	maintenances := resp.Data.Maintenance
// 	for i := range maintenances {
// 		if maintenances[i].Status != "COMPLETE - SUCCESSFUL" && maintenances[i].Status != "COMPLETE - UNSUCCESSFUL" {
// 			scheduledEqp = append(scheduledEqp, maintenances[i].Equipment.EquipmentID)
// 		}
// 		count += 1
// 	}
// 	return count, maintenances, scheduledEqp
// }

// func checkEquipment() {
// 	fmt.Println(time.Now())
// 	url := "http://host.docker.internal:4999/equipment"
// 	var resp EqpResp
// 	err := getJson(url, &resp)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	equipmentList := resp.Data.EquipmentData
// 	technicians := getTechnicians()

// for _, e := range equipmentList {
// 	status := false
// 	count, maintenances, scheduledEqp := getScheduledEqp()
// 	// if (e.ID == m.Equipment.EquipmentID && (m.Status == "COMPLETE - SUCCESSFUL" || m.Status == "COMPLETE - UNSUCCESSFUL") && e.EquipmentStatus == "Down" && !contains(scheduledEqp, e.ID) && count > 0) || (e.EquipmentStatus == "Down" && !contains(scheduledEqp, e.ID) && count > 0)
// 	if e.EquipmentStatus == "Down" && !contains(scheduledEqp, e.ID) && count > 0 {
// 		// add a day to current date.
// 		date := time.Now().AddDate(0, 0, 1)
// 		// while status == false, meaning havent schedule maintenance
// 		for !status {
// 			dateStr := date.Format("2006-01-02")
// 			availList := technicians
// 			busyTechs := getBusyTechs(dateStr)
// 			// remove technician from available list
// 			for i := range availList {
// 				if contains(busyTechs, availList[i].ID) {
// 					availList = append(availList[:i], availList[i+1:]...)
// 				}
// 			}

// 			if len(availList) > 0 {
// 				// invoke maintenance controller to schedule maintenance
// 				testarr := []string{}
// 				// e := e.Convert()
// 				var st SubmitTechnician
// 				st.ID, st.Name, st.Phone = availList[0].ID, availList[0].Name, availList[0].Phone
// 				details := map[string]interface{}{"equipment": e, "schedule_date": dateStr, "partlist": testarr, "technician": st}
// 				jsonData, err := json.Marshal(details)
// 				fmt.Println(details)
// 				if err != nil {
// 					fmt.Println(err)
// 					}

// 					// check if schedule_date is earlier than original maintenance time
// 					resp, err := http.Post("http://host.docker.internal:8080/schedule_maintenance",
// 						"application/json",
// 						bytes.NewBuffer(jsonData))
// 					if err != nil {
// 						log.Fatal(err)
// 					}

// 					var res map[string]interface{}
// 					json.NewDecoder(resp.Body).Decode(&res)
// 					fmt.Println(res["json"])

// 					status = true
// 				} else {

// 					date = date.AddDate(0, 0, 1)
// 					fmt.Println(maintenances)

// 				}
// 			}

// 		}

// 	}
// }

func testFunc(rw http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		rw.WriteHeader(405)
		rw.Write([]byte("Method Not Allowed"))
	}

	decoder := json.NewDecoder(req.Body)
	var data PostReqData
	err := decoder.Decode(&data)
	if err != nil {
		panic(err)
	}
	eqpID := data.EquipmentID
	// update equipment status
	// make post request
	e := getEqp(eqpID)
	e.EquipmentStatus = "Down"
	jsonData, err := json.Marshal(e)
	if err != nil {
		log.Fatal(err)
	}
	resp, err := http.Post("http://host.docker.internal:4999/equipment/" + eqpID,
		"application/json",
		bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatal(err)
	}
	var res map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&res)
	fmt.Println(res["json"])

	maintenances := getMaintenancesForEqp(eqpID)
	fmt.Println("===============")
	fmt.Println(maintenances)
	fmt.Println("===============")
	status := false
	date := time.Now().AddDate(0, 0, 1)
	// while status is false, for current date, check if any maintenance date is same as this
	for !status {
		dateStr := date.Format("2006-01-02")
		for _, m := range maintenances {
			if m.ScheduleDate == dateStr {
				fmt.Println("=== DATE MATCH ===")
				status = true
				break
			}
		}
		if status {
			fmt.Println("BREAKING")
			break
		}
		// invoke api to check , if code = 404

		availList := getTechnicians()
		busyTechs := getBusyTechs(dateStr)
		// remove technician from available list
		for i := range availList {
			if contains(busyTechs, availList[i].ID) {
				availList = append(availList[:i], availList[i+1:]...)
			}
		}

		if len(availList) > 0 {
			// invoke maintenance controller to schedule maintenance
			testarr := []string{}
			e := e.Convert()
			var st SubmitTechnician
			st.ID, st.Name, st.Phone = availList[0].ID, availList[0].Name, availList[0].Phone
			details := map[string]interface{}{"equipment": e, "schedule_date": dateStr, "partlist": testarr, "technician": st}
			jsonData, err := json.Marshal(details)
			fmt.Println(details)
			if err != nil {
				fmt.Println(err)
			}

			// check if schedule_date is earlier than original maintenance time
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
