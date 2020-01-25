package main

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"
)

//7:49
var checkIsMarketOpenMinute = 49
var checkIsMarketOpenHour = 7

//7:50
var conditionOneMinute = 50
var conditionOneHour = 7

//8:00
var conditionTwoMinute = 0
var conditionTwoHour = 8

//conditionMinuteHandleCalculateDownDay1 8:27 engage
var conditionMinuteHandleCalculateDownDay1 = 27
var conditionHourHandleCalculateDownDay1 = 8

//9:00
var conditionFourMinute = 0
var conditionFourHour = 9

//9:30
var conditionFiveMinute = 30
var conditionFiveHour = 9

//9:44
var conditionSixMinute = 44
var conditionSixHour = 9

//11:00
var conditionSevenMinute = 0
var conditionSevenHour = 11

//1:30
var conditionNineteenMinute = 30
var conditionNineteenHour = 13

//
var conditionTimeMinuteDow1 = checkIsMarketOpenMinute
var conditionTimeHourDow1 = checkIsMarketOpenHour

var conditionTimeMinuteDow2 = conditionOneMinute
var conditionTimeHourDow2 = conditionOneHour

var conditionTimeMinuteDow3 = conditionTwoMinute
var conditionTimeHourDow3 = conditionTwoHour

var conditionTimeMinuteDow4 = conditionMinuteHandleCalculateDownDay1
var conditionTimeHourDow4 = conditionHourHandleCalculateDownDay1

var conditionTimeMinuteDow5 = 45
var conditionTimeHourDow5 = 9

var conditionTimeMinuteDow6 = 0
var conditionTimeHourDow6 = 12

// var isDowStore = true
var checkIsMarketOpenBool = true
var boolOperate1 = true
var boolOperate2 = true
var boolOperateHandleCalculateDownDay1 = true
var boolOperate4 = true
var boolOperate5 = true
var boolOperate6 = true
var boolOperate19 = true

var boolOperateDow1 = true
var boolOperateDow2 = true
var boolOperateDow3 = true
var boolOperateDow4 = true
var boolOperateDow5 = true
var boolOperateDow6 = true

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
	cycleMapPool = map[string]*Cycle{}
	createCycle(15, 10000000, monitorSell, "monitorSell", []string{symbol, dropLoss, timeToSell})
	operatingCycle := cycleMapPool["monitorSell"]
	go startCycle(operatingCycle)
}

func processMonitorSellMarket(symbol string) {
	cycleMapPool = map[string]*Cycle{}
	createCycle(15, 10000000, monitorSellMarket, "monitorSellMarket", []string{symbol})
	operatingCycle := cycleMapPool["monitorSellMarket"]
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
	createCycle(300, 1000000, handleWhaleQueryStockList, "handleWhaleQueryStockList")
	operatingCycle := cycleMapPool["handleWhaleQueryStockList"]
	go startCycle(operatingCycle)
	// initialWhaleStockQueryPerformed = true
	// }
}

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
	// initialWhaleStockQueryPerformed = true
}

