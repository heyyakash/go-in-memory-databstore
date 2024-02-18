package helper

import (
	"net/http"

	"github.com/heyyakash/go-in-memory-datastore/datastore"
	"github.com/heyyakash/go-in-memory-datastore/models"
)

func CheckCondition(w http.ResponseWriter, command string, key string, toSet *bool) {
	switch command {
	case "NX":
		if exists := datastore.KeyValueStore.KeyExists(key); exists != false {
			*toSet = false
		}
	case "XX":
		if exists := datastore.KeyValueStore.KeyExists(key); exists != true {
			*toSet = false
		}
	default:
		ResponseGenerator(w, models.Response{Error: "Invalid 4th argument"}, http.StatusBadRequest)
		return
	}
}
