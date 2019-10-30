package main

import (
	"fmt"
	"reflect"
	"strings"
	"time"
)

var checkIsMarketOpenMinute = 24
var checkIsMarketOpenHour = 11

var conditionOneMinute = 25
var conditionOneHour = 11

var conditionTwoMinute = 0
var conditionTwoHour = 8

var conditionThreeMinute = 30
var conditionThreeHour = 8

var conditionFourMinute = 0
var conditionFourHour = 9

var conditionFiveMinute = 30
var conditionFiveHour = 9

var conditionSixMinute = 0
var conditionSixHour = 10

var conditionSevenMinute = 30
var conditionSevenHour = 10

var conditionEightMinute = 0
var conditionEightHour = 11

var conditionNineMinute = 30
var conditionNineHour = 11

var conditionTenMinute = 0
var conditionTenHour = 12

var conditionElevenMinute = 30
var conditionElevenHour = 12

var conditionTwelveMinute = 0
var conditionTwelveHour = 13

var conditionThirteenMinute = 30
var conditionThirteenHour = 13

var conditionFourteenMinute = 0
var conditionFourteenHour = 14

var conditionFifteenMinute = 30
var conditionFifteenHour = 14

var conditionSixteenMinute = 0
var conditionSixteenHour = 15

var conditionSeventeenMinute = 30
var conditionSeventeenHour = 15

var conditionEighteenMinute = 0
var conditionEighteenHour = 16

var conditionNineteenMinute = 26
var conditionNineteenHour = 11

var checkIsMarketOpenBool = true
var checkIsMarketOpenFollowUpBool = true
var boolOperate1 = true
var boolOperate2 = true
var boolOperate3 = true
var boolOperate4 = true
var boolOperate5 = true
var boolOperate6 = true
var boolOperate7 = true
var boolOperate8 = true
var boolOperate9 = true
var boolOperate10 = true
var boolOperate11 = true
var boolOperate12 = true
var boolOperate13 = true
var boolOperate14 = true
var boolOperate15 = true
var boolOperate16 = true
var boolOperate17 = true
var boolOperate18 = true
var boolOperate19 = true

var initialStockQueryPerformed = false
var initialWisemenStockQueryPerformed = false
var initialWhaleStockQueryPerformed = false

func processTimelineStart() {
	cycleMapPool = map[string]*Cycle{}
	createCycle(5, 10000000000000, handleTimelineConditionalTriggers, "handleTimelineConditionalTriggers")
	operatingCycle := cycleMapPool["handleTimelineConditionalTriggers"]
	go startCycle(operatingCycle)
}

func processWisemenQueryStockSet() {
	if initialWisemenStockQueryPerformed == true {
		fmt.Println("hit initialWisemenStockQueryPerformed == true")
		operatingCycle := cycleMapPool["handleWisemenQueryStockList"]
		go startCycle(operatingCycle)
	}
	if initialWisemenStockQueryPerformed == false {
		fmt.Println("hit initialWisemenStockQueryPerformed == false")
		createCycle(3, 1000000000000, handleWisemenQueryStockList, "handleWisemenQueryStockList")
		operatingCycle := cycleMapPool["handleWisemenQueryStockList"]
		go startCycle(operatingCycle)
		initialWisemenStockQueryPerformed = true
	}
}

func processWhaleQueryStockSet() {
	if initialWhaleStockQueryPerformed == true {
		go startCycle(cycleMapPool["handleWhaleQueryStockList"])
	}
	if initialWhaleStockQueryPerformed == false {
		createCycle(300, 1000000000000, handleWhaleQueryStockList, "handleWhaleQueryStockList")
		operatingCycle := cycleMapPool["handleWhaleQueryStockList"]
		go startCycle(operatingCycle)
		initialWhaleStockQueryPerformed = true
	}
}

// func intiateMonitorTradeWisemon() {
// 	// metrics := selectMetricsWisemen()
// 	// Select metrics make use of it, continue.
// 	// fmt.Println(metrics)
// 	//Delay before monitor cycle
// 	// time.Sleep(time.Duration(10) * time.Second)
// 	// fmt.Println("hit awesome")
// 	//single query is holding of symbol

// 	//if not delay, do iterate until true
// 	// for
// 	indexCheck := 1
// 	for indexCheck < 100000 {
// 		queryIsTradeCompleted()

// 		time.Sleep(time.Duration(3) * time.Second)
// 		indexCheck++
// 	}

// 	//evaluation if order is closed

// 	//
// }

