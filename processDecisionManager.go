package main

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"
)

var checkIsMarketOpenMinute = 49
var checkIsMarketOpenHour = 7

var conditionOneMinute = 50
var conditionOneHour = 7

var conditionTwoMinute = 0
var conditionTwoHour = 8

//conditionMinuteHandleCalculateDownDay1 8:29 engage
var conditionMinuteHandleCalculateDownDay1 = 29
var conditionHourHandleCalculateDownDay1 = 8

var conditionFourMinute = 0
var conditionFourHour = 9

var conditionFiveMinute = 30
var conditionFiveHour = 9

var conditionSixMinute = 9
var conditionSixHour = 44

var conditionNineteenMinute = 30
var conditionNineteenHour = 13

var checkIsMarketOpenBool = true
var boolOperate1 = true
var boolOperate2 = true
var boolOperateHandleCalculateDownDay1 = true
var boolOperate4 = true
var boolOperate5 = true
var boolOperate6 = true
var boolOperate19 = true

var initialStockQueryPerformed = false
var initialWisemenStockQueryPerformed = false
var initialWhaleStockQueryPerformed = false

func processTimelineStart() {
	cycleMapPool = map[string]*Cycle{}
	createCycle(5, 1000000000, handleTimelineConditionalTriggers, "handleTimelineConditionalTriggers")
	operatingCycle := cycleMapPool["handleTimelineConditionalTriggers"]
	go startCycle(operatingCycle)
}

func processMonitorSell(symbol string, dropLoss string, timeToSell string) {
	//Metrics from sell
	// create cycle for
	cycleMapPool = map[string]*Cycle{}
	createCycle(10, 100000000, monitorSell, "monitorSell", []string{symbol, dropLoss, timeToSell})
	operatingCycle := cycleMapPool["monitorSell"]
	go startCycle(operatingCycle)
}

func processWisemenQueryStockSet() {
	createCycle(3, 100000000, handleWisemenQueryStockList, "handleWisemenQueryStockList")
	operatingCycle := cycleMapPool["handleWisemenQueryStockList"]
	go startCycle(operatingCycle)
}

func processWhaleQueryStockSet() {
	// if initialWhaleStockQueryPerformed == true {
	// 	go startCycle(cycleMapPool["handleWhaleQueryStockList"])
	// }
	// if initialWhaleStockQueryPerformed == false {
	createCycle(300, 1000000000000, handleWhaleQueryStockList, "handleWhaleQueryStockList")
	operatingCycle := cycleMapPool["handleWhaleQueryStockList"]
	go startCycle(operatingCycle)
	// initialWhaleStockQueryPerformed = true
	// }
}

// func intiateMonitorTradeWisemon() {
// metrics := selectMetricsWisemen()
// Select metrics make use of it, continue.
// fmt.Println(metrics)
//Delay before monitor cycle
// time.Sleep(time.Duration(10) * time.Second)
// fmt.Println("hit awesome")
//single query is holding of symbol

//if not delay, do iterate until true
// for
// indexCheck := 1
// for indexCheck < 100000 {
// 	queryIsTradeCompleted()

// 	time.Sleep(time.Duration(3) * time.Second)
// 	indexCheck++
// }
// holdingStatus := calculateHoldingStatus()

//evaluation if order is closed

// }

// func processCheckIsBuyPeformed() {
// 	createCycle(3, 10000000000000, handleWisemenQueryStockList, "handleWisemenQueryStockList")
// 	operatingCycle := cycleMapPool["handleWisemenQueryStockList"]
// 	go startCycle(operatingCycle)
// }

func processTSPRefresh() {
	// go handleTSPRefresh()
}

func processDowWebscrape() {
	go handleDowWebscrape()
}

func processOverarchTopStock() {
	go handleOverarchTopStock()
	// go handleTwiWebscrape()
}
func processCheckIsTradeBought(symbol string) {
	createCycle(10, 100000, handleCheckIsTradeBought, "handleCheckIsTradeBought", []string{symbol})
	operatingCycle := cycleMapPool["handleCheckIsTradeBought"]
	go startCycle(operatingCycle)
	initialWhaleStockQueryPerformed = true
}

func checkIsDowStore(currentHour int, currentMinute int) {
	isStore := true
	//point1
	if isStore {
		if currentHour == 7 {
			if currentMinute == 49 {
				dowValue := handleDowWebscrape()
				insertDow(dowValue)
				isStore = false
			}
		}
	}
	if isStore == false {
		if currentHour == 7 {
			if currentMinute == 50 {
				isStore = true
			}
		}
	}
	//point2
	if isStore {
		if currentHour == 8 {
			if currentMinute == 30 {
				dowValue := handleDowWebscrape()
				insertDow(dowValue)
				isStore = false
			}
		}
	}
	if isStore == false {
		if currentHour == 8 {
			if currentMinute == 31 {
				isStore = true
			}
		}
	}
	//point3
	if isStore {
		if currentHour == 9 {
			if currentMinute == 45 {
				dowValue := handleDowWebscrape()
				insertDow(dowValue)
				isStore = false
			}
		}
	}
	if isStore == false {
		if currentHour == 9 {
			if currentMinute == 46 {
				isStore = true
			}
		}
	}

	//point4
	if isStore {
		if currentHour == 12 {
			if currentMinute == 0 {
				dowValue := handleDowWebscrape()
				insertDow(dowValue)
				isStore = false
			}
		}
	}
	if isStore == false {
		if currentHour == 12 {
			if currentMinute == 1 {
				isStore = true
			}
		}
	}
}

