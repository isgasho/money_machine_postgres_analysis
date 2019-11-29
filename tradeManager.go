package main

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"
)

func handleTradeWisemen(symbol string, limitPrice string) {
	fmt.Println("limitPrice")
	fmt.Println(limitPrice)
	desiredLimitPrice := 0.0
	if s, err := strconv.ParseFloat(limitPrice, 64); err == nil {
		desiredLimitPrice = s
	}
	fmt.Println("desiredLimitPrice")
	fmt.Println(desiredLimitPrice)
	//get balance
	response := queryBalance()
	balance := parseBalance(response)
	floatBalance := 0.0
	if s, err := strconv.ParseFloat(balance, 64); err == nil {
		floatBalance = s
	}
	fmt.Println("floatBalance")
	fmt.Println(floatBalance)

	//calculate qty to buy
	qty := calculateAmountOfStockToBuy(desiredLimitPrice, floatBalance)
	fmt.Println("before rounding down")
	fmt.Println(qty)
	qtyInt := roundDown(qty)
	fmt.Println("After rounding down")
	fmt.Println(qtyInt)
	// stringBalance := fmt.Sprintf("%f", floatBalance)
	stringPrice := fmt.Sprintf("%f", desiredLimitPrice)
	// stringQty := fmt.Sprintf("%f", qtyInt)
	stringQty := strconv.Itoa(qtyInt)
	//store trade entered information
	// informationAtTrade := InformationAtTrade{
	// 	Symbol: symbol,
	// }
	// handleInsertInformationAtTrade(symbol, "limit", "buy", "1.00")
	// insertTradeEnteredInformation(tradeEnteredInformation)

	// insertInformationAtTrade(informationAtTrade)

	//Submit buy limit to brokerage
	fmt.Println("symbol")
	fmt.Println(symbol)
	fmt.Println("stringQty")
	fmt.Println(stringQty)
	fmt.Println("stringPrice")
	fmt.Println(stringPrice)
	queryTradeBuyLimit(symbol, stringPrice, "1")
}

