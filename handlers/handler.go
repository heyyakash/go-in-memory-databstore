package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/heyyakash/go-in-memory-datastore/datastore"
	"github.com/heyyakash/go-in-memory-datastore/helper"
	"github.com/heyyakash/go-in-memory-datastore/models"
	"github.com/heyyakash/go-in-memory-datastore/qstore"
)

func HandlePostReq(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var req models.Request
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			helper.ResponseGenerator(w, models.Response{Error: "Invalid Command"}, http.StatusBadRequest)
			return
		}

		//spliting string into slice of args
		args := strings.Split(req.Command, " ")
		// Looking for command
		switch args[0] {

		// SET Commands
		case "SET":
			if len(args) < 3 {
				helper.ResponseGenerator(w, models.Response{Error: "Invalid SET Command"}, http.StatusBadRequest)
				return
			}

			// Set Values
			key := args[1]
			value := args[2]
			// Set to False by Default
			toSet := true
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
				helper.ResponseGenerator(w, models.Response{Error: "Invalid Commands"}, http.StatusBadRequest)
				return
			}
			// Check for condition after expirt
			if len(args) == 6 {
				helper.CheckCondition(w, args[5], key, &toSet)
			}

			helper.SetFunction(w, key, value, toSet, time)

		// GET comamnds
		case "GET":
			if len(args) < 2 || len(args) > 2 {
				helper.ResponseGenerator(w, models.Response{Error: "Invalid Command"}, http.StatusBadRequest)
				return
			}
			value, exists := datastore.KeyValueStore.GetValue(args[1])
			print(datastore.KeyValueStore.Store[args[1]])
			if !exists {
				helper.ResponseGenerator(w, models.Response{Error: "Key does not exists"}, http.StatusNotFound)
				return
			}
			helper.ResponseGenerator(w, models.Response{Value: value}, http.StatusOK)

		// QPUSH
		case "QPUSH":
			if len(args) < 3 {
				helper.ResponseGenerator(w, models.Response{Error: "Invalid command"}, http.StatusBadRequest)
				return
			}
			qname := args[1]
			queue, exists := qstore.QStore.QueueExists(qname)
			if !exists {
				qstore.QStore.CreateQueue(qname)
				qstore.QStore.Store[qname].Enqueue(args[2:])
				helper.ResponseGenerator(w, models.Response{Message: "Command Executed"}, http.StatusOK)
				return
			}
			queue.Enqueue(args[2:])
			helper.ResponseGenerator(w, models.Response{Message: "Command Executed"}, http.StatusOK)

		// QPOP
		case "QPOP":
			if len(args) < 2 || len(args) > 2 {
				helper.ResponseGenerator(w, models.Response{Error: "Invalid command"}, http.StatusBadRequest)
				return
			}
			qname := args[1]
			queue, exists := qstore.QStore.QueueExists(qname)
			if !exists {
				helper.ResponseGenerator(w, models.Response{Error: "Queue does not exist"}, http.StatusNotFound)
				return
			}

			val := queue.Dequeue()
			helper.ResponseGenerator(w, models.Response{Value: val}, http.StatusOK)

		//BQPOP
		case "BQPOP":
			if len(args) < 3 || len(args) > 3 {
				helper.ResponseGenerator(w, models.Response{Error: "Invalid command"}, http.StatusBadRequest)
				return
			}
			qname, qtime := args[1], 0
			queue, exists := qstore.QStore.QueueExists(qname)
			if !exists {
				helper.ResponseGenerator(w, models.Response{Error: "Queue does not exist"}, http.StatusNotFound)
				return
			}
			helper.CheckExpiry(w, args[2], &qtime)
			val, err := queue.BQPop(qtime)
			if err != nil {
				if err.Error() == "context deadline exceeded" {
					helper.ResponseGenerator(w, models.Response{Value: "null", Error: err.Error()}, http.StatusBadRequest)
					return
				}
				helper.ResponseGenerator(w, models.Response{Error: err.Error()}, http.StatusBadRequest)
				return
			}
			helper.ResponseGenerator(w, models.Response{Value: val}, http.StatusOK)

		// default
		default:
			helper.ResponseGenerator(w, models.Response{Error: "Invalid Command"}, http.StatusOK)
		}
	}

}
