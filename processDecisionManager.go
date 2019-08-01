package main

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

	createCycle(7, 10, queryMultiStockPull, stringListToQuery)
	operatingCycle := cyclePool[0]
	go startCycle(&operatingCycle)
}

//Iteration, subjectable to cancelling.

//Retrieve monitor stock list
//Stock symbol query list.

//Store Stock
