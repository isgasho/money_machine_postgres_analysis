package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	_ "github.com/lib/pq"
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

func parseBalanceQuery(queryString string) Holdings {
	splitDataQuery1 := strings.Split(queryString, "<sym>")[1]
	symbol := strings.Split(splitDataQuery1, "</sym>")[0]

	splitDataQuery1 = strings.Split(queryString, "<price>")[1]
	price := strings.Split(splitDataQuery1, "</price>")[0]

	splitDataQuery1 = strings.Split(queryString, "<qty>")[1]
	qty := strings.Split(splitDataQuery1, "</qty>")[0]

	holdings := Holdings{
		Symbol: symbol,
		Price:  price,
		Qty:    qty,
	}
	return holdings
}

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

		//calculated createdAt
		createdAt := createStockTimeStamp()

		var stock = Stock{
			Monitoring:   false,
			Symbol:       symbol,
			CreatedAt:    createdAt,
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
		stockList = append(stockList, stock)
	}
	return stockList
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

	fmt.Println(parseList[0])
	for i, v := range parseList {
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

func parseDowWebscrape(queryString string) (string, string, string) {
	currentDowValueQuery1 := strings.Split(queryString, "INDEXDJX: .DJI")[1]
	currentDowValueQuery2 := strings.Split(currentDowValueQuery1, "PM EDT")[0]
	currentDowValueQuery3 := strings.Split(currentDowValueQuery2, "</span>")[1]
	currentDowValueQuery4 := strings.Split(currentDowValueQuery3, ">")
	currentDowValueQuery5 := strings.TrimSpace(currentDowValueQuery4[(len(currentDowValueQuery4) - 1)])
	currentDowValue := currentDowValueQuery5
	// currentDowValue := "p"

	pointsChangedQuery1 := strings.Split(queryString, "INDEXDJX: .DJI")[1]
	pointsChangedQuery2 := strings.Split(pointsChangedQuery1, "PM EDT")[0]
	pointsChangedQuery3 := strings.Split(pointsChangedQuery2, "</span>")[4]
	pointsChangedQuery4 := strings.Split(pointsChangedQuery3, ">")
	pointsChangedQuery5 := strings.TrimSpace(pointsChangedQuery4[(len(pointsChangedQuery4) - 1)])
	pointsChanged := pointsChangedQuery5

	percentageChangedQuery1 := strings.Split(queryString, "INDEXDJX: .DJI")[1]
	percentageChangedQuery2 := strings.Split(percentageChangedQuery1, "PM EDT")[0]
	percentageChangedQuery3 := strings.Split(percentageChangedQuery2, " by ")[1]
	percentageChangedQuery4 := strings.Split(percentageChangedQuery3, "%")[0]
	percentageChange := percentageChangedQuery4

	// f, err := os.Create("output.html")
	// check(err)
	// n3, err1 := f.WriteString(currentDowValue + pointsChanged + percentageChange)
	// check(err1)
	// fmt.Printf("wrote %d bytes\n", n3)
	// f.Sync()

	// percentageChange := strings.TrimSpace(strings.Split(pointsChangedQuery3, "\\n")[1])
	// percentageChange := "b"
	// s1 := strings.Replace(percentageChange, "(", "", -1)
	// s2 := strings.Replace(s1, ")", "", -1)
	// percentageChange = strings.Replace(s2, "%", "", -1)
	return currentDowValue, pointsChanged, percentageChange
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