func checkIsDowStore(currentHour int, currentMinute int) {
	// boolOperateDow1
	// conditionTimeMinuteDow1
	// conditionTimeHourDow1
	if currentMinute == conditionTimeMinuteDow1 && currentHour == conditionTimeHourDow1 && boolOperateDow1 {
		dowValue := handleDowWebscrape()
		insertDow(dowValue)
		boolOperateDow1 = false
	}

	if currentMinute == conditionTimeMinuteDow2 && currentHour == conditionTimeHourDow2 && boolOperateDow2 {
		dowValue := handleDowWebscrape()
		insertDow(dowValue)
		boolOperateDow2 = false
	}

	if currentMinute == conditionTimeMinuteDow3 && currentHour == conditionTimeHourDow3 && boolOperateDow3 {
		dowValue := handleDowWebscrape()
		insertDow(dowValue)
		boolOperateDow3 = false
	}

	if currentMinute == conditionTimeMinuteDow4 && currentHour == conditionTimeHourDow4 && boolOperateDow4 {
		dowValue := handleDowWebscrape()
		insertDow(dowValue)
		boolOperateDow4 = false
	}

	//9:45
	if currentMinute == conditionTimeMinuteDow5 && currentHour == conditionTimeHourDow5 && boolOperateDow5 {
		dowValue := handleDowWebscrape()
		insertDow(dowValue)
		boolOperateDow5 = false
	}

	//12:00
	if currentMinute == conditionTimeMinuteDow6 && currentHour == conditionTimeHourDow6 && boolOperateDow6 {
		dowValue := handleDowWebscrape()
		insertDow(dowValue)
		boolOperateDow6 = false
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
		checkIsMarketOpenBool = false
		boolOperate19 = true
		systemStartProcesses()
		checKIsBrokerageResponding()
		healthCheck()
		//Wisemen algorithm same day calculation
		// handleDayRotation()
		conditionMarketClosed := selectMarketOpenAnalysis()
		if conditionMarketClosed[0].IsMarketClosed == "true" {
			setTimelineOperationsFalse()
			postMarketClosedEmail()
			fmt.Println("inside hit12")
		}

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
		handleTSPCollectionStatementPhase()
		// processWhaleQueryStockSet()
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

		// handleTSPCollectionStatementPhase1()
		healthCheck()
	}
	if currentTime.Minute() == conditionFourMinute && currentTime.Hour() == conditionFourHour && boolOperate4 {
		fmt.Println("hit4")
		boolOperate4 = false
		// listInformationAtTrade := selectInformationAtTrade()
		// if len(listInformationAtTrade) == 0 {
		// 	handleOverarchTopStock()
		// }
		handleTSPCollectionStatementPhase1()
	}
	if currentTime.Minute() == conditionFiveMinute && currentTime.Hour() == conditionFiveHour && boolOperate5 {
		fmt.Println("hit5")
		boolOperate5 = false
		// listInformationAtTrade := selectInformationAtTrade()
		// if len(listInformationAtTrade) == 0 {
		// 	handleOverarchTopStock()
		// }
	}
	if currentTime.Minute() == conditionSixMinute && currentTime.Hour() == conditionSixHour && boolOperate6 {
		fmt.Println("hit6")
		boolOperate6 = false
		// listInformationAtTrade := selectInformationAtTrade()
		// if len(listInformationAtTrade) == 0 {
		// 	handleOverarchTopStock()
		// }
	}
	if currentTime.Minute() == conditionNineteenMinute && currentTime.Hour() == conditionNineteenHour && boolOperate19 {
		fmt.Println("hit19")
		boolOperate19 = false
		listInformationAtTrade := selectInformationAtTrade()

		//handle ifMarketClosed with information about TRS.
		isMarketClosed := selectMarketOpenAnalysis()[0].IsMarketClosed
		if isMarketClosed == "true" {
			//handle TRS where market closed.
			createTradeResultStoreMarketClosed()
			fmt.Println("market closed")
		}
		if isMarketClosed != "true" {
			if len(listInformationAtTrade) != 2 {
				handleOverarchTopStock()
				handleNoBuyOnTradeDay()
				fmt.Println("No buy")
			}
		}

		//End of day dow scrape for next day analytics
		handleEndOfDayDowScrape()
		//At some point in this hour reset the pools, and reset the timeline.
		//clear cyclepool and reset timeline process
		resetCyclePools()
		handleDayReset()
		healthCheck()
		healthCheck()
		// resetTimeOperations()
	}
}

func resetTimeOperations() {
	startMinute := getCurrentMinute()
	startHour := getCurrentHour()

	//7:49
	checkIsMarketOpenMinute = startMinute
	checkIsMarketOpenHour = startHour
	fmt.Println("checkIsMarketOpenMinute")
	fmt.Println(checkIsMarketOpenMinute)
	fmt.Println("checkIsMarketOpenHour")
	fmt.Println(checkIsMarketOpenHour)

	//7:50
	conditionOneMinute = startMinute + 1
	conditionOneHour = startHour

	fmt.Println("conditionOneMinute")
	fmt.Println(conditionOneMinute)
	fmt.Println("conditionOneHour")
	fmt.Println(conditionOneHour)

	//8:00
	conditionTwoMinute = startMinute + 2
	conditionTwoHour = startHour

	fmt.Println("conditionTwoMinute")
	fmt.Println(conditionTwoMinute)
	fmt.Println("conditionTwoHour")
	fmt.Println(conditionTwoHour)

	//conditionMinuteHandleCalculateDownDay1 8:27 engage
	conditionMinuteHandleCalculateDownDay1 = startMinute + 3
	conditionHourHandleCalculateDownDay1 = startHour

	fmt.Println("conditionMinuteHandleCalculateDownDay1")
	fmt.Println(conditionMinuteHandleCalculateDownDay1)
	fmt.Println("conditionHourHandleCalculateDownDay1")
	fmt.Println(conditionHourHandleCalculateDownDay1)

	//9:00
	conditionFourMinute = startMinute + 4
	conditionFourHour = startHour

	fmt.Println("conditionFourMinute")
	fmt.Println(conditionFourMinute)
	fmt.Println("conditionFourHour")
	fmt.Println(conditionFourHour)

	//9:30
	conditionFiveMinute = startMinute + 5
	conditionFiveHour = startHour

	fmt.Println("conditionFiveMinute")
	fmt.Println(conditionFiveMinute)
	fmt.Println("conditionFiveHour")
	fmt.Println(conditionFiveHour)

	//9:44
	conditionSixMinute = startMinute + 6
	conditionSixHour = startHour

	fmt.Println("conditionSixMinute")
	fmt.Println(conditionSixMinute)
	fmt.Println("conditionSixHour")
	fmt.Println(conditionSixHour)

	//11:00
	conditionSevenMinute = startMinute + 7
	conditionSevenHour = startHour

	//1:30
	conditionNineteenMinute = startMinute + 8
	conditionNineteenHour = startHour

	//

	conditionTimeMinuteDow1 = checkIsMarketOpenMinute
	conditionTimeHourDow1 = checkIsMarketOpenHour

	conditionTimeMinuteDow2 = conditionOneMinute
	conditionTimeHourDow2 = conditionOneHour

	conditionTimeMinuteDow3 = conditionTwoMinute
	conditionTimeHourDow3 = conditionTwoHour

	conditionTimeMinuteDow4 = conditionMinuteHandleCalculateDownDay1
	conditionTimeHourDow4 = conditionHourHandleCalculateDownDay1

	conditionTimeMinuteDow5 = conditionSixMinute
	conditionTimeHourDow5 = conditionSixHour

	conditionTimeMinuteDow6 = conditionSevenMinute
	conditionTimeHourDow6 = conditionSevenHour

	fmt.Println("conditionNineteenMinute")
	fmt.Println(conditionNineteenMinute)
	fmt.Println("conditionNineteenHour")
	fmt.Println(conditionNineteenHour)
}

