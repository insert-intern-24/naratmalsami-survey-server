package service

import (
	"encoding/json"
	"net/http"
)

func GetHelloWorld(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{
		"message": "hello world",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
