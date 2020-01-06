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
	isUnsettledFunds := "true"
	//Reset dow day eval store before calculation
	truncateCashDayEvaluation()

	amountUnsettledFunds := calculateIsUnsettledFunds()
	floatAmountUnsettledFunds := 0.0
	if s, err := strconv.ParseFloat(amountUnsettledFunds, 64); err == nil {
		floatAmountUnsettledFunds = s
	}
	if floatAmountUnsettledFunds <= 100.0 {
		isUnsettledFunds = "false"
	}

	// fmt.Println("amountUnsettledFunds")
	// fmt.Println(amountUnsettledFunds)

	cashDayEvaluation := CashDayEvaluation{IsUnsettledFunds: isUnsettledFunds}
	// fmt.Println("cashDayEvaluation")
	// fmt.Println(cashDayEvaluation)
	insertCashDayEvaluation(cashDayEvaluation)
}

func handleCalculateDownDay() {
	isDownDay := "true"
	//Reset dow day eval store before calculation
	truncateDownDayEvaluation()
	//
	dowValue := handleDowWebscrape()
	endOfDayDowList := selectEndOfDayDow()

	//clear end of day dow if list not empty
	if len(endOfDayDowList) != 0 {
		truncateEndOfDayDow()
	}
	//if it equals 0

	//
	downDayEvaluation := DownDayEvaluation{}
	//if TRS is not empty
	if len(endOfDayDowList) != 0 {
		//get latest TRS
		endOfDayDowPulled := endOfDayDowList[len(endOfDayDowList)-1]
		listEndOfDayDow := []EndOfDayDow{endOfDayDowPulled}

		for i, endOfDayDow := range listEndOfDayDow {
			if endOfDayDow.EndOfDayDowValue != "" {
				fmt.Println("hit")
				fmt.Println("dowValue")
				fmt.Println(dowValue)
				fmt.Println("endOfDayDow.EndOfDayDowValue")
				fmt.Println(endOfDayDow.EndOfDayDowValue)
				if dowValue > endOfDayDow.EndOfDayDowValue {
					isDownDay = "false"
					downDayEvaluation = DownDayEvaluation{IsDownDay: isDownDay, Dow: dowValue, PreviousDow: endOfDayDow.EndOfDayDowValue} //tradeResultStore.Dow4}
					break
				}
				if dowValue <= endOfDayDow.EndOfDayDowValue {
					isDownDay = "true"
					downDayEvaluation = DownDayEvaluation{IsDownDay: isDownDay, Dow: dowValue, PreviousDow: endOfDayDow.EndOfDayDowValue} //tradeResultStore.Dow4}
					break
				}
			}
			i++
		}
	}
	if len(endOfDayDowList) == 0 {
		downDayEvaluation = DownDayEvaluation{IsDownDay: isDownDay, Dow: dowValue, PreviousDow: "does not exist"}
	}
	//store results in DB
	insertDownDayEvaluation(downDayEvaluation)
}

// if tradeResultStore.Dow6 != "" {
// 	if dowValue > tradeResultStore.Dow6 {
// 		isDownDay = "false"
// 		downDayEvaluation = DownDayEvaluation{IsDownDay: isDownDay, Dow: dowValue, PreviousDow: "6"} //tradeResultStore.Dow6}
// 		break
// 	}
// }
// if tradeResultStore.Dow5 != "" {
// 	if dowValue > tradeResultStore.Dow5 {
// 		isDownDay = "false"
// 		downDayEvaluation = DownDayEvaluation{IsDownDay: isDownDay, Dow: dowValue, PreviousDow: "5"} //tradeResultStore.Dow5}
// 		break
// 	}
// }

// handleDownDayEmail()

func overarchIsTradeDay() string {
	isTradeDay := "false"
	//query down day eval
	downDayEvalList := selectDownDayEvaluation()
	//query cash day eval
	cashDayEvalList := selectCashDayEvaluation()
	isDownDay := downDayEvalList[0].IsDownDay
	isUnsettledFunds := cashDayEvalList[0].IsUnsettledFunds
	if isDownDay == "false" {
		if isUnsettledFunds == "false" {
			isTradeDay = "true"
		}
	}
	postIsTradeDayEmail(isTradeDay, isDownDay, isUnsettledFunds)
	return isTradeDay
}

