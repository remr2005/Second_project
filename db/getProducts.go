package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

// без теста
func getProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("1")

	db, err := sql.Open("mysql", "root:godzila2005;@/Persons_2") // открытие БД

	if err != nil {
		fmt.Println(err)
	}
	defer db.Close() // освобождение памяти

	rows, err := db.Query("select * from Persons_2.Products_table") // получение всех товаров
	products := []product{}                                         // массив пользователей

	for rows.Next() {
		p := product{}
		err := rows.Scan(&p.Id, &p.Name, &p.Description, &p.Img_link, &p.Cost) // запись пользователя
		if err != nil {
			fmt.Println(err)
			continue
		}
		products = append(products, p) // добавление товара в массив
	}
	json.NewEncoder(w).Encode(products) // возвратить список пользователей на страницу
}
