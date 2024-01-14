package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// must test
func InsertProduct(id string, name string, description string, link string, cost string) bool {

	db, err := sql.Open("mysql", "root:godzila2005;@/Persons_2")
	if err != nil {
		return false
	}
	defer db.Close()

	// Выполнение SQL-запроса для добавления нового пользователя
	_, err = db.Exec("INSERT INTO Persons_2.Products_table (ID, NAMEP, DESCRIPTIONP, IMG_LINK, COST) VALUES (?, ?, ?, ?, ?)",
		id, name, description, link, cost)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
