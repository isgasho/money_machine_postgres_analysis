package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
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

	if requestType == "postEvalResultsWhale" {
		dropEvalResultsWhale()
		createEvalResultsWhale()

		fmt.Println("hit")
		fmt.Println(len(data))
		indexSamplePull := 0
		evalToAppend := EvalResultsWhale{}
		evalList := []EvalResultsWhale{}
		for i, v := range data {
			if indexSamplePull == 0 {
				evalToAppend.Symbol = v
			}
			if indexSamplePull == 1 {
				evalToAppend.IsPatternMet = v
			}
			if indexSamplePull == 2 {
				evalToAppend.IsBreachWorthy = v
				evalList = append(evalList, evalToAppend)
				evalToAppend = EvalResultsWhale{}
				indexSamplePull = 0
				continue
			}
			indexSamplePull++
			i++
		}
		// fmt.Println(evalList)

		for indexEvalResult, evalResult := range evalList {
			insertEvalResultsWhale(evalResult)
			indexEvalResult++
		}
		// storeResponse := DatabaseStockListResponse{}
		// js, err := json.Marshal(storeResponse)
		// if err != nil {
		// 	http.Error(w, err.Error(), http.StatusInternalServerError)
		// 	return
		// }
		w.Header().Set("Content-Type", "application/json")
		// w.Write(js)
	}

	//Select all monitor symbol
	// if requestType == "0" {
	// 	monitorSymbolList := selectTempSymbolHold()
	// 	monitorSymbolResponse := DatabaseMonitorSymbolListResponse{monitorSymbolList}
	// 	js, err := json.Marshal(monitorSymbolResponse)
	// 	if err != nil {
	// 		http.Error(w, err.Error(), http.StatusInternalServerError)
	// 		return
	// 	}
	// 	w.Header().Set("Content-Type", "application/json")
	// 	w.Write(js)
	// }
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
	// if requestType == "selectAllStockWhale" {
	// 	stockList := selectAllStockWhale()
	// 	// stockMatchList := filterStockEntriesWithinTimeset(stockList, range1, range2)

	// 	stockListResponse := DatabaseStockListResponse{stockList}
	// 	js, err := json.Marshal(stockListResponse)
	// 	if err != nil {
	// 		http.Error(w, err.Error(), http.StatusInternalServerError)
	// 		return
	// 	}
	// 	w.Header().Set("Content-Type", "application/json")
	// 	w.Write(js)
	// }

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
		insertMetricsWisemen(dataList[0], dataList[1], dataList[2], dataList[3], dataList[4], dataList[5], dataList[6], dataList[7], dataList[8])

		js, err := json.Marshal("success")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	}
	if requestType == "selectMetricsWisemen" {
		metricsWisemen := selectMetricsWisemen()
		// stockMatchList := filterStockEntriesWithinTimeset(stockList, range1, range2)

		metricsWisemenResponse := DatabaseMetricsWisemenResponse{MetricsWisemen: metricsWisemen}
		js, err := json.Marshal(metricsWisemenResponse)
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
	//
	// if requestType == "postMetricsWisemen" {
	// 	dataList := databaseQuery.Data
	// 	fmt.Println("dataList")
	// 	fmt.Println(dataList)
	// 	insertMetricsWisemen(dataList[0], dataList[1], dataList[2], dataList[3], dataList[4])

	// 	js, err := json.Marshal("success")
	// 	if err != nil {
	// 		http.Error(w, err.Error(), http.StatusInternalServerError)
	// 		return
	// 	}
	// 	w.Header().Set("Content-Type", "application/json")
	// 	w.Write(js)
	// }
	//Trade execution
	// if requestType == "isBuyWisemen" {
	// 	dataList := databaseQuery.Data
	// 	fmt.Println("isBuyWisemen")
	// 	fmt.Println(dataList)
	// 	// insertMetricsWisemen(dataList[0], dataList[1], dataList[2], dataList[3], dataList[4])

	// 	js, err := json.Marshal("success")
	// 	if err != nil {
	// 		http.Error(w, err.Error(), http.StatusInternalServerError)
	// 		return
	// 	}
	// 	w.Header().Set("Content-Type", "application/json")
	// 	w.Write(js)
	// }
	if requestType == "selectOrderInformationWisemen" {
		// dataList := databaseQuery.Data
		fmt.Println("selectOrderInformationWisemen")
		// fmt.Println(dataList)
		dataList := selectOrderInformationWisemen()

		databaseOrderInformationWisemenResponse := DatabaseOrderInformationWisemenResponse{OrderInformationWisemen: dataList}
		js, err := json.Marshal(databaseOrderInformationWisemenResponse)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	}

	if requestType == "tradeBuyWisemen" {
		dataList := databaseQuery.Data
		fmt.Println("tradeBuyWisemen")
		fmt.Println("dataList")
		fmt.Println(dataList)

		//process trade.
		handleTradeWisemen(dataList[0], dataList[1])
		time.Sleep(time.Duration(10) * time.Second)

		//Begin process monitoring for buy fulfilled.
		processCheckIsTradeBought(dataList[0])
		js, err := json.Marshal("success")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	}

	// if requestType == "initiateEarlySellWisemen" {
	// 	dataList := databaseQuery.Data
	// 	fmt.Println("initiateEarlySellWisemen")
	// 	fmt.Println("dataList")
	// 	fmt.Println(dataList)

	// 	//Handle process
	// 	intiateSellSystemProtocol()

	// 	handleSellWisemen(dataList[0], dataList[1])
	// 	js, err := json.Marshal("success")
	// 	if err != nil {
	// 		http.Error(w, err.Error(), http.StatusInternalServerError)
	// 		return
	// 	}
	// 	w.Header().Set("Content-Type", "application/json")
	// 	w.Write(js)
	// 	//Delay before monitor cycle
	// 	// time.Sleep(time.Duration(10) * time.Second)
	// 	// intiateMonitorTradeWisemon()
	// }

	if requestType == "postPriceSellDelimiterMetrics" {
		dataList := databaseQuery.Data
		fmt.Println("postPriceSellDelimiterMetrics")
		fmt.Println("dataList")
		fmt.Println(dataList)
		//sell at limit order submission.
		//
		handleSellLimitWisemen(dataList[0])
		//handle monitor
		processMonitorSell(dataList[0], dataList[1], dataList[2])

		js, err := json.Marshal("success")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	}

	if requestType == "postSellMarket" {
		dataList := databaseQuery.Data
		fmt.Println("postSellMarket")
		fmt.Println("dataList")
		fmt.Println(dataList)
		//process trade.
		//

		///get current holding qty
		handleSellAtMarket(dataList[0])
		// queryTradeSellMarket()
		// process
		// processMonitorSell(dataList[0], dataList[1], dataList[2])
		// processMonitorSell("20.20", "1430")
		// handleTradeWisemen(dataList[0], dataList[1])
		js, err := json.Marshal("success")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	}

	//

	// if requestType == "selectMetricsWisemen" {
	// 	dataList := databaseQuery.Data
	// 	fmt.Println("dataList")
	// 	fmt.Println(dataList)
	// 	selectMetricsWisemen(dataList[0], dataList[1], dataList[2], dataList[3], dataList[4])

	// 	js, err := json.Marshal("success")
	// 	if err != nil {
	// 		http.Error(w, err.Error(), http.StatusInternalServerError)
	// 		return
	// 	}
	// 	w.Header().Set("Content-Type", "application/json")
	// 	w.Write(js)
	// }

	// databaseResponse := DatabaseResponse{"Alex", []string{"snowboarding", "programming"}}
}

