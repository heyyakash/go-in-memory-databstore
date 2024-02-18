package helper

import (
	"net/http"
	"time"

	"github.com/heyyakash/go-in-memory-datastore/datastore"
	"github.com/heyyakash/go-in-memory-datastore/models"
)

func SetFunction(w http.ResponseWriter, key string, value string, toSet bool, t int) {
	if t == 0 && toSet {
		datastore.KeyValueStore.SetValue(key, value)
	} else if t > 0 && toSet {
		datastore.KeyValueStore.SetValue(key, value)
		go func() {
			timer := time.After(time.Duration(t) * time.Second)
			select {
			case <-timer:
				datastore.KeyValueStore.DeleteValue(key)
			}
		}()
	}
	print(datastore.KeyValueStore.Store[key])
	ResponseGenerator(w, models.Response{Value: "Command Executed"}, http.StatusOK)
	return
}
