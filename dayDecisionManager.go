package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// var holidayMap = make(map[string]bool) {}
var holidayMap = map[string]string{"New Years": "12 31", "MLK": "1 20", "Presidents": "2 17", "Good Friday": "4 10", "Memorial": "5 25", "Independence": "7 3", "Labor": "9 2", "Thanksgiving": "11 26", "Christmas": "12 25"}
var shortDayMap = map[string]string{"Pre Independence day": "7 2", "Post Thanksgiving Black Friday": "11 27", "Pre Christmas": "12 24"}

func checkIfHoliday() {
	//Check holiday list if case match, return true
	// parseDate()

	for k, v := range holidayMap {
		fmt.Println(k, v)
	}
	for k, v := range shortDayMap {
		fmt.Println(k, v)
	}
	// holidayMap[""] = true
	// fmt.Println(holidayMap)

	// n := map[string]int{"foo": 1, "bar": 2}
	// fmt.Println("map:", n)
}

// handleOverarchTopStock
func checKIsBrokerageResponding() {
	//Multi-stock query simulation
	response := queryIsBrokerageResponding()
	askTime := parseAskTimeQuery(response)
	isMarketClosed := false
	// fmt.Println("asktime is:", askTime)
	//Conditional if Ask time set to 0
	isMarketClosed = checkConditionalIsAskTime(askTime)
	if isMarketClosed == false {
		fmt.Println("isMarketClosed is false")
		marketOpenAnalysis := MarketOpenAnalysis{IsMarketClosed: "false"}
		insertMarketOpenAnalysis(marketOpenAnalysis)
	}
	if isMarketClosed == true {
		fmt.Println("marketClosed is true")
		marketOpenAnalysis := MarketOpenAnalysis{IsMarketClosed: "true"}
		insertMarketOpenAnalysis(marketOpenAnalysis)
	}
}

func checkConditionalIsAskTime(askTime string) bool {
	boolReturned := false
	// fmt.Println(askTime)
	if askTime == "00:00" {
		boolReturned = true
	}
	return boolReturned
}

func getDate() (int, int, int) {
	currentTime := time.Now()
	yr, mt, day := currentTime.Date()
	intMonth := int(mt)
	fmt.Println(yr, mt, day)
	// day = 27
	return yr, intMonth, day
}

func getCurrentHour() int {
	hour := time.Now().Hour()
	return hour
}

func getCurrentMinute() int {
	minute := time.Now().Minute()
	return minute
}

func getDayOfWeek() time.Weekday {
	weekday := time.Now().Weekday()
	// fmt.Println(weekday)      // "Tuesday"
	// fmt.Println(int(weekday)) // "2"
	return weekday
}

func handleDayRotation() {
	listDayTrackingRecord := selectDayTrackingRecord()
	if len(listDayTrackingRecord) != 0 {
		dayTrackingRecord := listDayTrackingRecord[0]
		if strings.Contains(dayTrackingRecord.IsWeekPassed, "true") {
			// dropDayTrackingRecord()
			// createDayTrackingRecord()
			truncateDayTrackingRecord()
		}
		if strings.Contains(dayTrackingRecord.IsWeekPassed, "false") {
			currentDayOfWeek := getDayOfWeek().String()
			intCurrentDayOfWeek, err := strconv.ParseInt(currentDayOfWeek, 10, 64)
			if err != nil {
				fmt.Println(err)
			}
			intLastDayOfWeekDayUpdate, err := strconv.ParseInt(dayTrackingRecord.LastDayOfWeekDayUpdate, 10, 64)
			if err != nil {
				fmt.Println(err)
			}
			if intCurrentDayOfWeek == intLastDayOfWeekDayUpdate {
				instanceDayTrackingRecord := dayTrackingRecord
				instanceDayTrackingRecord.IsWeekPassed = "true"
				// dropDayTrackingRecord()
				// createDayTrackingRecord()
				truncateDayTrackingRecord()
				insertDayTrackingRecord(instanceDayTrackingRecord)
			}
		}
	}
}
func processInsertDayTrackingRecord(symbol string) {
	listDayTrackingRecord := selectDayTrackingRecord()
	//if list is empty create entry
	if len(listDayTrackingRecord) == 0 {

	}
	//if list is not empty update day record
	if len(listDayTrackingRecord) != 0 {
		// dropDayTrackingRecord()
		// createDayTrackingRecord()
		truncateDayTrackingRecord()
		updateDayTrackingRecordSystem(symbol)
	}
}

func Abs(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}

// func handleOverarchUnsettledFunds() {
// 	isUnsettledFunds := isCashAccountCheckUnsettledFunds()
// 	if isUnsettledFunds == false {
// 		//no unsettled funds
// 	}
// }
//
func handleCalculateCashDay() {
	//Reset dow day eval store before calculation
	truncateCashDayEvaluation()
	isUnsettledFunds := "true"
	//query account
	isUnsettledFunds = "false"
	// }
	cashDayEvaluation := CashDayEvaluation{IsUnsettledFunds: isUnsettledFunds}
	insertCashDayEvaluation(cashDayEvaluation)
}

