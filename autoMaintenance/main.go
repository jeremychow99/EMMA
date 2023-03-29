package main

import (
	"log"
	"os"
	"net/http"

	"gopkg.in/robfig/cron.v2"
)


func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Auto Maintenance Service is running on port 4001."))
}
func init() {
    os.Setenv("TZ", "Asia/Singapore")
}

func runCronJobs() {
	s := cron.New()

	s.AddFunc("@every 10s", checkEquipment)

	s.Start()
    mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	log.Println("Starting server on :4001")
	err := http.ListenAndServe(":4001", mux)
	log.Fatal(err)
}

func main() {
	runCronJobs()
}
