package main

import (
	"log"
	"net/http"
	"os"
)

func init() {
	os.Setenv("TZ", "Asia/Singapore")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/mockIOT", autoSchedule)

	log.Println("Starting server on :4001")
	err := http.ListenAndServe(":4001", mux)
	log.Fatal(err)
}
