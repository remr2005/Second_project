package db

import (
	"app/domain"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// без теста
func GetProducts() ([]domain.Product, error) {
	db, err := sql.Open("mysql", "root:godzila2005;@/Persons_2") // открытие БД
	products := []domain.Product{}                               // массив пользователей
	if err != nil {
		fmt.Println(err)
		return products, err
	}
	defer db.Close() // освобождение памяти

	rows, err := db.Query("select * from Persons_2.Products_table") // получение всех товаров

	for rows.Next() {
		p := domain.Product{}
		err := rows.Scan(&p.Id, &p.Name, &p.Description, &p.Img_link, &p.Cost) // запись пользователя
		if err != nil {
			fmt.Println(err)
			return products, err
		}
		products = append(products, p) // добавление товара в массив
	}
	return products, nil
}
