package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func parseQuery(queryString string) Stock {
	splitDataQuery1 := strings.Split(queryString, "<quote>")[1]
	splitDataQuery2 := strings.Split(splitDataQuery1, "<symbol>")[1]
	symbol := strings.Split(splitDataQuery2, "</symbol>")[0]

	splitDataQuery2 = strings.Split(splitDataQuery1, "<bid>")[1]
	bid := strings.Split(splitDataQuery2, "</bid>")[0]

	splitDataQuery2 = strings.Split(splitDataQuery1, "<ask>")[1]
	ask := strings.Split(splitDataQuery2, "</ask>")[0]

	splitDataQuery2 = strings.Split(splitDataQuery1, "<last>")[1]
	last := strings.Split(splitDataQuery2, "</last>")[0]

	splitDataQuery2 = strings.Split(splitDataQuery1, "<pchg>")[1]
	pchg := strings.Split(splitDataQuery2, "</pchg>")[0]

	splitDataQuery2 = strings.Split(splitDataQuery1, "<pcls>")[1]
	pcls := strings.Split(splitDataQuery2, "</pcls>")[0]

	stock := Stock{
		Symbol: symbol,
		Bid:    bid,
		Ask:    ask,
		Last:   last,
		Pchg:   pchg,
		Pcls:   pcls,
	}
	return stock
}

// func parseBalanceQuery(queryString string) Holdings {
// 	splitDataQuery1 := strings.Split(queryString, "<sym>")[1]
// 	symbol := strings.Split(splitDataQuery1, "</sym>")[0]

// 	splitDataQuery1 = strings.Split(queryString, "<price>")[1]
// 	price := strings.Split(splitDataQuery1, "</price>")[0]

// 	splitDataQuery1 = strings.Split(queryString, "<qty>")[1]
// 	qty := strings.Split(splitDataQuery1, "</qty>")[0]

// 	holdings := Holdings{
// 		Symbol: symbol,
// 		Price:  price,
// 		Qty:    qty,
// 	}
// 	return holdings
// }

