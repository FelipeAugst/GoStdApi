package views

import (
	"encoding/json"
	"log"
	"net/http"
)

func ToJSON(w http.ResponseWriter, statusCode int, data any) {
	w.Header().Set("Content-Type", "application/json")
	if data == nil {
		return
	}
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Fatal(err)

	}

}
