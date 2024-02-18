package main

import (
	"net/http"

	"github.com/heyyakash/go-in-memory-datastore/datastore"
	"github.com/heyyakash/go-in-memory-datastore/handlers"
	"github.com/heyyakash/go-in-memory-datastore/qstore"
)

func main() {
	datastore.CreateNewStore()
	qstore.CreateNewQueueStore()
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.HandlePostReq)
	http.ListenAndServe(":8080", mux)
}