func parseStockSetQuery(queryString string) []Stock {
	splitDataQuery1 := strings.Split(queryString, "<quotes>")[1]
	splitDataQuery2 := strings.Split(splitDataQuery1, "</quotetype>")[1]
	splitDataQuery3 := strings.Split(splitDataQuery2, "</quotes>")[0]
	splitDataListQuery4 := strings.Split(splitDataQuery3, "<quote>")

	for i := range splitDataListQuery4 {
		if i == 0 {
			splitDataListQuery4 = append(splitDataListQuery4[:i], splitDataListQuery4[i+1:]...)
			break
		}
	}
	parseList := []string{}
	for i, v := range splitDataListQuery4 {
		i++
		stringParsed := strings.Split(v, "</quote>")[0]
		parseList = append(parseList, stringParsed)
	}

	var stockList = []Stock{}

	//
	// fmt.Println("parseList")
	// for i, v := range parseList {
	// 	fmt.Println(v)
	// 	i++
	// }

	for i, v := range parseList {
		//Create stock and append to composite
		// fmt.Println(i, v)
		i++

		symbolUnparsed := strings.Split(v, "<symbol>")[1]
		symbol := strings.Split(symbolUnparsed, "</symbol>")[0]

		bidUnparsed := strings.Split(v, "<bid>")[1]
		bid := strings.Split(bidUnparsed, "</bid>")[0]

		askUnparsed := strings.Split(v, "<ask>")[1]
		ask := strings.Split(askUnparsed, "</ask>")[0]

		lastUnparsed := strings.Split(v, "<last>")[1]
		last := strings.Split(lastUnparsed, "</last>")[0]

		pchgUnparsed := strings.Split(v, "<pchg>")[1]
		pchg := strings.Split(pchgUnparsed, "</pchg>")[0]
		// fmt.Println("pchg")
		// fmt.Println(pchg)

		pclsUnparsed := strings.Split(v, "<pcls>")[1]
		pcls := strings.Split(pclsUnparsed, "</pcls>")[0]

		opnUnparsed := strings.Split(v, "<opn>")[1]
		opn := strings.Split(opnUnparsed, "</opn>")[0]

		vlUnparsed := strings.Split(v, "<vl>")[1]
		vl := strings.Split(vlUnparsed, "</vl>")[0]

		pvolUnparsed := strings.Split(v, "<pvol>")[1]
		pvol := strings.Split(pvolUnparsed, "</pvol>")[0]

		volatility12Unparsed := strings.Split(v, "<volatility12>")[1]
		volatility12 := strings.Split(volatility12Unparsed, "</volatility12>")[0]

		wk52hiUnparsed := strings.Split(v, "<wk52hi>")[1]
		wk52hi := strings.Split(wk52hiUnparsed, "</wk52hi>")[0]

		wk52hidateUnparsed := strings.Split(v, "<wk52hidate>")[1]
		wk52hidate := strings.Split(wk52hidateUnparsed, "</wk52hidate>")[0]

		wk52loUnparsed := strings.Split(v, "<wk52lo>")[1]
		wk52lo := strings.Split(wk52loUnparsed, "</wk52lo>")[0]

		wk52lodateUnparsed := strings.Split(v, "<wk52lodate>")[1]
		wk52lodate := strings.Split(wk52lodateUnparsed, "</wk52lodate>")[0]

		hiUnparsed := strings.Split(v, "<hi>")[1]
		hi := strings.Split(hiUnparsed, "</hi>")[0]

		loUnparsed := strings.Split(v, "<lo>")[1]
		lo := strings.Split(loUnparsed, "</lo>")[0]

		prAdp50Unparsed := strings.Split(v, "<pr_adp_50>")[1]
		prAdp50 := strings.Split(prAdp50Unparsed, "</pr_adp_50>")[0]

		prAdp100Unparsed := strings.Split(v, "<pr_adp_100>")[1]
		prAdp100 := strings.Split(prAdp100Unparsed, "</pr_adp_100>")[0]

		prchgUnparsed := strings.Split(v, "<prchg>")[1]
		prchg := strings.Split(prchgUnparsed, "</prchg>")[0]

		adp50Unparsed := strings.Split(v, "<adp_50>")[1]
		adp50 := strings.Split(adp50Unparsed, "</adp_50>")[0]

		adp100Unparsed := strings.Split(v, "<adp_100>")[1]
		adp100 := strings.Split(adp100Unparsed, "</adp_100>")[0]

		adv30Unparsed := strings.Split(v, "<adv_30>")[1]
		adv30 := strings.Split(adv30Unparsed, "</adv_30>")[0]

		adv90Unparsed := strings.Split(v, "<adv_90>")[1]
		adv90 := strings.Split(adv90Unparsed, "</adv_90>")[0]

		//calculated createdAt
		createdAt := createStockTimeStamp()

		// isCurrentPriceHigherThanPreviousClose calculation
		isCurrentPriceHigherThanPreviousClose := "false"

		pclsFloat := 0.0
		lastFloat := 0.0
		//convert pcls and last to float.
		if s, err := strconv.ParseFloat(pcls, 64); err == nil {
			pclsFloat = s
		}
		if s1, err := strconv.ParseFloat(last, 64); err == nil {
			lastFloat = s1
		}

		if pclsFloat < lastFloat {
			isCurrentPriceHigherThanPreviousClose = "true"
		}

		var stock = Stock{
			Monitoring:                            false,
			Symbol:                                symbol,
			CreatedAt:                             createdAt,
			Bid:                                   bid,
			Ask:                                   ask,
			Last:                                  last,
			Pchg:                                  pchg,
			Pcls:                                  pcls,
			Opn:                                   opn,
			Vl:                                    vl,
			Pvol:                                  pvol,
			Volatility12:                          volatility12,
			Wk52hi:                                wk52hi,
			Wk52hidate:                            wk52hidate,
			Wk52lo:                                wk52lo,
			Wk52lodate:                            wk52lodate,
			Hi:                                    hi,
			Lo:                                    lo,
			PrAdp50:                               prAdp50,
			PrAdp100:                              prAdp100,
			Prchg:                                 prchg,
			Adp50:                                 adp50,
			Adp100:                                adp100,
			Adv30:                                 adv30,
			Adv90:                                 adv90,
			IsCurrentPriceHigherThanPreviousClose: isCurrentPriceHigherThanPreviousClose,
		}
		stockList = append(stockList, stock)
	}
	return stockList
}
func parseDBResponse(response string) DBResponseContainer {
	// fmt.Println("begin")
	instanceDBResponseContainer := DBResponseContainer{}
	listObjectString := strings.Split(response, "\"")[1]
	listParsedString := strings.Split(listObjectString, "|")
	copyList := listParsedString
	// fmt.Println("len(copyList)")
	// fmt.Println(len(copyList))

	// fmt.Println("response")
	// fmt.Println(response)
	for i, v := range listParsedString {
		// fmt.Println(v)
		if len(v) == 0 {
			copyList = removeElementFromListAtIndex(copyList, i)
		}
	}
	// fmt.Println("without")
	// fmt.Println("len(copyList)")
	// fmt.Println(len(copyList))

	// fmt.Println("copyList")
	// fmt.Println(copyList)
	for i, v := range copyList {
		listValues := strings.Split(v, ",")
		// fmt.Println("listValues")
		if len(listValues) != 0 {
			stringResponse := StringResponse{ListString: listValues}
			instanceDBResponseContainer.ListStringFromDB = append(instanceDBResponseContainer.ListStringFromDB, stringResponse)
		}
		// fmt.Println("listIn")
		// fmt.Println(len(listValues))
		// fmt.Println(listValues)
		i++
	}
	return instanceDBResponseContainer
}

func parseDBResponseDow(response string) DBResponseContainer {
	// fmt.Println("begin")
	instanceDBResponseContainer := DBResponseContainer{}
	listObjectString := strings.Split(response, "\"")[1]
	listParsedString := strings.Split(listObjectString, "|")
	copyList := listParsedString
	// fmt.Println("len(copyList)")
	// fmt.Println(len(copyList))

	// fmt.Println("response")
	// fmt.Println(response)
	for i, v := range listParsedString {
		// fmt.Println(v)
		if len(v) == 0 {
			copyList = removeElementFromListAtIndex(copyList, i)
		}
	}
	// fmt.Println("without")
	// fmt.Println("len(copyList)")
	// fmt.Println(len(copyList))

	// fmt.Println("copyList")
	// fmt.Println(copyList)
	for i, v := range copyList {
		// listValues := strings.Split(v, ",")
		// fmt.Println("listValues")
		// if len(listValues) != 0 {
		listValues := []string{v}
		stringResponse := StringResponse{ListString: listValues}
		instanceDBResponseContainer.ListStringFromDB = append(instanceDBResponseContainer.ListStringFromDB, stringResponse)
		// }
		// fmt.Println("listIn")
		// fmt.Println(len(listValues))
		// fmt.Println(listValues)
		i++
	}
	return instanceDBResponseContainer
}

