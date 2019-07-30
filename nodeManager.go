package main

import "fmt"

func topStockPull() {
}

func queryMultiStockPull(listSymbols []string) {
	json := `{
		"request_type": "top_stock_pull",
		"data": [`

	for i, v := range listSymbols {
		json += "\"" + v + "\""
		fmt.Println(i)
		fmt.Println(len(listSymbols))
		if i != (len(listSymbols) - 1) {
			json += ","
		}
		if i == (len(listSymbols) - 1) {
			json = json + `]}`
		}
	}

	fmt.Println(json)
	url := "http://localhost:3000/api/brokerage"
	post(url, json)
}
