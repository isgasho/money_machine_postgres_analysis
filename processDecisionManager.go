package main

import (
	"fmt"
	"strings"
	"time"
)

//Take top 3 stocks store them
func storeStocks() {
}

//If stocks don't already exist in monitoring add them.
func handleMonitoring() {
}

func checkIfMonitoring() {
	//Get all stocks being monitored
	//queryParser
}

//Handle variant processes and entry processes.
//Retrieve TSP

//type fn func(params ...interface{})

func processTimelineStart() {
	createCycle(5, 10000, handleTimelineConditionalTriggers)
	operatingCycle := cyclePool[0]
	go startCycle(&operatingCycle)
}

func processQueryStockSet() {
	// var stringListToQuery = []string{"AAPL", "AMD", "GE"}
	// queryMultiStockPull(stringListToQuery)

	// createCycle(3, 10, test, "dog", "frog", 2, false)
	// createCycle(3, 10, test, stringListToQuery)

	// createCycle(7, 1, handleTopStockPull, stringListToQuery)
	// operatingCycle := cyclePool[0]
	// go startCycle(&operatingCycle)

	// createCycle(7, 1, queryMultiStockPull, stringListToQuery)
	// operatingCycle := cyclePool[0]
	// go startCycle(&operatingCycle)

	// queryTestStoreQueryMonitored

	// createCycle(7, 1, queryStoreQueryMonitored, stringListToQuery)
	// operatingCycle := cyclePool[0]
	// go startCycle(&operatingCycle)

	// createCycle(7, 1, queryMonitoredStocks, stringListToQuery)
	// operatingCycle := cyclePool[0]
	// go startCycle(&operatingCycle)

	fmt.Println("cyclepool", cyclePool)
	createCycle(3, 10000, handleQueryStockList)
	operatingCycle := cyclePool[1]
	go startCycle(&operatingCycle)

	// createCycle(7, 1, handleQueryStockList)
	// operatingCycle := cyclePool[0]
	// go startCycle(&operatingCycle)

	// deleteMonitoredStock("Test1")
	// deleteMonitoredStock("Test2")
	// deleteMonitoredStock("Test3")

	// fmt.Println(selectAllMonitoringStock())
}
func processTSPRefresh() {
	createCycle(7, 1, handleTSPRefresh)
	operatingCycle := cyclePool[0]
	go startCycle(&operatingCycle)
}

var baselineSecond int
var baselineMinute int
var calculatingMinute int

var conditionOneSecond int
var conditionOneMinute int
var conditionOneHour int

var conditionTwoSecond int
var conditionTwoMinute int
var conditionTwoHour int

var conditionThreeSecond int
var conditionThreeMinute int
var conditionThreeHour int

var boolOperate1 = true
var boolOperate2 = true
var boolOperate3 = true
var timelineOperationIndex = 0

func handleTimelineConditionalTriggers(params ...interface{}) {
	currentTime := time.Now()
	fmt.Println(currentTime.Hour())
	fmt.Println(currentTime.Minute())
	fmt.Println(currentTime.Second())

	if timelineOperationIndex == 0 {
		baselineMinute = currentTime.Minute()
		// calculatingMinute = baselineMinute + 1

		conditionOneSecond = 5
		conditionOneMinute = baselineMinute + 1
		conditionOneHour = currentTime.Hour()

		conditionTwoSecond = 25
		conditionTwoMinute = baselineMinute + 1
		conditionTwoHour = currentTime.Hour()

		// conditionThreeSecond = 25
		// conditionThreeMinute = baselineMinute + 1
		// conditionThreeHour = currentTime.Hour()

		timelineOperationIndex++
	}
	//
	//Timeline events
	//
	//Initial monitoring pool
	if currentTime.Second() >= conditionOneSecond && currentTime.Minute() >= conditionOneMinute && currentTime.Hour() >= conditionOneHour && boolOperate1 {
		fmt.Println("hit1")
		boolOperate1 = false
		handleTSPRefresh()
		// processQueryStockSet()
		// fmt.Println(conditionOneSecond)
	}
	if currentTime.Second() >= conditionTwoSecond && currentTime.Minute() >= conditionTwoMinute && currentTime.Hour() >= conditionTwoHour && boolOperate2 {
		fmt.Println("hit2")
		boolOperate2 = false
		// handleTSPRefresh()
		// handleQueryStockList()
	}
	// if currentTime.Second() >= conditionThreeSecond && currentTime.Minute() >= conditionThreeMinute && currentTime.Hour() >= conditionThreeHour && boolOperate3 {
	// 	fmt.Println("hit3")
	// 	boolOperate3 = false
	// 	handleTSPRefresh()
	// }

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

			// fmt.Println(i1)
			// fmt.Println(len(monitorList))
			// fmt.Println((len(monitorList) - 1))

			if i1 == (len(monitorList) - 1) {
				fmt.Println("last symbol ", v.Symbol)
				boolStockMonitorMap[v.Symbol] = false
			}
			// i1++
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
		if v {
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

// func handleTopStockPull(params ...interface{}) {
// 	//
// 	//Monitor_Symbol and Analytics_Operation store
// 	//
// 	var queryResponse = queryTSP()
// 	//Top 3 stocks
// 	stockList := parseTopStockQuery(queryResponse)
// 	topRankList := []Stock{}
// 	//Store in symbol table
// 	for i, v := range stockList {
// 		if i < 3 {
// 			topRankList = append(topRankList, v)
// 		}
// 		insertMonitorSymbol(v)
// 	}
// 	stockRanking := topRankList[0].Symbol + "," + topRankList[1].Symbol + "," + topRankList[2].Symbol
// 	insertAnalyticsOperations(stockRanking)
// }

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
	fmt.Println(formattedSymbolList)

	//Form query to Node -> Brokerage
	queryResponse := queryMultiStockPull(formattedSymbolList)
	stockList := parseStockSetQuery(queryResponse)

	//Store Stocks in DB
	for i, v := range stockList {
		fmt.Println(v.Symbol)
		insertStock(v)
		i++
	}
}

//Iteration, subjectable to cancelling.

//Retrieve monitor stock list
//Stock symbol query list.

//Store Stock