func handleEndOfDayDowScrape() {
	truncateEndOfDayDow()
	dowValue := handleDowWebscrape()
	// dowValue := "28,701.38"

	insertEndOfDayDow(dowValue)
}

func createTradeResultStoreMarketClosed() {
	tradeResultStore := TradeResultStore{Result: "Market closed"}
	insertTradeResultStore(tradeResultStore)
}

func handleNoBuyOnTradeDay() {
	tradeResultStore := TradeResultStore{Result: "No buy on trade day"}
	insertTradeResultStore(tradeResultStore)
	postNoBuyOnTradeDayEmail()
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
	listVal := reflect.ValueOf(params[0])
	var listSymbolsInterface interface{} = listVal.Index(0).Interface()
	listSymbols := listSymbolsInterface.([]string)
	symbol := listSymbols[0]
	//Is holding for symbol present
	holdingWisemen := HoldingWisemen{Symbol: "default"}
	//if holding matches symbol seeking
	holdingList := getAllHolding()
	for i, v := range holdingList.ListHolding {
		if v.Symbol == symbol {
			holdingWisemen = HoldingWisemen{Symbol: symbol, Price: v.Price, Qty: v.Qty, OrderStatus: "pending eval"}
		}
		i++
	}
	//
	//Notice it's getting IAT where none entered
	holdingWisemen = calculateHoldingStatus(holdingWisemen)
	//Handle conditions for holding incomplete
	if holdingWisemen.OrderStatus == "completedFull" {
		fmt.Println("completedFull hit")
		//End cycle for monitoring
		//Remove previous IAT at order placement...
		truncateInformationAtTrade()
		//Insert IAT now that buy placed detected.
		//holdingWisemen.Qty to 1 value
		//break a
		formattedQty := formatQtyHolding(holdingWisemen.Qty)
		handleInsertInformationAtTrade(symbol, "limit", "buy", formattedQty)
		cancelCycle(cycleMapPool["handleCheckIsTradeBought"])
		response := postNeoBuyOrderResponse(holdingWisemen)
		fmt.Println(response)
	}

	//handle time delmiter 11
	// convert number delimiters to time
	// if minute is len 1, add 0
	// create time delimiter
	// hourDelimiter := strconv.Itoa(conditionSevenHour)
	// minuteDelimiter := strconv.Itoa(conditionSevenMinute)

	//set static apart...
	hourDelimiter := "8"
	minuteDelimiter := "50"

	if len(minuteDelimiter) == 1 {
		minuteDelimiter = "0" + minuteDelimiter
	}
	//
	timeDelimiter := hourDelimiter + minuteDelimiter
	fmt.Println("created timeDelimiter")
	fmt.Println(timeDelimiter)
	//update
	isTimeDelimiterMet := calculateIsTimeDelimiterMetSell(timeDelimiter)

	if isTimeDelimiterMet {
		//cancel cycle
		cancelCycle(cycleMapPool["handleCheckIsTradeBought"])

		//cancel order
		orderList := getAllOrders()
		// fmt.Println("len(orderList.ListOrders)")
		// fmt.Println(len(orderList.ListOrders))
		order := Order{}
		for i, v := range orderList.ListOrders {
			if v.Symbol == symbol {
				order = v
				break
			}
			i++
		}
		queryCancelOrder(order.SVI, symbol)

		//send email
		postCancellationBuyOrderEmail(symbol)
	}
}

func formatQtyHolding(qty string) string {
	formattedQty := strings.Split(qty, ".")[0]
	return formattedQty
}

