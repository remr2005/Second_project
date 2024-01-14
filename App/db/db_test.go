package db

import (
	"strconv"
	"testing"
)

func TestDelete(t *testing.T) {
	_ = InsertProduct("testid", "testname", "testdesc", "testlink", "1")
	res := DeleteProduct("testid")
	if !res {
		t.Error("Объект не удален")
		t.Fail()
	}
}

func TestGetProduct(t *testing.T) {
	_ = InsertProduct("testid", "testname", "testdesc", "testlink", "1")
	prod, err := GetProduct("testid")
	if err != nil {
		_ = DeleteProduct("testid")
		t.Error("Объект не получен")
		t.Fail()
	}
	str_cost := strconv.Itoa(prod.Cost)
	if prod.Id != "testid" || prod.Name != "testname" || prod.Description != "testdesc" || prod.Img_link != "testlink" || str_cost != "1" {
		_ = DeleteProduct("testid")
		t.Error("Объект получен не корректно", prod.Id, prod.Name, prod.Description, prod.Img_link, str_cost)
		t.Fail()
	}
	_ = DeleteProduct("testid")
}

func TestGetProducts(t *testing.T) {
	_, err := GetProducts()
	if err != nil {
		t.Error("Объекты не полученны")
		t.Fail()
	}
}

func TestInsert(t *testing.T) {
	res := InsertProduct("testid", "testname", "testdesc", "testlink", "1")
	if !res {
		_ = DeleteProduct("testid")
		t.Error("Объект не вставлен")
		t.Fail()
	}
	_ = DeleteProduct("testid")
}