func handleTimelineConditionalTriggers(params ...interface{}) {
	currentTime := time.Now()
	fmt.Println(currentTime.Hour())
	fmt.Println(currentTime.Minute())
	fmt.Println(currentTime.Second())
	fmt.Println(currentTime.Date())

	//Conditional operate
	if currentTime.Minute() == checkIsMarketOpenMinute && currentTime.Hour() == checkIsMarketOpenHour && checkIsMarketOpenBool {
		systemStartProcesses()
		checKIsBrokerageResponding()
		//Wisemen algorithm same day calculation
		// handleDayRotation()
		conditionMarketClosed := selectMarketOpenAnalysis()
		if conditionMarketClosed[0].IsMarketClosed == "true" {
			setTimelineOperationsFalse()
		}
		checkIsMarketOpenBool = false
		boolOperate19 = true
	}

	//dow store periodic
	checkIsDowStore(currentTime.Hour(), currentTime.Minute())

	//Initiate monitoring pool query on cycle
	//Periodic TSP refresh
	if currentTime.Minute() == conditionOneMinute && currentTime.Hour() == conditionOneHour && boolOperate1 {
		fmt.Println("hit1")
		boolOperate1 = false
		handleOverarchTopStock()
		processWisemenQueryStockSet()
		processWhaleQueryStockSet()
	}
	if currentTime.Minute() == conditionTwoMinute && currentTime.Hour() == conditionTwoHour && boolOperate2 {
		fmt.Println("hit2")
		boolOperate2 = false
		handleOverarchTopStock()
	}
	if currentTime.Minute() == conditionMinuteHandleCalculateDownDay1 && currentTime.Hour() == conditionHourHandleCalculateDownDay1 && boolOperateHandleCalculateDownDay1 {
		fmt.Println("handleCalculateDownDay1")
		boolOperateHandleCalculateDownDay1 = false
		handleOverarchTopStock()
		//handle down day calculation, later to be queried and checked by overarchIsTradeDay before purchases
		handleCalculateCashDay()
		handleCalculateDownDay()
	}
	if currentTime.Minute() == conditionFourMinute && currentTime.Hour() == conditionFourHour && boolOperate4 {
		fmt.Println("hit4")
		boolOperate4 = false
		handleOverarchTopStock()
	}
	if currentTime.Minute() == conditionFiveMinute && currentTime.Hour() == conditionFiveHour && boolOperate5 {
		fmt.Println("hit5")
		boolOperate5 = false
		handleOverarchTopStock()
	}
	if currentTime.Minute() == conditionSixMinute && currentTime.Hour() == conditionSixHour && boolOperate6 {
		fmt.Println("hit6")
		boolOperate6 = false
		handleOverarchTopStock()
	}
	if currentTime.Minute() == conditionNineteenMinute && currentTime.Hour() == conditionNineteenHour && boolOperate19 {
		fmt.Println("hit19")
		boolOperate19 = false
		//handle ifMarketClosed with information about TRS.
		isMarketClosed := selectMarketOpenAnalysis()[0].IsMarketClosed
		if isMarketClosed == "\"true\"" {
			//handle TRS where market closed.
			createTradeResultStoreMarketClosed()
		}
		//At some point in this hour reset the pools, and reset the timeline.
		//clear cyclepool and reset timeline process
		resetCyclePools()
		handleDayReset()
	}
}
func createTradeResultStoreMarketClosed() {
	tradeResultStore := TradeResultStore{Result: "Market closed"}
	insertTradeResultStore(tradeResultStore)
}

