package main

import (
	"net/http"

	"github.com/heyyakash/go-in-memory-datastore/datastore"
	"github.com/heyyakash/go-in-memory-datastore/handlers"
)

func main() {
	datastore.CreateNewStore()
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.HandlePostReq)
	http.ListenAndServe(":8080", mux)
}
