package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/maldahleh/stockspotter-data-service/models"
)

func handleRequest(rw http.ResponseWriter, req *http.Request) {
	var request models.DataRequest
	err := json.NewDecoder(req.Body).Decode(&request)
	if err != nil || request.Version == "" {
		request = models.DefaultRequest
	}

	resp, err := json.Marshal("{}")
	if err != nil {
		log.Println("Failed to marshal JSON", err)
		http.Error(rw, "request failed, try again later", http.StatusInternalServerError)
		return
	}

	_, err = rw.Write(resp)
	if err != nil {
		log.Println("HTTP write failure", err)
	}
}

func main() {
	http.HandleFunc("/", handleRequest)

	err := http.ListenAndServe(":8082", nil)
	log.Fatal("encountered error with web server", err)
}