func resetCyclePools() {
	for i, v := range cycleMapPool {
		cancelCycle(v)
		fmt.Println(i)
	}
	cycleMapPool = make(map[string]*Cycle)
	processTimelineStart()
}
func handleCheckIsTradeBought(params ...interface{}) {
	//Support for non-static time delimiter.
	// stringTimeDelimiterHour := ""
	listVal := reflect.ValueOf(params[0])
	var listSymbolsInterface interface{} = listVal.Index(0).Interface()
	listSymbols := listSymbolsInterface.([]string)
	symbol := listSymbols[0]
	//Is holding for symbol present
	holdingWisemen := HoldingWisemen{}
	//if holding matches symbol seeking
	holdingList := getAllHolding()
	for i, v := range holdingList.ListHolding {
		if v.Symbol == symbol {
			holdingWisemen = HoldingWisemen{Symbol: symbol, Price: v.Price, Qty: v.Qty, OrderStatus: "pending eval"}
		}
		i++
	}
	holdingWisemen = calculateHoldingStatus(holdingWisemen)
	if holdingWisemen.OrderStatus == "order not placed" {
		cancelCycle(cycleMapPool["handleCheckIsTradeBought"])
		postNeoBuyOrderResponse(holdingWisemen)
	}
	//Handle conditions for holding incomplete
	if holdingWisemen.OrderStatus == "completedFull" {
		fmt.Println("completedFull hit")
		//End cycle for monitoring
		handleInsertInformationAtTrade(symbol)
		cancelCycle(cycleMapPool["handleCheckIsTradeBought"])
		response := postNeoBuyOrderResponse(holdingWisemen)
		fmt.Println(response)
	}

	//handle time delimiter where no buy is completed.
	//
	//

	// if holdingWisemen.OrderStatus == "partial" {
	// 	fmt.Println("partial hit")
	// 	//in the impartial case it will iterate a global check variable,
	// 	//upon global variable reaching a delimter count, which represents monitor cycle time intervals.
	// 	//upon delmiter met, cancel and post to neo with holding "partial" status.
	// 	if getIntervalTradeMonitorDelimiter() == 4 {
	// 		cancelCycle(cycleMapPool["handleCheckIsTradeBought"])
	// 		orderContainer := getAllOrders()
	// 		orderForSymbol := Order{Symbol: "default"}
	// 		for i, v := range orderContainer.ListOrders {
	// 			if v.Symbol == symbol {
	// 				fmt.Println("hit ")
	// 				orderForSymbol = v
	// 				break
	// 			}
	// 			i++
	// 		}
	// 		if orderForSymbol.Symbol == "default" {
	// 			// queryCancelOrder()
	// 			// holdingWisemen.OrderStatus
	// 			postNeoBuyOrderResponse(holdingWisemen)
	// 		}
	// 		if orderForSymbol.Symbol != "default" {
	// 			queryCancelOrder(orderForSymbol.SVI)
	// 			postNeoBuyOrderResponse(holdingWisemen)
	// 		}
	// 	}
	// 	iterateIntervalTradeMonitorDelimiter()
	// }
}

// func handleOverarchTopStock(params ...interface{}) {
func handleOverarchTopStock() {
	twiStockList := twiWebscrape()
	// twiStockList := []Stock{}
	//High process for wisemen and whale
	highTransferanceProcess(twiStockList)
	//Low process for whale
	lowTransferanceProcess(twiStockList)
}
func highTransferanceProcess(twiStockList []Stock) {
	//TSP
	topStockPullStockList := topStockPull()

	for i, v := range topStockPullStockList {
		fmt.Println(v.Symbol)
		fmt.Println(v.Pchg)
		i++
	}

	listTempDuplicantFiltered := []Stock{}
	positiveTwiTSPList := []Stock{}
	positiveTwiStockList := []Stock{}
	isDuplicate := false
	isDuplicateTemp := false

	// //Handle positive values only in twiTSPList
	for i, v := range twiStockList {
		if v.IsCurrentPriceHigherThanPreviousClose == "true" {
			positiveTwiStockList = append(positiveTwiStockList, v)
		}
		i++
	}

	// //No duplicates in lists
	for indexTsp, tsp := range topStockPullStockList {
		isDuplicate = false
		for indexTwi, twi := range positiveTwiStockList {
			if twi.Symbol == tsp.Symbol {
				isDuplicate = true
				break
			}
			indexTwi++
		}
		if isDuplicate == false {
			positiveTwiTSPList = append(positiveTwiTSPList, tsp)
		}
		indexTsp++
	}
	// //Query temp
	tempSymbolHold := selectTempSymbolHoldHigh()
	// //Find duplicants in temp and appendedList
	for indexTwiTSP, twiTSP := range positiveTwiTSPList {
		isDuplicateTemp = false
		for indexTemp, temp := range tempSymbolHold {
			if twiTSP.Symbol == temp {
				isDuplicateTemp = true
				break
			}
			indexTemp++
		}
		if isDuplicateTemp == false {
			listTempDuplicantFiltered = append(listTempDuplicantFiltered, twiTSP)
		}
		indexTwiTSP++
	}

	i := 0
	topStockList := []Stock{}
	for i < 3 {
		// 	// remove highest index 3 times, to get top stocks.
		// 	//Pop top stock each iteration
		highestStockIndex := 0
		for indexTempDuplicantFiltered, tempDuplicantFiltered := range listTempDuplicantFiltered {
			if indexTempDuplicantFiltered == 0 {
				highestStockIndex = indexTempDuplicantFiltered
				continue
			}

			floatHighest := 0.0
			floatCurrent := 0.0
			if s, err := strconv.ParseFloat(listTempDuplicantFiltered[highestStockIndex].Pchg, 64); err == nil {
				floatHighest = s
			}
			if s1, err := strconv.ParseFloat(tempDuplicantFiltered.Pchg, 64); err == nil {
				floatCurrent = s1
			}

			if floatCurrent > floatHighest {
				highestStockIndex = indexTempDuplicantFiltered
			}
		}
		topStockList = append(topStockList, listTempDuplicantFiltered[highestStockIndex])
		if i < 2 {
			listTempDuplicantFiltered = removeElementString(listTempDuplicantFiltered, listTempDuplicantFiltered[highestStockIndex].Symbol)
		}
		i++
	}
	fmt.Println("topStockList")
	fmt.Println(topStockList)
	for i, v := range topStockList {
		insertTempSymbolHoldHigh(v.Symbol)
		i++
	}
	// //fill algorithm symbol holds
	handleWisemenFillHold()
	handleWhaleFillHoldHigh()
}