// func handleOverarchTopStock(params ...interface{}) {
func handleOverarchTopStock() {
	twiStockList := twiWebscrape()
	// fmt.Println(twiStockList)
	// fmt.Println("overarch after twi")
	// twiStockList := []Stock{}
	//High process for wisemen and whale
	highTransferanceProcess(twiStockList)
}

// if len(twiStockList) == 0 {
// 	go handleOverarchTopStockAync()
// 	return
// }

// for i, v := range twiStockList {
// 	fmt.Println(v.Symbol)
// 	fmt.Println(v.Pchg)
// 	i++
// }
// // // // //Low process for whale
// lowTransferanceProcess(twiStockList)

func handleOverarchTopStockAync() {
	fmt.Println("async operation activated")
	twiStockList := twiWebscrape()
	if len(twiStockList) == 0 {
		postNodeTSPAsyncFailureEmail()
		return
	}
	postNodeTSPAsyncSuccessEmail()
	//High process for wisemen and whale
	// highTransferanceProcess(twiStockList)
}

func highTransferanceProcess(twiStockList []Stock) {
	//TSP
	topStockPullStockList := topStockPull()
	listTempDuplicantFiltered := []Stock{}
	positiveTwiTSPList := []Stock{}
	positiveTwiStockList := []Stock{}
	isDuplicate := false
	isDuplicateTemp := false
	// for i,v := range topStockPullStockList()

	for i, v := range twiStockList {
		if v.IsCurrentPriceHigherThanPreviousClose == "true" {
			positiveTwiStockList = append(positiveTwiStockList, v)
		}
		i++
	}

	positiveTwiTSPList = positiveTwiStockList
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

	//
	// fmt.Println("listTempDuplicantFiltered")
	for i, v := range listTempDuplicantFiltered {
		fmt.Println(v)
		i++
	}
	for i < 3 {
		// 	// remove highest index 3 times, to get top stocks.
		// 	//Pop top stock each iteration
		fmt.Println("len(listTempDuplicantFiltered)")
		fmt.Println(len(listTempDuplicantFiltered))
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

	for i, v := range topStockList {
		insertTempSymbolHoldHigh(v.Symbol)
		i++
	}
	handleWisemenFillHold()
	// handleWhaleFillHoldHigh()
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

	// //Handle negative values only in twiTSPList
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
			fmt.Println("v that is lower")
			fmt.Println(v)
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
		// if strings.Contains(v.Symbol, ".") == false {
		// 	filteredStockList = append(filteredStockList, v)
		// }
		if strings.Contains(v.Symbol, ".") == false {
			if strings.Contains(v.Symbol, "'") == false {
				filteredStockList = append(filteredStockList, v)
			}
		}
		i++
	}
	return filteredStockList
}

func healthCheck() {
	isNeoResponse := "false"
	isPythonDBResponse := "false"

	//post to neo
	responseNeo := postNeoHealthCheck()
	// //If response from neo
	if responseNeo != "error received" {
		// fmt.Println(response)
		isNeoResponse = "true"
	}

	//post to python
	// responsePython := selectNews()
	response := postCommandDBSelect("SELECT news_info FROM news")

	// fmt.Println("response")
	// fmt.Println(response)

	if !strings.Contains(response, "<!DOCTYPE HTML") {
		// fmt.Println("wowzers")
		if response != "error received" {
			// fmt.Println("enters")
			isPythonDBResponse = "true"
		}
	}
	postHealthCheckNode(isNeoResponse, isPythonDBResponse)
	//res from check
	//nodemail
}

func purchaseUpdateSystem() {
}
func getTradeResultStoreList() []TradeResultStore {
	tradeResultStoreList := selectTradeResultStore("wisemen")
	//get latest tradeResultStore
	return tradeResultStoreList
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
	// fmt.Println("hit handleWisemenQueryStockList")
	//Query monitor_symbol
	symbolList := selectWisemenSymbolHold()
	// fmt.Println("symbolList")
	// fmt.Println(symbolList)

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
	dowValue := parseDowWebscrape(response)
	dowList := []Dow{Dow{CurrentDowValue: dowValue}}
	dowValue = formatDowListRemoveCommaValues(dowList)[0].CurrentDowValue
	return dowValue
}

