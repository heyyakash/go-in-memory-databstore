package helper

import (
	"errors"
	"net/http"

	"github.com/heyyakash/go-in-memory-datastore/datastore"
)

func CheckCondition(w http.ResponseWriter, command string, key string, toSet *bool) error {
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
		return errors.New("invalid 4th command")
	}
	return nil
}
