package frontend

import (
	"app/db"
	"testing"
)

func TestProductsHTML(t *testing.T) {
	_, err := ProductsHTML_builder()
	if err != nil {
		t.Error("Страница товаров не была прогружена")
		t.Fail()
	}
}

func TestProductHTML(t *testing.T) {
	_ = db.InsertProduct("testid", "testname", "testdesc", "testlink", "testcost")
	_, err := ProductHTML_builder("testid")
	if err != nil {
		db.DeleteProduct("testid")
		t.Error("Страница товаров не была прогружена")
		t.Fail()
	}
}