func twiWebscrape() []Stock {
	indexTwiWebscrape := 0
	response2 := ""
	stockList := []Stock{}
	symbolList := []string{}
	//keep trying
	for indexTwiWebscrape < 10 {
		response2 = queryWebscrapeTwi()
		fmt.Println("inside twi")
		if response2 != "try again failure" {
			// containsSymbolsTwiScrape
			symbolList = parseTwiWebscrape(response2)
			fmt.Println("symbolList internal")
			fmt.Println(symbolList)
			if len(symbolList) > 3 {
				break
			}
		}
		postNodeTSPFailureEmail()

		time.Sleep(time.Duration(1) * time.Second)
		fmt.Println("continuing")
		indexTwiWebscrape++
		if indexTwiWebscrape == 10 {
			fmt.Println("doing awesomeness")
			return []Stock{}
		}
	}
	fmt.Println("twiWebscrape")
	if len(symbolList) == 0 {
		symbolList = parseTwiWebscrape(response2)
		fmt.Println("symbolList == 0")
		fmt.Println(symbolList)
	}
	//
	// symbolList := []string{"AAPL", "MX.A", "MAX.O", "NVCN", "TOY.A"}
	// formattedSymbolList := filterTwiSymbolList(symbolList)
	if len(symbolList) != 0 {
		responseSymbolList := queryMultiStockPull(symbolList)
		stockList = parseStockSetQuery(responseSymbolList)
	}
	return stockList
}

func filterTwiSymbolList(symbolListToFilter []string) []string {
	symbolList := []string{}
	for i, v := range symbolListToFilter {
		if strings.Contains(v, ".") == false {
			symbolList = append(symbolList, v)
		}
		i++
	}
	return symbolList
}

// func twiWebscrapeAync(params ...interface{}) {
// 	indexTwiWebscrape := 0
// 	response2 := ""
// 	//keep trying
// 	for indexTwiWebscrape < 10 {
// 		response2 = queryWebscrapeTwi()

// 		if response2 != "try again failure" {
// 			break
// 		}
// 		postNodeTSPFailureEmail()

// 		time.Sleep(time.Duration(1) * time.Second)
// 		fmt.Println("continuing")
// 		indexTwiWebscrape++
// 		if indexTwiWebscrape == 10 {
// 			fmt.Println("doing awesomeness")
// 			return []Stock{}
// 		}
// 	}
// 	// catch response2 failure
// 	symbolList := parseTwiWebscrape(response2)
// 	responseSymbolList := queryMultiStockPull(symbolList)
// 	stockList := parseStockSetQuery(responseSymbolList)
// 	return stockList
// }

// for i, v := range stockList {
// 	fmt.Println(v.Symbol)
// 	fmt.Println("v.Last")
// 	fmt.Println(v.Last)
// 	fmt.Println("v.Pcls")
// 	fmt.Println(v.Pcls)
// 	fmt.Println("v.IsCurrentPriceHigherThanPreviousClose")
// 	fmt.Println(v.IsCurrentPriceHigherThanPreviousClose)
// 	i++
// }
// func handleEndOfDayAnalyticsOperations() {
// 	day := getDayOfWeek()
// 	//insert into table conditional
// }

func systemStartProcesses() {
	calculateShortDayAnalysis()
	handleInsertMetricsConditionalAlteration()
	//get balance begining of day
	storeBalanceValue()

}
func handleInsertMetricsConditionalAlteration() {
	shortDayAnalysis := selectShortDayAnalysis()[0]
	if shortDayAnalysis.IsShortDay == "true" {
		insertMetricsWisemen("20.00", "4.0", "20.0", "0", ".015", ".1", "1057")
		return
	}
	if shortDayAnalysis.IsShortDay == "false" {
		insertMetricsWisemen("20.00", "4.0", "20.0", "0", ".015", ".1", "1330")
		return
	}
}

func storeBalanceValue() {
	balance := queryBalance()
	parsedBalance := parseBalance(balance)
	insertAccountBalance(AccountBalance{Balance: parsedBalance})
}

func resetTempSymbolHold() {
	truncateTempSymbolHoldHigh()
}
func resetStockWisemenSymbolHold() {
	truncateWisemenSymbolHold()
}
func resetStockWisemen() {
	truncateStockWisemen()
}
func resetAltIntervalBuyWisemen() {
	truncateAltIntervalBuyWisemen()
}
func resetInformationAtTrade() {
	truncateInformationAtTrade()
}

func handleDayReset() {
	boolOperate1 = true
	boolOperate2 = true
	// boolOperate3 = true
	boolOperateHandleCalculateDownDay1 = true
	boolOperate4 = true
	boolOperate5 = true
	boolOperate6 = true
	boolOperateDow1 = true
	boolOperateDow2 = true
	boolOperateDow3 = true
	boolOperateDow4 = true
	boolOperateDow5 = true
	boolOperateDow6 = true
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

	//reset global cache.
	// globalCacheStockWisemenTopTier = []string{}

	truncateShortDayAnalysis()
	truncateMarketOpenAnalysis()
	truncateMetricsWisemen()
	truncateDow()
	resetTempSymbolHold()
	resetStockWisemenSymbolHold()
	resetStockWisemen()
	resetAltIntervalBuyWisemen()
	resetInformationAtTrade()
	fmt.Println("reset completed")
}

