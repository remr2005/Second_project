package db

import (
	"app/domain"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// no test
func GetProduct(id string) (domain.Product, error) {

	db, err := sql.Open("mysql", "root:godzila2005;@/Persons_2") // открытие БД

	if err != nil {
		fmt.Println(err)
	}
	defer db.Close() // освобождение памяти

	rows, err := db.Query("select * from Persons_2.Products_table") // получение всех товаров

	for rows.Next() {
		p := domain.Product{}
		err := rows.Scan(&p.Id, &p.Name, &p.Description, &p.Img_link, &p.Cost)
		if err != nil {
			return domain.Product{}, err
		}
		if p.Id == id {
			return p, nil
		}
	}
	return domain.Product{}, err
}