// func processCheckIsBuyPeformed() {
// 	createCycle(3, 10000000000000, handleWisemenQueryStockList, "handleWisemenQueryStockList")
// 	operatingCycle := cycleMapPool["handleWisemenQueryStockList"]
// 	go startCycle(operatingCycle)
// }

func processTSPRefresh() {
	go handleTSPRefresh()
}

func processFillHolds() {
	go handleFillHolds()
}

func processDowWebscrape() {
	go handleDowWebscrape()
}
func processCheckIsTradeBought(symbol string) {
	// go handleCheckIsTradeBought()
	//THe idea is to check every 5 seconds, and if a trade evaluation is positive,
	//or if the time delimiter for checking is met, then cancle this cycle and record results in DB.
	createCycle(10, 3, handleCheckIsTradeBought, "handleCheckIsTradeBought", []string{symbol})
	operatingCycle := cycleMapPool["handleCheckIsTradeBought"]
	go startCycle(operatingCycle)
	initialWhaleStockQueryPerformed = true
}

func handleTimelineConditionalTriggers(params ...interface{}) {
	currentTime := time.Now()
	fmt.Println(currentTime.Hour())
	fmt.Println(currentTime.Minute())
	fmt.Println(currentTime.Second())
	fmt.Println(currentTime.Date())

	//Conditional operate
	if currentTime.Minute() == checkIsMarketOpenMinute && currentTime.Hour() == checkIsMarketOpenHour && checkIsMarketOpenBool {
		checKIsBrokerageResponding()
		//Wisemen algorithm same day calculation
		// handleDayRotation()
		if isMarketClosed {
			setTimelineOperationsFalse()
		}
		checkIsMarketOpenBool = false
		boolOperate19 = true
	}
	//Initiate monitoring pool query on cycle
	//Periodic TSP refresh
	if currentTime.Minute() == conditionOneMinute && currentTime.Hour() == conditionOneHour && boolOperate1 {
		fmt.Println("hit1 init operations")
		boolOperate1 = false
		// processTSPRefresh()
		handleTSPRefresh()
		processDowWebscrape()
		processWisemenQueryStockSet()
		processWhaleQueryStockSet()
	}
	// if currentTime.Minute() == testOneMinute && currentTime.Hour() == testOneHour && boolOperate1 {
	// 	fmt.Println("hit1 init operations")
	// 	queryCycle.BooleanOperate = true
	// 	boolOperate1 = false
	// 	boolOperate19 = true
	// 	handleTSPRefresh()
	// 	processQueryStockSet()
	// 	initialOperationPerformed = true
	// }
	if currentTime.Minute() == conditionTwoMinute && currentTime.Hour() == conditionTwoHour && boolOperate2 {
		fmt.Println("hit2")
		boolOperate2 = false
		processTSPRefresh()
		processDowWebscrape()
	}
	if currentTime.Minute() == conditionThreeMinute && currentTime.Hour() == conditionThreeHour && boolOperate3 {
		fmt.Println("hit3")
		boolOperate3 = false
		processTSPRefresh()
		processDowWebscrape()
	}
	if currentTime.Minute() == conditionFourMinute && currentTime.Hour() == conditionFourHour && boolOperate4 {
		fmt.Println("hit4")
		boolOperate4 = false
		processTSPRefresh()
		processDowWebscrape()
	}
	if currentTime.Minute() == conditionFiveMinute && currentTime.Hour() == conditionFiveHour && boolOperate5 {
		fmt.Println("hit5")
		boolOperate5 = false
		processTSPRefresh()
		processDowWebscrape()
	}
	if currentTime.Minute() == conditionSixMinute && currentTime.Hour() == conditionSixHour && boolOperate6 {
		fmt.Println("hit6")
		boolOperate6 = false
		processTSPRefresh()
		processDowWebscrape()
	}
	if currentTime.Minute() == conditionSevenMinute && currentTime.Hour() == conditionSevenHour && boolOperate7 {
		fmt.Println("hit7")
		boolOperate7 = false
		processTSPRefresh()
		processDowWebscrape()
	}
	if currentTime.Minute() == conditionEightMinute && currentTime.Hour() == conditionEightHour && boolOperate8 {
		fmt.Println("hit8")
		boolOperate8 = false
		processTSPRefresh()
		processDowWebscrape()
	}
	if currentTime.Minute() == conditionNineMinute && currentTime.Hour() == conditionNineHour && boolOperate9 {
		fmt.Println("hit9")
		boolOperate9 = false
		processTSPRefresh()
		processDowWebscrape()
	}
	if currentTime.Minute() == conditionTenMinute && currentTime.Hour() == conditionTenHour && boolOperate10 {
		fmt.Println("hit10")
		boolOperate10 = false
		processTSPRefresh()
		processDowWebscrape()
	}
	if currentTime.Minute() == conditionElevenMinute && currentTime.Hour() == conditionElevenHour && boolOperate11 {
		fmt.Println("hit11")
		boolOperate11 = false
		processTSPRefresh()
		processDowWebscrape()
	}
	if currentTime.Minute() == conditionTwelveMinute && currentTime.Hour() == conditionTwelveHour && boolOperate12 {
		fmt.Println("hit12")
		boolOperate12 = false
		processTSPRefresh()
		processDowWebscrape()
	}
	if currentTime.Minute() == conditionThirteenMinute && currentTime.Hour() == conditionThirteenHour && boolOperate13 {
		fmt.Println("hit13")
		boolOperate13 = false
		processTSPRefresh()
		processDowWebscrape()
	}
	if currentTime.Minute() == conditionFourteenMinute && currentTime.Hour() == conditionFourteenHour && boolOperate14 {
		fmt.Println("hit14")
		boolOperate14 = false
		processTSPRefresh()
		processDowWebscrape()
	}
	if currentTime.Minute() == conditionFifteenMinute && currentTime.Hour() == conditionFifteenHour && boolOperate15 {
		fmt.Println("hit15")
		boolOperate15 = false
		processTSPRefresh()
		processDowWebscrape()
	}
	if currentTime.Minute() == conditionSixteenMinute && currentTime.Hour() == conditionSixteenHour && boolOperate16 {
		fmt.Println("hit16")
		boolOperate16 = false
		processTSPRefresh()
		processDowWebscrape()
	}
	if currentTime.Minute() == conditionSeventeenMinute && currentTime.Hour() == conditionSeventeenHour && boolOperate17 {
		fmt.Println("hit17")
		boolOperate17 = false
		processTSPRefresh()
		processDowWebscrape()
	}
	if currentTime.Minute() == conditionEighteenMinute && currentTime.Hour() == conditionEighteenHour && boolOperate18 {
		fmt.Println("hit18")
		boolOperate18 = false
		processTSPRefresh()
		processDowWebscrape()
	}
	if currentTime.Minute() == conditionNineteenMinute && currentTime.Hour() == conditionNineteenHour && boolOperate19 {
		fmt.Println("hit19")
		boolOperate19 = false

		handleWisemenQueryStockListCycle := cycleMapPool["handleWisemenQueryStockList"]
		if handleWisemenQueryStockListCycle.BooleanOperate {
			cancelCycle(handleWisemenQueryStockListCycle)
		}

		handleWhaleQueryStockListCycle := cycleMapPool["handleWhaleQueryStockList"]
		if handleWhaleQueryStockListCycle.BooleanOperate {
			cancelCycle(handleWhaleQueryStockListCycle)
		}

		//At some point in this hour reset the pools, and reset the timeline.
		resetCyclePools()

		// resetTempSymbolHold()
		// resetStockWisemenSymbolHold()
		// handleEndOfDayAnalyticsOperations()
		// handleDayReset()
	}

	//Conditional operate
	// if currentTime.Minute() == conditionThreeMinute && currentTime.Hour() == conditionThreeHour && boolOperate3 {
	// 	checKIsBrokerageResponding()

	// 	if isMarketClosed {
	// 		setTimelineOperationsFalse()
	// 	}
	// 	boolOperate3 = false
	// 	boolOperate19 = true
	// }
	// //Initiate monitoring pool query on cycle
	// //Periodic TSP refresh
	// if currentTime.Minute() == conditionFourMinute && currentTime.Hour() == conditionFourHour && boolOperate4 {
	// 	fmt.Println("hit1 init operations")
	// 	boolOperate4 = false
	// 	// processTSPRefresh()
	// 	// processDowWebscrape()
	// 	processQueryStockSet()
	// }

	// if currentTime.Minute() == conditionFiveMinute && currentTime.Hour() == conditionFiveHour && boolOperate19 {
	// 	fmt.Println("hit19")
	// 	boolOperate19 = false
	// 	if queryCycle.BooleanOperate {
	// 		queryCycle.BooleanOperate = false
	// 	}
	// 	handleEndOfDayAnalyticsOperations()
	// 	handleDayReset()
	// }

	// //

	// //Conditional operate
	// if currentTime.Minute() == conditionSixMinute && currentTime.Hour() == conditionSixHour && boolOperate6 {
	// 	checKIsBrokerageResponding()

	// 	if isMarketClosed {
	// 		setTimelineOperationsFalse()
	// 	}
	// 	boolOperate6 = false
	// 	boolOperate19 = true
	// }
	// //Initiate monitoring pool query on cycle
	// //Periodic TSP refresh
	// if currentTime.Minute() == conditionSevenMinute && currentTime.Hour() == conditionSevenHour && boolOperate7 {
	// 	fmt.Println("hit1 init operations")
	// 	boolOperate7 = false
	// 	processTSPRefresh()
	// 	processDowWebscrape()
	// 	processQueryStockSet()
	// }

	// if currentTime.Minute() == conditionEightMinute && currentTime.Hour() == conditionEightHour && boolOperate19 {
	// 	fmt.Println("hit19")
	// 	boolOperate19 = false
	// 	if queryCycle.BooleanOperate {
	// 		queryCycle.BooleanOperate = false
	// 	}
	// 	handleEndOfDayAnalyticsOperations()
	// 	handleDayReset()
	// }

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
	//Declare tradeBoughtEvaluation
	// tradeBoughtEvaluation := TradeBoughtEvaluation{}

	listVal := reflect.ValueOf(params[0])
	var listSymbolsInterface interface{} = listVal.Index(0).Interface()
	listSymbols := listSymbolsInterface.([]string)
	symbol := listSymbols[0]
	// fmt.Println(listSymbols[0])

	//Query
	//TradeBoughtEvaluation
	tradeBoughtEvaluation := queryIsTradeCompleted(symbol)

	//
	holdingWisemen := HoldingWisemen{}
	if tradeBoughtEvaluation.IsBought {
		for i, v := range tradeBoughtEvaluation.HoldingList {
			if v.Symbol == symbol {
				// type HoldingWisemen struct {
				// 	Symbol      string
				// 	Price       string
				// 	Qty         string
				// 	QtyBought   string
				// 	OrderStatus string
				// }
				holdingWisemen = HoldingWisemen{Symbol: symbol, Price: v.PurchasePrice, Qty: "0", QtyBought: v.Qty, OrderStatus: "pending eval"}
			}
			i++
		}
		insertHoldingWisemen(holdingWisemen)
	}

	//post to neo, holding information and qty.
	postNeoBuyOrderResponse(holdingWisemen)

	//insert information.

	//query order, if order still open pause until resolution.
	//Cycle but, will need buy simulation for this part.buy limit met, and then evaluation of order.

	// fmt.Println(isHoldingContained)
	// tradeBoughtEvaluation

	//store order

	//once bought
	//store order information

	//Evaluate
	// if strings.Contains(response, "<sym>") {
	// 	//Parse holdings, append to tradeBoughtEvaluation
	// 	// fmt.Println("hit true")
	// 	holdings := parseBalanceQuery(response)
	// 	tradeBoughtEvaluation.IsBought = true
	// 	tradeBoughtEvaluation.Holdings = holdings
	// }

	// //if positive evlauation store tradeBoughtEvaluation
	// if tradeBoughtEvaluation.IsBought {
	// 	//Store DB result
	// 	insertTradeBoughtEvaluation(tradeBoughtEvaluation)
	// 	operatingCycle := cycleMapPool["handleCheckIsTradeBought"]
	// 	cancelCycle(operatingCycle)
	// }

	// TradeBoughtEvaluation
}

