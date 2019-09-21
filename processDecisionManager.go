package main

import (
	"fmt"
	"strings"
	"time"
)

var checkIsMarketOpenMinute = 19
var checkIsMarketOpenHour = 14

var conditionOneMinute = 20
var conditionOneHour = 14

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

var conditionNineteenMinute = 1
var conditionNineteenHour = 16

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
	createCycle(5, 10000, handleTimelineConditionalTriggers, "handleTimelineConditionalTriggers")
	operatingCycle := cycleMapPool["handleTimelineConditionalTriggers"]
	go startCycle(operatingCycle)
}

func processWisemenQueryStockSet() {
	if initialWisemenStockQueryPerformed == true {
		fmt.Println("hit initialWisemenStockQueryPerformed == true")
		operatingCycle := cycleMapPool["handleWisemenQueryStockList"]
		// fmt.Println("operatingCycle.Name")
		// fmt.Println(operatingCycle.Name)
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

func processTSPRefresh() {
	go handleTSPRefresh()
}

func processFillHolds() {
	go handleFillHolds()
}

func processDowWebscrape() {
	go handleDowWebscrape()
}

func handleTimelineConditionalTriggers(params ...interface{}) {
	currentTime := time.Now()
	fmt.Println(currentTime.Hour())
	fmt.Println(currentTime.Minute())
	fmt.Println(currentTime.Second())
	fmt.Println(currentTime.Date())

	//
	//Timeline events
	//

	//Conditional operate
	if currentTime.Minute() == checkIsMarketOpenMinute && currentTime.Hour() == checkIsMarketOpenHour && checkIsMarketOpenBool {
		checKIsBrokerageResponding()

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
		processTSPRefresh()
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

		resetTempSymbolHold()
		resetStockWisemenSymbolHold()
		handleEndOfDayAnalyticsOperations()
		handleDayReset()
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

	//pull MonitorSymbol
	//if symbols do not exist in set add to monitorSymbol
	for i, v := range stockList {
		if len(topRankList) < 3 {
			if strings.Contains(v.Symbol, ".") == false {
				topRankList = append(topRankList, v)
			}
		}
		i++
	}

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

	//Here the temp hold will be filled and we can filter data to the other algorithm holds.
	//Do we want to select from temp hold? and then fill other holds?
	//Yes we want to fill holds retroactively.
	// for k, v:=
	processFillHolds()
	fmt.Println("topRankList")
	fmt.Println(topRankList)
	stockRanking := topRankList[0].Symbol + "," + topRankList[1].Symbol + "," + topRankList[2].Symbol
	insertAnalyticsOperations(stockRanking)
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
	//There's a disconnect here, from symbolList and stock
	//Really where do we add the symbol to symbolLis
	//If symbol in symbolList was added today.
	//Before enter stock, select all stock and see if symbol is present.

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
	stockList := parseStockSetQuery(queryResponse)
	noCalculationMatchFound := false
	isContinueCalculation := false
	//if symbol is presnt in list.
	//is symbol calculated for day.
	//if there is a symbol in the new stockList, that has not been calculated for,
	//meaning there is no match, noCalculationMatchFound,
	//if noCalculationMatchFound, then it needs to be accounted for and follow process for all stocks entering.
	listSymbolDayCalculatedWhale := []string{}
	for indexStock, stock := range stockList {
		isContinueCalculation = false
		for indexSymbolDayCalculated, symbolDayCalculated := range listSymbolDayCalculatedWhale {
			if stock.Symbol == symbolDayCalculated {
				isContinueCalculation = true
				break
			}
			indexSymbolDayCalculated++
		}
		if isContinueCalculation == false {
			noCalculationMatchFound = true
			break
		}
		indexStock++
	}

	//
	if noCalculationMatchFound == false {
		for i, stock := range stockList {
			insertStockWhale(stock)
			i++
		}
	}
	//
	if noCalculationMatchFound {
		isPresentInDB := false
		isDaySame := false
		//Select stock once
		stockListQueried := selectAllStockWhale()
		for indexStock, stock := range stockList {
			//iterate through each stock to be entered, and see if stockList matches.
			//If match then add as normal,
			isPresentInDB = false
			isDaySame = false
			for indexStockInDB, stockInDB := range stockListQueried {
				if stockInDB.Symbol == stock.Symbol {
					//compare
					isStocksComparedSameTimeStamps := compareTimeStamps(stock, stockInDB)
					if isStocksComparedSameTimeStamps {
						insertStockWhale(stock)
						isPresentInDB = true
						isDaySame = true
						break
					}
				}
				indexStockInDB++
			}
			//Is conditional that stock in DB is present. If stock in in DB is not present skip "isDaySame" process.
			//if no match, then special process, add symbol dynamically to list.
			if isPresentInDB == false || isDaySame == false {
				appendedStock := processAppendDayOfWeekToStock(stock)
				insertStockWhale(appendedStock)
			}
			indexStock++
		}
	}
}

func compareTimeStamps(stock1 Stock, stock2 Stock) bool {
	isTimeStampMatch := true
	//Split for day
	// 2019-09-19 17:47:10.944343-06
	stock1Split1 := strings.Split(stock1.CreatedAt, ":")[0]
	stock1Split2 := strings.Split(stock1Split1, "-")[1]
	day1 := strings.Split(stock1Split2, " ")[0]

	stock2Split1 := strings.Split(stock2.CreatedAt, ":")[0]
	stock2Split2 := strings.Split(stock2Split1, "-")[1]
	day2 := strings.Split(stock2Split2, " ")[0]

	//compare day,
	if day1 != day2 {
		isTimeStampMatch = false
	}
	return isTimeStampMatch
}
func processAppendDayOfWeekToStock(stock Stock) Stock {
	//handle on first field available value, append day of week.
	//get handle on day of week.
	// dayOfWeek := time.Date.
	instanceStock := stock
	currentTime := time.Now()
	instanceStock.Vl = "99.9"
	instanceStock.Vl = instanceStock.Vl + " " + currentTime.Weekday().String()
	fmt.Println(instanceStock.Vl)

	return instanceStock
}

func handleDowWebscrape(params ...interface{}) {
	response := queryWebscrape()
	fmt.Println(response)
	currentDowValue, pointsChanged, percentageChange := parseDowWebscrape(response)
	insertDow(currentDowValue, pointsChanged, percentageChange)
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
