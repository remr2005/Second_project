package db

import (
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Возвращает список всех пользователей
	r.HandleFunc("/get/persons", getMENSCHEN).Methods("GET")
	// Возвращает все товары
	r.HandleFunc("/get/products", getProducts).Methods("GET")

	log.Fatal(http.ListenAndServe(":8001", r))
}
