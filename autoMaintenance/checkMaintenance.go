package main

import (
	"fmt"
	"io"
	"net/http"
)

func checkMaintenanace() {
	// send GET request to Maintenance Service
	resp, err := http.Get("http://localhost:8080/")
	if err != nil {
		fmt.Println(err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	resp.Body.Close()

   // print response body
   fmt.Println(string(body))
}
