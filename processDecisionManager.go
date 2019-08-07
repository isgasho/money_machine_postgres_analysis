package main

import (
	"fmt"
	"strings"
	"time"
)

var queryCycle Cycle

func checkIfMonitoring() {
	//Get all stocks being monitored
	//queryParser
}

func processTimelineStart() {
	createCycle(5, 10000, handleTimelineConditionalTriggers)
	operatingCycle := cyclePool[0]
	go startCycle(&operatingCycle)
}

func processQueryStockSet() {
	fmt.Println("cyclepool", cyclePool)
	createCycle(3, 100000, handleQueryStockList)
	queryCycle = cyclePool[1]
	go startCycle(&queryCycle)
}

// var baselineSecond int
// var baselineMinute int
// var calculatingMinute int

var conditionOneMinute = 42
var conditionOneHour = 12

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

var conditionNineteenMinute = 43
var conditionNineteenHour = 12

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

var timelineOperationIndex = 0

func handleTimelineConditionalTriggers(params ...interface{}) {
	currentTime := time.Now()
	fmt.Println(currentTime.Hour())
	fmt.Println(currentTime.Minute())
	fmt.Println(currentTime.Second())

	if timelineOperationIndex == 0 {
		// baselineMinute = currentTime.Minute()
		// // calculatingMinute = baselineMinute + 1

		// conditionOneSecond = 5
		// conditionOneMinute = baselineMinute + 1
		// conditionOneHour = currentTime.Hour()

		// conditionTwoSecond = 25
		// conditionTwoMinute = baselineMinute + 1
		// conditionTwoHour = currentTime.Hour()

		// conditionThreeSecond = 25
		// conditionThreeMinute = baselineMinute + 1
		// conditionThreeHour = currentTime.Hour()
		timelineOperationIndex++
	}

	//
	//Timeline events
	//
	//Initial monitoring pool
	if currentTime.Minute() == conditionOneMinute && currentTime.Hour() == conditionOneHour && boolOperate1 {
		fmt.Println("hit1 init operations")
		boolOperate1 = false
		boolOperate19 = true
		handleTSPRefresh()
		processQueryStockSet()
	}
	if currentTime.Minute() == conditionTwoMinute && currentTime.Hour() == conditionTwoHour && boolOperate2 {
		fmt.Println("hit2")
		boolOperate2 = false
		handleTSPRefresh()
	}
	if currentTime.Minute() == conditionThreeMinute && currentTime.Hour() == conditionThreeHour && boolOperate3 {
		fmt.Println("hit3")
		boolOperate3 = false
		handleTSPRefresh()
	}
	if currentTime.Minute() == conditionFourMinute && currentTime.Hour() == conditionFourHour && boolOperate4 {
		fmt.Println("hit4")
		boolOperate4 = false
	}
	if currentTime.Minute() == conditionFiveMinute && currentTime.Hour() == conditionFiveHour && boolOperate5 {
		fmt.Println("hit5")
		boolOperate5 = false
		handleTSPRefresh()
	}
	if currentTime.Minute() == conditionSixMinute && currentTime.Hour() == conditionSixHour && boolOperate6 {
		fmt.Println("hit6")
		boolOperate6 = false
		handleTSPRefresh()
	}
	if currentTime.Minute() == conditionSevenMinute && currentTime.Hour() == conditionSevenHour && boolOperate7 {
		fmt.Println("hit7")
		boolOperate7 = false
		handleTSPRefresh()
	}
	if currentTime.Minute() == conditionEightMinute && currentTime.Hour() == conditionEightHour && boolOperate8 {
		fmt.Println("hit8")
		boolOperate8 = false
		handleTSPRefresh()
	}
	if currentTime.Minute() == conditionNineMinute && currentTime.Hour() == conditionNineHour && boolOperate9 {
		fmt.Println("hit9")
		boolOperate9 = false
		handleTSPRefresh()
	}
	if currentTime.Minute() == conditionTenMinute && currentTime.Hour() == conditionTenHour && boolOperate10 {
		fmt.Println("hit10")
		boolOperate10 = false
		handleTSPRefresh()
	}
	if currentTime.Minute() == conditionElevenMinute && currentTime.Hour() == conditionElevenHour && boolOperate11 {
		fmt.Println("hit11")
		boolOperate11 = false
		handleTSPRefresh()
	}
	if currentTime.Minute() == conditionTwelveMinute && currentTime.Hour() == conditionTwelveHour && boolOperate12 {
		fmt.Println("hit12")
		boolOperate12 = false
		handleTSPRefresh()
	}
	if currentTime.Minute() == conditionThirteenMinute && currentTime.Hour() == conditionThirteenHour && boolOperate13 {
		fmt.Println("hit13")
		boolOperate13 = false
		handleTSPRefresh()
	}
	if currentTime.Minute() == conditionFourteenMinute && currentTime.Hour() == conditionFourteenHour && boolOperate14 {
		fmt.Println("hit14")
		boolOperate14 = false
		handleTSPRefresh()
	}
	if currentTime.Minute() == conditionFifteenMinute && currentTime.Hour() == conditionFifteenHour && boolOperate15 {
		fmt.Println("hit15")
		boolOperate15 = false
		handleTSPRefresh()
	}
	if currentTime.Minute() == conditionSixteenMinute && currentTime.Hour() == conditionSixteenHour && boolOperate16 {
		fmt.Println("hit16")
		boolOperate16 = false
		handleTSPRefresh()
	}
	if currentTime.Minute() == conditionSeventeenMinute && currentTime.Hour() == conditionSeventeenHour && boolOperate17 {
		fmt.Println("hit17")
		boolOperate17 = false
		handleTSPRefresh()
	}
	if currentTime.Minute() == conditionEighteenMinute && currentTime.Hour() == conditionEighteenHour && boolOperate18 {
		fmt.Println("hit18")
		boolOperate18 = false
		handleTSPRefresh()
	}
	if currentTime.Minute() == conditionNineteenMinute && currentTime.Hour() == conditionNineteenHour && boolOperate19 {
		fmt.Println("hit19")
		boolOperate19 = false
		handleDayReset()
		fmt.Println("length: ", len(cyclePool))
		// operatingCycle := &cyclePool[1]
		// operatingCycle.BooleanOperate = false
		// operatingCycle := &cyclePool[1]
		queryCycle.BooleanOperate = false

		// cancelCycle(operatingCycle)
	}
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
		if i < 3 {
			topRankList = append(topRankList, v)
		}
	}

	//Query monitorSymbol
	monitorList := selectMonitorSymbol()

	boolStockMonitorMap := make(map[string]bool)

	fmt.Println("hit before stocklist")
	// for i, v := range topRankList {
	for i, v := range stockList {
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
	// fmt.Println("this are", boolStockMonitorList)

	//Insert symbol into monitor table if it doesn't exist there
	// for i, v := range boolStockMonitorMap {
	// 	if v == false {
	// 		// fmt.Println(topRankList[i].Symbol)
	// 		fmt
	// 		insertMonitorSymbol(topRankList[i])
	// 	}
	// }

	for k, v := range boolStockMonitorMap {
		// fmt.Printf("key[%s] value[%s]\n", k, v)
		if v == false {
			insertMonitorSymbol(k, false)
		}
	}

	// deleteMonitorSymbol("CMG")
	// deleteMonitorSymbol("CHE")
	// deleteMonitorSymbol("GRUB")

	stockRanking := topRankList[0].Symbol + "," + topRankList[1].Symbol + "," + topRankList[2].Symbol
	insertAnalyticsOperations(stockRanking)
	//Query follow-crossover should be handled by concurrent monitor cycle.
}

func handleQueryStockList(params ...interface{}) {
	//Query monitor_symbol
	symbolList := selectMonitorSymbol()

	//Parse format errors in symbols
	formattedSymbolList := []string{}
	for i, v := range symbolList {
		if strings.Contains(v, ".") {
			continue
		}
		formattedSymbolList = append(formattedSymbolList, v)
		i++
	}
	// fmt.Println(formattedSymbolList)

	//Form query to Node -> Brokerage
	queryResponse := queryMultiStockPull(formattedSymbolList)
	stockList := parseStockSetQuery(queryResponse)

	//Store Stocks in DB
	for i, v := range stockList {
		insertStock(v)
		i++
	}
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
}