func monitorSell(params ...interface{}) {
	listVal := reflect.ValueOf(params[0])
	var listSymbolsInterface interface{} = listVal.Index(0).Interface()
	listStrings := listSymbolsInterface.([]string)

	metrics := selectMetricsWisemen()[0]
	symbol := listStrings[0]
	priceDrop := listStrings[1]
	timeDelimiter := metrics.SellTime

	// calculateIsTimeDelimiterMetSell()
	// calculateIsTimeDelimiterMetSell
	// monitorSell()

	isSymbolPresentInHolding := calculateIsSymbolPresentInHolding(symbol)
	isDropPriceMet := calculateIsDropPriceMet(symbol, priceDrop)
	isTimeDelimiterMet := calculateIsTimeDelimiterMetSell(timeDelimiter)

	//Is holding is not present cancel cycle process
	//and move to wrap up.
	if isSymbolPresentInHolding == false {
		//cancel cycle
		//support nodemailerror
		//handle on history...
		alteredTransactionHistory := calculateTransactionHistory(TransactionHistory{Symbol: symbol})
		fmt.Println("alteredTransactionHistory.HistoryValueList")
		fmt.Println(alteredTransactionHistory.HistoryValueList)
		handleInsertInformationAtTrade(symbol, "limit", "sell", alteredTransactionHistory.HistoryValueList[1].Qty)
		transactionHistory := TransactionHistory{Symbol: symbol}
		wrapUpWisemenOutcome(transactionHistory)
		operatingCycle := cycleMapPool["monitorSell"]
		cancelCycle(operatingCycle)
	}

	if isTimeDelimiterMet {
		//if time delimiter met sell at market immediately
		holding := HoldingWisemen{}
		holdingList := getAllHolding()
		for i, v := range holdingList.ListHolding {
			if v.Symbol == symbol {
				// isBoolReturning = "true"
				holding = v
			}
			fmt.Println(v)
			i++
		}
		orderList := getAllOrders()
		fmt.Println("len(orderList.ListOrders)")
		fmt.Println(len(orderList.ListOrders))
		order := Order{}
		for i, v := range orderList.ListOrders {
			if v.Symbol == symbol {
				order = v
			}
			i++
		}
		queryCancelOrder(order.SVI)
		//cancel order, pause 10 seconds
		time.Sleep(time.Duration(10) * time.Second)
		queryTradeSellMarket(holding)
		handleInsertInformationAtTrade(symbol, "time delimiter", "sell", holding.Qty)
		time.Sleep(time.Duration(30) * time.Second)
		transactionHistory := TransactionHistory{Symbol: symbol}
		wrapUpWisemenOutcome(transactionHistory)
		operatingCycle := cycleMapPool["monitorSell"]
		cancelCycle(operatingCycle)
	}

	if isDropPriceMet {
		//if drop price met sell at market immediately
		holding := HoldingWisemen{}
		holdingList := getAllHolding()
		for i, v := range holdingList.ListHolding {
			if v.Symbol == symbol {
				// isBoolReturning = "true"
				holding = v
			}
			fmt.Println(v)
			i++
		}
		orderList := getAllOrders()
		fmt.Println("len(orderList.ListOrders)")
		fmt.Println(len(orderList.ListOrders))
		order := Order{}
		for i, v := range orderList.ListOrders {
			if v.Symbol == symbol {
				order = v
			}
			i++
		}
		queryCancelOrder(order.SVI)
		//cancel order, pause 10 seconds
		time.Sleep(time.Duration(10) * time.Second)
		queryTradeSellMarket(holding)
		//handle on holding
		handleInsertInformationAtTrade(symbol, "drop loss", "sell", holding.Qty)
		time.Sleep(time.Duration(30) * time.Second)
		transactionHistory := TransactionHistory{Symbol: symbol}
		wrapUpWisemenOutcome(transactionHistory)
		operatingCycle := cycleMapPool["monitorSell"]
		cancelCycle(operatingCycle)
	}
}
func handleInsertInformationAtTrade(symbol string, typeTrade string, side string, qty string) {
	//query stock
	response := queryMultiStockPull([]string{symbol})
	stockList := parseStockSetQuery(response)
	bid := stockList[0].Bid
	ask := stockList[0].Ask
	last := stockList[0].Last
	//dow
	// dow := handleDowWebscrape()
	dow := "26,000"
	//date
	year, month, day := getDate()
	yearString := strconv.Itoa(year)
	monthString := strconv.Itoa(month)
	dayString := strconv.Itoa(day)
	//time
	hour := getCurrentHour()
	minute := getCurrentMinute()
	hourString := strconv.Itoa(hour)
	minuteString := strconv.Itoa(minute)
	//instantiate
	informationAtTrade := InformationAtTrade{
		Symbol:    symbol,
		TypeTrade: typeTrade,
		Side:      side,
		Qty:       qty,
		Year:      yearString,
		Month:     monthString,
		Day:       dayString,
		Hour:      hourString,
		Minute:    minuteString,
		Dow:       dow,
		Bid:       bid,
		Ask:       ask,
		Last:      last,
	}
	fmt.Println("informationAtTrade")
	fmt.Println(informationAtTrade)
	//insert
	insertInformationAtTrade(informationAtTrade)
}
func calculateIsDropPriceMet(symbol string, dropPriceString string) bool {
	listSymbol := []string{symbol}
	isBoolReturning := false

	holding := HoldingWisemen{}
	holdingList := getAllHolding()
	for i, v := range holdingList.ListHolding {
		if v.Symbol == symbol {
			// isBoolReturning = "true"
			holding = v
		}
		fmt.Println(v)
		i++
	}

	response := queryMultiStockPull(listSymbol)
	stockList := parseStockSetQuery(response)
	priceFromQuery := stockList[0].Last
	//Support for bid vs ask variance from last.

	metrics := selectMetricsWisemen()[0]
	//Here we need a handle on drop metrics,...
	//Partial bind...
	metricPchgDrop := metrics.PriceLowPchgTrade
	floatPriceFromQuery := 0.0
	floatMetricPchgDrop := 0.0
	if s, err := strconv.ParseFloat(priceFromQuery, 64); err == nil {
		floatPriceFromQuery = s
	}
	if s, err := strconv.ParseFloat(metricPchgDrop, 64); err == nil {
		floatMetricPchgDrop = s
	}

	holdingPrice := 0.0

	if s, err := strconv.ParseFloat(holding.Price, 64); err == nil {
		holdingPrice = s
	}
	dropPrice := holdingPrice - (holdingPrice * floatMetricPchgDrop)

	//compare drop price to holding price.

	fmt.Println("dropPrice")
	fmt.Println(dropPrice)

	fmt.Println("dropPrice")
	fmt.Println(dropPrice)

	fmt.Println("floatPriceFromQuery")
	fmt.Println(floatPriceFromQuery)

	if floatPriceFromQuery <= dropPrice {
		// sellAtMarket
		isBoolReturning = true
	}
	return isBoolReturning
}

