package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

// no test
func getProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("1")

	id := r.URL.Query().Get("id")

	db, err := sql.Open("mysql", "root:godzila2005;@/Persons_2") // открытие БД

	if err != nil {
		fmt.Println(err)
	}
	defer db.Close() // освобождение памяти

	rows, err := db.Query("select * from Persons_2.Products_table") // получение всех товаров

	for rows.Next() {
		p := product{}
		err := rows.Scan(&p.Id, &p.Name, &p.Description, &p.Img_link, &p.Cost)
		if err != nil {
			fmt.Println(err)
			continue
		}
		if p.Id == id {
			json.NewEncoder(w).Encode(p)
		}
	} // возвратить список пользователей на страницу
}
