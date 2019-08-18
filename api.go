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

		// stockList := selectAllStockOfSymbol(data)
		dowList := selectDow()
		for i, v := range dowList {
			fmt.Println(v.CreatedAt)
			i++
		}
		// databaseResponse := DatabaseResponse{monitorSymbolList[0]}
		// databaseStockListResponse := DatabaseStockListResponse{[]Stock{monitorSymbolList[0], monitorSymbolList[1]}}
		// js, err := json.Marshal(databaseStockListResponse)
		// if err != nil {
		// 	http.Error(w, err.Error(), http.StatusInternalServerError)
		// 	return
		// }
		// w.Header().Set("Content-Type", "application/json")
		// w.Write(js)

		// filterEntriesWithinTimeset(stockList)
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
	// processTimelineStart()
	// checKIsBrokerageResponding()

	// dowList := selectDow()
	// dowMatchList := filterDowEntriesWithinTimeset(dowList, "2019-06-10T23:11:39", "2019-08-16T11:00:29")

	// fmt.Println(dowMatchList)

	stockList := selectAllStockOfSymbol("CAT")

	fmt.Println(stockList[0].CreatedAt)
	fmt.Println(stockList[1].CreatedAt)
	fmt.Println(stockList[5].CreatedAt)
	fmt.Println(stockList[(len(stockList) - 2)].CreatedAt)
	fmt.Println(stockList[(len(stockList) - 1)].CreatedAt)

	stockMatchList := filterStockEntriesWithinTimeset(stockList, "2019-08-05T13:32:12", "2019-08-16T11:00:31")

	// not case where less than second, but match on point.

	fmt.Println("break")
	fmt.Println(len(stockMatchList))
	fmt.Println(stockMatchList[0].CreatedAt)
	fmt.Println(stockMatchList[(len(stockMatchList) - 1)].CreatedAt)

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
