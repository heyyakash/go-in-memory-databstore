package helper

import (
	"log"
	"net/http"

	"github.com/heyyakash/go-in-memory-datastore/datastore"
	"github.com/heyyakash/go-in-memory-datastore/models"
)

func CheckCondition(w http.ResponseWriter, command string, key string, toSet *bool) {
	switch command {
	case "NX":
		if exists := datastore.KeyValueStore.KeyExists(key); exists == false {
			*toSet = true
		}
	case "XX":
		if exists := datastore.KeyValueStore.KeyExists(key); exists == true {
			log.Print("Here")
			*toSet = true
		}
	default:
		ResponseGenerator(w, models.Response{Message: "Invalid 4th argument"}, http.StatusBadRequest)
		return
	}
}
