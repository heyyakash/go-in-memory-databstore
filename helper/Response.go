package helper

import (
	"encoding/json"
	"net/http"
)

func ResponseGenerator(w http.ResponseWriter, message interface{}, status int) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(message)
}
