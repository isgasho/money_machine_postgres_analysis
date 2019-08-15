package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var isTimeMonitoringLoop bool

func coolPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the coolPage!")
	fmt.Println("Endpoint Hit: coolPage")
}

func stockQuery(rw http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	// var brokerageQuery BrokerageQuery
	var requestCase string
	err := decoder.Decode(&requestCase)
	if err != nil {
		panic(err)
		fmt.Println("Stock error1")
	}
	fmt.Println(requestCase)
}

func databaseQuery(w http.ResponseWriter, req *http.Request) {
	var databaseQuery DatabaseQuery
	err := json.NewDecoder(req.Body).Decode(&databaseQuery)
	check(err)
	requestType := databaseQuery.RequestType
	data := databaseQuery.Data
	rangeForData := databaseQuery.RangeForData

	fmt.Println(requestType)
	fmt.Println(data)
	fmt.Println(rangeForData)

	//Select all monitor symbol
	if requestType == "0" {
		fmt.Println("case 0")
		// monitorSymbolList := selectMonitorSymbol()
		// databaseResponse := DatabaseResponse{monitorSymbolList[0]}
		// js, err := json.Marshal(databaseResponse)
		// if err != nil {
		// 	http.Error(w, err.Error(), http.StatusInternalServerError)
		// 	return
		// }
		// w.Header().Set("Content-Type", "application/json")
		// w.Write(js)
	}
	//Select all stock where symbol == data
	if requestType == "1" {
		fmt.Println("case 1", data)
		//Handle data
		// stockList := selectAllStockOfSymbol(data)

		//Return ID list, and then stock objects
		// databaseResponse := DatabaseResponse{stockList[0]}
		// js, err := json.Marshal(databaseResponse)
		// if err != nil {
		// 	http.Error(w, err.Error(), http.StatusInternalServerError)
		// 	return
		// }
		// w.Header().Set("Content-Type", "application/json")
		// w.Write(js)
	}
	//Select all stock where symbol == data
	if requestType == "2" {
		fmt.Println("case 2")
		//Handle range within created time.

		monitorSymbolList := selectAllStockOfSymbol(data)
		// databaseResponse := DatabaseResponse{monitorSymbolList[0]}
		databaseStockListResponse := DatabaseStockListResponse{[]Stock{monitorSymbolList[0], monitorSymbolList[1]}}
		js, err := json.Marshal(databaseStockListResponse)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	}

	// databaseResponse := DatabaseResponse{"Alex", []string{"snowboarding", "programming"}}
}

func handleRequests() {
	http.HandleFunc("/databaseQuery", databaseQuery)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	//Open server API connections
	//Begin Select data retrieval for particular processes.
	// go handleRequests()

	//Begin processTimeline upon condition isMarketClosed == false
	processTimelineStart()
	// checKIsBrokerageResponding()
	fmt.Scanln()
	fmt.Println("done")
}

func initTimeMonitoring() {
	// time excution process looped on regular basis
	// triggered every few seconds
	i := 0
	for i < 3 { //isTimeMonitoringLoop {
		// timeConditionExecutionProcess()
		// if i == 3 {
		// 	fmt.Println("is false")
		// 	isTimeMonitoringLoop = false
		// }
		fmt.Println("Awesome sauce")
		// time.Sleep(3 * time.Second)
		i++
	}
}