func setTimelineOperationsFalse() {
	boolOperate1 = false
	boolOperate2 = false
	boolOperateHandleCalculateDownDay1 = false
	boolOperate4 = false
	boolOperate5 = false
	boolOperate6 = false

	boolOperateDow1 = false
	boolOperateDow2 = false
	boolOperateDow3 = false
	boolOperateDow4 = false
	boolOperateDow5 = false
	boolOperateDow6 = false
	boolOperateHandleCalculateDownDay1 = false
}

func wrapUpWisemenOutcome(transactionHistory TransactionHistory) {
	fmt.Println("hit")
	alteredTransactionHistory := calculateTransactionHistory(transactionHistory)
	//get although this is reset insertInformationAtTrade for buy and sell for day
	listMatchingSymbolInformationAtTrade := handleInformationAtTradeDayListArbitration(alteredTransactionHistory.Symbol)

	// fmt.Println("listMatchingSymbolInformationAtTrade")
	// fmt.Println(listMatchingSymbolInformationAtTrade)
	// listMatchingSymbolInformationAtTrade := []InformationAtTrade{}

	fmt.Println("inside listMatchingSymbolInformationAtTrade")
	//query stocks
	stockList := selectStockWisemenBySymbol(transactionHistory.Symbol)

	highestLastIndex := 0
	highestLast := 0.0
	//iterate through pchg list, find highest
	for i, stock := range stockList {
		floatLast := 0.0
		if s, err := strconv.ParseFloat(stock.Last, 64); err == nil {
			floatLast = s
		}

		if i == 0 {
			highestLast = floatLast
			continue
		}
		if floatLast > highestLast {
			highestLast = floatLast
			highestLastIndex = i
		}
	}
	highestStock := stockList[highestLastIndex]
	timeCreatedHigh := highestStock.TimeCreated
	fmt.Println(highestStock.Last)
	fmt.Println(timeCreatedHigh)

	fmt.Println("highestStock")
	fmt.Println(highestStock.Last)
	fmt.Println("timeCreatedHigh")
	fmt.Println(timeCreatedHigh)
	//

	lowestLastIndex := 0
	lowestLast := 0.0
	//iterate through pchg list, find highest
	for i, stock := range stockList {
		floatLast := 0.0
		if s, err := strconv.ParseFloat(stock.Last, 64); err == nil {
			floatLast = s
		}
		if i == 0 {
			lowestLast = floatLast
			continue
		}
		if floatLast < lowestLast {
			lowestLast = floatLast
			lowestLastIndex = i
		}
	}
	lowestStock := stockList[lowestLastIndex]
	timeCreatedLow := lowestStock.TimeCreated
	//
	fmt.Println("lowestStock")
	fmt.Println(lowestStock.Last)
	fmt.Println("timeCreatedLow")
	fmt.Println(timeCreatedLow)

	dowList := selectDow()
	//Should not be called if no trade completed
	if len(listMatchingSymbolInformationAtTrade) == 0 {
		//no trade occured handle TradeResultStore
		if len(dowList) == 0 {
			tradeResultStore := TradeResultStore{
				AlgorithmUsed: "wisemen",
				Result:        "No trade len(dowList) == 0",
			}
			insertTradeResultStore(tradeResultStore)
		}
		if len(dowList) == 4 {
			tradeResultStore := TradeResultStore{
				AlgorithmUsed: "wisemen",
				Result:        "No trade len(dowList) == 4",
				Dow1:          dowList[0].CurrentDowValue,
				Dow2:          dowList[1].CurrentDowValue,
				Dow3:          dowList[2].CurrentDowValue,
				Dow4:          dowList[3].CurrentDowValue,
			}
			insertTradeResultStore(tradeResultStore)
		}
		if len(dowList) == 5 {
			tradeResultStore := TradeResultStore{
				AlgorithmUsed: "wisemen",
				Result:        "No trade len(dowList) == 5",
				Dow1:          dowList[0].CurrentDowValue,
				Dow2:          dowList[1].CurrentDowValue,
				Dow3:          dowList[2].CurrentDowValue,
				Dow4:          dowList[3].CurrentDowValue,
				Dow5:          dowList[4].CurrentDowValue,
			}
			insertTradeResultStore(tradeResultStore)
		}
		if len(dowList) == 6 {
			tradeResultStore := TradeResultStore{
				AlgorithmUsed: "wisemen",
				Result:        "No trade len(dowList) == 6",
				Dow1:          dowList[0].CurrentDowValue,
				Dow2:          dowList[1].CurrentDowValue,
				Dow3:          dowList[2].CurrentDowValue,
				Dow4:          dowList[3].CurrentDowValue,
				Dow5:          dowList[4].CurrentDowValue,
				Dow6:          dowList[5].CurrentDowValue,
			}
			insertTradeResultStore(tradeResultStore)
		}
	}

	//if buy and sell exists...InformationAtTrade
	if len(listMatchingSymbolInformationAtTrade) == 2 {
		buyHistoryValuePrice := alteredTransactionHistory.HistoryValueList[0].Price
		sellHistoryValuePrice := alteredTransactionHistory.HistoryValueList[1].Price
		// buyHistoryValuePrice := "11.39"  //alteredTransactionHistory.HistoryValueList[0].Price
		// sellHistoryValuePrice := "12.39" //alteredTransactionHistory.HistoryValueList[1].Price

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

		// fmt.Println("changeAmount")
		// fmt.Println(changeAmount)

		floatPercentageChangeAmount := changeAmount / floatBuyHistoryValuePrice
		stringChangeAmount := fmt.Sprintf("%f", floatPercentageChangeAmount)

		// fmt.Println("floatPercentageChangeAmount")
		// fmt.Println(stringChangeAmount)

		stringChangeAmount = transformPercentageToPercentageVisual(stringChangeAmount)
		result := "negative"
		//if sell was less than optimal
		if floatPercentageChangeAmount >= 0 {
			result = "positive "
		}
		//add handle result, if market trade, droploss or time delimiter
		//
		if listMatchingSymbolInformationAtTrade[1].TypeTrade != "limit" {
			result += listMatchingSymbolInformationAtTrade[1].TypeTrade + " "
		}

		boughtTime := listMatchingSymbolInformationAtTrade[0].Hour + " " + listMatchingSymbolInformationAtTrade[0].Minute
		sellTime := listMatchingSymbolInformationAtTrade[1].Hour + " " + listMatchingSymbolInformationAtTrade[1].Minute

		// boughtTime := "13 14"
		// sellTime := "14 15"

		//getHoldingBuy

		if len(dowList) == 6 {
			tradeResultStore := TradeResultStore{
				AlgorithmUsed:           "wisemen",
				Result:                  result,
				BoughtPrice:             buyHistoryValuePrice,
				SellPrice:               sellHistoryValuePrice,
				ChangeAmount:            stringChangeAmount,
				StockSymbol:             alteredTransactionHistory.Symbol,
				Qty:                     alteredTransactionHistory.HistoryValueList[0].Qty,
				TimeTradeBuy:            boughtTime,
				TimeTradeSell:           sellTime,
				HighestPricePointForDay: highestStock.Last,
				TimeHighestPricePoint:   timeCreatedHigh,
				LowestPricePointForDay:  lowestStock.Last,
				TimeLowestPricePoint:    timeCreatedLow,
				Dow1:                    dowList[0].CurrentDowValue,
				Dow2:                    dowList[1].CurrentDowValue,
				Dow3:                    dowList[2].CurrentDowValue,
				Dow4:                    dowList[3].CurrentDowValue,
				Dow5:                    dowList[4].CurrentDowValue,
				Dow6:                    dowList[5].CurrentDowValue,
			}
			// fmt.Println("tradeResultStore")
			// fmt.Println(tradeResultStore)
			insertTradeResultStore(tradeResultStore)
			postEmailTradeResultStore(tradeResultStore)
		}

		//
		if len(dowList) == 5 {
			tradeResultStore := TradeResultStore{
				AlgorithmUsed:           "wisemen",
				Result:                  result,
				BoughtPrice:             buyHistoryValuePrice,
				SellPrice:               sellHistoryValuePrice,
				ChangeAmount:            stringChangeAmount,
				StockSymbol:             alteredTransactionHistory.Symbol,
				Qty:                     alteredTransactionHistory.HistoryValueList[0].Qty,
				TimeTradeBuy:            boughtTime,
				TimeTradeSell:           sellTime,
				HighestPricePointForDay: highestStock.Last,
				TimeHighestPricePoint:   timeCreatedHigh,
				LowestPricePointForDay:  lowestStock.Last,
				TimeLowestPricePoint:    timeCreatedLow,
				Dow1:                    dowList[0].CurrentDowValue,
				Dow2:                    dowList[1].CurrentDowValue,
				Dow3:                    dowList[2].CurrentDowValue,
				Dow4:                    dowList[3].CurrentDowValue,
				Dow5:                    dowList[4].CurrentDowValue,
			}
			// fmt.Println("tradeResultStore")
			// fmt.Println(tradeResultStore)
			insertTradeResultStore(tradeResultStore)
			postEmailTradeResultStore(tradeResultStore)
		}
		//
		if len(dowList) == 4 {
			tradeResultStore := TradeResultStore{
				AlgorithmUsed:           "wisemen",
				Result:                  result,
				BoughtPrice:             buyHistoryValuePrice,
				SellPrice:               sellHistoryValuePrice,
				ChangeAmount:            stringChangeAmount,
				StockSymbol:             alteredTransactionHistory.Symbol,
				Qty:                     alteredTransactionHistory.HistoryValueList[0].Qty,
				TimeTradeBuy:            boughtTime,
				TimeTradeSell:           sellTime,
				HighestPricePointForDay: highestStock.Last,
				TimeHighestPricePoint:   timeCreatedHigh,
				LowestPricePointForDay:  lowestStock.Last,
				TimeLowestPricePoint:    timeCreatedLow,
				Dow1:                    dowList[0].CurrentDowValue,
				Dow2:                    dowList[1].CurrentDowValue,
				Dow3:                    dowList[2].CurrentDowValue,
				Dow4:                    dowList[3].CurrentDowValue,
			}
			// fmt.Println("tradeResultStore")
			// fmt.Println(tradeResultStore)
			insertTradeResultStore(tradeResultStore)
			postEmailTradeResultStore(tradeResultStore)
		}
		if len(dowList) < 4 {
			tradeResultStore := TradeResultStore{
				AlgorithmUsed:           "wisemen",
				Result:                  result,
				BoughtPrice:             buyHistoryValuePrice,
				SellPrice:               sellHistoryValuePrice,
				ChangeAmount:            stringChangeAmount,
				StockSymbol:             alteredTransactionHistory.Symbol,
				Qty:                     alteredTransactionHistory.HistoryValueList[0].Qty,
				TimeTradeBuy:            boughtTime,
				TimeTradeSell:           sellTime,
				HighestPricePointForDay: highestStock.Last,
				TimeHighestPricePoint:   timeCreatedHigh,
				LowestPricePointForDay:  lowestStock.Last,
				TimeLowestPricePoint:    timeCreatedLow,
			}
			insertTradeResultStore(tradeResultStore)
			postEmailTradeResultStore(tradeResultStore)
		}
	}
}

