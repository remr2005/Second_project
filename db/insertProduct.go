package main

import (
	"database/sql"
	"net/http"
)

func insertPerson(w http.ResponseWriter, r *http.Request) {

	db, err := sql.Open("mysql", "root:godzila2005;@/Persons_2")
	if err != nil {
		panic(err)
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
		panic(err)
	}
}