func lowTransferanceProcess(twiStockList []Stock) {
	//TSP
	//Support for find lowest TSP
	// topStockPullStockList := topStockPull()
	//[]Stock{Stock{Symbol: "test1", Pchg: "3.00"}, Stock{Symbol: "test2", Pchg: "6.00"}, Stock{Symbol: "test3", Pchg: "5.00"}}

	//Twi
	// twiStockList := twiWebscrape()
	// twiTSPList := twiStockList
	// listTempDuplicantFiltered := []Stock{}
	// isDuplicate := false
	// isDuplicateTemp := false

	// negativeTwiTSPList := []Stock{}
	negativeTwiStockList := []Stock{}

	//No duplicates in lists

	// //Handle positive values only in twiTSPList
	for i, v := range twiStockList {
		fmt.Println("v.Symbol")
		fmt.Println(v.Symbol)
		fmt.Println("v.Pchg")
		fmt.Println(v.Pchg)
		fmt.Println("v.Last")
		fmt.Println(v.Last)
		fmt.Println("v.Pcls")
		fmt.Println(v.Pcls)
		fmt.Println("v.IsCurrentPriceHigherThanPreviousClose")
		fmt.Println(v.IsCurrentPriceHigherThanPreviousClose)
		if v.IsCurrentPriceHigherThanPreviousClose == "false" {
			negativeTwiStockList = append(negativeTwiStockList, v)
		}
		i++
	}

	fmt.Println("negativeTwiStockList")
	fmt.Println(negativeTwiStockList)

	for i, v := range negativeTwiStockList {
		insertTempSymbolHoldLow(v.Symbol)
		i++
	}

	// insert temp
	//fill hold
	handleWhaleFillHoldLow()
}

func getIntervalTradeMonitorDelimiter() int {
	return intervalTradeMonitorDelimiter
}
func iterateIntervalTradeMonitorDelimiter() {
	intervalTradeMonitorDelimiter++
}

func checkWhaleDelimiterMetHigh() bool {
	isWhaleDelimiterMet := false
	symbolList := selectWhaleSymbolHoldHigh()
	if len(symbolList) >= 200 {
		isWhaleDelimiterMet = true
	}
	return isWhaleDelimiterMet
}

func checkWhaleDelimiterMetLow() bool {
	isWhaleDelimiterMet := false
	symbolList := selectWhaleSymbolHoldLow()
	if len(symbolList) >= 200 {
		isWhaleDelimiterMet = true
	}
	return isWhaleDelimiterMet
}

func topStockPull() []Stock {
	var queryResponse = queryTSP()
	stockList := parseTopStockQuery(queryResponse)
	filteredStockList := []Stock{}
	for i, v := range stockList {
		if strings.Contains(v.Symbol, ".") == false {
			filteredStockList = append(filteredStockList, v)
		}
		i++
	}
	return filteredStockList
}

func healthCheck() {
	isNeoResponse := "false"
	//post to neo
	response := postNeoHealthCheck()
	//If response from neo
	if response != "error received" {
		// fmt.Println(response)
		isNeoResponse = "true"
	}
	postHealthCheckNode(isNeoResponse)
	//res from check
	//nodemail
}
func purchaseUpdateSystem() {
}
func getLatestTradeResultStore() TradeResultStore {
	tradeResultStoreList := selectTradeResultStore("wisemen")
	//get latest tradeResultStore
	tradeResultStore := tradeResultStoreList[len(tradeResultStoreList)-1]
	return tradeResultStore
}