func removeElementFromListAtIndex(listEntered []string, val int) []string {
	// var i int
	listAltered := listEntered
	listAltered = listAltered[:val+copy(listAltered[val:], listAltered[val+1:])]
	return listAltered
}

func createStockTimeStamp() string {
	// currentTime := time.Now()
	year, month, day := time.Now().Date()
	fmt.Println(year)
	monthInt := int(month)
	monthString := strconv.Itoa(monthInt)
	dayString := strconv.Itoa(day)
	// fmt.Println(day)
	// timeStamp := currentTime.Date()
	timeStamp := monthString + " " + dayString + " " //:= string(int(month)) + string(day)
	// stringAppend :=
	return timeStamp
}

func parseAskTimeQuery(queryString string) string {
	splitDataQuery1 := strings.Split(queryString, "<quotes>")[1]
	splitDataQuery2 := strings.Split(splitDataQuery1, "</quotetype>")[1]
	splitDataQuery3 := strings.Split(splitDataQuery2, "</quotes>")[0]
	splitDataListQuery4 := strings.Split(splitDataQuery3, "<quote>")

	for i := range splitDataListQuery4 {
		if i == 0 {
			splitDataListQuery4 = append(splitDataListQuery4[:i], splitDataListQuery4[i+1:]...)
			break
		}
	}
	parseList := []string{}
	for i, v := range splitDataListQuery4 {
		i++
		stringParsed := strings.Split(v, "</quote>")[0]
		parseList = append(parseList, stringParsed)
	}
	askTimeUnparsed := strings.Split(parseList[0], "<ask_time>")[1]
	askTime := strings.Split(askTimeUnparsed, "</ask_time>")[0]

	return askTime
}

func parseMonitoredStockQuery(queryString string) []Stock {
	splitDataQuery1 := strings.Split(queryString, "<quotes>")[1]
	splitDataQuery2 := strings.Split(splitDataQuery1, "</quotetype>")[1]
	splitDataQuery3 := strings.Split(splitDataQuery2, "</quotes>")[0]
	splitDataListQuery4 := strings.Split(splitDataQuery3, "<quote>")

	for i := range splitDataListQuery4 {
		if i == 0 {
			splitDataListQuery4 = append(splitDataListQuery4[:i], splitDataListQuery4[i+1:]...)
			break
		}
	}

	// fmt.Println(splitDataListQuery4[0])

	parseList := []string{}
	for i, v := range splitDataListQuery4 {
		i++
		stringParsed := strings.Split(v, "</quote>")[0]
		parseList = append(parseList, stringParsed)
	}

	var stockList = []Stock{}

	// fmt.Println(parseList[0])
	for i, v := range parseList {
		//Create stock and append to composite
		// fmt.Println(i, v)
		i++

		symbolUnparsed := strings.Split(v, "<symbol>")[1]
		symbol := strings.Split(symbolUnparsed, "</symbol>")[0]

		bidUnparsed := strings.Split(v, "<bid>")[1]
		bid := strings.Split(bidUnparsed, "</bid>")[0]

		askUnparsed := strings.Split(v, "<ask>")[1]
		ask := strings.Split(askUnparsed, "</ask>")[0]

		lastUnparsed := strings.Split(v, "<last>")[1]
		last := strings.Split(lastUnparsed, "</last>")[0]

		pchgUnparsed := strings.Split(v, "<pchg>")[1]
		pchg := strings.Split(pchgUnparsed, "</pchg>")[0]

		pclsUnparsed := strings.Split(v, "<pcls>")[1]
		pcls := strings.Split(pclsUnparsed, "</pcls>")[0]

		opnUnparsed := strings.Split(v, "<opn>")[1]
		opn := strings.Split(opnUnparsed, "</opn>")[0]

		vlUnparsed := strings.Split(v, "<vl>")[1]
		vl := strings.Split(vlUnparsed, "</vl>")[0]

		pvolUnparsed := strings.Split(v, "<pvol>")[1]
		pvol := strings.Split(pvolUnparsed, "</pvol>")[0]

		volatility12Unparsed := strings.Split(v, "<volatility12>")[1]
		volatility12 := strings.Split(volatility12Unparsed, "</volatility12>")[0]

		wk52hiUnparsed := strings.Split(v, "<wk52hi>")[1]
		wk52hi := strings.Split(wk52hiUnparsed, "</wk52hi>")[0]

		wk52hidateUnparsed := strings.Split(v, "<wk52hidate>")[1]
		wk52hidate := strings.Split(wk52hidateUnparsed, "</wk52hidate>")[0]

		wk52loUnparsed := strings.Split(v, "<wk52lo>")[1]
		wk52lo := strings.Split(wk52loUnparsed, "</wk52lo>")[0]

		wk52lodateUnparsed := strings.Split(v, "<wk52lodate>")[1]
		wk52lodate := strings.Split(wk52lodateUnparsed, "</wk52lodate>")[0]

		hiUnparsed := strings.Split(v, "<hi>")[1]
		hi := strings.Split(hiUnparsed, "</hi>")[0]

		loUnparsed := strings.Split(v, "<lo>")[1]
		lo := strings.Split(loUnparsed, "</lo>")[0]

		prAdp50Unparsed := strings.Split(v, "<pr_adp_50>")[1]
		prAdp50 := strings.Split(prAdp50Unparsed, "</pr_adp_50>")[0]

		prAdp100Unparsed := strings.Split(v, "<pr_adp_100>")[1]
		prAdp100 := strings.Split(prAdp100Unparsed, "</pr_adp_100>")[0]

		prchgUnparsed := strings.Split(v, "<prchg>")[1]
		prchg := strings.Split(prchgUnparsed, "</prchg>")[0]

		adp50Unparsed := strings.Split(v, "<adp_50>")[1]
		adp50 := strings.Split(adp50Unparsed, "</adp_50>")[0]

		adp100Unparsed := strings.Split(v, "<adp_100>")[1]
		adp100 := strings.Split(adp100Unparsed, "</adp_100>")[0]

		adv30Unparsed := strings.Split(v, "<adv_30>")[1]
		adv30 := strings.Split(adv30Unparsed, "</adv_30>")[0]

		adv90Unparsed := strings.Split(v, "<adv_90>")[1]
		adv90 := strings.Split(adv90Unparsed, "</adv_90>")[0]

		// fmt.Println(dayID)
		var stock = Stock{
			Monitoring:   true,
			Symbol:       symbol,
			Bid:          bid,
			Ask:          ask,
			Last:         last,
			Pchg:         pchg,
			Pcls:         pcls,
			Opn:          opn,
			Vl:           vl,
			Pvol:         pvol,
			Volatility12: volatility12,
			Wk52hi:       wk52hi,
			Wk52hidate:   wk52hidate,
			Wk52lo:       wk52lo,
			Wk52lodate:   wk52lodate,
			Hi:           hi,
			Lo:           lo,
			PrAdp50:      prAdp50,
			PrAdp100:     prAdp100,
			Prchg:        prchg,
			Adp50:        adp50,
			Adp100:       adp100,
			Adv30:        adv30,
			Adv90:        adv90,
		}
		// fmt.Println(stock.Symbol, stock.Bid, stock.Ask)
		stockList = append(stockList, stock)
	}
	return stockList
}

