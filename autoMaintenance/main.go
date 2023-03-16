package main


import (
	"fmt"
	"log"
	"net/http"

	"gopkg.in/robfig/cron.v2"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello"))
}

func hello(name string) {
	message := fmt.Sprintf("Hi, %v", name)
	fmt.Println(message)
}

func runCronJobs() {

	s := cron.New()

	s.AddFunc("@every 1s", checkMaintenanace)

	s.Start()
    mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}


func main() {
    runCronJobs()
	fmt.Scanln()
}