func calculateIsTimeDelimiterMetSell(timeDelimiter string) bool {
	isTimeDelimiterMet := false
	//get current hour
	currentHour := getCurrentHour()
	//get current hour
	currentMinute := getCurrentMinute()

	fmt.Println("currentHour")
	fmt.Println(currentHour)

	fmt.Println("currentMinute")
	fmt.Println(currentMinute)

	stringCurrentHour := strconv.Itoa(currentHour)
	stringCurrentMinute := strconv.Itoa(currentMinute)

	//append 0 if current minute is single digit
	if len(stringCurrentMinute) == 1 {
		stringCurrentMinute = "0" + stringCurrentMinute
	}

	timeCompositeCurrentString := stringCurrentHour + stringCurrentMinute
	fmt.Println("timeDelimiter")
	fmt.Println(timeDelimiter)
	//compare time to delimiter

	fmt.Println("timeCompositeCurrentString")
	fmt.Println(timeCompositeCurrentString)
	// fmt.Println(timeDelimiter)
	if timeCompositeCurrentString == timeDelimiter {
		isTimeDelimiterMet = true
	}
	return isTimeDelimiterMet
}

// func removeElement(listEntered []int, val int) []int {
// 	var i int
// 	listAltered := listEntered
// 	for {
// 		if i == len(listAltered) {
// 			break
// 		}

// 		if listAltered[i] == val {
// 			listAltered = listAltered[:i+copy(listAltered[i:], listAltered[i+1:])]
// 			i = 0
// 		}
// 		i++
// 	}
// 	return listAltered
// }

func handleTimeDelimiterMetSell() {

}

func handleDropPriceMet() {

}
func calculateIsSymbolPresentInHolding(symbol string) bool {
	isBoolReturning := false
	holdingList := getAllHolding()
	for i, v := range holdingList.ListHolding {
		if v.Symbol == symbol {
			isBoolReturning = true
		}
		fmt.Println(v)
		i++
	}
	return isBoolReturning
}

func handleHistoryDayListArbitration(symbol string) []HistoryValue {
	listMatchingSymbolHistoryValue := []HistoryValue{}
	response := queryHistory()
	historyList := parseHistory(response)
	listHistoryValues := createListHistoryValuesForWisemen(historyList)
	// fmt.Println("listHistoryValues")
	// fmt.Println(listHistoryValues)
	//date
	yearCurrent, monthCurrent, dayCurrent := getDate()
	// dayCurrent = 8
	//sort values by day...
	//store values of today only...
	for i, v := range listHistoryValues {
		i++
		year := strings.Split(v.Date, " ")[0]
		month := strings.Split(v.Date, " ")[1]
		day := strings.Split(v.Date, " ")[2]
		intYear, err := strconv.Atoi(year)
		intMonth, err := strconv.Atoi(month)
		intDay, err := strconv.Atoi(day)
		if err != nil {
			fmt.Println(err)
		}
		//get trades done today...
		// fmt.Println("intYear")
		// fmt.Println(intYear)
		// fmt.Println("yearCurrent")
		// fmt.Println(yearCurrent)
		// fmt.Println("intMonth")
		// fmt.Println(intMonth)
		// fmt.Println("intDay")
		// fmt.Println(intDay)
		// fmt.Println("dayCurrent")
		// fmt.Println(dayCurrent)
		// fmt.Println("v.Symbol")
		// fmt.Println(v.Symbol)
		if intYear == yearCurrent {
			// fmt.Println("hit...o")
			if intMonth == monthCurrent {
				// fmt.Println("hit...i")
				if intDay == dayCurrent {
					// fmt.Println("hit...")
					if v.Symbol == symbol {
						// fmt.Println("values in")
						// fmt.Println(v)
						listMatchingSymbolHistoryValue = append(listMatchingSymbolHistoryValue, v)
					}
				}
			}
		}
	}
	return listMatchingSymbolHistoryValue
}

