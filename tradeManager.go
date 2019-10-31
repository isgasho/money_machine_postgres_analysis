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

func handleSellWisemen(symbol string, limitPrice string, quantity string) {
	// Need handle on quantity, symbol, price to sell.
	//Overarch handle sell system metrics read data flow and set.

	//Read order bought information.
	//Is order bought information entered yet?
	//or call holding... for a dynamic flow... yes.
	//Future support for advanced recording system.

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
	//[["symbol",orderID]]
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

func floatToString(input_num float64) string {
	// to convert a float number to a string
	return strconv.FormatFloat(input_num, 'f', 6, 64)
}