func checkWhaleDelimiterMet() bool {
	isWhaleDelimiterMet := false
	symbolList := selectWhaleSymbolHold()
	if len(symbolList) >= 200 {
		isWhaleDelimiterMet = true
	}
	return isWhaleDelimiterMet
}

func handleTSPRefresh(params ...interface{}) {
	//TSP
	var queryResponse = queryTSP()
	//Top 3 stocks
	stockList := parseTopStockQuery(queryResponse)
	topRankList := []Stock{}

	//Store of symbols will affect both wisemen and whale.
	//Temp,
	//if symbols do not exist in set add to monitorSymbol
	for indexStock, stock := range stockList {
		if len(topRankList) < 3 {
			if strings.Contains(stock.Symbol, ".") == false {
				//compare price, proceed to add until met...
				//pull wisemen metrics, price delimiter,
				metricsWisemen := selectMetricsWisemen()[0]
				if stock.Last >= metricsWisemen.DesiredPriceRangeLow && stock.Last <= metricsWisemen.DesiredPriceRangeHigh {
					topRankList = append(topRankList, stock)
				}
			}

		}
		indexStock++
	}

	//Store for monitorList, handle temp hold and
	//Query monitorSymbol
	monitorList := selectTempSymbolHold()
	if len(monitorList) == 0 {
		for i, v := range topRankList {
			insertTempSymbolHold(v.Symbol, false)
			i++
		}
	}
	if len(monitorList) != 0 {
		boolStockMonitorMap := make(map[string]bool)
		// for i, v := range topRankList {
		for i, v := range topRankList {
			for i1, v1 := range monitorList {
				if v.Symbol == v1 {
					fmt.Println(v.Symbol)
					boolStockMonitorMap[v.Symbol] = true
					break
				}
				if i1 == (len(monitorList) - 1) {
					// fmt.Println("last symbol ", v.Symbol)
					boolStockMonitorMap[v.Symbol] = false
				}
			}
			i++
		}
		for k, v := range boolStockMonitorMap {
			// fmt.Printf("key[%s] value[%s]\n", k, v)
			if v == false {
				insertTempSymbolHold(k, false)
			}
		}
	}
	// processFillHolds()
	handleFillHolds()
	// stockRanking := topRankList[0].Symbol + "," + topRankList[1].Symbol + "," + topRankList[2].Symbol

	// for i,v:= range topRankList
	insertAnalyticsOperations(topRankList)
	//Query follow-crossover should be handled by concurrent monitor cycle.
}

