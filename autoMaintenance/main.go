package main

import (
	pb "autoMaintenance/proto"
	"context"

	"log"
	"net/http"

	"gopkg.in/robfig/cron.v2"
)

func doCheck(c pb.SchedulerClient) {
	log.Println("doCheck invoked")
	res, err := c.ScheduleMaintenance(context.Background(), &pb.Equipment{
		Name: "test",
	})
	if err != nil {
		log.Fatalf("could not check: %v\n", err)
	}

	log.Printf(res.Status)
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Auto Maintenance Service is running on port 4001."))
}

func runCronJobs() {
	s := cron.New()

	s.AddFunc("@every 1s", checkEquipment)

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
