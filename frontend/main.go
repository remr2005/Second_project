package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	// Страница продуктов 
	r.HandleFunc("/ProductStore", mainProducts).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", r))
}
