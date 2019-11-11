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
var isMarketClosed = false

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

func checKIsBrokerageResponding() {
	//Multi-stock query simulation
	response := queryIsBrokerageResponding()
	askTime := parseAskTimeQuery(response)
	fmt.Println("asktime is:", askTime)
	//Conditional if Ask time set to 0
	checkConditionalIsAskTime(askTime)
	if isMarketClosed == false {
		fmt.Println("isMarketClosed is false")
	}
	if isMarketClosed == true {
		fmt.Println("marketClosed is true")
	}
}

func checkConditionalIsAskTime(askTime string) {
	fmt.Println(askTime)
	if askTime == "00:00" {
		isMarketClosed = true
	}
}

func getDate() (int, int, int) {
	currentTime := time.Now()
	yr, mt, day := currentTime.Date()
	intMonth := int(mt)
	fmt.Println(yr, mt, day)
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
	fmt.Println(weekday)      // "Tuesday"
	fmt.Println(int(weekday)) // "2"
	return weekday
}

func handleDayRotation() {
	listDayTrackingRecord := selectDayTrackingRecord()
	if len(listDayTrackingRecord) != 0 {
		dayTrackingRecord := listDayTrackingRecord[0]
		if strings.Contains(dayTrackingRecord.IsWeekPassed, "true") {
			dropDayTrackingRecord()
			createDayTrackingRecord()
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
				dropDayTrackingRecord()
				createDayTrackingRecord()
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
		dropDayTrackingRecord()
		createDayTrackingRecord()
		updateDayTrackingRecordSystem(symbol)
	}
}

func Abs(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}
