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
	range1 := databaseQuery.Range1
	range2 := databaseQuery.Range2

	// "2019-08-06T00:32:12"
	// "2019-08-16T11:00:30"

	// "2019-08-06T00:32:12"
	// "2019-08-16T11:00:30"

	fmt.Println(requestType)
	fmt.Println(data)
	// fmt.Println(rangeForData)

	//Select all monitor symbol
	if requestType == "0" {
		monitorSymbolList := selectTempSymbolHold()
		monitorSymbolResponse := DatabaseMonitorSymbolListResponse{monitorSymbolList}
		js, err := json.Marshal(monitorSymbolResponse)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	}
	//Select all dow within range
	if requestType == "1" {
		dowList := selectDow()
		dowMatchList := filterDowEntriesWithinTimeset(dowList, range1, range2)

		dowListResponse := DatabaseDowListResponse{dowMatchList}
		js, err := json.Marshal(dowListResponse)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	}
	//Select all stock where symbol == data within range
	// if requestType == "2" {
	// 	stockList := selectAllStockOfSymbol(data)
	// 	stockMatchList := filterStockEntriesWithinTimeset(stockList, range1, range2)

	// 	stockListResponse := DatabaseStockListResponse{stockMatchList}
	// 	js, err := json.Marshal(stockListResponse)
	// 	if err != nil {
	// 		http.Error(w, err.Error(), http.StatusInternalServerError)
	// 		return
	// 	}
	// 	w.Header().Set("Content-Type", "application/json")
	// 	w.Write(js)
	// }

	if requestType == "selectAllStockWisemen" {
		stockList := selectAllStockWisemen()
		stockListResponse := DatabaseStockListResponse{stockList}
		js, err := json.Marshal(stockListResponse)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	}

	//Whale operations
	if requestType == "selectAllStockWhale" {
		stockList := selectAllStockWhale()
		// stockMatchList := filterStockEntriesWithinTimeset(stockList, range1, range2)

		stockListResponse := DatabaseStockListResponse{stockList}
		js, err := json.Marshal(stockListResponse)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	}

	//Metrics whale
	if requestType == "insertMetricsWhale" {
		// data := databaseQuery.Data
		dataList := databaseQuery.Data
		insertMetricsWhale(dataList[0], dataList[1], dataList[2], dataList[3], dataList[4])

		js, err := json.Marshal("success")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	}
	if requestType == "selectMetricsWhale" {
		metricsWhale := selectMetricsWhale()
		// stockMatchList := filterStockEntriesWithinTimeset(stockList, range1, range2)

		metricsWhaleResponse := DatabaseMetricsWhaleResponse{MetricsWhale: metricsWhale}
		js, err := json.Marshal(metricsWhaleResponse)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	}
	if requestType == "dropMetricsWhale" {
		dropMetricsWhale()
		// stockMatchList := filterStockEntriesWithinTimeset(stockList, range1, range2)
		// stockListResponse := DatabaseStockListResponse{stockList}
		js, err := json.Marshal("success")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	}
	if requestType == "createMetricsWhale" {
		createMetricsWhale()
		js, err := json.Marshal("success")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	}

	//Metrics wisemen
	if requestType == "insertMetricsWisemen" {
		// data := databaseQuery.Data
		dataList := databaseQuery.Data
		insertMetricsWisemen(dataList[0], dataList[1], dataList[2], dataList[3], dataList[4])

		js, err := json.Marshal("success")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	}
	if requestType == "selectMetricsWisemen" {
		metricsWhale := selectMetricsWisemen()
		// stockMatchList := filterStockEntriesWithinTimeset(stockList, range1, range2)

		metricsWhaleResponse := DatabaseMetricsWhaleResponse{MetricsWhale: metricsWhale}
		js, err := json.Marshal(metricsWhaleResponse)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	}
	if requestType == "dropMetricsWisemen" {
		dropMetricsWisemen()
		// stockMatchList := filterStockEntriesWithinTimeset(stockList, range1, range2)
		// stockListResponse := DatabaseStockListResponse{stockList}
		js, err := json.Marshal("success")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	}
	if requestType == "createMetricsWisemen" {
		createMetricsWisemen()
		js, err := json.Marshal("success")
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
	// processAppendDayOfWeekToStock(Stock{})

	// processWisemenQueryStockSet()

	//initial insert whale metrics.

	//Begin processTimeline upon condition isMarketClosed == false
	// insertMetricsWhale("hey", "hey1", "hey2", "hey3", "hey4")

	// response := selectMetricsWhale()
	// fmt.Println(response)

	// dropMetricsWhale()
	// createMetricsWhale()

	// processTimelineStart()

	// handleTSPRefresh()
	// processDowWebscrape()
	// processWisemenQueryStockSet()
	// processWhaleQueryStockSet()
	// fmt.Println(createStockTimeStamp())

	// handleWhaleQueryStockList()

	results := selectStockWhale("CMG")
	// fmt.Println(results)
	processDetectStockWhaleWhereDayIsAtIndex(results)

	// handleTSPRefresh()
	// handleFillHolds()

	// strings.Contains("something", "some")
	// isBoolConditionMet := strings.Contains("ABC.L", ".")
	// if isBoolConditionMet {
	// 	fmt.Println("true")
	// }
	// if isBoolConditionMet == false {
	// 	fmt.Println("false")
	// }

	// cycleMapPool = map[string]*Cycle{}
	// processWisemenQueryStockSet()
	// time.Sleep(time.Duration(10) * time.Second)
	// operatingCycle := cycleMapPool["handleWisemenQueryStockList"]
	// cancelCycle(operatingCycle)
	// time.Sleep(time.Duration(10) * time.Second)
	// // operatingCycle.BooleanOperate = true
	// processWisemenQueryStockSet()

	// fmt.Println(selectAllStockWisemen())

	// processFillHolds()
	// handleWisemenQueryStockList()
	// processWisemenQueryStockSet()
	// processWhaleQueryStockSet()

	// handleWisemenQueryStockList()
	// handleWhaleQueryStockList()

	// checKIsBrokerageResponding()
	// insertWhaleSymbolHold("AAPL", false)
	// checkWhaleDelimiterMet()
	// dropWisemenSymbolHold()
	// dropTempSymbolHold()
	// createWisemenSymbolHold()
	// createTempSymbolHold()

	// dowList := selectDow()
	// dowMatchList := filterDowEntriesWithinTimeset(dowList, "2019-06-10T23:11:39", "2019-08-16T11:00:29")

	// fmt.Println(dowMatchList)

	// fmt.Println(stockList[0].CreatedAt)
	// fmt.Println(stockList[1].CreatedAt)
	// fmt.Println(stockList[5].CreatedAt)
	// fmt.Println(stockList[(len(stockList) - 2)].CreatedAt)
	// fmt.Println(stockList[(len(stockList) - 1)].CreatedAt)

	// not case where less than second, but match on point.

	// fmt.Println("break")
	// fmt.Println(len(stockMatchList))
	// fmt.Println(stockMatchList[0].CreatedAt)
	// fmt.Println(stockMatchList[(len(stockMatchList) - 3)].CreatedAt)
	// fmt.Println(stockMatchList[(len(stockMatchList) - 2)].CreatedAt)
	// fmt.Println(stockMatchList[(len(stockMatchList) - 1)].CreatedAt)

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
