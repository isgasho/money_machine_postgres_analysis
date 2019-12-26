package main

import (
	"fmt"
	"strconv"
)

// type AlgorithmEvaluationForDay struct {
// 	Name          string
// 	TimeStart     string
// 	TimeEnd       string
// 	IsCompleted   string
// 	IsProfitable  string
// 	BalanceBefore string
// 	BalanceAfter  string
// }

func beginAlgorithmRecordingWisemen(symbol string) {
	minute := strconv.Itoa(getCurrentMinute())
	hour := strconv.Itoa(getCurrentHour())
	currentTIme := hour + " " + minute
	balance := queryBalance()
	parsedBalance := parseBalance(balance)
	algorithmEvaluationForDay := AlgorithmEvaluationForDay{
		Name:          "Wisemen",
		Symbol:        symbol,
		TimeStart:     currentTIme,
		TimeEnd:       "unknown",
		IsCompleted:   "false",
		IsProfitable:  "unknown",
		BalanceBefore: parsedBalance,
		BalanceAfter:  "unknown",
	}
	// fmt.Println("algorithmEvaluationForDay")
	fmt.Println(algorithmEvaluationForDay)
	// insertAlgorithmEvaluationForDay(algorithmEvaluationForDay)
}
func completeAlgorithmRecordingWisemen() {
	// minute := strconv.Itoa(getCurrentMinute())
	// hour := strconv.Itoa(getCurrentHour())
	// currentTIme := minute + " " + hour
	// balance := queryBalance()
	// parsedBalance := parseBalance(balance)

	// algorithmEvaluationForDay := selectAlgorithmEvaluationForDay("test")
	// fmt.Println(algorithmEvaluationForDay)
}

func writeToTradeResultCSV() {

}

//Write to csv file upload csv file to drive.