func handleRequests() {
	http.HandleFunc("/databaseQuery", databaseQuery)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func createRecordSystemMonthContainer() RecordSystemMonthContainer {
	recordSystemMonthContainer := RecordSystemMonthContainer{}
	//Static populate days in month
	recordSystemMonth := RecordSystemMonth{IntMonthOfYear: 11, IntNumberOfDays: 30}
	recordSystemMonth1 := RecordSystemMonth{IntMonthOfYear: 12, IntNumberOfDays: 31}
	recordSystemMonth2 := RecordSystemMonth{IntMonthOfYear: 1, IntNumberOfDays: 31}
	recordSystemMonth3 := RecordSystemMonth{IntMonthOfYear: 2, IntNumberOfDays: 29}
	recordSystemMonth4 := RecordSystemMonth{IntMonthOfYear: 3, IntNumberOfDays: 31}
	recordSystemMonth5 := RecordSystemMonth{IntMonthOfYear: 4, IntNumberOfDays: 30}
	recordSystemMonth6 := RecordSystemMonth{IntMonthOfYear: 5, IntNumberOfDays: 31}
	recordSystemMonth7 := RecordSystemMonth{IntMonthOfYear: 6, IntNumberOfDays: 30}
	recordSystemMonth8 := RecordSystemMonth{IntMonthOfYear: 7, IntNumberOfDays: 31}
	recordSystemMonth9 := RecordSystemMonth{IntMonthOfYear: 8, IntNumberOfDays: 31}
	recordSystemMonth10 := RecordSystemMonth{IntMonthOfYear: 9, IntNumberOfDays: 30}
	recordSystemMonth11 := RecordSystemMonth{IntMonthOfYear: 10, IntNumberOfDays: 31}

	recordSystemMonthContainer.RecordSystemMonthList = append(recordSystemMonthContainer.RecordSystemMonthList, recordSystemMonth, recordSystemMonth1, recordSystemMonth2, recordSystemMonth3, recordSystemMonth4, recordSystemMonth5, recordSystemMonth6, recordSystemMonth7, recordSystemMonth8, recordSystemMonth9, recordSystemMonth10, recordSystemMonth11)
	return recordSystemMonthContainer
}

func calculateIsResetDayRecord() {
	recordSystemMonthContainer := createRecordSystemMonthContainer()
	for i, v := range recordSystemMonthContainer.RecordSystemMonthList {
		fmt.Println(v.IntMonthOfYear)
		fmt.Println(v.IntNumberOfDays)
		i++
	}
	year, month, day := getDate()
	fmt.Println(year)
	fmt.Println(month)
	fmt.Println(day)
	//depending on cuurent month get number of days
	//get record month and day of trade record.
	//determine 7 days from that value
	//if current day is greater thab that value reset.
}
func main() {
	go handleRequests()

	// detectDownDay()
	// handleCalculateDownDay()
	// handleCalculateCashDay()

	// resetTempSymbolHold()
	// resetStockWisemenSymbolHold()
	// resetStockWisemen()

	// queryStopTwi()
	// queryStartTwi()
	// processOverarchTopStock()
	// processWisemenQueryStockSet()

	//before operation query
	// isBool := overarchIsTradeDay()
	// fmt.Println(isBool)
	// dropDow()
	// createDow()
	// insertDow("23000")
	// insertDow("23000")
	// insertDow("23000")
	// insertDow("23000")
	// cashAccountCheckUnsettledFunds()

	// dropMetricsWisemen()
	// createMetricsWisemen()
	// informationAtTrade := InformationAtTrade{}
	// insertInformationAtTrade(informationAtTrade)
	// handleInsertInformationAtTrade("MTW")
	// handleInsertInformationAtTrade("MTW")
	//Create artificial or go with bad metrics... in this case bad metrics is fine.

	// handleHistoryDayListArbitration()
	// transactionHistory := TransactionHistory{Symbol: "MTW"}
	// calculateTransactionHistory(transactionHistory)
	//
	// dropDow()
	// createDow()
	// insertDow("34000")
	// insertDow("34000")
	// insertDow("34000")
	// insertDow("34000")

	// dowList := selectDow()
	// fmt.Println(dowList)
	// transactionHistory := TransactionHistory{Symbol: "MTW"}
	// wrapUpWisemenOutcome(transactionHistory)
	// handleDowWebscrape
	// listTradeResultStore := selectTradeResultStore("wisemen")
	// fmt.Println(listTradeResultStore)
	//
	// handleInformationAtTradeDayListArbitration("MTW")

	// holdingWisemen := HoldingWisemen{Symbol: "symbol", Price: "v.Price", Qty: "v.Qty", OrderStatus: "pending eval"}
	// insertHoldingWisemen(holdingWisemen)
	// holdingWisemen1 := selectHoldingWisemen()
	// fmt.Println(holdingWisemen1)

	// dropHoldingWisemen()
	// createHoldingWisemen()

	// processCheckIsTradeBought("VICI")

	// time.Sleep(time.Duration(15) * time.Second)
	// fmt.Println("len(cycleMapPool)")
	// fmt.Println(len(cycleMapPool))
	// cancelCycle(cycleMapPool["handleCheckIsTradeBought"])
	// fmt.Println("len(cycleMapPool)")
	// fmt.Println(len(cycleMapPool))
	// func processCheckIsTradeBought(symbol string) {
	// 	// go handleCheckIsTradeBought()
	// 	//THe idea is to check every 5 seconds, and if a trade evaluation is positive,
	// 	//or if the time delimiter for checking is met, then cancle this cycle and record results in DB.
	// 	createCycle(10, 100000, handleCheckIsTradeBought, "handleCheckIsTradeBought", []string{symbol})
	// 	operatingCycle := cycleMapPool["handleCheckIsTradeBought"]
	// 	go startCycle(operatingCycle)
	// 	initialWhaleStockQueryPerformed = true
	// }

	// systemStartProcesses()
	// handleOverarchTopStock()
	// twiWebscrape()

	// dataList := []string{"MTW"}
	// handleCheckIsTradeBought()
	// processCheckIsTradeBought(dataList[0])

	// handleCheckIsTradeBought([]string{"MTW"})

	// calculateIsResetDayRecord()
	// monthString, dayInt := getDate()
	// fmt.Println(monthString, dayInt)
	// isSellShowingInHistory("CRC")
	// fmt.Println(isSymbolPresentInHolding("VICI"))
	// queryMultiStockPull()

	// monitorSell("VICI")
	//dropoff price, time delimiter
	// processMonitorSell("VICI", "20.20", "1430")

	// monitorSell
	// calculateIsTimeDelimiterMetSell("1330")

	// isDropPriceMet := calculateIsDropPriceMet("VICI")
	// // dropPrice := calculateIsDropPriceMet("VICI")
	// fmt.Println(isDropPriceMet)
	// transactionHistory := TransactionHistory{Symbol: "VICI"}
	// wrapUpWisemenOutcome(transactionHistory)

	// queryHolding()
	// container := getAllHolding()
	// for i, v := range container.ListHolding {
	// 	fmt.Println(v)
	// 	i++
	// }
	//if latest entry
	// parseLatestHistory()

	//sell procedure.

	// healthCheck()
	// query
	// queryTSP()
	// handleTSPRefresh()

	// dropTempSymbolHoldHigh()
	// createTempSymbolHoldHigh()
	// dropWisemenSymbolHold()
	// createWisemenSymbolHold()

	// processOverarchTopStock()
	// time.Sleep(time.Duration(40) * time.Second)
	// processWisemenQueryStockSet()

	// processDowWebscrape()
	// processWisemenQueryStockSet()
	// processWhaleQueryStockSet()

	// listStocks := twiWebscrape()
	// // fmt.Println(listStocks)
	// for i, v := range listStocks {
	// 	fmt.Println(v.Symbol)
	// 	fmt.Println(v.Pchg)
	// 	i++
	// }

	// dropWisemenSymbolHold()
	// createWisemenSymbolHold()

	// listTempDuplicantFiltered := []Stock{Stock{Symbol: "test1", Pchg: "5.0"}, Stock{Symbol: "test2", Pchg: "10.0"}, Stock{Symbol: "test3", Pchg: "20.0"}}
	// i := 0
	// topStockList := []Stock{}

	// listTempDuplicantFiltered = removeElement(listTempDuplicantFiltered, listTempDuplicantFiltered[2].Symbol)

	// fmt.Println(listTempDuplicantFiltered)
	//Filter for top pchg top 3
	// for i < 3 {
	// 	// 	// remove highest index 3 times, to get top stocks.
	// 	// 	//Pop top stock each iteration
	// 	highestStockIndex := 0
	// 	for indexTempDuplicantFiltered, tempDuplicantFiltered := range listTempDuplicantFiltered {
	// 		if indexTempDuplicantFiltered == 0 {
	// 			highestStockIndex = indexTempDuplicantFiltered
	// 			continue
	// 		}

	// 		floatHighest := 0.0
	// 		floatCurrent := 0.0
	// 		if s, err := strconv.ParseFloat(listTempDuplicantFiltered[highestStockIndex].Pchg, 64); err == nil {
	// 			floatHighest = s
	// 		}
	// 		if s1, err := strconv.ParseFloat(tempDuplicantFiltered.Pchg, 64); err == nil {
	// 			floatCurrent = s1
	// 		}

	// 		if floatCurrent > floatHighest {
	// 			fmt.Println("previousHighest")
	// 			fmt.Println(listTempDuplicantFiltered[highestStockIndex].Pchg)
	// 			highestStockIndex = indexTempDuplicantFiltered
	// 			fmt.Println("index")
	// 			fmt.Println(i)
	// 			fmt.Println("listTempDuplicantFiltered[highestStockIndex].Pchg")
	// 			fmt.Println(listTempDuplicantFiltered[highestStockIndex].Pchg)
	// 			fmt.Println(tempDuplicantFiltered.Pchg)
	// 		}
	// 	}
	// 	topStockList = append(topStockList, listTempDuplicantFiltered[highestStockIndex])
	// 	fmt.Println(listTempDuplicantFiltered[highestStockIndex])
	// 	if i < 2 {
	// 		listTempDuplicantFiltered = removeElement(listTempDuplicantFiltered, listTempDuplicantFiltered[highestStockIndex].Symbol)
	// 	}
	// 	// fmt.Println("listAltered")
	// 	// fmt.Println(listTempDuplicantFiltered)
	// 	i++
	// }
	// fmt.Println("topStockList")
	// fmt.Println(topStockList)

	// func removeElement(nums []int, val int) int {
	// listNumbers := []Stock{Stock{Symbol: "1"}, Stock{Symbol: "2"}, Stock{Symbol: "3"}}
	// response := removeElement(listNumbers, "1")
	// fmt.Println(response)
	// fmt.Println(listNumbers)
	// processDowWebscrape()
	// processTwiWebscrape()
	// queryWebscrapeTwi
	// handleSellWisemen("CRC")
	// // splitFloatAfterSecondDecimalPlace(251.298100)

	// // dropMetricsWisemen()
	// // createMetricsWisemen()

	// stockList := getCurrentPriceStatsForStock([]string{"AAPL"})
	// for i, v := range stockList {
	// 	fmt.Println("v.Bid")
	// 	fmt.Println(v.Bid)
	// 	fmt.Println("v.Ask")
	// 	fmt.Println(v.Ask)
	// 	fmt.Println("v.Last")
	// 	fmt.Println(v.Last)
	// 	i++
	// }

	// //change order
	// containerOrder := getAllOrders()
	// order := Order{}
	// for i, v := range containerOrder.ListOrders {
	// 	if v.Symbol == "VICI" {
	// 		order = v
	// 	}
	// 	i++
	// }
	// fmt.Println(order)
	// queryTradeChangeLimit(order.SVI, order.Symbol, "23.50", order.Qty)

	// beginAlgorithmRecordingWisemen("test")
	// completeAlgorithmRecordingWisemen()

	// selectAlgorithmEvaluationForDay("test")
	// intValue := roundDown(3.43)
	// fmt.Println(intValue)
	// response := queryBalance()
	// fmt.Println(response)
	// response := queryOrders()
	// fmt.Println(response)
	// processCheckIsTradeBought("VICI")
	// balances := queryBalance()
	// fmt.Println(balances)

	// orderContainer := getAllOrders()

	//get all open orders.

	// order := Order{SVI: "SVI-6085037565"}
	// for i, v := range orderContainer.ListOrders {
	// 	order = v
	// 	i++
	// }
	// // var order = req.body.data.order;
	// fmt.Println(order)
	// func queryCancelOrder(symbol string, limitPrice string, qty string) string {
	// queryCancelOrder(order.SVI)

	// queryOrder := queryOrders()
	// fmt.Println(queryOrder)
	//Open server API connections
	//Begin Select data retrieval for particular processes.
	//

	// intiateMonitorTradeWisemon()
	// <span class="IsqQVc NprOob iXPM7ggEYSKk-zJFzKq8ukm8">26,958.06</span>
	// queryWebscrape()
	// processDowWebscrape()
	// parseDowWebscrape("")
	// processCheckIsTradeBought("VICI")

	// dropMetricsWisemen()
	// createMetricsWisemen()

	// getAllOrders()
	// dropTradeEnteredInformation()
	// createTradeEnteredInformation()

	// deleteTradeEnteredInformation("test")
	// dropTradeEnteredInformation()
	// createTradeEnteredInformation()
	// process check on cycle if balance
	// processCheckIsTradeBought()

	// queryCheckBalance()
	// queryTradeBuyLimit()

	// func calculateAmountOfStockToBuy(pricePointOfStock float32, balance float32) {

	// amountOfBuy := calculateAmountOfStockToBuy(3.34, 5025.94)
	// fmt.Println(amountOfBuy)

	// handleTradeWisemen()

	// response := queryTradeCheckBalance()
	// fmt.Println(response)

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
	// deleteAllStockOfSymbolInWhale("CMG")

	//delete all of symbol from whale.

	//Test if //if isPresentInDB {
	//if days are different

	//

	// Discover where the process is failing and repair it.

	// results := selectStockWhale("CMG")
	// // fmt.Println(results[0].Vl)
	// // fmt.Println(results)
	// for i, v := range results {
	// 	fmt.Println(v.Vl)
	// 	i++
	// }

	//How to add if it's a different day?
	//

	//
	// processDetectStockWhaleWhereDayIsAtIndex(results)

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