func handleFillHolds(params ...interface{}) {
	tempSymbolHoldList := selectTempSymbolHold()

	whaleDelimiterMet := checkWhaleDelimiterMet()
	for i, tempSymbol := range tempSymbolHoldList {
		//condition meet if whale symbol or wisemen symbol already exists.

		isSymbolExistsInWisemen := false
		isSymbolExistsInWhale := false

		//select from wisemen
		wisemenSymbolList := selectWisemenSymbolHold()
		//select from whale
		whaleSymbolList := selectWisemenSymbolHold()

		//iterrate set isSymbolExistsInWiseMen
		for i, wisemenSymbol := range wisemenSymbolList {
			if wisemenSymbol == tempSymbol {
				isSymbolExistsInWisemen = true
			}
			i++
		}

		for i, whaleSymbol := range whaleSymbolList {
			if whaleSymbol == tempSymbol {
				isSymbolExistsInWhale = true
			}
			i++
		}

		if isSymbolExistsInWisemen == false {
			//insert for wisemen
			insertWisemenSymbolHold(tempSymbol, false)
		}
		if isSymbolExistsInWhale == false {
			//check process for whale
			if whaleDelimiterMet == false {
				insertWhaleSymbolHold(tempSymbol, false)
			}
			i++
		}

	}
	dropTempSymbolHold()
	createTempSymbolHold()
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
	//Query monitor_symbol
	symbolList := selectWhaleSymbolHold()
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
		stockList := selectStockWhale(symbol)
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
					insertStockWhale(stockBrokerage)
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
			insertStockWhale(appendedStock)
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

func handleDowWebscrape(params ...interface{}) {
	response := queryWebscrape()
	fmt.Println("hit awesome")
	currentDowValue := parseDowWebscrape(response)
	// fmt.Println(currentDowValue)
	// fmt.Println(pointsChanged)
	// fmt.Println(percentageChange)
	insertDow(currentDowValue)
}

func handleEndOfDayAnalyticsOperations() {
	day := getDayOfWeek()
	//insert into table conditional
	insertEndOfDayAnalyticsOperations(isMarketClosed, day.String())
}

func resetTempSymbolHold() {
	dropTempSymbolHold()
	createTempSymbolHold()
}
func resetStockWisemenSymbolHold() {
	dropWisemenSymbolHold()
	createWisemenSymbolHold()
}

func handleDayReset() {
	boolOperate1 = true
	boolOperate2 = true
	boolOperate3 = true
	boolOperate4 = true
	boolOperate5 = true
	boolOperate6 = true
	boolOperate7 = true
	boolOperate8 = true
	boolOperate9 = true
	boolOperate10 = true
	boolOperate11 = true
	boolOperate12 = true
	boolOperate13 = true
	boolOperate14 = true
	boolOperate15 = true
	boolOperate16 = true
	boolOperate17 = true
	boolOperate18 = true
	checkIsMarketOpenBool = true
	isMarketClosed = false
}

func setTimelineOperationsFalse() {
	boolOperate1 = false
	boolOperate2 = false
	boolOperate3 = false
	boolOperate4 = false
	boolOperate5 = false
	boolOperate6 = false
	boolOperate7 = false
	boolOperate8 = false
	boolOperate9 = false
	boolOperate10 = false
	boolOperate11 = false
	boolOperate12 = false
	boolOperate13 = false
	boolOperate14 = false
	boolOperate15 = false
	boolOperate16 = false
	boolOperate17 = false
	boolOperate18 = false
}
