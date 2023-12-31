package main

import (
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Возвращает список всех пользователей
	r.HandleFunc("/db/get/persons", getMENSCHEN).Methods("GET")
	// Возвращает все товары
	r.HandleFunc("/db/get/products", getProducts).Methods("GET")
	// Добавляет товар
	r.HandleFunc("/db/ins/products", insertPerson).Methods("GET")

	log.Fatal(http.ListenAndServe(":8001", r))
}
