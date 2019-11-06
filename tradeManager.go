package main

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
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
	tradeEnteredInformation := TradeEnteredInformation{
		Symbol:      symbol,
		Price:       stringPrice,
		OrderStatus: "pending",
		Qty:         stringQty,
		QtyBought:   "0",
	}
	insertTradeEnteredInformation(tradeEnteredInformation)
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
	symbol := listStrings[0]
	// priceDropPercentage := listStrings[1]
	timeDelimiter := listStrings[2]

	// Hook
	// fmt.Println(priceDrop)
	// fmt.Println(timeDelimiter)
	// isDropPriceMet := false
	// isTimeDelimiterMet := false
	// // queryMultiStockPull()
	// //
	isDropPriceMet := calculateIsDropPriceMet(symbol)
	// //
	isTimeDelimiterMet := calculateIsTimeDelimiterMetSell(timeDelimiter)

	// fmt.Println(isDropPriceMet)
	// fmt.Println(isTimeDelimiterMet)
	// //require feed, query.

	// //check on time

	// //
	// // handleTimeDelimiterMetSell
	if isTimeDelimiterMet {
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
		postSellAtMarket(holding)
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
		postSellAtMarket(holding)
	}

}

func calculateIsDropPriceMet(symbol string) bool {
	// listSymbol := []string{symbol}
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

	priceFromQuery := "23.02"
	metrics := selectMetricsWisemen()[0]
	//Here we need a handle on drop metrics,...
	//Partial bind...
	metricPchgDrop := metrics.PriceLowPchg
	// fmt.Println(metricPchgDrop)

	floatPriceFromQuery := 0.0
	floatMetricPchgDrop := 0.0
	if s, err := strconv.ParseFloat(priceFromQuery, 64); err == nil {
		floatPriceFromQuery = s
	}
	if s, err := strconv.ParseFloat(metricPchgDrop, 64); err == nil {
		floatMetricPchgDrop = s
	}

	dropPrice := floatPriceFromQuery - (floatPriceFromQuery * floatMetricPchgDrop)
	holdingPrice := 0.0

	if s, err := strconv.ParseFloat(holding.Price, 64); err == nil {
		holdingPrice = s
	}
	//compare drop price to holding price.

	if holdingPrice <= dropPrice {
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
func isSymbolPresentInHolding(symbol string) string {
	//Question, when is sold fully out.
	//When is there no holding,
	//Is there an order still open
	//Who cares if it's still open.
	//
	//There is no other case.
	//It's either gone or not.
	//If it's not
	//Given time delimiter, or pchg Dropoff
	//Get holding QTY
	//set limit to sell.
	//
	isBoolReturning := "false"
	holdingList := getAllHolding()
	for i, v := range holdingList.ListHolding {
		if v.Symbol == symbol {
			isBoolReturning = "true"
		}
		fmt.Println(v)
		i++
	}
	return isBoolReturning
}

func isSellShowingInHistory(symbol string) {
	//Question, when is sold fully out.
	//When is there no holding,
	//Is there an order still open
	//Who cares if it's still open.

	//
	//There is no other case.
	//It's either gone or not.
	//If it's not
	//Given time delimiter, or pchg Dropoff
	//Get holding QTY
	//set limit to sell.
	//

	// holdingList := getAllHolding()
	// for i, v := range holdingList.ListHolding {
	// 	if
	// 	fmt.Println(v)
	// 	i++
	// }

	//Future support for complex timing on trades, same symbols.
	//
	//Get handle on history.
	response := queryHistory()
	//UpgradeHistory to only have list.
	historyList := parseHistory(response)
	listHistoryValues := createListHistoryValuesForWisemen(historyList)

	//Given historySell
	historySell := HistoryValue{}
	for i, v := range listHistoryValues {
		alteredHistoryValue := HistoryValue{}
		// fmt.Println(i)
		// fmt.Println(v)
		stringedInterval := strconv.Itoa(i)
		alteredHistoryValue = calculateSellHistoryMatchesSymbol(v, symbol, stringedInterval)
		if alteredHistoryValue.IsCalculationTrue == "true" {
			historySell = v
			break
		}
	}

	//Given historyBuy
	historyBuy := HistoryValue{}
	for i, v := range listHistoryValues {
		alteredHistoryValue := HistoryValue{}
		// fmt.Println(i)
		// fmt.Println(v)
		stringedInterval := strconv.Itoa(i)

		alteredHistoryValue = calculateBuyHistoryMatchesSymbol(v, symbol, stringedInterval)
		if alteredHistoryValue.IsCalculationTrue == "true" {
			historyBuy = v
			break
		}
	}
	fmt.Println("sell")
	fmt.Println(historySell)
	fmt.Println("buy")
	fmt.Println(historyBuy)
	//conditional if QTY fully sold...
	//Handling partitioned buys will be difficult...
	//Because the QTY will be less? or will it?
	//There's no way of knowing, not even a way of setting that case in the wild.
	//
	//But can create a conditional watch, handle outcomes and holdings.
	//
	//create the net and the environment to catch that data...
	//
	//Here future support for conditional -> partial store and retrieval.

	//

	//For right now just handle conditional with presumption that full orders commited both ways.
	//
	//But what are we even looking for. If we are simply looking for none in holding...
	//If none in holding then bought...
	//...We are looking to see if fully sold.
	//A fully sold condition will show no in holding...
	//
	//If some still in Holding than we have a partial, or still open order and condition...
	//
	//That partial detection is needed for this calculation.
	//

	//
	//Handle full detection...

	//get holding...

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
	isDelimiterNeeded := false
	delimiter := 0
	listHistoryFiltered := listHistory[:len(listHistory)-1]
	lenListHistory := len(listHistoryFiltered)
	// lengthListLessThanFive := []string{}
	delimiter = lenListHistory - 5
	//test len less than 5
	// for i, v := range listHistoryFiltered {
	// 	if i >= delimiter {
	// 		lengthListLessThanFive = append(lengthListLessThanFive, v)
	// 	}
	// }
	if lenListHistory > 5 {
		delimiter = lenListHistory - 5
		isDelimiterNeeded = true
	}
	for i, v := range listHistoryFiltered {
		fmt.Println(i)
		fmt.Println(v)
		if isDelimiterNeeded {
			if i >= delimiter {
				//sym
				symQuery := strings.Split(v, "</sym>")[0]
				symFromHistory := strings.Split(symQuery, "<sym>")[1]

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
				historyValue := HistoryValue{Symbol: symFromHistory, Side: sideFromHistory, Qty: quantityFromHistory, Price: priceFromHistory}
				listHistoryValues = append(listHistoryValues, historyValue)
			}
			continue
		}
		symQuery := strings.Split(v, "</sym>")[0]
		symFromHistory := strings.Split(symQuery, "<sym>")[1]

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
		historyValue := HistoryValue{Symbol: symFromHistory, Side: sideFromHistory, Qty: quantityFromHistory, Price: priceFromHistory}
		listHistoryValues = append(listHistoryValues, historyValue)
	}
	return listHistoryValues
}

func systemReadPreviousHistory() {

}

// func handleSellWisemen(symbol string) {
// 	//Overarch handle sell system metrics read data flow and set.
// 	//Read order bought information.
// 	//Is order bought information entered yet?
// 	//or call holding... for a dynamic flow... yes.
// 	//Future support for advanced recording system.

// 	//source symbol pull
// 	//pull holding for symbol
// 	// holding
// 	// handleSellWisemen
// 	containerHolding := getAllHolding()
// 	holdingToSell := HoldingWisemen{Symbol: "default"}
// 	for i, v := range containerHolding.ListHolding {
// 		if symbol == v.Symbol {
// 			holdingToSell = v
// 		}
// 		i++
// 	}

// 	//Handle order not found
// 	if holdingToSell.Symbol == "default" {
// 		fmt.Println("Holding not present")
// 	}

// 	if holdingToSell.Symbol != "default" {
// 		//get price information
// 		// stockList := getCurrentPriceStatsForStock([]string{holdingToSell.Symbol})
// 		// for i, v := range stockList {
// 		// 	// fmt.Println("stock")

// 		// 	// fmt.Println("bid")
// 		// 	// fmt.Println(v.Bid)
// 		// 	// fmt.Println("ask")
// 		// 	// fmt.Println(v.Ask)
// 		// 	// fmt.Println("last")
// 		// 	// fmt.Println(v.Last)

// 		// 	i++
// 		// }

// 		//Get metric delimiter
// 		//
// 		metricsWisemen := selectMetricsWisemen()[0]
// 		holdingPrice := 0.0
// 		metricsDesiredPchg := 0.0
// 		//String to float
// 		if s, err := strconv.ParseFloat(holdingToSell.Price, 64); err == nil {
// 			holdingPrice = s
// 		}

// 		//String to float
// 		if s, err := strconv.ParseFloat(metricsWisemen.DesiredPchg, 64); err == nil {
// 			metricsDesiredPchg = s
// 		}

// 		fmt.Println("holdingPrice")
// 		fmt.Println(holdingPrice)
// 		limitPrice := holdingPrice + (holdingPrice * metricsDesiredPchg)
// 		fmt.Println(limitPrice)
// 		fmt.Println(limitPrice)

// 		stringLimitPrice := fmt.Sprintf("%f", limitPrice)
// 		fmt.Println("stringLimitPrice")
// 		fmt.Println(stringLimitPrice)

// 		stringLimitPrice = floatToString(splitFloatAfterSecondDecimalPlace(stringToFloat(stringLimitPrice)))
// 		fmt.Println("stringLimitPrice")
// 		fmt.Println(stringLimitPrice)
// 		// fmt.Println("holdingPrice")
// 		// fmt.Println(holdingPrice)
// 		// fmt.Println("metricsDesiredPchg")
// 		// fmt.Println(metricsDesiredPchg)
// 		queryTradeSellLimit(holdingToSell.Symbol, stringLimitPrice, holdingToSell.QtyBought)
// 	}

// 	//sell

// 	//Pull holdings

// 	//Read metrics and set system

// 	//Data flow for one stock and analysis

// 	//handleSellWisemen set

// 	//
// 	// queryTradeSellLimit(symbol, stringPrice, quantity)

// 	// fmt.Println("limitPrice")
// 	// fmt.Println(limitPrice)
// 	// desiredLimitPrice := 0.0
// 	// if s, err := strconv.ParseFloat(limitPrice, 64); err == nil {
// 	// 	desiredLimitPrice = s
// 	// }

// 	// fmt.Println("desiredLimitPrice")
// 	// fmt.Println(desiredLimitPrice)
// 	// //get balance
// 	// response := queryBalance()
// 	// balance := parseBalance(response)

// 	// floatBalance := 0.0
// 	// if s, err := strconv.ParseFloat(balance, 64); err == nil {
// 	// 	floatBalance = s
// 	// }

// 	// fmt.Println("floatBalance")
// 	// fmt.Println(floatBalance)
// 	// //
// 	// //calculate qty to buy
// 	// qty := calculateAmountOfStockToBuy(desiredLimitPrice, floatBalance)
// 	// fmt.Println("before rounding down")
// 	// fmt.Println(qty)

// 	// qtyInt := roundDown(qty)
// 	// fmt.Println("After rounding down")
// 	// fmt.Println(qtyInt)

// 	// stringBalance := fmt.Sprintf("%f", floatBalance)
// 	// stringPrice := fmt.Sprintf("%f", desiredLimitPrice)
// 	// // stringQty := fmt.Sprintf("%f", qtyInt)
// 	// stringQty := strconv.Itoa(qtyInt)

// 	//store trade entered information
// 	// tradeEnteredInformation := TradeEnteredInformation{
// 	// 	Symbol:      symbol,
// 	// 	Price:       stringPrice,
// 	// 	OrderStatus: "pending",
// 	// 	Qty:         stringQty,
// 	// 	QtyBought:   "0",
// 	// }
// 	// insertTradeEnteredInformation(tradeEnteredInformation)

// 	//Submit buy limit to brokerage
// 	// fmt.Println("symbol")
// 	// fmt.Println(symbol)
// 	// fmt.Println("stringQty")
// 	// fmt.Println(stringQty)
// 	// fmt.Println("stringPrice")
// 	// fmt.Println(stringPrice)

// }

// intiateSellSystemProtocol(){

// }
func getCurrentPriceStatsForStock(symbolList []string) []Stock {
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
	// defaultHoldingStatus := "undetermined"
	holdingWisemen.OrderStatus = "undetermined"
	isPartialUnfinished := false
	isCompletedFull := false
	//calculate
	//Get order
	containerOrders := getAllOrders()
	order := Order{Symbol: "default"}
	for i, v := range containerOrders.ListOrders {
		if v.Symbol == holdingWisemen.Symbol {
			order = v
		}
		i++
	}

	//contingent that order.qty still exists.
	//If order does not exist should pull trade information.
	//Handle conditional where order is not placed.
	if order.Symbol == "default" {
		//cancel process return
		holdingWisemen.OrderStatus = "order not placed"
		return holdingWisemen
	}

	fmt.Println("order.Qty")
	fmt.Println(order.Qty)

	fmt.Println("holdingWisemen.QtyBought")
	fmt.Println(holdingWisemen.QtyBought)
	//compare order qty to bought qty.

	if order.Qty == holdingWisemen.QtyBought {
		isCompletedFull = true
	}
	if order.Qty > holdingWisemen.QtyBought {
		isPartialUnfinished = true
	}

	//update holdingWisemen status
	//return holdingWisemen
	if isCompletedFull {
		holdingWisemen.OrderStatus = "completedFull"
	}
	if isPartialUnfinished {
		holdingWisemen.OrderStatus = "partial"
	}
	return holdingWisemen
}
func calculateIsTradeBoughtSuccessful(symbol string) {
	isSuccessful := detectIsTradeBoughtSuccessful(symbol)
	if isSuccessful {
		updateDayTrackingRecordSystem(symbol)
	}
	return
}

func updateDayTrackingRecordSystem(symbol string) {
	dayOfWeek := getDayOfWeek()
	//create record
	dayTrackingRecord := DayTrackingRecord{Symbol: symbol, DayOfWeekCreated: dayOfWeek.String(), DayOfWeekDayIteration: "0", LastDayOfWeekDayUpdate: getDayOfWeek().String(), AmountOfTrades: "0", IsWeekPassed: "false"}
	insertDayTrackingRecord(dayTrackingRecord)
}

func detectIsTradeBoughtSuccessful(symbol string) bool {
	isSuccessful := false
	entryList := selectTradeBoughtEvaluation()
	//Support for multiple stream operations of trading.
	for i, v := range entryList {
		if v.Symbol == symbol {
			isSuccessful = true
		}
		i++
	}
	return isSuccessful
}

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
