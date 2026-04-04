package handler

import (
	"encoding/json"
	"net/http"
)

func ShortenHandler(w http.ResponseWriter, r *http.Request) {
	
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Hello World",
	})
}
