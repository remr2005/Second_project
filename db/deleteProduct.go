package main

import (
	"database/sql"
	"fmt"
	"net/http"
)

// must test
func deleteProduct(w http.ResponseWriter, r *http.Request) {
	// Github айди
	git_id := r.URL.Query().Get("ID")

	// Открытие БД
	db, err := sql.Open("mysql", "root:godzila2005;@/Persons_2")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	// Выполнение SQL-запроса для удаления пользователя по GITID
	_, err = db.Exec("delete from Persons_2.Products_table where ID = ?", git_id)
	if err != nil {
		fmt.Println(err)
		fmt.Fprint(w, `0`)
		return
	}
	fmt.Fprint(w, `1`)
}