// func handleFillHolds(params ...interface{}) {
func handleWisemenFillHold() {
	tempSymbolHoldList := selectTempSymbolHoldHigh()
	// whaleDelimiterMet := checkWhaleDelimiterMet()
	for i, tempSymbol := range tempSymbolHoldList {
		//condition meet if whale symbol or wisemen symbol already exists.
		isSymbolExistsInWisemen := false
		//select from wisemen
		wisemenSymbolList := selectWisemenSymbolHold()

		//iterrate set isSymbolExistsInWiseMen
		for indexWisemenSymbol, wisemenSymbol := range wisemenSymbolList {
			if wisemenSymbol == tempSymbol {
				isSymbolExistsInWisemen = true
			}
			indexWisemenSymbol++
		}
		if isSymbolExistsInWisemen == false {
			insertWisemenSymbolHold(tempSymbol)
		}
		i++
	}
}
func handleWhaleFillHoldHigh() {
	tempSymbolHoldList := selectTempSymbolHoldHigh()
	whaleDelimiterMet := checkWhaleDelimiterMetHigh()
	for i, tempSymbol := range tempSymbolHoldList {
		//condition meet if whale symbol or wisemen symbol already exists.
		isSymbolExistsInWhale := false
		//select from whale
		whaleSymbolList := selectWhaleSymbolHoldHigh()
		for indexWhaleSymbol, whaleSymbol := range whaleSymbolList {
			if whaleSymbol == tempSymbol {
				isSymbolExistsInWhale = true
			}
			indexWhaleSymbol++
		}
		if isSymbolExistsInWhale == false {
			//check process for whale delimiter
			if whaleDelimiterMet == false {
				insertWhaleSymbolHoldHigh(tempSymbol)
			}
		}
		i++
	}
}
func handleWhaleFillHoldLow() {
	tempSymbolHoldList := selectTempSymbolHoldLow()
	whaleDelimiterMet := checkWhaleDelimiterMetLow()
	for i, tempSymbol := range tempSymbolHoldList {
		//condition meet if whale symbol or wisemen symbol already exists.
		isSymbolExistsInWhale := false
		//select from whale
		whaleSymbolList := selectWhaleSymbolHoldLow()
		for indexWhaleSymbol, whaleSymbol := range whaleSymbolList {
			if whaleSymbol == tempSymbol {
				isSymbolExistsInWhale = true
			}
			indexWhaleSymbol++
		}
		if isSymbolExistsInWhale == false {
			//check process for whale delimiter
			if whaleDelimiterMet == false {
				insertWhaleSymbolHoldLow(tempSymbol)
			}
		}
		i++
	}
}

func handleWisemenQueryStockList(params ...interface{}) {
	fmt.Println("hit handleWisemenQueryStockList")
	//Query monitor_symbol
	symbolList := selectWisemenSymbolHold()
	fmt.Println("symbolList")
	fmt.Println(symbolList)

	//Parse format errors in symbols
	formattedSymbolList := []string{}
	for i, v := range symbolList {
		if strings.Contains(v, ".") {
			continue
		}
		formattedSymbolList = append(formattedSymbolList, v)
		i++
	}
	//Form query to Node -> Brokerage
	queryResponse := queryMultiStockPull(formattedSymbolList)
	stockList := parseStockSetQuery(queryResponse)

	//Store Stocks in DB
	for i, v := range stockList {
		insertStockWisemen(v)
		i++
	}
}

