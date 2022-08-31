package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

func JsonRes(w http.ResponseWriter, statusCode int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if payload == nil {
		return
	}

	if err := json.NewEncoder(w).Encode(payload); err != nil {
		log.Fatal(err)
	}
}