func parseTopStockQuery(queryString string) []Stock {
	splitDataQuery1 := strings.Split(queryString, "<quotes>")[1]
	splitDataQuery2 := strings.Split(splitDataQuery1, "</quotes>")[0]
	// fmt.Println(splitDataQuery2)

	splitDataListQuery3 := strings.Split(splitDataQuery2, "<quote>")
	// fmt.Println(splitDataListQuery3)
	// fmt.Println(splitDataListQuery3)

	// fmt.Printf("Pop %d\n", splitDataListQuery3.Pop())

	// s := []string{"one", "two", "three"}
	splitDataListQuerySpaceIndexRemoved := splitDataListQuery3

	// Find and remove "two"
	for i := range splitDataListQuerySpaceIndexRemoved {
		// if v == "two" {
		// fmt.Println(v)
		// fmt.Println(i)

		if i == 0 {
			splitDataListQuerySpaceIndexRemoved = append(splitDataListQuerySpaceIndexRemoved[:i], splitDataListQuerySpaceIndexRemoved[i+1:]...)
			break
		}
	}

	// fmt.Println(splitDataListQuerySpaceIndexRemoved[0]) // Prints [one three]
	// list := []string{}

	parseList := []string{}
	for i, v := range splitDataListQuerySpaceIndexRemoved {
		// fmt.Println(i, v)
		i++
		stringParsed := strings.Split(v, "</quote>")[0]
		parseList = append(parseList, stringParsed)
		// splitDataListQuery4 := strings.Split(splitDataListQuerySpaceIndexRemoved, "<quote>")
	}

	var stockList = []Stock{}

	// fmt.Println(parseList[0])
	for i, v := range parseList {
		// isCurrentPriceHigherThanPreviousClose := false
		//Create stock and append to composite
		// fmt.Println(i, v)
		i++
		symbolParsed := strings.Split(v, "<symbol>")[1]
		symbolParsed1 := strings.Split(symbolParsed, "</symbol>")[0]
		// fmt.Println(symbolParsed1)

		pchgParsed := strings.Split(v, "<pchg>")[1]
		pchgParsed1 := strings.Split(pchgParsed, "</pchg>")[0]
		// fmt.Println(pchgParsed1)

		rankParsed := strings.Split(v, "<rank>")[1]
		rankParsed1 := strings.Split(rankParsed, "</rank>")[0]

		var stock = Stock{
			Symbol: symbolParsed1,
			Pchg:   pchgParsed1,
			Rank:   rankParsed1,
		}

		// fmt.Println(stock.Symbol, stock.Pchg, stock.Rank)

		stockList = append(stockList, stock)
		// stringParsed := strings.Split(v, "</quote>")[0]
		// parseList1 = append(parseList1, symbolParsed1)

		// return stockList
		// splitDataListQuery4 := strings.Split(splitDataListQuerySpaceIndexRemoved, "<quote>")
	}
	return stockList
}

const numbers = "1234567890.,"

func containsNumbers(s string) bool {
	for _, char := range s {
		if strings.Contains(numbers, strings.ToLower(string(char))) {
			return true
		}
	}
	return false
}