func wrapUpWisemenOutcomeNoBuy(transactionHistory TransactionHistory) {
	downDayEvaluationList := selectDownDayEvaluation()
	cashDayEvaluationList := selectCashDayEvaluation()

	downDayEval := downDayEvaluationList[len(downDayEvaluationList)-1]
	dow := 0.0
	dowPrevious := 0.0
	// downDayEvalPchg := 0.0

	//handle where does not pass neo...
	//query Neo no handle
	//handle select list on AIB
	listAltIntervalBuyWisemen := selectAltIntervalBuyWisemen()
	reason := "No purchase "
	//Handle where previous dow does not exist.

	for i, v := range listAltIntervalBuyWisemen {
		reason += v.ReasonCancelation + " "
		i++
	}

	if downDayEval.PreviousDow == "does not exist" {
		if cashDayEvaluationList[0].IsUnsettledFunds == "true" {
			reason += "Unsettled funds present"
		}
	}

	if downDayEval.PreviousDow != "does not exist" {
		if s, err := strconv.ParseFloat(downDayEval.Dow, 64); err == nil {
			dow = s
		}
		if s1, err := strconv.ParseFloat(downDayEval.PreviousDow, 64); err == nil {
			dowPrevious = s1
		}
		if dow < dowPrevious {
			reason += "Dow is less than previous day close dow: " + downDayEval.Dow + " < " + downDayEval.PreviousDow + " "
		}
		if cashDayEvaluationList[0].IsUnsettledFunds == "true" {
			reason += "Unsettled funds present"
		}
	}
	//store time
	// dowList := selectDow()
	tradeResultStore := TradeResultStore{
		AlgorithmUsed: "wisemen",
		Result:        reason,
		// Dow1:          dowList[0].CurrentDowValue,
		// Dow2:          dowList[1].CurrentDowValue,
		// Dow3:          dowList[2].CurrentDowValue,
		// Dow4:          dowList[3].CurrentDowValue,
	}
	insertTradeResultStore(tradeResultStore)
}