func handleInformationAtTradeDayListArbitration(symbol string) []InformationAtTrade {
	listMatchingSymbolInformationAtTrade := []InformationAtTrade{}
	listInformationAtTrade := selectInformationAtTrade()
	// for i, v := range listInformationAtTrade {

	// 	// listInformationAtTrade[i].CreatedAt = formatCreatedAtFromInformationAtTrade(v.CreatedAt)
	// }
	fmt.Println("listInformationAtTrade")
	fmt.Println(listInformationAtTrade)
	//date
	yearCurrent, monthCurrent, dayCurrent := getDate()

	//dayCurrent = 8
	//sort values by day...
	//store values of today only...
	for i, v := range listInformationAtTrade {
		i++
		// year := strings.Split(v.Year, " ")[0]
		// month := strings.Split(v.Month, " ")[1]
		// day := strings.Split(v.Day, " ")[2]
		intYear, err := strconv.Atoi(v.Year)
		intMonth, err := strconv.Atoi(v.Month)
		intDay, err := strconv.Atoi(v.Day)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("yearCurrent")
		fmt.Println(yearCurrent)
		fmt.Println("intYear")
		fmt.Println(intYear)

		fmt.Println("monthCurrent")
		fmt.Println(monthCurrent)
		fmt.Println("intMonth")
		fmt.Println(intMonth)

		fmt.Println("dayCurrent")
		fmt.Println(dayCurrent)
		fmt.Println("intDay")
		fmt.Println(intDay)

		// fmt.Println(year)
		//get trades completed this day
		if intYear == yearCurrent {
			if intMonth == monthCurrent {
				if intDay == dayCurrent {
					fmt.Println("hit internal")
					fmt.Println("v.Symbol")
					fmt.Println(v.Symbol)

					fmt.Println("symbol")
					fmt.Println(symbol)
					if v.Symbol == symbol {
						fmt.Println("values in")
						fmt.Println(v)
						listMatchingSymbolInformationAtTrade = append(listMatchingSymbolInformationAtTrade, v)
					}
				}
			}
		}
	}
	return listMatchingSymbolInformationAtTrade
}

func calculateTransactionHistory(transactionHistory TransactionHistory) TransactionHistory {
	alteredTransactionHistory := transactionHistory
	listValuesMatchedHistoryDayList := handleHistoryDayListArbitration(alteredTransactionHistory.Symbol)
	//if condition: a sell exists, from previous day and algorithm with the same symbol, before buy...
	//remove sell from consideration index...
	//at this point trading is completed and wrap up engaged.
	buySellList := []HistoryValue{}

	fmt.Println("listValuesMatchedHistoryDayList")
	fmt.Println(listValuesMatchedHistoryDayList)
	for i, v := range listValuesMatchedHistoryDayList {
		if i == 0 {
			if v.Side == "-1" {
				continue
			}
		}
		buySellList = append(buySellList, v)
	}
	fmt.Println("buySellList")
	fmt.Println(buySellList)
	//future support for multiple trade in same day of algorithm.
	if len(buySellList) != 0 {
		alteredTransactionHistory.HistoryValueList = append(alteredTransactionHistory.HistoryValueList, buySellList[0])
		alteredTransactionHistory.HistoryValueList = append(alteredTransactionHistory.HistoryValueList, buySellList[1])
	}
	return alteredTransactionHistory
}

