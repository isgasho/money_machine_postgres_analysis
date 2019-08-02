package main

import "fmt"

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
//Retrieve 3-set stock query

func processQueryStockSet() {
	var stringListToQuery = []string{"AAPL", "AMD", "GE"}
	// queryMultiStockPull(stringListToQuery)

	// createCycle(3, 10, test, "dog", "frog", 2, false)
	// createCycle(3, 10, test, stringListToQuery)

	// createCycle(7, 1, queryMultiStockPull, stringListToQuery)
	// operatingCycle := cyclePool[0]
	// go startCycle(&operatingCycle)

	// queryTestStoreQueryMonitored

	createCycle(7, 1, queryStoreQueryMonitored, stringListToQuery)
	operatingCycle := cyclePool[0]
	go startCycle(&operatingCycle)

	// createCycle(7, 1, queryMonitoredStocks, stringListToQuery)
	// operatingCycle := cyclePool[0]
	// go startCycle(&operatingCycle)

	deleteMonitoredStock("Test1")
	deleteMonitoredStock("Test2")
	deleteMonitoredStock("Test3")

	fmt.Println(selectAllMonitoringStock())
}

//Iteration, subjectable to cancelling.

//Retrieve monitor stock list
//Stock symbol query list.

//Store Stock