const punctuation = ".,"

func containsPunctuation(s string) bool {
	for _, char := range s {
		if strings.Contains(punctuation, strings.ToLower(string(char))) {
			return true
		}
	}
	return false
}

func containsMinimumSeriesNumbers(s string) bool {
	containerNumberRange := ContainerNumberRange{}
	webscrapeNumberRange := WebscrapeNumberRange{}
	for i, char := range s {
		// type ContainerNumberRange struct {
		// 	ListNumberRange []WebscrapeNumberRange
		// }
		// type WebscrapeNumberRange struct {
		// 	NumberRange []string
		// }
		if i == 0 {
			webscrapeNumberRange := WebscrapeNumberRange{}
			containerNumberRange.ListNumberRange = append(containerNumberRange.ListNumberRange, webscrapeNumberRange)
		}
		if len(containerNumberRange.ListNumberRange[len(containerNumberRange.ListNumberRange)-1].NumberRange) == 0 {
			if strings.Contains(numbers, strings.ToLower(string(char))) {
				// get latest WebscrapeNumberRange append to NumberRange
				containerNumberRange.ListNumberRange[len(containerNumberRange.ListNumberRange)-1].NumberRange = append(containerNumberRange.ListNumberRange[len(containerNumberRange.ListNumberRange)-1].NumberRange, string(char))
				continue
			}
		}
		if len(containerNumberRange.ListNumberRange[len(containerNumberRange.ListNumberRange)-1].NumberRange) > 0 {
			if strings.Contains(numbers, strings.ToLower(string(char))) {
				containerNumberRange.ListNumberRange[len(containerNumberRange.ListNumberRange)-1].NumberRange = append(containerNumberRange.ListNumberRange[len(containerNumberRange.ListNumberRange)-1].NumberRange, string(char))
				continue
			}
			if !strings.Contains(numbers, strings.ToLower(string(char))) {
				webscrapeNumberRange = WebscrapeNumberRange{}
				containerNumberRange.ListNumberRange = append(containerNumberRange.ListNumberRange, webscrapeNumberRange)
			}
		}
	}
	// listValidStrings := []string{}
	for i, v := range containerNumberRange.ListNumberRange {
		if len(v.NumberRange) == 9 {
			fmt.Println(v.NumberRange)
			// listValidStrings = append(listValidStrings, v.NumberRange)
			return true
		}
		i++
	}

	return false
}

func calculateIsMatchingDelimiterWebscrape(delimiterStringList []string, testingString string) bool {
	//calculate match delimiter
	if strings.Contains(testingString, delimiterStringList[0]) {
		fmt.Println(testingString)
		return true
	}
	if strings.Contains(testingString, delimiterStringList[1]) {
		fmt.Println(testingString)
		return true
	}
	return false
}

func calculateIndexMatchClosestToDelimiter(listSpanClasses []string, listIndexPossibleMatches []int) string {
	delimiterStringList := []string{"\"Up by ", "\"Down by "}
	// containerForDistance := []int{}

	// fmt.println()
	fmt.Println("listIndexPossibleMatches")
	fmt.Println(listIndexPossibleMatches)
	listWisemenMatchClosestToDelimiter := []WisemenMatchClosestToDelimiter{}
	for indexIndexPossibleMatch, indexPossibleMatch := range listIndexPossibleMatches {
		//starting from span class
		//iterate until delimiter met. Count distance traveled.
		distanceInt := 0
		for indexSpanClass, spanClass := range listSpanClasses {
			if indexSpanClass >= indexPossibleMatch {
				// fmt.Println("working spanClass")
				// fmt.Println(spanClass)
				if calculateIsMatchingDelimiterWebscrape(delimiterStringList, spanClass) {
					//for wisemenMatchClosestToDelimiter
					listWisemenMatchClosestToDelimiter = append(listWisemenMatchClosestToDelimiter, WisemenMatchClosestToDelimiter{SplitStringValue: indexPossibleMatch, DistanceFromDelimiter: distanceInt})
					// containerForDistance = append(containerForDistance, distanceInt)
					break
				}
				distanceInt++
			}
		}
		indexIndexPossibleMatch++
	}

	lowestIndex := 0
	// fmt.Println("v.DistanceFromDelimiter")
	for i, v := range listWisemenMatchClosestToDelimiter {
		if i == 0 {
			continue
		}
		if v.DistanceFromDelimiter < listWisemenMatchClosestToDelimiter[lowestIndex].DistanceFromDelimiter {
			lowestIndex = i
		}

		// fmt.Println(v.DistanceFromDelimiter)
		// fmt.Println(v.SplitStringValue)
		// i++

		// if i == 6 {
		// 	fmt.Println(v.SplitStringValue)
		// }
		i++
	}
	// fmt.Println("lowestIndex")
	// fmt.Println(listWisemenMatchClosestToDelimiter[lowestIndex].DistanceFromDelimiter)
	// fmt.Println(listSpanClasses[listWisemenMatchClosestToDelimiter[lowestIndex].SplitStringValue])
	// fmt.Println("listWisemenMatchClosestToDelimiter")
	// fmt.Println(len(listWisemenMatchClosestToDelimiter))

	return listSpanClasses[listWisemenMatchClosestToDelimiter[lowestIndex].SplitStringValue] //listWisemenMatchClosestToDelimiter[lowestIndex].SplitStringValue
}

