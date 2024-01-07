package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func mainProducts(w http.ResponseWriter, r *http.Request) {
	url := "http://localhost:8001/db/get/products"
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
	products := make([]product, 0)
	_ = json.NewDecoder(res.Body).Decode(&products)

	fmt.Fprint(w, `<!DOCTYPE html>
	<html lang="ru">
	<head>
		<meta charset="UTF-8">
		<title>Shop</title>
	</head>
	<body>
		<style>html{
			font-family: Arial, Helvetica, sans-serif;
			background-color: rgba(128, 112, 178, 0.608);
		}
		
		img{
			width: 40%
		}
		button{
			font-size: 1vw;
			background-color: rgba(88, 88, 88, 0.45);
			border-radius: 15px;
			border:0px;
			height: 15%;
			width: 20%;
			list-style: none;
			position: relative;
			transition: 0.6s;
			cursor: pointer;
		}
			button:hover{
				background-color: rgba(88, 88, 88, 0.87);
				top: 0px;
				left: 0px;
				color: azure;
				box-shadow: 4px 4px 4px rgb(0, 0, 0);
				text-shadow: 0px 0px 5px azure;
				transition: 0.6s;
			}
			button:active{
				top: 4px;
				left: 4px;
				transition: 0.4s;
				box-shadow: none;
			}
		.wrapper_header{
			text-align: center;
			display: flex;
			justify-content: center;
			align-items: center;
		
		}
		.header{
			background-color: transparent;
			border-radius: 15px;
			border-color: azure;
			box-shadow: 0px 0px 10px azure;
			width: 30%;
			height: 25%;
			font-size: 2vw;
			cursor: default;
		}
		
		.line1, .line2, .line3{
			text-align: center;
			display: flex;
			justify-content: center;
			align-items: center;
			margin-top: 3%;
		}
		</style>
		<div id="clock" style="color:black; right: 2%; position: absolute; font-size: 1vw;">
		</div>
		<div class="wrapper_header">
			<div class="header" id="12">
				<p>ArianShop</p>
			</div>
		</div>
		<div class="lineBtn">`)
	for i := 0; i < len(products); i++ {
		if i%3 == 0 {
			fmt.Fprint(w, `<form class="line1" target="_blank">`)
			fmt.Fprint(w, `<button class="cpu" style="margin-right: 17%;" formaction="">`)
			fmt.Fprint(w, products[i].Name)
			fmt.Fprint(w, `<p><img src="`)
			fmt.Fprint(w, products[i].Img_link)
			fmt.Fprint(w, `"></p></button>`)
		} else if i%3 == 2 {
			fmt.Fprint(w, `<button class="cpu" style="margin-right: 17%;" formaction="">`)
			fmt.Fprint(w, products[i].Name)
			fmt.Fprint(w, `<p><img src="`)
			fmt.Fprint(w, products[i].Img_link)
			fmt.Fprint(w, `"></p></button>`)
			fmt.Fprint(w, `</form>`)
		} else {
			fmt.Fprint(w, `<button class="cpu" style="margin-right: 17%;" formaction="">`)
			fmt.Fprint(w, products[i].Name)
			fmt.Fprint(w, `<p><img src="`)
			fmt.Fprint(w, products[i].Img_link)
			fmt.Fprint(w, `"></p></button>`)
		}
	}
	if len(products)%3 != 0 {
		fmt.Fprint(w, `</form>`)
	}
	fmt.Fprint(w, `</div>
    <script>function displayTime() {
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
