package main

import (
	"app/server"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Возвращает все товары
	r.HandleFunc("/db/get/products", server.GetProducts_host).Methods("POST")
	// Возвращает все товары
	r.HandleFunc("/db/get/product", server.GetProduct_host).Methods("POST")
	// Добавляет товар
	r.HandleFunc("/db/ins/products", server.AddPr_host).Methods("POST")
	// Удалить товар
	r.HandleFunc("/db/del/products", server.DeletePr_host).Methods("POST")
	// Страница товаров
	r.HandleFunc("/ProductStore", server.Products_page).Methods("GET")
	// Страница товара
	r.HandleFunc("/ProductStore/product", server.Product_page).Methods("GET")

	log.Fatal(http.ListenAndServe(":8001", r))
}
