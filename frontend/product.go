package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Product(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	url := "http://localhost:8001//db/get/product?id=" + id
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	var Product product
	_ = json.NewDecoder(res.Body).Decode(&Product)
	fmt.Fprint(w, `<!DOCTYPE html>
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
				<p><strong>`)
	fmt.Fprint(w, Product.Name)
	fmt.Fprint(w, `</strong></p>
			</div>
		</div>
		<div class="photo">
			<img src="`)
	fmt.Fprint(w, Product.Img_link)
	fmt.Fprint(w, `" style="width: 15%;">
		</div>
		
		<div class="info">`)
	fmt.Fprint(w, `Цена:`)
	fmt.Fprint(w, Product.Cost)
	fmt.Fprint(w, `<br>`)
	fmt.Fprint(w, Product.Description)
	fmt.Fprint(w, `</div>
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
	</html>`)
}
