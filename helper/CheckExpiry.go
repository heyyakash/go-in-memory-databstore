package helper

import (
	"net/http"
	"strconv"

	"github.com/heyyakash/go-in-memory-datastore/models"
)

func CheckExpiry(w http.ResponseWriter, command string, time *int) {
	i, err := strconv.Atoi(command)
	if err != nil {
		ResponseGenerator(w, models.Response{Error: "Invalid time provided"}, http.StatusBadRequest)
	}
	*time = i
}
