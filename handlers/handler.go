package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/heyyakash/go-in-memory-datastore/datastore"
	"github.com/heyyakash/go-in-memory-datastore/helper"
	"github.com/heyyakash/go-in-memory-datastore/models"
)

func HandlePostReq(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var req models.Request
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			helper.ResponseGenerator(w, models.Response{Message: "Invalid Command"}, http.StatusBadRequest)
			return
		}

		//spliting string into slice of args
		args := strings.Split(req.Command, " ")
		// Looking for command
		switch args[0] {

		case "SET":
			if len(args) < 3 {
				helper.ResponseGenerator(w, models.Response{Message: "Invalid SET Command"}, http.StatusBadRequest)
				return
			}

			// Set Values
			key := args[1]
			value := args[2]
			// Set to False by Default
			toSet := false
			// Set time to 0 by default
			time := 0
			// Check for  conditions
			if len(args) == 4 {
				helper.CheckCondition(w, args[3], key, &toSet)
			}

			// Check for expriry
			if len(args) >= 5 && args[3] == "EX" {
				helper.CheckExpiry(w, args[4], &time)
			} else if len(args) >= 5 && args[3] != "EX" {
				helper.ResponseGenerator(w, models.Response{Message: "Invalid Commands"}, http.StatusBadRequest)
				return
			}
			// Check for condition after expirt
			if len(args) == 6 {
				helper.CheckCondition(w, args[5], key, &toSet)
			}
			if len(args) == 3 {
				toSet = true
			}

			helper.SetFunction(w, key, value, toSet, time)

		case "GET":
			if len(args) < 2 || len(args) > 2 {
				helper.ResponseGenerator(w, models.Response{Message: "Invalid Command"}, http.StatusBadRequest)
				return
			}
			value, exists := datastore.KeyValueStore.GetValue(args[1])
			print(datastore.KeyValueStore.Store[args[1]])
			if exists == false {
				helper.ResponseGenerator(w, models.Response{Message: "Key does not exists"}, http.StatusNotFound)
				return
			}
			helper.ResponseGenerator(w, models.Response{Message: value}, http.StatusOK)

		}

	}
}