func calculateBuyHistoryMatchesSymbol(historyValue HistoryValue, symbol string, intervalInList string) HistoryValue {
	isBoolResult := "false"

	alteredHistoryValue := historyValue
	if historyValue.Symbol == symbol {
		if historyValue.Side == "1" {
			alteredHistoryValue.IntervalInList = intervalInList
			isBoolResult = "true"
		}
	}
	alteredHistoryValue.IsCalculationTrue = isBoolResult
	return alteredHistoryValue
}

func calculateSellHistoryMatchesSymbol(historyValue HistoryValue, symbol string, intervalInList string) HistoryValue {
	isBoolResult := "false"

	alteredHistoryValue := historyValue
	if historyValue.Symbol == symbol {
		if historyValue.Side == "2" {
			alteredHistoryValue.IntervalInList = intervalInList
			isBoolResult = "true"
		}
	}
	alteredHistoryValue.IsCalculationTrue = isBoolResult
	return alteredHistoryValue
}

func createListHistoryValuesForWisemen(listHistory []string) []HistoryValue {
	listHistoryValues := []HistoryValue{}
	listHistoryFiltered := listHistory[:len(listHistory)-1]

	for i, v := range listHistoryFiltered {
		fmt.Println(i)
		fmt.Println("v")
		fmt.Println(v)
		isTrade := false
		if strings.Contains(v, ">Trade<") {
			isTrade = true
		}
		if isTrade {
			//sym
			symQuery := strings.Split(v, "</sym>")[0]
			symFromHistory := strings.Split(symQuery, "<sym>")[1]

			//date
			dateQuery := strings.Split(v, "</date>")[0]
			dateFromHistory := strings.Split(dateQuery, "<date>")[1]
			//parse date to usuable format
			date := formatDateFromHistory(dateFromHistory)

			//side
			sideQuery := strings.Split(v, "</side>")[0]
			sideFromHistory := strings.Split(sideQuery, "<side>")[1]

			//parse qty
			quantityQuery := strings.Split(v, "</quantity>")[0]
			quantityFromHistory := strings.Split(quantityQuery, "<quantity>")[1]

			//parse price
			priceQuery := strings.Split(v, "</price>")[0]
			priceFromHistory := strings.Split(priceQuery, "<price>")[1]

			//HistoryValue
			historyValue := HistoryValue{Symbol: symFromHistory, Date: date, Side: sideFromHistory, Qty: quantityFromHistory, Price: priceFromHistory}
			listHistoryValues = append(listHistoryValues, historyValue)
		}
	}
	return listHistoryValues
}

// isDelimiterNeeded := false
// delimiter := 0
// fmt.Println("listHistoryFiltered")
// fmt.Println(listHistoryFiltered)
// lenListHistory := len(listHistoryFiltered)
// // lengthListLessThanFive := []string{}
// delimiter = lenListHistory - 5
//test len less than 5
// for i, v := range listHistoryFiltered {
// 	fmt.Println(i)
// 	fmt.Println(v)
// }
// 	if i >= delimiter {
// 		lengthListLessThanFive = append(lengthListLessThanFive, v)
// 	}
// }
// if lenListHistory > 5 {
// 	delimiter = lenListHistory - 5
// 	isDelimiterNeeded = true
// }

// 	// 		}
// 	// 		continue
// 	// 	}
// 	// 	symQuery := strings.Split(v, "</sym>")[0]
// 	// 	symFromHistory := strings.Split(symQuery, "<sym>")[1]

// 	// 	//side
// 	// 	sideQuery := strings.Split(v, "</side>")[0]
// 	// 	sideFromHistory := strings.Split(sideQuery, "<side>")[1]

// 	// 	//parse qty
// 	// 	quantityQuery := strings.Split(v, "</quantity>")[0]
// 	// 	quantityFromHistory := strings.Split(quantityQuery, "<quantity>")[1]

// 	// 	//parse price
// 	// 	priceQuery := strings.Split(v, "</price>")[0]
// 	// 	priceFromHistory := strings.Split(priceQuery, "<price>")[1]

// 	// 	//HistoryValue
// 	// 	historyValue := HistoryValue{Symbol: symFromHistory, Side: sideFromHistory, Qty: quantityFromHistory, Price: priceFromHistory}
// 	// 	listHistoryValues = append(listHistoryValues, historyValue)
// 	i++