func calculateShortDayAnalysis() {
	shortDayAnalysis := ShortDayAnalysis{}
	listShortCalendarDay := selectShortCalendarDay()
	isShortDay := "false"
	year, month, day := getDate()

	stringMonth := strconv.Itoa(month)
	stringDay := strconv.Itoa(day)
	// stringMonth = "7"
	// stringDay = "2"

	fmt.Println(year)
	for i, v := range listShortCalendarDay {
		listSplitDate := strings.Split(v.DateOfOccurance, " ")
		if listSplitDate[0] == stringMonth {
			if listSplitDate[1] == stringDay {
				isShortDay = "true"
			}
		}
		i++
	}
	shortDayAnalysis.IsShortDay = isShortDay
	insertShortDayAnalysis(shortDayAnalysis)

	if shortDayAnalysis.IsShortDay == "true" {
		conditionNineteenMinute = 59
		conditionNineteenHour = 10
	}
}

func handleTSPCollectionStatementPhase1() {
	//This is at a certain time after dow and wisemen selection.
	//
	stringTSPCollectionStatementCache := calculateTSPCollectionStatementString()

	//write to cache
	globalTSPCollectionStatementCache = append(globalTSPCollectionStatementCache, stringTSPCollectionStatementCache)
}
func handleTSPCollectionStatementPhase2() {
	//retrieve from cache
	stringTSPCollectionStatementCache := globalTSPCollectionStatementCache[0]

	stringTSPCollectionStatementCache1 := calculateTSPCollectionStatementString()

	stringTSPCollectionStatementCache += stringTSPCollectionStatementCache1
	//clear cache
	globalTSPCollectionStatementCache = []string{}
	fmt.Println("globalTSPCollectionStatementCache cleared")
	fmt.Println(globalTSPCollectionStatementCache)
	//persist
	instanceTSPCollectionStatement := TSPCollectionStatement{DataCache: stringTSPCollectionStatementCache}
	fmt.Println(instanceTSPCollectionStatement)
	insertTSPCollectionStatement(instanceTSPCollectionStatement)
}

func calculateTSPCollectionStatementString() string {
	stringTSPCollectionStatementCache := ""

	marketOpenAnalysis := selectMarketOpenAnalysis()[0]
	stringTSPCollectionStatementCache += marketOpenAnalysis.IsMarketClosed + "~"
	//
	dowList := selectDow()

	formatedDowList := formatDowListRemoveCommaValues(dowList)

	for i, v := range formatedDowList {
		stringTSPCollectionStatementCache += v.CurrentDowValue

		if i == (len(dowList) - 1) {
			stringTSPCollectionStatementCache += "~"
			break
		}
		stringTSPCollectionStatementCache += " "
	}

	wisemenSymbolList := selectWisemenSymbolHold()

	//retrieve symbol values
	stocKResponse := queryMultiStockPull(wisemenSymbolList)
	stockList := parseStockSetQuery(stocKResponse)

	for i, v := range stockList {
		stockSymbol := v.Symbol
		stockLast := v.Last
		stockPchg := v.Pchg
		stringTSPCollectionStatementCache += stockSymbol + " " + stockLast + " " + stockPchg

		if i == (len(stockList) - 1) {
			stringTSPCollectionStatementCache += "~"
			break
		}
		stringTSPCollectionStatementCache += " "
		i++
	}

	fmt.Println("stringTSPCollectionStatementCache")
	fmt.Println(stringTSPCollectionStatementCache)
	return stringTSPCollectionStatementCache
}

//

// currentDay :=
// for i, v := range listShortCalendarDay {

// 	listSplitDate := strings.Split(v.DateOfOccurance, " ")
// 	if listSplitDate[0] == {
// 		if listSplitDate[1] {
// 		}
// 	}
// 	i++
// }

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
