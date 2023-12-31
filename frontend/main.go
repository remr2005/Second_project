package main

import (
	"gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	// Возвращает список всех пользователей
	r.HandleFunc("/ProductStore", mainProducts).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", r))
}
