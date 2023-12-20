package github.com/remr2005/Secod_project/main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	log.Fatal(http.ListenAndServe(":8001", r))
}
