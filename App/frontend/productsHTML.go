package frontend

import "app/db"

func ProductsHTML_builder() (string, error) {
	products, err := db.GetProducts()
	if err != nil {
		return "", err
	}
	str := `<!DOCTYPE html>
	<html lang="ru">
	<head>
		<meta charset="UTF-8">
		<title>ArianShop</title>
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
		<div class="lineBtn">`
	for i := 0; i < len(products); i++ {
		if i%3 == 0 {
			str += `<form class="line1" target="_blank">`
			str += `<button class="cpu" style="margin-right: 17%;">`
			str += `<a href="`
			str += "/ProductStore/product?id=" + products[i].Id
			str += `">`
			str += products[i].Name
			str += `<p><img src="`
			str += products[i].Img_link
			str += `"></p></a></button>`
		} else if i%3 == 2 {
			str += `<button class="cpu" style="margin-right: 17%;">`
			str += `<a href="`
			str += "/ProductStore/product?id=" + products[i].Id
			str += `">`
			str += products[i].Name
			str += `<p><img src="`
			str += products[i].Img_link
			str += `"></p></a></button>`
			str += `</form>`
		} else {
			str += `<button class="cpu" style="margin-right: 17%;">`
			str += `<a href="`
			str += "/ProductStore/product?id=" + products[i].Id
			str += `">`
			str += products[i].Name
			str += `<p><img src="`
			str += products[i].Img_link
			str += `"></p></a></button>`
		}
	}
	if len(products)%3 != 0 {
		str += `</form>`
	}
	str += `</div>
		<script>function displayTime() {
			let element = document.getElementById("clock"); // Найти элемент с id="clock" 
			let now = new Date();    // Получить текущее время
			element.innerHTML = now.toLocaleTimeString(); // Отобразить время
			setTimeout(displayTime, 1000);    // Вызывать функцию каждую секунду
		}
		window.onload = displayTime;
		</script>
	</body>
	</html>`
	return str, nil
}