func handleWhaleQueryStockList(params ...interface{}) {
	//Branch system for high, and for low
	//for high
	//
	symbolList := selectWhaleSymbolHoldHigh()
	//Parse format errors in symbols
	formattedSymbolList := []string{}
	for i, v := range symbolList {
		if strings.Contains(v, ".") {
			continue
		}
		formattedSymbolList = append(formattedSymbolList, v)
		i++
	}
	//Given a symbolList, query brokerage, stock list response.
	queryResponse := queryMultiStockPull(formattedSymbolList)
	stockListBrokerage := parseStockSetQuery(queryResponse)

	//so monitorSymbol given
	//for each symbol query whale
	listOfDatabaseStockListResponse := []DatabaseStockListResponse{}
	for index, symbol := range formattedSymbolList {
		stockList := selectStockWhaleHigh(symbol)
		if len(stockList) == 0 {
			continue
		}
		databaseStockListResponse := DatabaseStockListResponse{StockList: stockList}
		listOfDatabaseStockListResponse = append(listOfDatabaseStockListResponse, databaseStockListResponse)
		index++
	}
	isPresentInDB := false
	isDaySame := false
	for indexStockBrokerage, stockBrokerage := range stockListBrokerage {
		isPresentInDB = false
		isDaySame = false

		for indexDatabaseStockListResponse, databaseStockListResponse := range listOfDatabaseStockListResponse {

			fmt.Println(databaseStockListResponse.StockList)
			lastIndexedStock := databaseStockListResponse.StockList[(len(databaseStockListResponse.StockList) - 1)]
			if lastIndexedStock.Symbol == stockBrokerage.Symbol {
				isPresentInDB = true
			}
			if isPresentInDB {
				isStocksComparedSameTimeStamps := compareTimeStamps(stockBrokerage, lastIndexedStock)
				//if present and day is the same
				if isStocksComparedSameTimeStamps {
					insertStockWhaleHigh(stockBrokerage)
					isDaySame = true
					break
				}
			}
			indexDatabaseStockListResponse++
		}
		//if symbol in DB, but different days
		if isDaySame == false {
			appendedStock := processAppendDayOfWeekToStock(stockBrokerage)
			//When stock is inserted it does not take into account the day.
			insertStockWhaleHigh(appendedStock)
		}
		indexStockBrokerage++
	}

	//for low
	//
	symbolListLow := selectWhaleSymbolHoldLow()
	//Parse format errors in symbols
	formattedSymbolListLow := []string{}
	for i, v := range symbolListLow {
		if strings.Contains(v, ".") {
			continue
		}
		formattedSymbolListLow = append(formattedSymbolListLow, v)
		i++
	}
	//Given a symbolList, query brokerage, stock list response.
	queryResponseLow := queryMultiStockPull(formattedSymbolListLow)
	stockListBrokerageLow := parseStockSetQuery(queryResponseLow)

	//so monitorSymbol given
	//for each symbol query whale
	listOfDatabaseStockListResponseLow := []DatabaseStockListResponse{}
	for index, symbol := range formattedSymbolListLow {
		stockList := selectStockWhaleLow(symbol)
		if len(stockList) == 0 {
			continue
		}
		databaseStockListResponseLow := DatabaseStockListResponse{StockList: stockList}
		listOfDatabaseStockListResponseLow = append(listOfDatabaseStockListResponseLow, databaseStockListResponseLow)
		index++
	}
	isPresentInDBLow := false
	isDaySameLow := false
	for indexStockBrokerage, stockBrokerage := range stockListBrokerageLow {
		isPresentInDBLow = false
		isDaySameLow = false

		for indexDatabaseStockListResponse, databaseStockListResponse := range listOfDatabaseStockListResponseLow {

			fmt.Println(databaseStockListResponse.StockList)
			lastIndexedStock := databaseStockListResponse.StockList[(len(databaseStockListResponse.StockList) - 1)]
			if lastIndexedStock.Symbol == stockBrokerage.Symbol {
				isPresentInDBLow = true
			}
			if isPresentInDBLow {
				isStocksComparedSameTimeStamps := compareTimeStamps(stockBrokerage, lastIndexedStock)
				//if present and day is the same
				if isStocksComparedSameTimeStamps {
					insertStockWhaleLow(stockBrokerage)
					isDaySameLow = true
					break
				}
			}
			indexDatabaseStockListResponse++
		}
		//if symbol in DB, but different days
		if isDaySameLow == false {
			appendedStock := processAppendDayOfWeekToStock(stockBrokerage)
			//When stock is inserted it does not take into account the day.
			insertStockWhaleLow(appendedStock)
		}
		indexStockBrokerage++
	}
}

func processDetectStockWhaleWhereDayIsAtIndex(stockList []Stock) {
	for index, stock := range stockList {

		fmt.Println(stock.CreatedAt)
		// isContained := strings.Contains(stock.CreatedAt, "?")
		// if isContained {
		// 	fmt.Println(stock.CreatedAt)
		// }
		index++
	}
}

func compareTimeStamps(stockBrokerage Stock, stockDB Stock) bool {
	isTimeStampMatch := true
	//Split for day
	// 2019-09-19 17:47:10.944343-06
	fmt.Println("stock1.CreatedAt")
	fmt.Println(stockBrokerage.CreatedAt)
	fmt.Println("stock2.CreatedAt")
	fmt.Println(stockDB.CreatedAt)

	stock1Split1 := strings.Split(stockBrokerage.CreatedAt, " ")[1]
	// stock1Split2 := strings.Split(stock1Split1, "-")[1]
	day1 := stock1Split1 //strings.Split(stock1Split1, " ")[1]

	fmt.Println("day1")
	fmt.Println(day1)

	stock2Split1 := strings.Split(stockDB.CreatedAt, ":")[0]
	stock2Split2 := strings.Split(stock2Split1, "-")[2]
	stock2Split3 := strings.Split(stock2Split2, "T")[0]
	day2 := strings.Split(stock2Split3, " ")[0]

	fmt.Println("day2")
	fmt.Println(day2)
	//compare day,
	if day1 != day2 {
		isTimeStampMatch = false
	}
	return isTimeStampMatch
}

func processAppendDayOfWeekToStock(stock Stock) Stock {
	//if adding stock, and different day identified, add stock and append first stock entry
	//with weekday
	instanceStock := stock
	currentTime := time.Now()
	instanceStock.Vl = instanceStock.Vl + " ?" + currentTime.Weekday().String()
	// fmt.Println(instanceStock.Vl)
	return instanceStock
}

func handleDowWebscrape() string {
	response := queryWebscrape()
	currentDowValue := parseDowWebscrape(response)
	return currentDowValue
}