func handleCalculateDownDay() {
	//Reset dow day eval store before calculation
	truncateDownDayEvaluation()
	isDownDay := "true"
	//is dow in the red
	//handle on dow...
	dowValue := handleDowWebscrape()
	tradeResultStoreList := getTradeResultStoreList()
	downDayEvaluation := DownDayEvaluation{}
	//if TRS is not empty
	if len(tradeResultStoreList) != 0 {
		//get latest TRS
		tradeResultStorePulled := tradeResultStoreList[len(tradeResultStoreList)-1]
		listCalculationTradeResultStore := []TradeResultStore{tradeResultStorePulled}

		for i, tradeResultStore := range listCalculationTradeResultStore {
			if tradeResultStore.Dow6 != "" {
				if dowValue > tradeResultStore.Dow6 {
					isDownDay = "false"
					downDayEvaluation = DownDayEvaluation{IsDownDay: isDownDay, Dow: dowValue, PreviousDow: "6"} //tradeResultStore.Dow6}
					break
				}
			}
			if tradeResultStore.Dow5 != "" {
				if dowValue > tradeResultStore.Dow5 {
					isDownDay = "false"
					downDayEvaluation = DownDayEvaluation{IsDownDay: isDownDay, Dow: dowValue, PreviousDow: "5"} //tradeResultStore.Dow5}
					break
				}
			}
			if tradeResultStore.Dow4 != "" {
				if dowValue > tradeResultStore.Dow4 {
					isDownDay = "false"
					downDayEvaluation = DownDayEvaluation{IsDownDay: isDownDay, Dow: dowValue, PreviousDow: "4"} //tradeResultStore.Dow4}
					break
				}
			}
			if tradeResultStore.Dow4 == "" {
				downDayEvaluation = DownDayEvaluation{IsDownDay: isDownDay, Dow: dowValue, PreviousDow: "3"} //tradeResultStore.Dow4}
				break
			}
			i++
		}
	}
	if len(tradeResultStoreList) == 0 {
		downDayEvaluation = DownDayEvaluation{IsDownDay: isDownDay, Dow: dowValue, PreviousDow: "does not exist"}
	}
	//store results in DB
	insertDownDayEvaluation(downDayEvaluation)
}

func overarchIsTradeDay() bool {
	isTradeDay := false
	//query down day eval
	// downDayEvalList := selectDownDayEvaluation()
	// //query cash day eval
	// cashDayEvalList := selectCashDayEvaluation()
	// if downDayEvalList[0].IsDownDay == "false" {
	// if cashDayEvalList[0].IsUnsettledFunds == "false" {
	isTradeDay = true
	// }
	// }
	return isTradeDay
}

// func handleCalculateCashDay() {
// 	//Reset dow day eval store before calculation
// 	dropCashDayEvaluation()
// 	createCashDayEvaluation()
// 	isCashDay := "true"
// 	isCashDown := false
// 	isTopStockAbovePchgDelimiter := false
// 	//is dow in the red
// 	//handle on dow...
// 	dowValue := handleDowWebscrape()

// 	// pchgFromPreviousDay...pull from TRS dow4
// 	//query trs from previous day...
// 	tradeResultStore := getLatestTradeResultStore()
// 	if dowValue < tradeResultStore.Dow4 {
// 		isDowDown = true
// 	}
// 	// //query highest pchg, is pchg greater than delmiter...
// 	// //do TSP, doesn't matter if it's recurrent,
// 	// //TSP get topstock...
// 	wisemenSymbolList := selectWisemenSymbolHold()
// 	response := queryMultiStockPull(wisemenSymbolList)
// 	stockList := parseStockSetQuery(response)
// 	// //sort for highest...
// 	highestStockIndex := 0
// 	for indexStock, stock := range stockList {
// 		if indexStock == 0 {
// 			highestStockIndex = indexStock
// 			continue
// 		}
// 		floatHighest := 0.0
// 		floatCurrent := 0.0
// 		if s, err := strconv.ParseFloat(stockList[highestStockIndex].Pchg, 64); err == nil {
// 			floatHighest = s
// 		}
// 		if s1, err := strconv.ParseFloat(stock.Pchg, 64); err == nil {
// 			floatCurrent = s1
// 		}
// 		if floatCurrent > floatHighest {
// 			highestStockIndex = indexStock
// 		}
// 	}
// 	//if topStop greater than delmiter.
// 	// topStock
// 	//
// 	floatTopStockPchg := 0.0
// 	if s, err := strconv.ParseFloat(stockList[highestStockIndex].Pchg, 64); err == nil {
// 		floatTopStockPchg = s
// 	}
// 	//static wismenDownDayDelimiter 10
// 	wismenDownDayDelimiter := 10.00
// 	//support for dynamic delimiter store
// 	//is highest pchg greater than delmiter.
// 	if floatTopStockPchg >= wismenDownDayDelimiter {
// 		isTopStockAbovePchgDelimiter = true
// 	}
// 	if isTopStockAbovePchgDelimiter {
// 		if isDowDown == false {
// 			isDownDay = "false"
// 		}
// 	}
// 	// fmt.Println("isDowDown")
// 	// fmt.Println(isDowDown)
// 	// fmt.Println("isTopStockAbovePchgDelimiter")
// 	// fmt.Println(isTopStockAbovePchgDelimiter)
// 	// fmt.Println("isDownDay")
// 	// fmt.Println(isDownDay)
// 	//store results in DB
// 	downDayEvaluation := DownDayEvaluation{IsDownDay: isDownDay}
// 	insertDownDayEvaluation(downDayEvaluation)
// }
