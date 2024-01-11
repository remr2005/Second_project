package main

import (
	"database/sql"
	"fmt"
	"net/http"
)

// must test
func insertPerson(w http.ResponseWriter, r *http.Request) {

	db, err := sql.Open("mysql", "root:godzila2005;@/Persons_2")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	id := r.URL.Query().Get("id")
	name := r.URL.Query().Get("name")
	description := r.URL.Query().Get("description")
	link := r.URL.Query().Get("link")
	cost := r.URL.Query().Get("cost")

	// Выполнение SQL-запроса для добавления нового пользователя
	_, err = db.Exec("INSERT INTO Persons_2.Products_table (ID, NAMEP, DESCRIPTIONP, IMG_LINK, COST) VALUES (?, ?, ?, ?, ?)",
		id, name, description, link, cost)
	if err != nil {
		fmt.Println(err)
		fmt.Fprint(w, `0`)
		return
	}
	fmt.Fprint(w, `1`)
}
