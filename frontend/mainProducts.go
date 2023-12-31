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
		<link rel="stylesheet" type="text/css" href="shop_style.css">
		<title>Shop</title>
	</head>
	<body>
		<div id="clock" style="color:black; right: 2%; position: absolute; font-size: 1vw;">
		</div>
		<div class="wrapper_header">
			<div class="header" id="12">
				<p>ArianShop</p>
			</div>
		</div>
		<div class="lineBtn">`)
	for i := 0; i < len(products)/3; i++ {
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
    <script src="shop_script.js"></script>
</body>
</html>`)

}
