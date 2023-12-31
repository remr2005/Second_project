package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

func getMENSCHEN(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("1")

	db, err := sql.Open("mysql", "root:godzila2005;@/Persons_2") // открытие БД

	if err != nil {
		panic(err)
	}
	defer db.Close() // освобождение памяти

	rows, err := db.Query("select * from Persons_2.Persons_table") // получение всех пользователей
	persons := []person{}                                          // массив пользователей

	for rows.Next() {
		p := person{}
		err := rows.Scan(&p.id, &p.name, &p.money) // запись пользователя
		if err != nil {
			fmt.Println(err)
			continue
		}
		persons = append(persons, p) // добавление пользователя в массив
	}
	json.NewEncoder(w).Encode(persons) // возвратить список пользователей на страницу
}