func parseDowWebscrape(queryString string) string {
	listSpanClasses := strings.Split(queryString, "<span")
	listIndexPossibleMatches := []int{}
	for i, v := range listSpanClasses {
		isBool := containsMinimumSeriesNumbers(v)
		if isBool {
			listIndexPossibleMatches = append(listIndexPossibleMatches, i)
		}
	}
	stringClosestMatch := calculateIndexMatchClosestToDelimiter(listSpanClasses, listIndexPossibleMatches)

	currentDowValueQuery1 := strings.Split(stringClosestMatch, "</span>")[0]
	currentDowValueQuery2 := strings.Split(currentDowValueQuery1, "\">")[1]
	// fmt.Println(currentDowValueQuery2)
	currentDowValue := currentDowValueQuery2
	return currentDowValue
}

func parseTwiWebscrape(queryString string) []string {
	listSpanClasses := strings.Split(queryString, "</span>") //"<span")
	listPossibleMatches := []string{}
	listIndexCuts := []int{}
	listMatches := []string{}
	listSymbolBeforeParse := []string{}
	listSymbols := []string{}
	filterDotList := []string{}
	// for i, v := range listSpanClasses {
	// 	isBool := containsMinimumSeriesNumbers(v)
	// 	if isBool {
	// 		listIndexPossibleMatches = append(listIndexPossibleMatches, i)
	// 	}
	// }
	// stringClosestMatch := calculateIndexMatchClosestToDelimiter(listSpanClasses, listIndexPossibleMatches)

	fmt.Println("hit twiparse")
	for i, v := range listSpanClasses {
		// if strings.Contains(v, "</div></a><a href=\"/symbol/") {
		// 	listPossibleMatches = append(listPossibleMatches, v)
		// }
		if strings.Contains(v, "/symbol/") {
			listPossibleMatches = append(listPossibleMatches, v)
		}
		i++
	}
	// listIndexPossibleMatches
	// fmt.Println(len(listPossibleMatches))
	for i, v := range listPossibleMatches {
		if strings.Contains(v, "/svg>") {
			fmt.Println("inside yo")
			listIndexCuts = append(listIndexCuts, i)
		}
		i++
	}
	// fmt.Println(len(listPossibleMatches))
	fmt.Println("hit")
	// fmt.Println(listPossibleMatches[listIndexCuts[len(listIndexCuts)-1]])

	// fmt.Println(listIndexCuts[len(listIndexCuts)-1])
	for i, v := range listPossibleMatches {
		// stringIndex := listIndexCuts[len(listIndexCuts)-1]
		// valueIndex, err := //strconv.Atoi(stringIndex)
		// if err != nil {
		// 	fmt.Println(err)
		// }

		// listIndexCuts[len(listIndexCuts)-1]

		if i >= listIndexCuts[len(listIndexCuts)-1] {
			//</div></a><a href=\"/symbol/
			if strings.Contains(v, "</div></a><a href=\\\"/symbol/") {
				listMatches = append(listMatches, v)
			}
		}
	}
	// fmt.Println(len(listMatches))
	// fmt.Println("listMatches")
	// fmt.Println(listMatches)

	for i, v := range listMatches {
		// fmt.Println(v)
		// if len(v) < 150 {
		// fmt.Println(v)
		listSymbolBeforeParse = append(listSymbolBeforeParse, v)
		// }
		i++
	}

	for i, v := range listSymbolBeforeParse {
		if len(v) > 150 {
			currentTwiValueQuery1 := strings.Split(v, "/symbol/")[1]
			currentTwiValueQuery2 := strings.Split(currentTwiValueQuery1, "\\\"><div")[0]
			// fmt.Println("currentTwiValueQuery2")
			// fmt.Println(currentTwiValueQuery2)
			listSymbols = append(listSymbols, currentTwiValueQuery2)
		}
		if len(v) < 150 {
			currentTwiValueQuery1 := strings.Split(v, ">")
			listSymbols = append(listSymbols, currentTwiValueQuery1[len(currentTwiValueQuery1)-1])
		}
		i++
	}
	for i, v := range listSymbols {
		// fmt.Println(v)
		if strings.Contains(v, ".") == false {
			if strings.Contains(v, "'") == false {
				filterDotList = append(filterDotList, v)
			}
		}
		i++
	}
	return filterDotList
}

func parseBalance(queryString string) string {
	// <accountvalue>5026.06</accountvalue>
	splitDataQuery1 := strings.Split(queryString, "<accountvalue>")[2]
	balance := strings.Split(splitDataQuery1, "</accountvalue>")[0]
	// fmt.Println(balance)
	// balance := strings.Split(splitDataQuery2, "</quotes>")[0]
	// balance := ""
	return balance
}

// type AccountBrokerage struct {
// 	Total          string
// 	CashAvailable  string
// 	UnsettledFunds string
// }

func parseAccountBrokerage(queryString string) AccountBrokerage {
	accountBrokerage := AccountBrokerage{}
	splitDataQuery1 := strings.Split(queryString, "<unsettledfunds>")[1]
	unsettledfunds := strings.Split(splitDataQuery1, "</unsettledfunds>")[0]
	accountBrokerage.UnsettledFunds = unsettledfunds

	splitDataQuery2 := strings.Split(queryString, "<cashavailable>")[1]
	cashavailable := strings.Split(splitDataQuery2, "</cashavailable>")[0]
	accountBrokerage.CashAvailable = cashavailable

	splitDataQuery3 := strings.Split(queryString, "<total>")[1]
	total := strings.Split(splitDataQuery3, "</total>")[0]
	accountBrokerage.Total = total

	return accountBrokerage
}

