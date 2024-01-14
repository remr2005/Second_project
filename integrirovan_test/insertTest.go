package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

func stringifi(a io.Reader) bool {
	res, err := io.ReadAll(a)
	if err != nil {
		fmt.Println(err)
		return false
	}
	if string(res) == "1" {
		return true
	} else {
		return false
	}
}

func test_insert(id string, name string, desc string, link string, cost string) bool {
	url := "http://localhost:8001/db/ins/products?id=" + id + "&name=" + name + "&description=" + desc + "&link=" + link + "&cost=" + cost
	method := "POST"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return false
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return false
	}
	if stringifi(res.Body) {
		fmt.Println("InsertTest   ", id, " пройден")
	} else {
		fmt.Println("InsertTest   ", id, " не пройден")
	}
	url = "http://localhost:8001/db/get/product?id=" + id
	method = "POST"

	client = &http.Client{}
	req, err = http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return false
	}
	res, err = client.Do(req)
	if err != nil {
		fmt.Println(err)
		return false
	}

	var product_ product
	_ = json.NewDecoder(res.Body).Decode(&product_)
	int_cost, err := strconv.Atoi(cost)
	url = "http://localhost:8001/db/del/products?ID=" + id
	method = "POST"

	client = &http.Client{}
	req, err = http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return false
	}
	res, err = client.Do(req)
	if err != nil {
		fmt.Println(err)
		return false
	}
	if stringifi(res.Body) {
		fmt.Println("DeleteTest   ", id, " пройден")
	} else {
		fmt.Println("DeleteTest   ", id, " не пройден")
	}
	if int_cost == product_.Cost && link == product_.Img_link && desc == product_.Description && name == product_.Name && id == product_.Id {
		return true
	} else {
		return false
	}

}
