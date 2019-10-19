package main

import (
	"fmt"
	"strconv"
)

func handleTradeWisemen() {

	//given desired delimiter
	//Desired limit price point should be pulled from Db or passed.
	//Depends on process basis.
	//analytics based, so should be...
	//stored through anlaytics...
	symbol := "PEACE"
	desiredLimitPrice := 3.35

	//get balance
	response := queryTradeCheckBalance()
	balance := parseBalance(response)
	// fmt.Println(balance)
	floatBalance := 0.0
	if s, err := strconv.ParseFloat(balance, 64); err == nil {
		// fmt.Println(s)
		floatBalance = s
	}
	//
	//calculate qty to buy
	floatPrice := calculateAmountOfStockToBuy(desiredLimitPrice, floatBalance)
	fmt.Println(floatPrice)

	stringBalance := fmt.Sprintf("%f", floatBalance)
	stringPrice := fmt.Sprintf("%f", floatPrice)

	//create buy query
	// queryTradeBuyLimit
	queryTradeBuyLimit(symbol, stringPrice, stringBalance)
	// createBuyQuery()
	//on response calculate success or failure.

	//system splint for watching for sell
}

// func createBuyQuery(){
// 	json:=
// 	return
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
	buffer := balance * .01
	bufferedBalance := balance - buffer
	//Want to leave enough account balance to buffer for variance in buying.
	amountToBuy := bufferedBalance / pricePointOfStock
	amountToBuy = amountToBuy - buffer

	return amountToBuy
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
	// type DayTrackingRecord struct {
	// 	CreatedAt              string
	// 	Symbol                 string
	// 	DayOfWeekCreated       string
	// 	DayOfWeekDayIteration  string
	// 	LastDayOfWeekDayUpdate string
	// 	AmountOfTrades         string
	// }
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
