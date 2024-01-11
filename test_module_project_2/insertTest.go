package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
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
	method := "GET"

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
	method = "GET"

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
	var str string = `<!DOCTYPE html>
	<html lang="ru">
	<head>
		<meta charset="UTF-8">
		<title>Page</title>
		<style>
			html{
		font-family: Arial, Helvetica, sans-serif;
		background-color: rgba(128, 112, 178, 0.608);
	}
	
	img{
		width: 10%
	}
	.wrapper_header{
		text-align: center;
		display: flex;
		justify-content: center;
		align-items: center;
	}
	.header{
		border-radius: 10px;
		font-size: 2vw;
		width: 20%;
		box-shadow: 0px 0px 10px azure;
	}
	.photo{
		position: relative;
		text-align: center;
		top: 20px;
	}
	.info{
		font-size: 1.3vw;
		position: absolute;
		left: 10%;
		top: 55%;
		width: 80%;
		border-radius: 10px;
		box-shadow: 0px 0px 30px rgb(187, 91, 169);
		background-color: rgba(161, 137, 223, 0.681);
	}
	
	
		</style>
	</head>
	<body>
		<div id="clock" style="color: black; position: absolute; right: 2%; font-size: 1vw;"></div>
		<div class="wrapper_header">
			<div class="header">
				<p><strong>` + name + `</strong></p>
			</div>
		</div>
		<div class="photo">
			<img src="` + link + `" style="width: 15%;">
		</div>
		
		<div class="info">Цена:` + cost + `<br>` + desc + `</div>
		<script >
			function displayTime() {
		let element = document.getElementById("clock"); // Найти элемент с id="clock" 
		let now = new Date();    // Получить текущее время
		element.innerHTML = now.toLocaleTimeString(); // Отобразить время
		setTimeout(displayTime, 1000);    // Вызывать функцию каждую секунду
	}
	window.onload = displayTime;
		</script>
	</body>
	</html>`
	if frontend_test(str, id) {
		fmt.Println("FrontendTest ", id, " пройден")
	} else {
		fmt.Println("FrontendTest ", id, " не пройден")
	}
	url = "http://localhost:8001/db/del/products?ID=" + id
	method = "GET"

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

// тест фронтэнда, странички продукта
func frontend_test(str1 string, id string) bool {
	url := "http://localhost:8000/ProductStore/product?id=" + id
	method := "GET"

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
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return false
	}
	if string(body) == str1 {
		return true
	} else {
		return false
	}
}
