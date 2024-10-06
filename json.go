package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func ResponseWithError(w http.ResponseWriter, code int, msg string) {

	if code > 499 {
		log.Printf("ResponseWithError: %s", msg)
	}

	type errorMessage struct {
		Error      string `json:"error"`
		StatusCode int    `json:"statusCode"`
	}

	ResponseWithJson(w, code, errorMessage{Error: msg, StatusCode: code})
}

func ResponseWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)

	if err != nil {
		log.Printf("Error marshal JSON: %v", payload)
		w.WriteHeader(500)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
