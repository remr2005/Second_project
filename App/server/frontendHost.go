package server

import (
	"app/frontend"
	"fmt"
	"net/http"
)

func Product_page(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	str, err := frontend.ProductHTML_builder(id)
	if err != nil {
		fmt.Fprint(w, err)
		return
	}
	fmt.Fprint(w, str)
}

func Products_page(w http.ResponseWriter, r *http.Request) {
	str, err := frontend.ProductsHTML_builder()
	if err != nil {
		fmt.Fprint(w, err)
		return
	}
	fmt.Fprint(w, str)
}
