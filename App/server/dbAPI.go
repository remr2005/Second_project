package server

import (
	"app/db"
	"encoding/json"
	"fmt"
	"net/http"
)

func DeletePr_host(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	res := db.DeleteProduct(id)
	if res {
		fmt.Fprint(w, `1`)
	} else {
		fmt.Fprint(w, `0`)
	}
}

func AddPr_host(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	name := r.URL.Query().Get("name")
	description := r.URL.Query().Get("description")
	link := r.URL.Query().Get("link")
	cost := r.URL.Query().Get("cost")

	res := db.InsertProduct(id, name, description, link, cost)
	if res {
		fmt.Fprint(w, `1`)
	} else {
		fmt.Fprint(w, `0`)
	}
}

func GetProducts_host(w http.ResponseWriter, r *http.Request) {
	products, err := db.GetProducts()
	if err != nil {
		fmt.Fprint(w, err)
	}
	json.NewEncoder(w).Encode(products)
}

func GetProduct_host(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	products, err := db.GetProduct(id)
	if err != nil {
		fmt.Fprint(w, err)
	}
	json.NewEncoder(w).Encode(products)
}