func parseHistory(queryString string) []string {
	splitDataQuery1 := strings.Split(queryString, "</transaction>")
	historyList := []string{}
	for i, v := range splitDataQuery1 {
		if len(v) < 15 {
			continue
		}
		historyList = append(historyList, v)
		i++
	}
	return historyList
}

func parseOrders(query string) ContainerOrders {
	splitDataQuery := strings.Split(query, "</FIXML>]]>")
	//remove last index which is a server message
	splitDataQuery = splitDataQuery[:len(splitDataQuery)-1]

	// fmt.Println("splitDataQuery")
	// fmt.Println(splitDataQuery)
	containerOrders := ContainerOrders{}
	for i, v := range splitDataQuery {
		symParsed := strings.Split(v, "Sym=")
		symParsed1 := strings.Split(symParsed[1], "\"/>")
		symParsed2 := strings.Split(symParsed1[0], "\"")
		symParsed3 := strings.Replace(symParsed2[1], "\\", "", -1)
		orderCreated := Order{Symbol: symParsed3}

		//parse qty
		qtyParsed := strings.Split(v, "OrdQty Qty=")
		qtyParsed1 := strings.Split(qtyParsed[1], "\"/>")
		qtyParsed2 := strings.Split(qtyParsed1[0], "\"")
		qtyParsed3 := strings.Replace(qtyParsed2[1], "\\", "", -1)
		orderCreated.Qty = qtyParsed3
		// fmt.Println("Qty in")
		// fmt.Println(qtyParsed3)
		// fmt.Println(qtyParsed)
		//parse SVI ex) SVI-6084382688
		// sviParsed := strings.Split(v, "SVI-")

		sviParsed := strings.Split(v, "OrdID=")
		sviParsed1 := strings.Split(sviParsed[1], "\"/>")
		sviParsed2 := strings.Split(sviParsed1[0], "\"")
		sviParsed3 := strings.Replace(sviParsed2[1], "\\", "", -1)
		orderCreated.SVI = sviParsed3
		// // orderCreated.SVI = sviParsed
		// fmt.Println(sviParsed[1])

		//if order does not contain Canceled by user or RejRsn then limit is pending successfully
		if strings.Contains(v, "Canceled by user") {
			orderCreated.OrderStatus = "Canceled"
			containerOrders.ListOrders = append(containerOrders.ListOrders, orderCreated)
			continue
		}
		if strings.Contains(v, "RejRsn") {
			orderCreated.OrderStatus = "RejRsn"
			containerOrders.ListOrders = append(containerOrders.ListOrders, orderCreated)
			continue
		}
		orderCreated.OrderStatus = "Successful"
		containerOrders.ListOrders = append(containerOrders.ListOrders, orderCreated)
		i++
	}

	// for i, v := range containerOrders.ListOrders {
	// 	fmt.Println()
	// 	fmt.Println(v)
	// 	i++
	// }
	return containerOrders
}

func parseAllHolding(query string) ContainerHolding {
	//repair multi holding
	containerHolding := ContainerHolding{}
	splitDataQuery1 := strings.Split(query, "<accountholdings>")[1]
	splitDataQuery2 := strings.Split(splitDataQuery1, "</accountholdings>")
	if len(splitDataQuery2) > 0 {
		splitDataQuery2 = splitDataQuery2[:len(splitDataQuery2)-1]
	}
	splitHolding := strings.Split(query, "</holding>")

	splitHolding = splitHolding[:len(splitHolding)-1]

	for i, v := range splitHolding {
		// fmt.Println()
		// fmt.Println(v)
		symbolString1 := strings.Split(v, "<sym>")[1]
		symbolString2 := strings.Split(symbolString1, "</sym>")
		symbol := symbolString2[0]

		purchasePriceString1 := strings.Split(v, "<purchaseprice>")[1]
		purchasePriceString2 := strings.Split(purchasePriceString1, "</purchaseprice>")
		purchasePrice := purchasePriceString2[0]

		qtyString1 := strings.Split(v, "<qty>")[1]
		qtyString2 := strings.Split(qtyString1, "</qty>")
		qty := qtyString2[0]

		holding := HoldingWisemen{Symbol: symbol, Price: purchasePrice, Qty: qty}
		containerHolding.ListHolding = append(containerHolding.ListHolding, holding)
		i++
	}
	return containerHolding
}

