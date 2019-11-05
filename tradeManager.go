package main

import (
	"fmt"
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
	//
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

func isSellShowingInHistory(symbol string) {
	//Future support for complex timing on trades, same symbols.
	//

	//Presume the latest order is the symbol...
	//Think of any cases where this is not true.
	//
	//Query order information
	//
	// orderList := getAllOrders()
	// ordersBySymbolList := []Order{}
	// fmt.Println(orderList)
	// for i, v := range orderList.ListOrders {
	// 	if v.Symbol == symbol {
	// 		ordersBySymbolList = append(ordersBySymbolList, v)
	// 	}
	// 	i++
	// }

	// if len(ordersBySymbolList) != 0 {
	// latestOrder := ordersBySymbolList[(len(ordersBySymbolList) - 1)]
	// fmt.Println(latestOrder)
	//Get handle on history.
	response := queryHistory()
	// latestHistory, history1 := parseLatestHistory(response)

	//UpgradeHistory to only have list.
	historyList := parseHistory(response)
	// handleListAppendedValues
	listHistoryValues := createListHistoryValuesForWisemen(historyList)
	//We want a list of appended values for
	fmt.Println(listHistoryValues)

	for i, v := range listHistoryValues {
		fmt.Println(i)
		fmt.Println(v)
	}
	// fmt.Println(latestEntry)
	//Parse for symbol in latestEntry
	//

	// splitDataQuery := strings.Split(latestHistory, "</sym>")[0]
	// symFromHistory := strings.Split(splitDataQuery, "<sym>")[1]

	// //parse side
	// sideQuery := strings.Split(latestHistory, "</side>")[0]
	// sideFromHistory := strings.Split(sideQuery, "<side>")[1]

	// //parse qty
	// quantityQuery := strings.Split(latestHistory, "</quantity>")[0]
	// quantityFromHistory := strings.Split(quantityQuery, "<quantity>")[1]

	// //history1
	// splitDataQuery1 := strings.Split(history1, "</sym>")[0]
	// symFromHistory1 := strings.Split(splitDataQuery1, "<sym>")[1]

	// //parse side
	// sideQuery1 := strings.Split(history1, "</side>")[0]
	// sideFromHistory1 := strings.Split(sideQuery1, "<side>")[1]

	// //parse qty
	// quantityQuery1 := strings.Split(history1, "</quantity>")[0]
	// quantityFromHistory1 := strings.Split(quantityQuery1, "<quantity>")[1]

	//Match symbol
	// if symFromHistory == symbol {
	// 	if sideFromHistory == "2" {
	// 		//Handle side
	// 		//if side then sell was done for symbol.
	// 		//Read results
	// 		//Read either balance or history
	// 		//Need trade result
	// 		//Analyze

	// 		//pull trade order information and if change entered.
	// 		//Support if trade entered.

	// 		//previous history will have buy history before it.
	// 		//get history 7.9
	// 		//Only problem is if the entier

	// 		// if
	// 		//System to read previous consecutive historys if quantity does not match
	// 		if symFromHistory1 == symFromHistory {
	// 			//if not go back further. if further limit reached then stop.
	// 		}
	// 		//Indicating partial not completed. Except the changeorders.
	// 		//How to determine change orders.
	// 		// systemReadPreviousHistory()
	// 		//Query trade store, was a changeorder entered.
	// 	}
	// }
	//Need difference between sell order placed and holding...
}

func createListHistoryValuesForWisemen(listHistory []string) []HistoryValue {
	listHistoryValues := []HistoryValue{}
	isDelimiterNeeded := false
	delimiter := 0
	lenListHistory := len(listHistory)
	listHistoryFiltered := listHistory[:len(listHistory)-1]

	if lenListHistory > 5 {
		delimiter = lenListHistory - 5
		isDelimiterNeeded = true
	}
	// splitDataQuery2 = splitDataQuery2[:len(splitDataQuery2)-1]
	for i, v := range listHistoryFiltered {
		if isDelimiterNeeded {
			if i >= delimiter {
				fmt.Println(v)
				fmt.Println(i)
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

		//sym

		//go back only 5 orders.

		// if
		// if

		//splitDataQuery2 = splitDataQuery2[:len(splitDataQuery2)-1]

		//given len(splitDataQuery2)-1
		//go back 5 orders

		// fmt.Println(len(symFromHistory))
		// if i == 1 {
		// 	break
		// }

		// i++
	}
	return listHistoryValues
}

func systemReadPreviousHistory() {

}
func handleSellWisemen(symbol string) {
	//Overarch handle sell system metrics read data flow and set.
	//Read order bought information.
	//Is order bought information entered yet?
	//or call holding... for a dynamic flow... yes.
	//Future support for advanced recording system.

	//source symbol pull
	//pull holding for symbol
	// holding
	// handleSellWisemen
	containerHolding := getAllHolding()
	holdingToSell := HoldingWisemen{Symbol: "default"}
	for i, v := range containerHolding.ListHolding {
		if symbol == v.Symbol {
			holdingToSell = v
		}
		i++
	}

	//Handle order not found
	if holdingToSell.Symbol == "default" {
		fmt.Println("Holding not present")
	}

	if holdingToSell.Symbol != "default" {
		//get price information
		// stockList := getCurrentPriceStatsForStock([]string{holdingToSell.Symbol})
		// for i, v := range stockList {
		// 	// fmt.Println("stock")

		// 	// fmt.Println("bid")
		// 	// fmt.Println(v.Bid)
		// 	// fmt.Println("ask")
		// 	// fmt.Println(v.Ask)
		// 	// fmt.Println("last")
		// 	// fmt.Println(v.Last)

		// 	i++
		// }

		//Get metric delimiter
		//
		metricsWisemen := selectMetricsWisemen()[0]
		holdingPrice := 0.0
		metricsDesiredPchg := 0.0
		//String to float
		if s, err := strconv.ParseFloat(holdingToSell.Price, 64); err == nil {
			holdingPrice = s
		}

		//String to float
		if s, err := strconv.ParseFloat(metricsWisemen.DesiredPchg, 64); err == nil {
			metricsDesiredPchg = s
		}

		fmt.Println("holdingPrice")
		fmt.Println(holdingPrice)
		limitPrice := holdingPrice + (holdingPrice * metricsDesiredPchg)
		fmt.Println(limitPrice)
		fmt.Println(limitPrice)

		stringLimitPrice := fmt.Sprintf("%f", limitPrice)
		fmt.Println("stringLimitPrice")
		fmt.Println(stringLimitPrice)

		stringLimitPrice = floatToString(splitFloatAfterSecondDecimalPlace(stringToFloat(stringLimitPrice)))
		fmt.Println("stringLimitPrice")
		fmt.Println(stringLimitPrice)
		// fmt.Println("holdingPrice")
		// fmt.Println(holdingPrice)
		// fmt.Println("metricsDesiredPchg")
		// fmt.Println(metricsDesiredPchg)
		queryTradeSellLimit(holdingToSell.Symbol, stringLimitPrice, holdingToSell.QtyBought)
	}

	//sell

	//Pull holdings

	//Read metrics and set system

	//Data flow for one stock and analysis

	//handleSellWisemen set

	//
	// queryTradeSellLimit(symbol, stringPrice, quantity)

	// fmt.Println("limitPrice")
	// fmt.Println(limitPrice)
	// desiredLimitPrice := 0.0
	// if s, err := strconv.ParseFloat(limitPrice, 64); err == nil {
	// 	desiredLimitPrice = s
	// }

	// fmt.Println("desiredLimitPrice")
	// fmt.Println(desiredLimitPrice)
	// //get balance
	// response := queryBalance()
	// balance := parseBalance(response)

	// floatBalance := 0.0
	// if s, err := strconv.ParseFloat(balance, 64); err == nil {
	// 	floatBalance = s
	// }

	// fmt.Println("floatBalance")
	// fmt.Println(floatBalance)
	// //
	// //calculate qty to buy
	// qty := calculateAmountOfStockToBuy(desiredLimitPrice, floatBalance)
	// fmt.Println("before rounding down")
	// fmt.Println(qty)

	// qtyInt := roundDown(qty)
	// fmt.Println("After rounding down")
	// fmt.Println(qtyInt)

	// stringBalance := fmt.Sprintf("%f", floatBalance)
	// stringPrice := fmt.Sprintf("%f", desiredLimitPrice)
	// // stringQty := fmt.Sprintf("%f", qtyInt)
	// stringQty := strconv.Itoa(qtyInt)

	//store trade entered information
	// tradeEnteredInformation := TradeEnteredInformation{
	// 	Symbol:      symbol,
	// 	Price:       stringPrice,
	// 	OrderStatus: "pending",
	// 	Qty:         stringQty,
	// 	QtyBought:   "0",
	// }
	// insertTradeEnteredInformation(tradeEnteredInformation)

	//Submit buy limit to brokerage
	// fmt.Println("symbol")
	// fmt.Println(symbol)
	// fmt.Println("stringQty")
	// fmt.Println(stringQty)
	// fmt.Println("stringPrice")
	// fmt.Println(stringPrice)

}

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
