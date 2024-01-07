package main

import (
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Возвращает все товары
	r.HandleFunc("/db/get/products", getProducts).Methods("GET")
	// Возвращает все товары
	r.HandleFunc("/db/get/product", getProduct).Methods("GET")
	// Добавляет товар
	r.HandleFunc("/db/ins/products", insertPerson).Methods("GET")
	r.HandleFunc("/db/del/products", deleteProduct).Methods("GET")

	log.Fatal(http.ListenAndServe(":8001", r))
}