func twiWebscrape() []Stock {

	// func queryWebscrapeTwi() string {
	// 	json := `{
	// 		"request_type": "webscrapeTwi"
	// 		}`
	// 	url := "http://localhost:3000/api/brokerage"
	// 	response := post(url, json)
	// 	return response
	// }
	//stop twi server
	// response := queryStopTwi()
	// fmt.Println(response)
	// //start twi server
	// response1 := queryStartTwi()
	// fmt.Println(response1)

	response2 := queryWebscrapeTwi()
	// fmt.Println(response2)
	symbolList := parseTwiWebscrape(response2)
	responseSymbolList := queryMultiStockPull(symbolList)
	stockList := parseStockSetQuery(responseSymbolList)
	// stockList := []Stock{}
	for i, v := range stockList {
		fmt.Println(v)
		i++
	}
	return stockList
}

// func handleEndOfDayAnalyticsOperations() {
// 	day := getDayOfWeek()
// 	//insert into table conditional
// }

func systemStartProcesses() {
	//reset existing twi server
	queryStopTwi()
	queryStartTwi()
	//reset dow
	// dropDow()
	// createDow()
	truncateDow()
	//
	resetTempSymbolHold()
	//
	resetStockWisemenSymbolHold()
	//
	resetStockWisemen()
	//
	resetAltIntervalBuyWisemen()
}

func resetTempSymbolHold() {
	truncateTempSymbolHoldHigh()
}
func resetStockWisemenSymbolHold() {
	truncateWisemenSymbolHold()
}
func resetStockWisemen() {
	// dropStockWisemen()
	// createStockWisemen()
	truncateStockWisemen()
}

func resetAltIntervalBuyWisemen() {
	// dropAltIntervalBuyWisemen()
	// createAltIntervalBuyWisemen()
	truncateAltIntervalBuyWisemen()
}

func handleDayReset() {
	boolOperate1 = true
	boolOperate2 = true
	// boolOperate3 = true
	boolOperate4 = true
	boolOperate5 = true
	boolOperate6 = true
	// boolOperate7 = true
	// boolOperate8 = true
	// boolOperate9 = true
	// boolOperate10 = true
	// boolOperate11 = true
	// boolOperate12 = true
	// boolOperate13 = true
	// boolOperate14 = true
	// boolOperate15 = true
	// boolOperate16 = true
	// boolOperate17 = true
	// boolOperate18 = true
	checkIsMarketOpenBool = true
}

func setTimelineOperationsFalse() {
	boolOperate1 = false
	boolOperate2 = false
	// boolOperate3 = false
	boolOperate4 = false
	boolOperate5 = false
	boolOperate6 = false
	// boolOperate7 = false
	// boolOperate8 = false
	// boolOperate9 = false
	// boolOperate10 = false
	// boolOperate11 = false
	// boolOperate12 = false
	// boolOperate13 = false
	// boolOperate14 = false
	// boolOperate15 = false
	// boolOperate16 = false
	// boolOperate17 = false
	// boolOperate18 = false
	boolOperateHandleCalculateDownDay1 = false
}

