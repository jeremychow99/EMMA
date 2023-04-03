package main

import (
	"log"
	"os"
	"net/http"

)


func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Auto Maintenance Service is running on port 4001."))
}
func init() {
    os.Setenv("TZ", "Asia/Singapore")
}

func runServer() {
    mux := http.NewServeMux()
	mux.HandleFunc("/home", home)
	mux.HandleFunc("/mockIOT", autoSchedule)

	log.Println("Starting server on :4001")
	err := http.ListenAndServe(":4001", mux)
	log.Fatal(err)
}

func main() {
	runServer()
}
