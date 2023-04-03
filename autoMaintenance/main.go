package main

import (
	"log"
	"os"
	"net/http"

)


func init() {
    os.Setenv("TZ", "Asia/Singapore")
}

func runServer() {
    mux := http.NewServeMux()
	mux.HandleFunc("/mockIOT", autoSchedule)

	log.Println("Starting server on :4001")
	err := http.ListenAndServe(":4001", mux)
	log.Fatal(err)
}

func main() {
	runServer()
}