func transformPercentageToPercentageVisual(stringValue string) string {
	// floatValue := 0.0
	// if s, err := strconv.ParseFloat(stringValue, 64); err == nil {
	// 	floatValue = s
	// }
	//0.00955

	//split at period,
	//
	//.9%

	symbolString1 := strings.Split(stringValue, ".")[1]
	// fmt.Println(symbolString1)
	newString := ""
	for i, v := range symbolString1 {
		if i == 2 {
			newString += "."
		}
		newString += string(v)
	}

	// symbolString2 := strings.Split(symbolString1, "</sym>")
	// symbol := symbolString2[0]
	returnString := ""
	for i, v := range newString {
		if i == 0 {
			if string(v) == "0" {
				continue
			}
		}
		returnString += string(v)
	}
	// fmt.Println("returnString")
	// fmt.Println(returnString)

	//
	zeroIndexCount := 0
	isPeriodReached := false
	returnString1 := ""
	for i, v := range returnString {
		if isPeriodReached {
			zeroIndexCount += 1
			if zeroIndexCount == 3 {
				break
			}
		}
		if string(v) == "." {
			isPeriodReached = true
		}
		returnString1 += string(v)
		i++
	}
	returnString1 += "%"

	if returnString1 == "0.00%" {
		returnString1 = "less than 1%"
	}
	return returnString1
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