func formatDateFromHistory(date string) string {
	splitDate := strings.Split(date, "-")
	year := splitDate[0]
	month := splitDate[1]
	daySplit := strings.Split(splitDate[2], "T")[0]
	formattedHistory := year + " " + month + " " + daySplit
	return formattedHistory
}
func formatCreatedAtFromInformationAtTrade(createdAt string) string {
	splitDate := strings.Split(createdAt, "-")
	year := splitDate[0]
	month := splitDate[1]
	daySplit := strings.Split(splitDate[2], "T")[0]
	formattedHistory := year + " " + month + " " + daySplit
	return formattedHistory
}

func systemReadPreviousHistory() {

}

func handleSellAtMarket(symbol string) {
	// handleSellAtMarket

	//Query holding get QTY...
	holdingWisemen := HoldingWisemen{}
	holdingWisemenContainer := getAllHolding()

	for i, v := range holdingWisemenContainer.ListHolding {
		if v.Symbol == symbol {
			holdingWisemen = v
		}
		i++
	}
	//
	// fmt.Println(holdingWisemen)
	queryTradeSellMarket(holdingWisemen)
}

func handleSellLimitWisemen(symbol string) {
	containerHolding := getAllHolding()
	holdingToSell := HoldingWisemen{Symbol: "default"}
	for i, v := range containerHolding.ListHolding {
		if symbol == v.Symbol {
			holdingToSell = v
		}
		i++
	}
	// //Handle hold not found
	if holdingToSell.Symbol == "default" {
		fmt.Println("Holding not present")
	}
	if holdingToSell.Symbol != "default" {
		// get price information
		//get metric delimiter
		metricsWisemen := selectMetricsWisemen()[0]
		//
		holdingPrice := 0.0
		priceHighPchgTrade := 0.0
		//string to float
		if s, err := strconv.ParseFloat(holdingToSell.Price, 64); err == nil {
			holdingPrice = s
		}
		if s1, err := strconv.ParseFloat(metricsWisemen.PriceHighPchgTrade, 64); err == nil {
			priceHighPchgTrade = s1
		}
		fmt.Println("priceHighPchgTrade")
		fmt.Println(priceHighPchgTrade)
		//calculate limit price...
		limitPrice := holdingPrice + (holdingPrice * priceHighPchgTrade)

		stringLimitPrice := fmt.Sprintf("%f", limitPrice)

		stringLimitPrice = floatToString(splitFloatAfterSecondDecimalPlace(stringToFloat(stringLimitPrice)))
		queryTradeSellLimit(holdingToSell.Symbol, stringLimitPrice, holdingToSell.Qty)
	}
}

func getCurrentStockFromQuery(symbolList []string) []Stock {
	// func queryMultiStockPull(symbolList []string) string {
	queryResponse := queryMultiStockPull(symbolList)
	stockList := parseStockSetQuery(queryResponse)
	return stockList
}
func calculateAmountOfStockToBuy(pricePointOfStock float64, balance float64) float64 {
	// pricePointOfStock being target limit to buy
	//Calculate amount of shares to buy at given balance and bias
	// biasPercentage
	// biasPercentage := 0.0
	// //BiasPrice
	// biasPrice := 0
	//10% 5000 500
	//1% 5000 50
	buffer := balance * .015
	bufferedBalance := balance - buffer
	fmt.Println(bufferedBalance)
	//Want to leave enough account balance to buffer for variance in buying.
	amountToBuy := bufferedBalance / pricePointOfStock
	// amountToBuy = amountToBuy - buffer

	return amountToBuy
}

func getAllOrders() ContainerOrders {
	queryResponse := queryOrders()
	containerOrders := parseOrders(queryResponse)
	return containerOrders
}
func getAllHolding() ContainerHolding {
	queryResponse := queryHolding()
	containerHolding := parseAllHolding(queryResponse)
	return containerHolding
}

// func getOrder(symbol string) {
// 	//Handle if more than one order
// 	// pull orders
// 	queryOrders(){}
// 	//parse
// }
func cancelOrder(symbol string) {
	//cancel order by id
}

