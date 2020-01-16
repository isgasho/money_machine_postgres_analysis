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

	if requestType == "postEvalResultsWhale" {
		// dropEvalResultsWhale()
		// createEvalResultsWhale()

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

		// for indexEvalResult, evalResult := range evalList {
		// 	// insertEvalResultsWhale(evalResult)
		// 	indexEvalResult++
		// }
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

	if requestType == "selectStockWisemen" {
		// stockList := selectStockWisemen()
		symbolList := selectWisemenSymbolHold()
		responseList := queryMultiStockPull(symbolList)
		stockList := parseStockSetQuery(responseList)
		//
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
	if requestType == "truncateMetricsWhale" {
		truncateMetricsWhale()
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

	//Metrics wisemen
	if requestType == "insertMetricsWisemen" {
		// data := databaseQuery.Data
		dataList := databaseQuery.Data
		insertMetricsWisemen(dataList[0], dataList[1], dataList[2], dataList[3], dataList[4], dataList[5], dataList[6])

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
		// dropMetricsWisemen()
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
		// createMetricsWisemen()
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
	// if requestType == "selectOrderInformationWisemen" {
	// 	// dataList := databaseQuery.Data
	// 	fmt.Println("selectOrderInformationWisemen")
	// 	// fmt.Println(dataList)
	// 	dataList := selectOrderInformationWisemen()

	// 	databaseOrderInformationWisemenResponse := DatabaseOrderInformationWisemenResponse{OrderInformationWisemen: dataList}
	// 	js, err := json.Marshal(databaseOrderInformationWisemenResponse)
	// 	if err != nil {
	// 		http.Error(w, err.Error(), http.StatusInternalServerError)
	// 		return
	// 	}
	// 	w.Header().Set("Content-Type", "application/json")
	// 	w.Write(js)
	// }

	if requestType == "tradeBuyWisemen" {
		dataList := databaseQuery.Data
		overarchTradeWisemen(dataList)
		js, err := json.Marshal("success")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	}

	if requestType == "altIntervalBuyWisemen" {
		dataList := databaseQuery.Data
		fmt.Println("altIntervalBuyWisemen")
		fmt.Println("dataList")
		fmt.Println(dataList)
		listAltIntervalBuyWisemen := selectAltIntervalBuyWisemen()
		js, err := json.Marshal(listAltIntervalBuyWisemen[0].IsAltIntervalOperation)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	}

	// if requestType == "postNoBuyWisemen" {
	// 	dataList := databaseQuery.Data
	// 	fmt.Println("postNoBuyWisemen")
	// 	fmt.Println("dataList")
	// 	fmt.Println(dataList)
	// 	// listAltIntervalBuyWisemen := selectAltIntervalBuyWisemen()

	// 	//TransactionHistory reason for failure append to model.
	// 	transactionHistory := TransactionHistory{Symbol: dataList[0]}
	// 	wrapUpWisemenOutcomeNoBuy(transactionHistory)
	// 	js, err := json.Marshal("success")
	// 	if err != nil {
	// 		http.Error(w, err.Error(), http.StatusInternalServerError)
	// 		return
	// 	}
	// 	w.Header().Set("Content-Type", "application/json")
	// 	w.Write(js)
	// }

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
		// dataList[0] = "GSX"
		// dataList[1] = "21.50"
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

	if requestType == "pullMarketOpenAnaylsis" {
		// dataList := databaseQuery.Data
		//
		marketOpenAnaylsisIsMarketClosed := selectMarketOpenAnalysis()[0].IsMarketClosed
		js, err := json.Marshal(marketOpenAnaylsisIsMarketClosed)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	}

	if requestType == "metricsWisemenUpdateHighDelimiter" {
		dataList := databaseQuery.Data
		fmt.Println("metricsWisemenUpdateHighDelimiter")
		fmt.Println("dataList")
		fmt.Println(dataList)
		//process trade
		//
		truncateMetricsWisemen()
		insertMetricsWisemen("20.00", "4.0", "8.0", "0", ".02", ".1", "1330")
		js, err := json.Marshal("success")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	}

	if requestType == "selectInformationAtTrade" {
		dataList := databaseQuery.Data
		fmt.Println("selectInformationAtTrade")
		fmt.Println("dataList")
		fmt.Println(dataList)
		//process trade
		//
		// truncateMetricsWisemen()
		// insertMetricsWisemen("20.00", "4.0", "8.0", "0", ".02", ".1", "1330")
		response := "empty"
		listInformationAtTrade := selectInformationAtTrade()
		if len(listInformationAtTrade) != 0 {
			response = "entries"
		}
		js, err := json.Marshal(response)
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

	// handleCalculateDownDay()
	// handleCalculateCashDay()
	// globalCacheStockWisemenTopTier()
	// resetTimeOperations()
	// postMarketClosedEmail()
	processTimelineStart()

	//
	// handleCalculateCashDay()
	// overarchIsTradeDay()
	// postIsTradeDayEmail()
	// handleCalculateCashDay()
	// limitPrice := calculateLimitPriceForBuy("AGTC")
	// fmt.Println("limitPrice")
	// fmt.Println(limitPrice)

	// handleTradeWisemen("AGTC", "0.0")
	// hourDelimiter := "22"
	// minuteDelimiter := "11"

	// if len(minuteDelimiter) == 1 {
	// 	minuteDelimiter = "0" + minuteDelimiter
	// }
	// //
	// timeDelimiter := hourDelimiter + minuteDelimiter
	// fmt.Println("created timeDelimiter")
	// fmt.Println(timeDelimiter)
	// //update
	// isTimeDelimiterMet := calculateIsTimeDelimiterMetSell(timeDelimiter)

	// fmt.Println("isTimeDelimiterMet")
	// fmt.Println(isTimeDelimiterMet)

	// truncateTempSymbolHoldHigh()
	// truncateWisemenSymbolHold()
	// intervalTest := 0
	// for intervalTest < 100 {
	// 	handleOverarchTopStock()
	// 	time.Sleep(time.Duration(5) * time.Second)
	// 	intervalTest++
	// }
	// handleOverarchTopStock()
	// handleEndOfDayDowScrape()
	// handleNoBuyOnTradeDay()
	// resetTimeOperations()

	// dowList := selectDow()
	// formattedDowList := formatDowListRemoveCommaValues(dowList)

	// fmt.Println("formattedDowList")
	// fmt.Println(formattedDowList)
	//

	// truncateDownDayEvaluation()

	// selectDownDayEvaluation()

	// dow := handleDowWebscrape()
	// insertDow(dow)
	// dowVal := selectDow()
	// fmt.Println(dowVal)

	// insertEndOfDayDow("28,703.38")
	// truncateDow()

	// dowValue := handleDowWebscrape()
	// insertDow(dowValue)

	// insertDow(dowValue)

	// insertDow(dowValue)

	// insertDow(dowValue)

	// handleEndOfDayDowScrape()
	// handleCalculateDownDay()
	// truncateDownDayEvaluation()
	// truncateTSPCollectionStatement()
	// truncateTSPCollectionStatement()
	// handleTSPCollectionStatementPhase1()
	// handleTSPCollectionStatementPhase2()

	// tspCollectionStatement := selectTSPCollectionStatement()[0]

	// fmt.Println("tspCollectionStatement")
	// fmt.Println(tspCollectionStatement.DataCache)

	//

	// handleTSPCollectionStatementPhase1()
	// handleTSPCollectionStatementPhase2()

	// // truncateTSPCollectionStatement()

	// tspCollectionStatement1 := selectTSPCollectionStatement()[0]

	// fmt.Println("tspCollectionStatement1")
	// fmt.Println(tspCollectionStatement1.DataCache)

	//
	// handleTradeWisemen("ZYNE", "2.00")

	// postEmailTradeResultStore(TradeResultStore{})
	// processCheckIsTradeBought("SES")
	// twiWebscrape()

	// for i, v := range stockList {
	// 	// [fmt.Println(v.Symbol)
	// 	i++
	// }
	// fmt.Println(len(stockList))

	// i := 0
	// for i < 100 {
	// 	fmt.Println(stockList[i])
	// 	i++
	// }
	// stockList := selectStockWisemen()

	// insertDayReviewStatement(DayReviewStatement{DowValues: "DowValues", IsMarketClosed: "IsMarketClosed", IsDownDay: "IsDownDay", IsCashDay: "IsCashDay", PrevEndDow: "PrevEndDow", WisemenSymbolValues: "WisemenSymbolValues"})

	// truncateDayReviewStatement()
	// dayReviewStatement := selectDayReviewStatement()
	// fmt.Println("dayReviewStatement")
	// fmt.Println(dayReviewStatement)

	//
	// resetTimeOperations()
	// systemStartProcesses()
	// truncateWisemenSymbolHold()
	// processOverarchTopStock()
	// handleOverarchTopStock()
	// symbolList := []string{"AAPL", "SNE"}
	// response := queryMultiStockPull(symbolList)
	// stockList := parseStockSetQuery(response)

	// for i, v := range stockList {
	// 	fmt.Println(v.Symbol)
	// 	fmt.Println(v.Pchg)
	// 	i++
	// }

	// twiStockList := twiWebscrape()
	// for i, v := range twiStockList {
	// 	fmt.Println(v.Symbol)
	// 	fmt.Println(v.Pchg)
	// 	i++
	// }
	//
	// topStockPullStockList := topStockPull()
	// for i, v := range topStockPullStockList {
	// 	fmt.Println(v.Symbol)
	// 	fmt.Println(v.Pchg)
	// 	i++
	// }

	// handleEndOfDayDowScrape()
	// dataList := []string{"NVCN", "5.46"}
	// handleTradeWisemen(dataList[0], dataList[1])

	// symbol := "NVCN"
	// orderList := getAllOrders()
	// fmt.Println("len(orderList.ListOrders)")
	// fmt.Println(len(orderList.ListOrders))
	// order := Order{}
	// for i, v := range orderList.ListOrders {
	// 	if v.Symbol == symbol {
	// 		order = v
	// 		break
	// 	}
	// 	i++
	// }
	// queryCancelOrder(order.SVI, symbol)

	//
	// processTimelineStart()
	// handleEndOfDayDowScrape()
	// handleCalculateDownDay()
	//
	// dataList := []string{"NVCN", "5.46"}
	// handleTradeWisemen(dataList[0], dataList[1])
	//
	// symbol := "NVCN"
	// orderList := getAllOrders()
	// fmt.Println("len(orderList.ListOrders)")
	// fmt.Println(len(orderList.ListOrders))
	// order := Order{}
	// for i, v := range orderList.ListOrders {
	// 	if v.Symbol == symbol {
	// 		order = v
	// 		break
	// 	}
	// 	i++
	// }

	// fmt.Println(order.SVI)
	//

	// postCancellationBuyOrderEmail(symbol)

	// queryCancelOrder(order.SVI)

	// time.Sleep(time.Duration(10) * time.Second)
	// // //Begin process monitoring for buy fulfilled.
	// processCheckIsTradeBought(dataList[0])

	// timeDelimiter := "1101"
	// isBool := calculateIsTimeDelimiterMetSell(timeDelimiter)
	// fmt.Println("isBool")
	// fmt.Println(isBool)

	fmt.Scanln()
	fmt.Println("done")
}

//
// tradeResultStore := TradeResultStore{
// 	AlgorithmUsed:           "wisemen",
// 	Result:                  "awesome good",
// 	BoughtPrice:             "12.23",
// 	SellPrice:               "13.31",
// 	ChangeAmount:            "1.08",
// 	StockSymbol:             "ZYNE",
// 	Qty:                     "39",
// 	TimeTradeBuy:            "12 13",
// 	TimeTradeSell:           "13 14",
// 	HighestPricePointForDay: "13.45",
// 	TimeHighestPricePoint:   "12 10",
// 	LowestPricePointForDay:  "10.13",
// 	TimeLowestPricePoint:    "8 45",
// 	Dow1:                    "24,000",
// 	Dow2:                    "24,000",
// 	Dow3:                    "24,000",
// 	Dow4:                    "24,000",
// }
// postEmailTradeResultStore(tradeResultStore)

// handleEndOfDayDowScrape()
// handleCalculateDownDay()
// handleCalculateCashDay()
// overarchIsTradeDay()

// truncateEndOfDayDow()
// handleEndOfDayDowScrape()
// truncateEndOfDayDow()
// insertEndOfDayDow("26,000")
// endOfDayDow := selectEndOfDayDow()[0]
// fmt.Println("endOfDayDow")
// fmt.Println(endOfDayDow.EndOfDayDowValue)

// handleCalculateDownDay()
//
//
// truncateTradeResultStore()
// handleCalculateCashDay()
//
// insertTradeResultStore(TradeResultStore{Dow1: "23,000.02", Dow2: "23,000.02", Dow3: "23,000.02", Dow4: "29,000.02"})
// handleCalculateDownDay()

// isTradeDay := overarchIsTradeDay()
// fmt.Println("isTradeDay")
// fmt.Println(isTradeDay)

// insertTradeResultStore(TradeResultStore{StockSymbol: "BARK", Qty: "24"})
// trs := selectTradeResultStore("BARK")[0]

// postEmailTradeResultStore(tradeResultStore)
// fmt.Println(trs.StockSymbol)
// fmt.Println(trs.Qty)
// alteredTransactionHistory := calculateTransactionHistory(TransactionHistory{Symbol: "RAD"})
// fmt.Println(alteredTransactionHistory.HistoryValueList[0].Qty)
// fmt.Println(alteredTransactionHistory.HistoryValueList[0].Side)
// processCheckIsTradeBought("TGTX")
// value := formatQtyHolding("1.00")
// fmt.Println(value)
// processMonitorSell("GSX", "21.85", "1330")

// dataList := []string{"1", "2", "3"}
// dataList[0] = "0"
// for i, v := range dataList {
// 	fmt.Println(v)
// 	i++
// }

// systemStartProcesses()
// symbol := "MBOT"
// stringQty := "1"
// handleInsertInformationAtTrade(symbol, "limit", "buy", stringQty)

// holdingWisemen := HoldingWisemen{Symbol: "default"}
// holdingList := getAllHolding()
// for i, v := range holdingList.ListHolding {
// 	if v.Symbol == symbol {
// 		holdingWisemen = HoldingWisemen{Symbol: symbol, Price: v.Price, Qty: v.Qty, OrderStatus: "pending eval"}
// 	}
// 	i++
// }
// holdingWisemenReturned := calculateHoldingStatus(holdingWisemen)
// fmt.Println("holdingWisemenReturned")
// fmt.Println(holdingWisemenReturned)

// alteredTransactionHistory := calculateTransactionHistory(TransactionHistory{Symbol: symbol})
// handleInsertInformationAtTrade(symbol, "limit", "sell", alteredTransactionHistory.HistoryValueList[1].Qty)

//
// handleTradeWisemen("RAD", "16.14")
// listShortCalendarDay := selectShortCalendarDay()
// fmt.Println(listShortCalendarDay)
// fmt.Println("listShortCalendarDay")
// calculateShortDayAnalysis()
// shortDayAnalysis := selectShortDayAnalysis()[0]
// fmt.Println("shortDayAnalysis")
// fmt.Println(shortDayAnalysis)

// systemStartProcesses()
// metrics := selectMetricsWisemen()[0]
// // symbol := listStrings[0]
// // priceDrop := listStrings[1]
// timeDelimiter := metrics.SellTime

// fmt.Println("timeDelimiter")
// fmt.Println(timeDelimiter)
// fmt.Println("conditionNineteenMinute")
// fmt.Println(conditionNineteenMinute)
// fmt.Println("conditionNineteenHour")
// fmt.Println(conditionNineteenHour)

//
//
// insertShortDayAnalysis(ShortDayAnalysis{IsShortDay: "true"})
//ZYNE
// monitorSell()
// processMonitorSell("ZYNE", "6.06", "1130")

// processMonitorSell
// handleOverarchTopStock()
//
// limitPrice = strconv.FormatFloat(splitValue, 'f', -1, 64)
// fmt.Println(limitPrice)
// response := queryHistory()
// historyList := parseHistory(response)
// listHistoryValues := createListHistoryValuesForWisemen(historyList)

// historyValue := HistoryValue{Symbol: "RAD", Date: "2019 12 19", Side: "1", Qty: "1", Price: "10.1"}
// historyValue1 := HistoryValue{Symbol: "RAD", Date: "2019 12 19", Side: "2", Qty: "-1", Price: "10.3"}

// listHistoryValues = append(listHistoryValues, historyValue)
// listHistoryValues = append(listHistoryValues, historyValue1)

// fmt.Println(listHistoryValues)
// transactionHistory := TransactionHistory{Symbol: "RAD"}
// calculateTransactionHistory(transactionHistory)

// truncateInformationAtTrade()
// // informationAtTrade := InformationAtTrade{}
// handleInsertInformationAtTrade("RAD", "limit", "buy", "2.00")
// //
// handleInsertInformationAtTrade("RAD", "limit", "sell", "2.00")

// insertInformationAtTrade(informationAtTrade)
// insertInformationAtTrade(informationAtTrade1)

// truncateTradeResultStore()
// transactionHistory := TransactionHistory{Symbol: "RAD"}
// wrapUpWisemenOutcome(transactionHistory)

// dowList := selectDow()
// fmt.Println(dowList)

// listResults := handleHistoryDayListArbitration("VNCE")
// fmt.Println(listResults)

// processOverarchTopStock()

// processCheckIsTradeBought("QYNE")

// holdingList := getAllHolding()
// fmt.Println(holdingList)

//Work

// truncateInformationAtTrade()
// handleInsertInformationAtTrade("TGTX", "limit", "buy", "1.00")

// holdingWisemen := HoldingWisemen{Symbol: "TGTX", Price: "10.28", OrderStatus: "completedFull"}
// postNeoBuyOrderResponse(holdingWisemen)
// holdings := getAllHolding()
// fmt.Println(holdings)

// processMonitorSellMarket("TGTX")

// orderList := getAllOrders()
// fmt.Println("len(orderList.ListOrders)")
// fmt.Println(len(orderList.ListOrders))
// order := Order{}
// for i, v := range orderList.ListOrders {
// 	if v.Symbol == "TGTX" {
// 		order = v
// 		break
// 	}
// 	i++
// }
// // operatingCycle := cycleMapPool["monitorSell"]
// // cancelCycle(operatingCycle)

// fmt.Println(order.SVI)
// queryCancelOrder(order.SVI)

// fmt.Println("sell time")
// fmt.Println(metrics)
// metrics := "1211"
// metrics := selectMetricsWisemen()[0].SellTime

// boolTest := calculateIsTimeDelimiterMetSell(metrics)
// if boolTest {
// 	println("success")
// }

// if boolTest == false {
// 	println("false")
// }

// stock := Stock{Symbol: "TGTX", Last: "10.28"}
// insertStockWisemen(stock)
// select
// postNodeTSPFailureEmail()
// storeBalanceValue()
// metrics := selectMetricsWisemen()
// fmt.Println(metrics)
// twiWebscrape()
// dowValue := handleDowWebscrape()
// fmt.Println(dowValue)
// parseDowWebscrape()
// fmt.Println("Init")
// systemStartProcesses()
// handleOverarchTopStock()
// handleDowWebscrape()
// handleCalculateDownDay()

// handleOverarchTopStock()
// //handle down day calculation, later to be queried and checked by overarchIsTradeDay before purchases
// handleCalculateCashDay()
// handleCalculateDownDay()
// healthCheck()

// truncateInformationAtTrade()
// systemStartProcesses()
// // dowValue := handleDowWebscrape()

// startDayBalance := insertAccountBalance("7,000.00")

// storeBalanceValue()
// dowValue := "26,300"
// // insertDow(dowValue)
// // insertDow(dowValue)
// truncateDow()
// insertDow(dowValue)
// insertDow(dowValue)
// insertDow(dowValue)
// insertDow(dowValue)
// // truncateTradeResultStore()
// // handleInsertInformationAtTrade("HEPA", "limit", "buy", "1.00")
// // // // // handleInsertInformationAtTrade("HEPA", "limit", "sell", "1.00")
// truncateStockWisemen()
// insertStockWisemen(Stock{Symbol: "ZYNE", Last: "5.12"})
// insertStockWisemen(Stock{Symbol: "ZYNE", Last: "5.13"})
// insertStockWisemen(Stock{Symbol: "ZYNE", Last: "5.15"})
// // processMonitorSell("HEPA", "0.0", "1330")

// wrapUpWisemenOutcome(TransactionHistory{Symbol: "ZYNE"})
// truncateDownDayEvaluation()
// handleCalculateDownDay()

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
