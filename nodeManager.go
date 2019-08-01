package main

import (
	"reflect"
)

func topStockPull() {
}

// type fn func(params ...interface{})

func queryMultiStockPull(params ...interface{}) {

	listVal := reflect.ValueOf(params[0])
	var listSymbolsInterface interface{} = listVal.Index(0).Interface()

	listSymbols := listSymbolsInterface.([]string)

	symbol1 := listSymbols[0]
	symbol2 := listSymbols[1]
	symbol3 := listSymbols[2]

	// query_multi_stock
	json := `{
		"request_type": "test14",
		"data": [
		`

	json += "\"" + symbol1 + "\","
	json += "\"" + symbol2 + "\","
	json += "\"" + symbol3 + "\""
	// json += ","
	json = json + `]}`

	// fmt.Println(json)
	url := "http://localhost:3000/api/brokerage"
	// response := ""
	post(url, json)
	// fmt.Println(response)

}
