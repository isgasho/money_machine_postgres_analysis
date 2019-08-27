package main

import (
	"fmt"
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

func parseDate() {
	currentTime := time.Now()
	yr, mt, day := currentTime.Date()
	// dateSplit := strings.Split(day, " ")
	fmt.Println(yr)
	fmt.Println(mt, day)
}

func getDayOfWeek() time.Weekday {
	// currentTime := time.Now()
	// day := fmt.Sprintf("%b", currentTime.Day())
	// day := currentTime.Weekday()

	// weekday := time.Now().Weekday()
	// fmt.Println("hit day " + string(day)
	weekday := time.Now().Weekday()
	fmt.Println(weekday)      // "Tuesday"
	fmt.Println(int(weekday)) // "2"

	return weekday
}