func wrapUpWisemenOutcome(transactionHistory TransactionHistory) {
	//Post wisemen outcome.
	//
	metrics := selectMetricsWisemen()[0]
	alteredTransactionHistory := calculateTransactionHistory(transactionHistory)
	//get insertInformationAtTrade for buy and sell
	listMatchingSymbolInformationAtTrade := handleInformationAtTradeDayListArbitration(alteredTransactionHistory.Symbol)

	fmt.Println("alteredTransactionHistory")
	fmt.Println(alteredTransactionHistory)
	fmt.Println("listMatchingSymbolInformationAtTrade")
	fmt.Println(listMatchingSymbolInformationAtTrade)

	//Support for handle multiple InformationAtTrade during day...
	//typically only would be two... but in case of two of more...
	//Support for more than 2 trades
	//if no trade occured...

	dowList := selectDow()
	if len(listMatchingSymbolInformationAtTrade) == 0 {
		//no trade occured handle TradeResultStore
		if len(dowList) == 0 {
			tradeResultStore := TradeResultStore{
				AlgorithmUsed: "wisemen",
				Result:        "No trade",
			}
			insertTradeResultStore(tradeResultStore)
		}
		if len(dowList) == 4 {
			tradeResultStore := TradeResultStore{
				AlgorithmUsed: "wisemen",
				Result:        "No trade",
				Dow1:          dowList[0].CurrentDowValue,
				Dow2:          dowList[1].CurrentDowValue,
				Dow3:          dowList[2].CurrentDowValue,
				Dow4:          dowList[3].CurrentDowValue,
			}
			insertTradeResultStore(tradeResultStore)
		}
	}

	//if buy and sell exists...InformationAtTrade
	if len(listMatchingSymbolInformationAtTrade) == 2 {
		buyHistoryValuePrice := alteredTransactionHistory.HistoryValueList[0].Price
		sellHistoryValuePrice := alteredTransactionHistory.HistoryValueList[1].Price

		floatBuyHistoryValuePrice := 0.0
		floatSellHistoryValuePrice := 0.0
		if s, err := strconv.ParseFloat(buyHistoryValuePrice, 64); err == nil {
			floatBuyHistoryValuePrice = s
		}
		if s1, err := strconv.ParseFloat(sellHistoryValuePrice, 64); err == nil {
			floatSellHistoryValuePrice = s1
		}
		//calculate result
		//if buy and sell, and if changeAmount meet delimiter,
		changeAmount := floatSellHistoryValuePrice - floatBuyHistoryValuePrice
		stringChangeAmount := fmt.Sprintf("%f", 123.456)

		fmt.Println("changeAmount")
		fmt.Println(changeAmount)
		//handle on metrics delimiter...
		metricsPriceHighPchg := metrics.PriceHighPchg
		floatMetricsPriceHighPchg := 0.0
		if s2, err := strconv.ParseFloat(metricsPriceHighPchg, 64); err == nil {
			floatMetricsPriceHighPchg = s2
		}
		optimal := floatBuyHistoryValuePrice + (floatBuyHistoryValuePrice * floatMetricsPriceHighPchg)

		result := "negative"
		//if sell was less than optimal
		// isAlgorithmProfitable := false
		if floatSellHistoryValuePrice >= optimal {
			result = "positive"
		}
		fmt.Println(stringChangeAmount)
		fmt.Println(result)

		fmt.Println("listMatchingSymbolInformationAtTrade")
		fmt.Println(listMatchingSymbolInformationAtTrade[0])

		boughtTime := listMatchingSymbolInformationAtTrade[0].Hour + " " + listMatchingSymbolInformationAtTrade[0].Minute
		sellTime := listMatchingSymbolInformationAtTrade[1].Hour + " " + listMatchingSymbolInformationAtTrade[1].Minute

		if len(dowList) == 4 {
			tradeResultStore := TradeResultStore{
				AlgorithmUsed: "wisemen",
				Result:        result,
				ChangeAmount:  stringChangeAmount,
				StockSymbol:   alteredTransactionHistory.Symbol,
				TimeTradeBuy:  boughtTime,
				TimeTradeSell: sellTime,
				Dow1:          dowList[0].CurrentDowValue,
				Dow2:          dowList[1].CurrentDowValue,
				Dow3:          dowList[2].CurrentDowValue,
				Dow4:          dowList[3].CurrentDowValue,
			}
			fmt.Println(tradeResultStore)
			insertTradeResultStore(tradeResultStore)
		}
		if len(dowList) != 4 {
			tradeResultStore := TradeResultStore{
				AlgorithmUsed: "wisemen",
				Result:        result,
				ChangeAmount:  stringChangeAmount,
				StockSymbol:   alteredTransactionHistory.Symbol,
				TimeTradeBuy:  boughtTime,
				TimeTradeSell: sellTime,
			}
			fmt.Println(tradeResultStore)
			insertTradeResultStore(tradeResultStore)
		}
	}
}

func wrapUpWisemenOutcomeNoBuy(transactionHistory TransactionHistory) {
	downDayEvaluationList := selectDownDayEvaluation()
	cashDayEvaluationList := selectCashDayEvaluation()

	downDayEval := downDayEvaluationList[len(downDayEvaluationList)-1]
	dow := 0.0
	dowPrevious := 0.0
	dowPchg := 0.0

	if s, err := strconv.ParseFloat(downDayEval.Dow, 64); err == nil {
		dow = s
	}
	if s1, err := strconv.ParseFloat(downDayEval.PreviousDow, 64); err == nil {
		dowPrevious = s1
	}
	if s2, err := strconv.ParseFloat(downDayEval.GreatestPchg, 64); err == nil {
		dowPchg = s2
	}

	reason := "No purchase "
	if dow < dowPrevious {
		reason += "Dow is less than previous day close dow: " + downDayEval.Dow + " < " + downDayEval.PreviousDow + " "
	}
	if dowPchg < 10.00 {
		reason += "Greatest stock pchg does not meet delimiter required: " + downDayEval.GreatestPchg + " < " + "10.00 "
	}

	if cashDayEvaluationList[0].IsUnsettledFunds == "true" {
		reason += "Unsettled funds present"
	}

	//store time
	dowList := selectDow()
	tradeResultStore := TradeResultStore{
		AlgorithmUsed: "wisemen",
		Result:        reason,
		Dow1:          dowList[0].CurrentDowValue,
		Dow2:          dowList[1].CurrentDowValue,
		Dow3:          dowList[2].CurrentDowValue,
		Dow4:          dowList[3].CurrentDowValue,
	}
	insertTradeResultStore(tradeResultStore)
}