func calculateHoldingStatus(holdingWisemen HoldingWisemen) HoldingWisemen {
	isPartialUnfinished := false
	isCompletedFull := false
	// //Populate order container
	// containerOrders := getAllOrders()
	// order := Order{Symbol: "default"}
	// containerHoldings := getAllHolding()
	// //get informationAtTrade
	// listInformationAtTrade := selectInformationAtTrade()
	// holding := Holding{}
	holdingWisemenReturned := holdingWisemen
	holdingWisemenReturned.OrderStatus = "undetermined"
	// for i,v := range listInformationAtTrade {
	// 	// if v.Symbol == holdingWisemen.Symbol {
	// // 	order = v
	// // }
	// if v.
	// 	i++
	// }

	//get informatin at trade
	informationAtTradeList := selectInformationAtTrade()
	informationAtTrade := informationAtTradeList[0]
	//compare holding qty to information at trade qty.
	fmt.Println("informationAtTrade")
	fmt.Println(informationAtTrade)
	fmt.Println("informationAtTrade.qty")
	fmt.Println(informationAtTrade.Qty)

	fmt.Println("holdingWisemenReturned.Qty")
	fmt.Println(holdingWisemenReturned.Qty)
	//compare order qty to bought qty.
	if informationAtTrade.Qty == holdingWisemenReturned.Qty {
		isCompletedFull = true
	}
	if informationAtTrade.Qty > holdingWisemenReturned.Qty {
		isPartialUnfinished = true
	}
	//
	if isCompletedFull {
		holdingWisemenReturned.OrderStatus = "completedFull"
	}
	if isPartialUnfinished {
		holdingWisemenReturned.OrderStatus = "partial"
	}
	return holdingWisemenReturned
}

// func calculateIsTradeBoughtSuccessful(symbol string) {
// 	isSuccessful := detectIsTradeBoughtSuccessful(symbol)
// 	if isSuccessful {
// 		updateDayTrackingRecordSystem(symbol)
// 	}
// 	return
// }
func updateDayTrackingRecordSystem(symbol string) {
	dayOfWeek := getDayOfWeek()
	//create record
	dayTrackingRecord := DayTrackingRecord{Symbol: symbol, DayOfWeekCreated: dayOfWeek.String(), DayOfWeekDayIteration: "0", LastDayOfWeekDayUpdate: getDayOfWeek().String(), AmountOfTrades: "0", IsWeekPassed: "false"}
	insertDayTrackingRecord(dayTrackingRecord)
}

// func detectIsTradeBoughtSuccessful(symbol string) bool {
// 	isSuccessful := false
// 	entryList := selectTradeBoughtEvaluation()
// 	//Support for multiple stream operations of trading.
// 	for i, v := range entryList {
// 		if v.Symbol == symbol {
// 			isSuccessful = true
// 		}
// 		i++
// 	}
// 	return isSuccessful
// }

func roundDown(floatValue float64) int {
	stringValue := fmt.Sprintf("%f", floatValue)
	stringSlice := strings.Split(stringValue, ".")
	// intValue, err := strconv.ParseInt(stringSlice[0], 10, 64)
	intValue, err := strconv.Atoi(stringSlice[0])
	if err != nil {
		fmt.Println(err)
	}
	// intValue := 9
	// floatValue
	//splice at "."
	//string value to int.
	return intValue
}

func splitFloatAfterSecondDecimalPlace(floatValue float64) float64 {

	// convert to string
	s := fmt.Sprintf("%f", floatValue)
	// split by period
	stringSlice := strings.Split(s, ".")
	//split 2 delimiter
	fmt.Println(stringSlice)
	returningFloat := floatValue

	valueAfterPeriod := stringSlice[1]
	if len(valueAfterPeriod) > 2 {
		fmt.Println("hit")
		diffenceIndex := 2
		for diffenceIndex < 100 {
			if diffenceIndex == len(valueAfterPeriod) {
				break
			}
			diffenceIndex++
		}
		fmt.Println(diffenceIndex)
		i := 2
		for i < diffenceIndex {
			valueAfterPeriod = valueAfterPeriod[:len(valueAfterPeriod)-1]
			i++
		}
		stringedFloatValue := stringSlice[0] + "." + valueAfterPeriod
		fmt.Println(stringedFloatValue)
		returningFloat = stringToFloat(stringedFloatValue)
	}
	return returningFloat
}

