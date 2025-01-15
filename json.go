package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// Send an HTTP response with a JSON payload.
func responseWithJSON(w http.ResponseWriter, code int, payload interface{}) error {
	dat, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	_, writeErr := w.Write(dat)
	return writeErr
}

func responseWithError(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Println("Responding with 5XX error:", msg)
	}
	type errResponse struct {
		Error string `json:"error"`
	}
	responseWithJSON(w, code, errResponse{Error: msg})
}