func parseHoldings(query string) []Holding {
	splitDataQuery1 := strings.Split(query, "<accountholdings>")[1]
	splitDataQuery2 := strings.Split(splitDataQuery1, "</accountholdings>")
	// fmt.Println("splitDataQuery1")
	// fmt.Println(splitDataQuery1[1])
	// fmt.Println(splitDataQuery2[0])
	// fmt.Println(len(splitDataQuery2))
	if len(splitDataQuery2) > 0 {
		splitDataQuery2 = splitDataQuery2[:len(splitDataQuery2)-1]
	}
	// fmt.Println(len(splitDataQuery2))

	listHoldings := []Holding{}
	for i, v := range splitDataQuery2 {
		symbolString1 := strings.Split(splitDataQuery2[0], "<sym>")
		symbolString2 := strings.Split(symbolString1[1], "</sym>")
		// fmt.Println(symbolString2[0])
		symbol := symbolString2[0]

		purchasePriceString1 := strings.Split(splitDataQuery2[0], "<purchaseprice>")
		purchasePriceString2 := strings.Split(purchasePriceString1[1], "</purchaseprice>")
		// fmt.Println(purchasePriceString2[0])
		purchasePrice := purchasePriceString2[0]

		qtyString1 := strings.Split(splitDataQuery2[0], "<qty>")
		qtyString2 := strings.Split(qtyString1[1], "</qty>")
		// fmt.Println(purchasePriceString2[0])
		qty := qtyString2[0]
		// fmt.Println(qty)

		holding := Holding{Symbol: symbol, PurchasePrice: purchasePrice, Qty: qty}
		listHoldings = append(listHoldings, holding)

		fmt.Println(v)

		i++
	}
	return listHoldings
}

// <?xml version="1.0" encoding="UTF-8"?><response id="126565f9-ee57-4117-aff4-dcbf19f4d673"><elapsedtime>0</elapsedtime>
// <quotes>
//   <quotetype>Real Time -  market data real time, National Best Bid and Offer</quotetype>
//   <quote>
//     <adp_100>196.7560</adp_100>
//     <adp_200>185.8540</adp_200>
//     <adp_50>199.2840</adp_50>
//     <adv_21>18696278</adv_21>
//     <adv_30>19827715</adv_30>
//     <adv_90>26735159</adv_90>
//     <ask>207.08</ask>
//     <ask_time>17:16</ask_time>
//     <asksz>5</asksz>
//     <basis>na</basis>
//     <beta>1.071</beta>
//     <bid>206.95</bid>
//     <bid_time>17:15</bid_time>
//     <bidsz>1</bidsz>
//     <bidtick>d</bidtick>
//     <chg>1.6500</chg>
//     <chg_sign>d</chg_sign>
//     <chg_t>na</chg_t>
//     <cl>208.67</cl>
//     <contract_size>na</contract_size>
//     <cusip>na</cusip>
//     <date>2019-07-25</date>
//     <datetime>2019-07-25T15:59:00-04:00</datetime>
//     <days_to_expiration>na</days_to_expiration>
//     <div>0.77</div>
//     <divexdate>20190510</divexdate>
//     <divfreq>Q</divfreq>
//     <divpaydt>20190516</divpaydt>
//     <dollar_value>2880148662.79</dollar_value>
//     <eps>11.9</eps>
//     <exch>NASD</exch>
//     <exch_desc>NASDAQ</exch_desc>
//     <hi>209.24</hi>
//     <iad>3.08</iad>
//     <idelta>na</idelta>
//     <igamma>na</igamma>
//     <imp_volatility>na</imp_volatility>
//     <incr_vl>1261226</incr_vl>
//     <irho>na</irho>
//     <issue_desc>na</issue_desc>
//     <itheta>na</itheta>
//     <ivega>na</ivega>
//     <last>207.02</last>
//     <lo>206.73</lo>
//     <name>APPLE INC</name>
//     <op_delivery>na</op_delivery>
//     <op_flag>1</op_flag>
//     <op_style>na</op_style>
//     <op_subclass>na</op_subclass>
//     <openinterest>na</openinterest>
//     <opn>208.89</opn>
//     <opt_val>na</opt_val>
//     <pchg>0.79</pchg>
//     <pchg_sign>na</pchg_sign>
//     <pcls>208.67</pcls>
//     <pe>17.5353</pe>
//     <phi>209.15</phi>
//     <plo>207.17</plo>
//     <popn>207.67</popn>
//     <pr_adp_100>196.6192</pr_adp_100>
//     <pr_adp_200>185.6875</pr_adp_200>
//     <pr_adp_50>198.4545</pr_adp_50>
//     <pr_date>2019-07-24</pr_date>
//     <pr_openinterest>na</pr_openinterest>
//     <prbook>8.89</prbook>
//     <prchg>1.00</prchg>
//     <prem_mult>na</prem_mult>
//     <put_call>na</put_call>
//     <pvol>0</pvol>
//     <qcond>0</qcond>
//     <rootsymbol>na</rootsymbol>
//     <secclass>0</secclass>
//     <sesn>na</sesn>
//     <sho>4601075000</sho>
//     <strikeprice>na</strikeprice>
//     <symbol>AAPL</symbol>
//     <tcond>29</tcond>
//     <timestamp>1564089376</timestamp>
//     <tr_num>123348</tr_num>
//     <tradetick>e</tradetick>
//     <trend>na</trend>
//     <under_cusip>na</under_cusip>
//     <undersymbol>na</undersymbol>
//     <vl>13862806</vl>
//     <volatility12>0.3137</volatility12>
//     <vwap>207.76</vwap>
//     <wk52hi>233.47</wk52hi>
//     <wk52hidate>20181003</wk52hidate>
//     <wk52lo>142.00</wk52lo>
//     <wk52lodate>20190103</wk52lodate>
//     <xdate>na</xdate>
//     <xday>na</xday>
//     <xmonth>na</xmonth>
//     <xyear>na</xyear>
//     <yield>1.47601</yield>
//   </quote>
// </quotes><error>Success</error></response>