func floatToString(input_num float64) string {
	// to convert a float number to a string
	return strconv.FormatFloat(input_num, 'f', 6, 64)
}
func stringToFloat(value string) float64 {
	returnFloat := 0.0
	if s, err := strconv.ParseFloat(value, 64); err == nil {
		returnFloat = s
	}
	return returnFloat
}

func removeElementString(listEntered []Stock, symbol string) []Stock {
	var i int
	listAltered := listEntered
	for {
		if i == len(listAltered) {
			break
		}
		if listAltered[i].Symbol == symbol {
			listAltered = listAltered[:i+copy(listAltered[i:], listAltered[i+1:])]
			i = 0
		}
		i++
	}
	return listAltered
}

func removeElementInt(listEntered []int, val int) []int {
	var i int
	listAltered := listEntered
	for {
		if i == len(listAltered) {
			break
		}

		if listAltered[i] == val {
			listAltered = listAltered[:i+copy(listAltered[i:], listAltered[i+1:])]
			i = 0
		}
		i++
	}
	return listAltered
}

func overarchTradeWisemen(dataList []string) {
	isTradeDay := overarchIsTradeDay()
	fmt.Println(dataList)
	//handle no neo found..
	// overarchTradeWisemen
	if dataList[0] == "none chosen" {
		fmt.Println("none chosen")
		//persist none chosen, in ReasonCancelation in AIB
		//proceed to alt interval.
		altIntervalList := selectAltIntervalBuyWisemen()
		if len(altIntervalList) == 0 {
			//Support for isTradeDay in returned response, for multi cancelation scenarios.
			altIntervalBuyWisemen := AltIntervalBuyWisemen{Symbol: dataList[0], IsAltIntervalOperation: "true", ReasonCancelation: "Neo none chosen"}
			insertAltIntervalBuyWisemen(altIntervalBuyWisemen)
		}
		if len(altIntervalList) != 0 {
			altIntervalBuyWisemen := AltIntervalBuyWisemen{Symbol: dataList[0], IsAltIntervalOperation: "true", ReasonCancelation: "Neo none chosen"}
			insertAltIntervalBuyWisemen(altIntervalBuyWisemen)
			transactionHistory := TransactionHistory{Symbol: dataList[0]}
			wrapUpWisemenOutcomeNoBuy(transactionHistory)
		}
	}

	if dataList[0] != "none chosen" {
		if isTradeDay {
			//process trade.
			fmt.Println("internal overarchTradeWisemen")
			handleTradeWisemen(dataList[0], dataList[1])
			time.Sleep(time.Duration(10) * time.Second)
			// //Begin process monitoring for buy fulfilled.
			processCheckIsTradeBought(dataList[0])
		}
		//handle if not isTradeDay
		if isTradeDay == false {
			//insert AltIntervalBuyWisemen
			//interval 1
			altIntervalList := selectAltIntervalBuyWisemen()
			if len(altIntervalList) == 0 {
				//Support for isTradeDay in returned response, for multi cancelation scenarios.
				altIntervalBuyWisemen := AltIntervalBuyWisemen{Symbol: dataList[0], IsAltIntervalOperation: "true", ReasonCancelation: "isTradeDayFalse"}
				insertAltIntervalBuyWisemen(altIntervalBuyWisemen)
			}
			if len(altIntervalList) != 0 {
				altIntervalBuyWisemen := AltIntervalBuyWisemen{Symbol: dataList[0], IsAltIntervalOperation: "true", ReasonCancelation: "isTradeDayFalse"}
				insertAltIntervalBuyWisemen(altIntervalBuyWisemen)
				transactionHistory := TransactionHistory{Symbol: dataList[0]}
				wrapUpWisemenOutcomeNoBuy(transactionHistory)
			}
		}
	}
}
