package main

import (
	"log"
	"net/http"
)

func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	if err := responseWithJSON(w, http.StatusOK, struct{}{}); err != nil {
		log.Printf("Failed to send response: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
