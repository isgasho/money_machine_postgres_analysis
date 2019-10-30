package main

import (
	"fmt"
	"math"
	"strconv"
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
	// fmt.Println(floatPrice)

	qty = math.Round(qty)

	// stringBalance := fmt.Sprintf("%f", floatBalance)
	stringPrice := fmt.Sprintf("%f", desiredLimitPrice)
	stringQty := fmt.Sprintf("%f", qty)

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
	queryTradeBuyLimit(symbol, stringPrice, stringQty)
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
	buffer := balance * .01
	bufferedBalance := balance - buffer
	//Want to leave enough account balance to buffer for variance in buying.
	amountToBuy := bufferedBalance / pricePointOfStock
	amountToBuy = amountToBuy - buffer

	return amountToBuy
}

func getAllOrders() ContainerOrders {
	//[["symbol",orderID]]
	queryResponse := queryOrders()
	containerOrders := parseOrders(queryResponse)
	return containerOrders
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

func calculateHoldingStatus(holdingWisemen HoldingWisemen) string {
	defaultHoldingStatus := "undetermined"
	isPartialUnfinished := true
	isCompletedFull := true
	//calculate
	//Need original order sent in and compare to bought.
	//Get order
	// orderContainer := getAllOrders()
	//Need order submitted, compare to holding wisemen
	// orderContainer.ListOrders
	//Get order for
	// for i, v := range orderContainer.ListOrders {

	// }

	//Get the order compare...
	//return status
	if isPartialUnfinished {
		return "isPartialUnfinished"
	}
	if isCompletedFull {
		return "isCompletedFull"
	}
	return defaultHoldingStatus
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
