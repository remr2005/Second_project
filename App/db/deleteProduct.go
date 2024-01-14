package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// must test
func DeleteProduct(id string) bool {
	// Github айди
	git_id := id

	// Открытие БД
	db, err := sql.Open("mysql", "root:godzila2005;@/Persons_2")
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer db.Close()

	// Выполнение SQL-запроса для удаления пользователя по GITID
	_, err = db.Exec("delete from Persons_2.Products_table where ID = ?", git_id)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
