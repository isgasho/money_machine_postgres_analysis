package main

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	host   = "localhost"
	port   = 5432
	user   = "postgres"
	dbname = "money_machine"
)

// func insertDay(dayOfWeek string) {
// 	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
// 		"dbname=%s sslmode=disable",
// 		host, port, user, dbname)
// 	db, err := sql.Open("postgres", psqlInfo)
// 	if err != nil {
// 		fmt.Println("Create Error 1")
// 	}
// 	defer db.Close()

// 	sqlStatement := `
// 		INSERT INTO day (day_of_Week)
// 		VALUES ($1)
// 		RETURNING id, created_at, day_of_Week
// 		`
// 	var day Day
// 	row := db.QueryRow(sqlStatement, dayOfWeek)
// 	err1 := row.Scan(&day.ID, &day.CreatedAt, &day.DayOfWeek)
// 	if err1 != nil {
// 		fmt.Println("Create Error 2")
// 	}
// 	fmt.Println(day.ID, day.DayOfWeek, day.CreatedAt)
// }
// func setDay() {
// }
// func selectDay() {
// }
// func deleteDay() {
// }

// func insertNews(newsEntry News) {
// 	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
// 		"dbname=%s sslmode=disable",
// 		host, port, user, dbname)
// 	db, err := sql.Open("postgres", psqlInfo)
// 	if err != nil {
// 		fmt.Println("Create Error 1")
// 	}
// 	defer db.Close()

// 	sqlStatement := `
// 		INSERT INTO news (day_id, news_info)
// 		VALUES ($1, $2)
// 		RETURNING id, day_id, created_at, news_info
// 		`
// 	var news News
// 	row := db.QueryRow(sqlStatement, newsEntry.DayID, newsEntry.NewsInfo)
// 	err1 := row.Scan(&news.ID, &news.DayID, &news.CreatedAt, &news.NewsInfo)
// 	if err1 != nil {
// 		fmt.Println("Create Error 2")
// 	}
// 	fmt.Println(news.ID, news.DayID, news.CreatedAt, news.NewsInfo)
// }
// func setNews() {
// }
// func selectNews() {
// }
// func deleteNews() {
// }

// func dropDow() {
// 	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
// 		"dbname=%s sslmode=disable",
// 		host, port, user, dbname)
// 	db, err := sql.Open("postgres", psqlInfo)
// 	if err != nil {
// 		fmt.Println("Read Error 1")
// 		panic(err)
// 	}
// 	defer db.Close()

// 	res, err1 := db.Exec("drop table dow")
// 	if err1 != nil {
// 		fmt.Println("Delete Error 2")
// 	}
// 	count, err2 := res.RowsAffected()
// 	if err2 != nil {
// 		fmt.Println("Delete Error 3")
// 	}
// 	fmt.Println(count)
// }

// func createDow() {
// 	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
// 		"dbname=%s sslmode=disable",
// 		host, port, user, dbname)
// 	db, err := sql.Open("postgres", psqlInfo)
// 	if err != nil {
// 		fmt.Println("Read Error 1")
// 		panic(err)
// 	}
// 	defer db.Close()

// 	res, err1 := db.Exec(`CREATE TABLE dow
// 	(
// 	   id SERIAL PRIMARY KEY,
// 	   created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
// 	   current_dow_value VARCHAR
// 	);`)

// 	if err1 != nil {
// 		fmt.Println("Delete Error 2")
// 	}
// 	count, err2 := res.RowsAffected()
// 	if err2 != nil {
// 		fmt.Println("Delete Error 3")
// 	}
// 	fmt.Println(count)
// }

//

//

// func deleteStock(symbolToDel string) {
// 	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
// 		"dbname=%s sslmode=disable",
// 		host, port, user, dbname)
// 	db, err := sql.Open("postgres", psqlInfo)
// 	if err != nil {
// 		fmt.Println("Read Error 1")
// 		panic(err)
// 	}
// 	defer db.Close()

// 	res, err1 := db.Exec("DELETE FROM stock WHERE symbol=$1", symbolToDel)
// 	if err1 != nil {
// 		fmt.Println("Delete Error 2")
// 	}
// 	count, err2 := res.RowsAffected()
// 	if err2 != nil {
// 		fmt.Println("Delete Error 3")
// 	}
// 	fmt.Println(count)
// }

// func selectAllStockOfSymbol(symbolToSearch string) []Stock {
// 	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
// 		"dbname=%s sslmode=disable",
// 		host, port, user, dbname)
// 	db, err := sql.Open("postgres", psqlInfo)
// 	if err != nil {
// 		fmt.Println("Read Error 1")
// 		panic(err)
// 	}
// 	defer db.Close()

// 	rows, err1 := db.Query("SELECT id, created_at, monitoring, symbol, bid, ask, last, pchg, pcls, opn, vl, pvol, volatility12, wk52hi, wk52hidate, wk52lo, wk52lodate, hi, lo, pr_adp_50, pr_adp_100, prchg, adp_50, adp_100, adv_30, adv_90 FROM stock WHERE symbol=$1", symbolToSearch)
// 	if err1 != nil {
// 		fmt.Println(err1)
// 	}
// 	defer rows.Close()
// 	stockList := make([]Stock, 0)

// 	for rows.Next() {
// 		var stock Stock
// 		if err2 := rows.Scan(&stock.ID, &stock.CreatedAt, &stock.Monitoring, &stock.Symbol, &stock.Bid, &stock.Ask, &stock.Last, &stock.Pchg, &stock.Pcls, &stock.Opn, &stock.Vl, &stock.Pvol, &stock.Volatility12, &stock.Wk52hi, &stock.Wk52hidate, &stock.Wk52lo, &stock.Wk52lodate, &stock.Hi, &stock.Lo, &stock.PrAdp50, &stock.PrAdp100, &stock.Prchg, &stock.Adp50, &stock.Adp100, &stock.Adv30, &stock.Adv90); err2 != nil {
// 			fmt.Println("err2")
// 		}
// 		stockList = append(stockList, stock)
// 	}
// 	return stockList
// }

// func selectEvalResultsWhale() []string {
// 	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
// 		"dbname=%s sslmode=disable",
// 		host, port, user, dbname)
// 	db, err := sql.Open("postgres", psqlInfo)
// 	if err != nil {
// 		fmt.Println("Read Error 1")
// 		panic(err)
// 	}
// 	defer db.Close()

// 	rows, err1 := db.Query("SELECT symbol FROM eval_results_whale")
// 	if err1 != nil {
// 		fmt.Println(err1)
// 	}
// 	defer rows.Close()
// 	symbolList := make([]string, 0)

// 	for rows.Next() {
// 		var symbol string
// 		if err2 := rows.Scan(&symbol); err2 != nil {
// 			fmt.Println("err2")
// 		}
// 		symbolList = append(symbolList, symbol)
// 	}
// 	return symbolList
// }

// func dropEvalResultsWhale() {
// 	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
// 		"dbname=%s sslmode=disable",
// 		host, port, user, dbname)
// 	db, err := sql.Open("postgres", psqlInfo)
// 	if err != nil {
// 		fmt.Println("Read Error 1")
// 		panic(err)
// 	}
// 	defer db.Close()

// 	res, err1 := db.Exec("drop table eval_results_whale")
// 	if err1 != nil {
// 		fmt.Println("Delete Error 2")
// 	}
// 	count, err2 := res.RowsAffected()
// 	if err2 != nil {
// 		fmt.Println("Delete Error 3")
// 	}
// 	fmt.Println(count)
// }

// func createEvalResultsWhale() {
// 	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
// 		"dbname=%s sslmode=disable",
// 		host, port, user, dbname)
// 	db, err := sql.Open("postgres", psqlInfo)
// 	if err != nil {
// 		fmt.Println("Read Error 1")
// 		panic(err)
// 	}
// 	defer db.Close()

// 	res, err1 := db.Exec(`CREATE TABLE eval_results_whale
// 	(
// 		id SERIAL PRIMARY KEY,
// 		created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
// 		symbol VARCHAR,
// 		is_breach_worthy VARCHAR,
// 		is_pattern_met VARCHAR
// 	 );`)

// 	if err1 != nil {
// 		fmt.Println("Delete Error 2")
// 	}
// 	count, err2 := res.RowsAffected()
// 	if err2 != nil {
// 		fmt.Println("Delete Error 3")
// 	}
// 	fmt.Println(count)
// }

// func deleteTradeEnteredInformation(symbolToDel string) {
// 	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
// 		"dbname=%s sslmode=disable",
// 		host, port, user, dbname)
// 	db, err := sql.Open("postgres", psqlInfo)
// 	if err != nil {
// 		fmt.Println("Read Error 1")
// 		panic(err)
// 	}
// 	defer db.Close()

// 	res, err1 := db.Exec("DELETE FROM trade_entered_information WHERE symbol=$1", symbolToDel)
// 	if err1 != nil {
// 		fmt.Println("Delete Error 2")
// 	}
// 	count, err2 := res.RowsAffected()
// 	if err2 != nil {
// 		fmt.Println("Delete Error 3")
// 	}
// 	fmt.Println(count)
// }

//
func queryIsTradeCompleted(symbol string) TradeBoughtEvaluation {
	tradeBoughtEvaluation := TradeBoughtEvaluation{}
	response := queryHolding()
	holdingList := parseHoldings(response)

	isHoldingSymbol := false
	for i, v := range holdingList {
		if v.Symbol == symbol {
			isHoldingSymbol = true
		}
		i++
	}
	tradeBoughtEvaluation.HoldingList = holdingList
	tradeBoughtEvaluation.IsBought = isHoldingSymbol
	return tradeBoughtEvaluation
}

// insertAltIntervalBuyWisemen
func insertAltIntervalBuyWisemen(altIntervalBuyWisemen AltIntervalBuyWisemen) {
	listValues := []string{altIntervalBuyWisemen.Symbol, altIntervalBuyWisemen.IsAltIntervalOperation, altIntervalBuyWisemen.ReasonCancelation}
	postCommandDBInsert("INSERT INTO alt_interval_buy_wisemen (symbol, is_alt_interval_operation, reason_cancelation) VALUES (", listValues)
}
func selectAltIntervalBuyWisemen() []AltIntervalBuyWisemen {
	listAltIntervalBuyWisemen := []AltIntervalBuyWisemen{}
	response := postCommandDBSelect("SELECT symbol, is_alt_interval_operation, reason_cancelation FROM alt_interval_buy_wisemen")
	container := parseDBResponse(response)
	fmt.Println(container.ListStringFromDB)
	fmt.Println(len(container.ListStringFromDB))
	for i, v := range container.ListStringFromDB {
		stringToStrip := v.ListString[0]
		strippedString := strings.Join(strings.Fields(stringToStrip), "")
		altIntervalBuyWisemen := AltIntervalBuyWisemen{Symbol: strippedString, IsAltIntervalOperation: v.ListString[1], ReasonCancelation: v.ListString[2]}
		listAltIntervalBuyWisemen = append(listAltIntervalBuyWisemen, altIntervalBuyWisemen)
		i++
	}
	return listAltIntervalBuyWisemen
}

func truncateAltIntervalBuyWisemen() {
	postCommandDBTruncate("TRUNCATE table alt_interval_buy_wisemen")
}

//day
func insertDay(dayOfWeek string) {
	listValues := []string{dayOfWeek}
	// 		INSERT INTO day (day_of_Week)
	// 		VALUES ($1)
	// 		RETURNING id, created_at, day_of_Week
	postCommandDBInsert("INSERT INTO day (day_of_Week) VALUES (", listValues)
}
func selectDay() []Day {
	listDay := []Day{}
	response := postCommandDBSelect("SELECT day_of_Week FROM day")
	container := parseDBResponse(response)
	fmt.Println(container.ListStringFromDB)
	fmt.Println(len(container.ListStringFromDB))
	for i, v := range container.ListStringFromDB {
		stringToStrip := v.ListString[0]
		strippedString := strings.Join(strings.Fields(stringToStrip), "")
		day := Day{DayOfWeek: strippedString}
		listDay = append(listDay, day)
		i++
	}
	return listDay
}
func truncateDay() {
	postCommandDBTruncate("TRUNCATE table day")
}

//news
func insertNews(news_info string) {
	listValues := []string{news_info}
	// INSERT INTO news (day_id, news_info)
	// 	VALUES ($1, $2)
	// 	RETURNING id, day_id, created_at, news_info
	postCommandDBInsert("INSERT INTO news (news_info) VALUES (", listValues)
}
func selectNews() []News {
	listNews := []News{}
	response := postCommandDBSelect("SELECT news_info FROM news")
	container := parseDBResponse(response)
	fmt.Println(container.ListStringFromDB)
	fmt.Println(len(container.ListStringFromDB))
	for i, v := range container.ListStringFromDB {
		stringToStrip := v.ListString[0]
		strippedString := strings.Join(strings.Fields(stringToStrip), "")
		news := News{NewsInfo: strippedString}
		listNews = append(listNews, news)
		i++
	}
	return listNews
}
func truncateNews() {
	postCommandDBTruncate("TRUNCATE table news")
}

//dow
func insertDow(current_dow_value string) {
	listValues := []string{current_dow_value}
	// INSERT INTO news (day_id, news_info)
	// 	VALUES ($1, $2)
	// 	RETURNING id, day_id, created_at, news_info
	postCommandDBInsert("INSERT INTO dow (current_dow_value) VALUES (", listValues)
}
func selectDow() []Dow {
	listDow := []Dow{}
	response := postCommandDBSelect("SELECT current_dow_value FROM dow")
	container := parseDBResponse(response)
	fmt.Println(container.ListStringFromDB)
	fmt.Println(len(container.ListStringFromDB))
	for i, v := range container.ListStringFromDB {
		stringToStrip := v.ListString[0]
		strippedString := strings.Join(strings.Fields(stringToStrip), "")
		dow := Dow{CurrentDowValue: strippedString}
		listDow = append(listDow, dow)
		i++
	}
	return listDow
}
func truncateDow() {
	postCommandDBTruncate("TRUNCATE table dow")
}

//MarketOpenAnalysis
func insertMarketOpenAnalysis(marketOpenAnalysis MarketOpenAnalysis) {
	listValues := []string{marketOpenAnalysis.IsMarketClosed}
	// 	CREATE TABLE market_open_analysis
	// (
	//    id SERIAL PRIMARY KEY,
	//    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
	//    is_market_closed VARCHAR
	// );
	postCommandDBInsert("INSERT INTO market_open_analysis (is_market_closed) VALUES (", listValues)
}
func selectMarketOpenAnalysis() []MarketOpenAnalysis {
	listMarketOpenAnalysis := []MarketOpenAnalysis{}
	response := postCommandDBSelect("SELECT is_market_closed FROM market_open_analysis")
	container := parseDBResponse(response)
	fmt.Println(container.ListStringFromDB)
	fmt.Println(len(container.ListStringFromDB))
	for i, v := range container.ListStringFromDB {
		// stringToStrip := v.ListString[0]
		// strippedString := strings.Join(strings.Fields(stringToStrip), "")
		// stringToStrip1 := v.ListString[1]
		// strippedString1 := strings.Join(strings.Fields(stringToStrip1), "")
		marketOpenAnalysis := MarketOpenAnalysis{IsMarketClosed: v.ListString[0]}
		listMarketOpenAnalysis = append(listMarketOpenAnalysis, marketOpenAnalysis)
		i++
	}
	return listMarketOpenAnalysis
}
func truncateMarketOpenAnalysis() {
	postCommandDBTruncate("TRUNCATE table market_open_analysis")
}

//MarketOpenAnalysis
func insertStockWisemen(stock Stock) {
	hour := getCurrentHour()
	minute := getCurrentMinute()
	stringHour := strconv.Itoa(hour)
	stringMinute := strconv.Itoa(minute)
	timeCreated := stringHour + " " + stringMinute
	listValues := []string{stock.Symbol, stock.Bid, stock.Ask, stock.Last, stock.Pchg, stock.Pcls, stock.Opn, stock.Vl, timeCreated}
	postCommandDBInsert("INSERT INTO stock_wisemen (symbol, bid, ask, last, pchg, pcls, opn, vl, time_created) VALUES (", listValues)
}
func selectStockWisemen() []Stock {
	listStock := []Stock{}
	response := postCommandDBSelect("SELECT symbol, bid, ask, last, pchg, pcls, opn, vl, time_created FROM stock_wisemen")
	container := parseDBResponse(response)
	fmt.Println(container.ListStringFromDB)
	fmt.Println(len(container.ListStringFromDB))
	for i, v := range container.ListStringFromDB {
		stock := Stock{Symbol: v.ListString[0], Bid: v.ListString[1], Ask: v.ListString[2], Last: v.ListString[3], Pchg: v.ListString[4], Pcls: v.ListString[5], Opn: v.ListString[6], Vl: v.ListString[7], TimeCreated: v.ListString[8]}
		listStock = append(listStock, stock)
		i++
	}
	return listStock
}
func truncateStockWisemen() {
	postCommandDBTruncate("TRUNCATE table stock_wisemen")
}

//StockWhale High
func insertStockWhaleHigh(stock Stock) {
	hour := getCurrentHour()
	minute := getCurrentMinute()
	stringHour := strconv.Itoa(hour)
	stringMinute := strconv.Itoa(minute)
	timeCreated := stringHour + " " + stringMinute
	listValues := []string{stock.Symbol, stock.Bid, stock.Ask, stock.Last, stock.Pchg, stock.Pcls, stock.Opn, stock.Vl, timeCreated}
	postCommandDBInsert("INSERT INTO stock_whale_high (symbol, bid, ask, last, pchg, pcls, opn, vl, time_created) VALUES (", listValues)
}
func selectStockWhaleHigh(symbol string) []Stock {
	listStock := []Stock{}
	response := postCommandDBSelect("SELECT symbol, bid, ask, last, pchg, pcls, opn, vl, time_created FROM stock_whale_high")
	container := parseDBResponse(response)
	fmt.Println(container.ListStringFromDB)
	fmt.Println(len(container.ListStringFromDB))
	for i, v := range container.ListStringFromDB {
		stock := Stock{Symbol: v.ListString[0], Bid: v.ListString[1], Ask: v.ListString[2], Last: v.ListString[3], Pchg: v.ListString[4], Pcls: v.ListString[5], Opn: v.ListString[6], Vl: v.ListString[7], TimeCreated: v.ListString[8]}
		listStock = append(listStock, stock)
		i++
	}
	return listStock
}
func truncateStockWhaleHigh() {
	postCommandDBTruncate("TRUNCATE table stock_whale_high")
}

//StockWhale Low
func insertStockWhaleLow(stock Stock) {
	hour := getCurrentHour()
	minute := getCurrentMinute()
	stringHour := strconv.Itoa(hour)
	stringMinute := strconv.Itoa(minute)
	timeCreated := stringHour + " " + stringMinute
	listValues := []string{stock.Symbol, stock.Bid, stock.Ask, stock.Last, stock.Pchg, stock.Pcls, stock.Opn, stock.Vl, timeCreated}
	postCommandDBInsert("INSERT INTO stock_whale_low (symbol, bid, ask, last, pchg, pcls, opn, vl, time_created) VALUES (", listValues)
}
func selectStockWhaleLow(symbol string) []Stock {
	listStock := []Stock{}
	response := postCommandDBSelect("SELECT symbol, bid, ask, last, pchg, pcls, opn, vl, time_created FROM stock_whale_low")
	container := parseDBResponse(response)
	fmt.Println(container.ListStringFromDB)
	fmt.Println(len(container.ListStringFromDB))
	for i, v := range container.ListStringFromDB {
		stock := Stock{Symbol: v.ListString[0], Bid: v.ListString[1], Ask: v.ListString[2], Last: v.ListString[3], Pchg: v.ListString[4], Pcls: v.ListString[5], Opn: v.ListString[6], Vl: v.ListString[7], TimeCreated: v.ListString[8]}
		listStock = append(listStock, stock)
		i++
	}
	return listStock
}
func truncateStockWhaleLow() {
	postCommandDBTruncate("TRUNCATE table stock_whale_low")
}

//wisemen_symbol_hold
func insertWisemenSymbolHold(symbol string) {
	listValues := []string{symbol}
	postCommandDBInsert("INSERT INTO wisemen_symbol_hold (symbol) VALUES (", listValues)
}
func selectWisemenSymbolHold() []string {
	symbolList := []string{}
	response := postCommandDBSelect("SELECT symbol FROM wisemen_symbol_hold")
	container := parseDBResponse(response)
	fmt.Println(container.ListStringFromDB)
	fmt.Println(len(container.ListStringFromDB))
	for i, v := range container.ListStringFromDB {
		symbolList = append(symbolList, v.ListString[0])
		i++
	}
	return symbolList
}
func truncateWisemenSymbolHold() {
	postCommandDBTruncate("TRUNCATE table wisemen_symbol_hold")
}

//whale_symbol_hold_high
//WhaleSymbolHoldHigh
func insertWhaleSymbolHoldHigh(symbol string) {
	listValues := []string{symbol}
	postCommandDBInsert("INSERT INTO whale_symbol_hold_high (symbol) VALUES (", listValues)
}
func selectWhaleSymbolHoldHigh() []string {
	symbolList := []string{}
	response := postCommandDBSelect("SELECT symbol FROM whale_symbol_hold_high")
	container := parseDBResponse(response)
	fmt.Println(container.ListStringFromDB)
	fmt.Println(len(container.ListStringFromDB))
	for i, v := range container.ListStringFromDB {
		symbolList = append(symbolList, v.ListString[0])
		i++
	}
	return symbolList
}
func truncateWhaleSymbolHoldHigh() {
	postCommandDBTruncate("TRUNCATE table whale_symbol_hold_high")
}

//whale_symbol_hold_low
//WhaleSymbolHoldLow
func insertWhaleSymbolHoldLow(symbol string) {
	listValues := []string{symbol}
	postCommandDBInsert("INSERT INTO whale_symbol_hold_low (symbol) VALUES (", listValues)
}
func selectWhaleSymbolHoldLow() []string {
	symbolList := []string{}
	response := postCommandDBSelect("SELECT symbol FROM whale_symbol_hold_low")
	container := parseDBResponse(response)
	fmt.Println(container.ListStringFromDB)
	fmt.Println(len(container.ListStringFromDB))
	for i, v := range container.ListStringFromDB {
		symbolList = append(symbolList, v.ListString[0])
		i++
	}
	return symbolList
}
func truncateWhaleSymbolHoldLow() {
	postCommandDBTruncate("TRUNCATE table whale_symbol_hold_low")
}

//temp_symbol_hold_high
//TempSymbolHoldHigh
func insertTempSymbolHoldHigh(symbol string) {
	listValues := []string{symbol}
	postCommandDBInsert("INSERT INTO temp_symbol_hold_high (symbol) VALUES (", listValues)
}
func selectTempSymbolHoldHigh() []string {
	symbolList := []string{}
	response := postCommandDBSelect("SELECT symbol FROM temp_symbol_hold_high")
	container := parseDBResponse(response)
	fmt.Println(container.ListStringFromDB)
	fmt.Println(len(container.ListStringFromDB))
	for i, v := range container.ListStringFromDB {
		symbolList = append(symbolList, v.ListString[0])
		i++
	}
	return symbolList
}
func truncateTempSymbolHoldHigh() {
	postCommandDBTruncate("TRUNCATE table temp_symbol_hold_high")
}

//temp_symbol_hold_low
func insertTempSymbolHoldLow(symbol string) {
	listValues := []string{symbol}
	postCommandDBInsert("INSERT INTO temp_symbol_hold_low (symbol) VALUES (", listValues)
}
func selectTempSymbolHoldLow() []string {
	symbolList := []string{}
	response := postCommandDBSelect("SELECT symbol FROM temp_symbol_hold_low")
	container := parseDBResponse(response)
	fmt.Println(container.ListStringFromDB)
	fmt.Println(len(container.ListStringFromDB))
	for i, v := range container.ListStringFromDB {
		symbolList = append(symbolList, v.ListString[0])
		i++
	}
	return symbolList
}
func truncateTempSymbolHoldselectAltIntervalBuyWisemenLow() {
	postCommandDBTruncate("TRUNCATE table temp_symbol_hold_low")
}

//down_day_evaluation
func insertDownDayEvaluation(downDayEvaluation DownDayEvaluation) {
	listValues := []string{downDayEvaluation.IsDownDay, downDayEvaluation.Dow, downDayEvaluation.PreviousDow, downDayEvaluation.GreatestPchg}
	postCommandDBInsert("INSERT INTO down_day_evaluation (is_down_day, dow, previous_dow, greatest_pchg) VALUES (", listValues)
}
func selectDownDayEvaluation() []DownDayEvaluation {
	downDayEvaluationList := []DownDayEvaluation{}
	response := postCommandDBSelect("SELECT is_down_day, dow, previous_dow, greatest_pchg FROM down_day_evaluation")
	container := parseDBResponse(response)
	fmt.Println(container.ListStringFromDB)
	fmt.Println(len(container.ListStringFromDB))
	for i, v := range container.ListStringFromDB {
		downDayEvaluation := DownDayEvaluation{IsDownDay: v.ListString[0], Dow: v.ListString[1], PreviousDow: v.ListString[2], GreatestPchg: v.ListString[3]}
		downDayEvaluationList = append(downDayEvaluationList, downDayEvaluation)
		i++
	}
	return downDayEvaluationList
}
func truncateDownDayEvaluation() {
	postCommandDBTruncate("TRUNCATE table down_day_evaluation")
}

//cash_day_evaluation
func insertCashDayEvaluation(cashDayEvaluation CashDayEvaluation) {
	listValues := []string{cashDayEvaluation.IsUnsettledFunds}
	postCommandDBInsert("INSERT INTO cash_day_evaluation (is_unsettled_funds) VALUES (", listValues)
}
func selectCashDayEvaluation() []CashDayEvaluation {
	cashDayEvaluationList := []CashDayEvaluation{}
	response := postCommandDBSelect("SELECT is_unsettled_funds FROM cash_day_evaluation")
	container := parseDBResponse(response)
	fmt.Println(container.ListStringFromDB)
	fmt.Println(len(container.ListStringFromDB))
	for i, v := range container.ListStringFromDB {
		cashDayEvaluation := CashDayEvaluation{IsUnsettledFunds: v.ListString[0]}
		cashDayEvaluationList = append(cashDayEvaluationList, cashDayEvaluation)
		i++
	}
	return cashDayEvaluationList
}
func truncateCashDayEvaluation() {
	postCommandDBTruncate("TRUNCATE table cash_day_evaluation")
}

//insertMetricsWhale
func insertMetricsWhale(desired_price_range_high string, desired_price_range_low string, desired_pchg, desired_pchg_variance_value string, desired_volatility_variance_value string) {
	listValues := []string{desired_price_range_high, desired_price_range_low, desired_pchg, desired_pchg_variance_value, desired_volatility_variance_value}
	postCommandDBInsert("INSERT INTO metrics_whale (desired_price_range_high, desired_price_range_low, desired_pchg, desired_pchg_variance_value, desired_volatility_variance_value) VALUES (", listValues)
}
func selectMetricsWhale() []MetricsWhale {
	metricsWhaleList := []MetricsWhale{}
	response := postCommandDBSelect("SELECT desired_price_range_high, desired_price_range_low, desired_pchg, desired_pchg_variance_value, desired_volatility_variance_value FROM metrics_whale")
	container := parseDBResponse(response)
	fmt.Println(container.ListStringFromDB)
	fmt.Println(len(container.ListStringFromDB))
	for i, v := range container.ListStringFromDB {
		metricsWhale := MetricsWhale{DesiredPriceRangeHigh: v.ListString[0], DesiredPriceRangeLow: v.ListString[1], DesiredPchg: v.ListString[2], DesiredPchgVarianceValue: v.ListString[3], DesiredVolatilityVarianceValue: v.ListString[4]}
		metricsWhaleList = append(metricsWhaleList, metricsWhale)
		i++
	}
	return metricsWhaleList
}
func truncateMetricsWhale() {
	postCommandDBTruncate("TRUNCATE table metrics_whale")
}

//insertMetricsWisemen
func insertMetricsWisemen(desired_price_range_high string, desired_price_range_low string, price_high_pchg_algo_decision string, price_low_pchg_algo_decision string, price_high_pchg_trade string, price_low_pchg_trade string, desired_pchg_variance_value string, desired_volatility_variance_value string, trade_buy_monitor_delay_seconds string, trade_buy_monitor_delay_query_seconds string, trade_buy_monitor_delay_iteration_count string) {
	listValues := []string{desired_price_range_high, desired_price_range_low, price_high_pchg_algo_decision, price_low_pchg_algo_decision, price_high_pchg_trade, price_low_pchg_trade, desired_pchg_variance_value, desired_volatility_variance_value, trade_buy_monitor_delay_seconds, trade_buy_monitor_delay_query_seconds, trade_buy_monitor_delay_iteration_count}
	postCommandDBInsert("INSERT INTO metrics_wisemen (desired_price_range_high, desired_price_range_low, price_high_pchg_algo_decision, price_low_pchg_algo_decision, price_high_pchg_trade, price_low_pchg_trade, desired_pchg_variance_value, desired_volatility_variance_value, trade_buy_monitor_delay_seconds, trade_buy_monitor_delay_query_seconds, trade_buy_monitor_delay_iteration_count) VALUES (", listValues)
}
func selectMetricsWisemen() []MetricsWisemen {
	metricsWisemenList := []MetricsWisemen{}
	response := postCommandDBSelect("SELECT desired_price_range_high, desired_price_range_low, price_high_pchg_algo_decision, price_low_pchg_algo_decision, price_high_pchg_trade, price_low_pchg_trade, desired_pchg_variance_value, desired_volatility_variance_value, trade_buy_monitor_delay_seconds, trade_buy_monitor_delay_query_seconds, trade_buy_monitor_delay_iteration_count FROM metrics_wisemen")
	container := parseDBResponse(response)
	fmt.Println(container.ListStringFromDB)
	fmt.Println(len(container.ListStringFromDB))
	for i, v := range container.ListStringFromDB {
		metricsWisemen := MetricsWisemen{DesiredPriceRangeHigh: v.ListString[0], DesiredPriceRangeLow: v.ListString[1], PriceHighPchgAlgoDecision: v.ListString[2], PriceLowPchgAlgoDecision: v.ListString[3], PriceHighPchgTrade: v.ListString[4], PriceLowPchgTrade: v.ListString[5], DesiredPchgVarianceValue: v.ListString[6], DesiredVolatilityVarianceValue: v.ListString[7], TradeBuyMonitorDelaySeconds: v.ListString[8], TradeBuyMonitorDelayQuerySeconds: v.ListString[9], TradeBuyMonitorDelayIterationCount: v.ListString[10]}
		metricsWisemenList = append(metricsWisemenList, metricsWisemen)
		i++
	}
	return metricsWisemenList
}
func truncateMetricsWisemen() {
	postCommandDBTruncate("TRUNCATE table metrics_wisemen")
}

//crit
//insertTradeResultStore
func insertTradeResultStore(tradeResultStore TradeResultStore) {
	listValues := []string{tradeResultStore.AlgorithmUsed, tradeResultStore.Result, tradeResultStore.BoughtPrice, tradeResultStore.SellPrice, tradeResultStore.ChangeAmount, tradeResultStore.StockSymbol, tradeResultStore.TimeStart, tradeResultStore.TimeEnd, tradeResultStore.TimeTradeBuy, tradeResultStore.TimeTradeSell, tradeResultStore.HighestPricePointForDay, tradeResultStore.TimeHighestPricePoint, tradeResultStore.LowestPricePointForDay, tradeResultStore.TimeLowestPricePoint, tradeResultStore.Dow1, tradeResultStore.Dow2, tradeResultStore.Dow3, tradeResultStore.Dow4}
	postCommandDBInsert("INSERT INTO trade_result_store (algorithm_used, result, bought_price, sell_price, change_amount, stock_symbol, time_start, time_end, time_trade_buy, time_trade_sell, highest_price_point_for_day, time_highest_price_point, lowest_price_point_for_day, time_lowest_price_point, dow1, dow2, dow3, dow4) VALUES (", listValues)
}
func selectTradeResultStore(symbol string) []TradeResultStore {
	tradeResultStoreList := []TradeResultStore{}
	response := postCommandDBSelect("SELECT algorithm_used, result, bought_price, sell_price, change_amount, stock_symbol, time_start, time_end, time_trade_buy, time_trade_sell, highest_price_point_for_day, time_highest_price_point, lowest_price_point_for_day, time_lowest_price_point, dow1, dow2, dow3, dow4 FROM trade_result_store")
	container := parseDBResponse(response)
	fmt.Println(container.ListStringFromDB)
	fmt.Println(len(container.ListStringFromDB))
	for i, v := range container.ListStringFromDB {
		tradeResultStore := TradeResultStore{AlgorithmUsed: v.ListString[0], Result: v.ListString[1], BoughtPrice: v.ListString[2], SellPrice: v.ListString[3], ChangeAmount: v.ListString[4], StockSymbol: v.ListString[5], TimeStart: v.ListString[6], TimeEnd: v.ListString[7], TimeTradeBuy: v.ListString[8], TimeTradeSell: v.ListString[9], HighestPricePointForDay: v.ListString[10], TimeHighestPricePoint: v.ListString[11], LowestPricePointForDay: v.ListString[12], TimeLowestPricePoint: v.ListString[13], Dow1: v.ListString[14], Dow2: v.ListString[15], Dow3: v.ListString[16], Dow4: v.ListString[17]}
		tradeResultStoreList = append(tradeResultStoreList, tradeResultStore)
		i++
	}
	return tradeResultStoreList
}
func truncateTradeResultStore() {
	postCommandDBTruncate("TRUNCATE table trade_result_store")
}

//insertDayTrackingRecord
func insertDayTrackingRecord(dayTrackingRecord DayTrackingRecord) {
	listValues := []string{dayTrackingRecord.Symbol, dayTrackingRecord.DayOfWeekCreated, dayTrackingRecord.DayOfWeekDayIteration, dayTrackingRecord.LastDayOfWeekDayUpdate, dayTrackingRecord.AmountOfTrades, dayTrackingRecord.IsWeekPassed}
	postCommandDBInsert("INSERT INTO day_tracking_record (symbol, day_of_week_created, day_of_week_day_iteration, last_day_of_week_day_update, amount_of_trades, is_week_passed) VALUES (", listValues)
}
func selectDayTrackingRecord() []DayTrackingRecord {
	dayTrackingRecordList := []DayTrackingRecord{}
	response := postCommandDBSelect("SELECT symbol, day_of_week_created, day_of_week_day_iteration, last_day_of_week_day_update, amount_of_trades, is_week_passed FROM day_tracking_record")
	container := parseDBResponse(response)
	fmt.Println(container.ListStringFromDB)
	fmt.Println(len(container.ListStringFromDB))
	for i, v := range container.ListStringFromDB {
		dayTrackingRecord := DayTrackingRecord{Symbol: v.ListString[0], DayOfWeekCreated: v.ListString[1], DayOfWeekDayIteration: v.ListString[2], LastDayOfWeekDayUpdate: v.ListString[3], AmountOfTrades: v.ListString[4], IsWeekPassed: v.ListString[5]}
		dayTrackingRecordList = append(dayTrackingRecordList, dayTrackingRecord)
		i++
	}
	return dayTrackingRecordList
}
func truncateDayTrackingRecord() {
	postCommandDBTruncate("TRUNCATE table information_at_trade")
}

//insertInformationAtTrade
func insertInformationAtTrade(informationAtTrade InformationAtTrade) {
	listValues := []string{informationAtTrade.Symbol, informationAtTrade.TypeTrade, informationAtTrade.Qty, informationAtTrade.Hour, informationAtTrade.Minute, informationAtTrade.Dow, informationAtTrade.Bid, informationAtTrade.Ask, informationAtTrade.Last}
	postCommandDBInsert("INSERT INTO information_at_trade (symbol, type_trade, qty, hour, minute, dow, bid, ask, last) VALUES (", listValues)
}
func selectInformationAtTrade() []InformationAtTrade {
	informationAtTradeList := []InformationAtTrade{}
	response := postCommandDBSelect("SELECT symbol, type_trade, qty, hour, minute, dow, bid, ask, last FROM information_at_trade")
	container := parseDBResponse(response)
	fmt.Println(container.ListStringFromDB)
	fmt.Println(len(container.ListStringFromDB))
	for i, v := range container.ListStringFromDB {
		informationAtTrade := InformationAtTrade{Symbol: v.ListString[0], TypeTrade: v.ListString[1], Qty: v.ListString[2], Hour: v.ListString[3], Minute: v.ListString[4], Dow: v.ListString[5], Bid: v.ListString[6], Ask: v.ListString[7], Last: v.ListString[8]}
		informationAtTradeList = append(informationAtTradeList, informationAtTrade)
		i++
	}
	return informationAtTradeList
}
func truncateInformationAtTrade() {
	postCommandDBTruncate("TRUNCATE table information_at_trade")
}

// sqlStatement := `
// INSERT INTO dow (current_dow_value)
// VALUES ($1)
// RETURNING id
// `
// var dow Dow
// row := db.QueryRow(sqlStatement, currentDowValue)
// err1 := row.Scan(&dow.ID)
// if err1 != nil {
// fmt.Println("Create Error 2")
// }
// fmt.Println(dow.ID)
// }
// func selectDow() []Dow {
// psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
// "dbname=%s sslmode=disable",
// host, port, user, dbname)
// db, err := sql.Open("postgres", psqlInfo)
// if err != nil {
// fmt.Println("Read Error 1")
// panic(err)
// }
// defer db.Close()

// rows, err1 := db.Query("SELECT id, created_at, current_dow_value FROM dow")

// psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
// 	"dbname=%s sslmode=disable",
// 	host, port, user, dbname)
// db, err := sql.Open("postgres", psqlInfo)
// if err != nil {
// 	fmt.Println("Create Error 1")
// }
// // type AltIntervalBuyWisemen struct {
// // 	CreatedAt              string
// // 	Symbol                 string
// // 	IsAltIntervalOperation string
// //  }

// // 	CREATE TABLE alt_interval_buy_wisemen
// // (
// //    id SERIAL PRIMARY KEY,
// //    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
// //    symbol VARCHAR,
// //    is_alt_interval_operation VARCHAR
// // );
// defer db.Close()
// sqlStatement := `
// INSERT INTO alt_interval_buy_wisemen (symbol, is_alt_interval_operation)
// VALUES ($1,$2)
// 	RETURNING id
// 	`
// var id int

// row := db.QueryRow(sqlStatement, altIntervalBuyWisemen.Symbol, altIntervalBuyWisemen.IsAltIntervalOperation)
// err1 := row.Scan(&id)
// if err1 != nil {
// 	fmt.Println("Create Error 2")
// }

//parse response needed

// 	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
// 		"dbname=%s sslmode=disable",
// 		host, port, user, dbname)
// 	db, err := sql.Open("postgres", psqlInfo)
// 	if err != nil {
// 		fmt.Println("Read Error 1")
// 		panic(err)
// 	}
// 	defer db.Close()
// 	rows, err1 := db.Query("SELECT created_at, symbol, is_alt_interval_operation FROM alt_interval_buy_wisemen")
// 	if err1 != nil {
// 		fmt.Println(err1)
// 	}
// 	defer rows.Close()
// 	altIntervalBuyWisemenList := make([]AltIntervalBuyWisemen, 0)

// 	for rows.Next() {
// 		var altIntervalBuyWisemen AltIntervalBuyWisemen
// 		if err2 := rows.Scan(&altIntervalBuyWisemen.CreatedAt, &altIntervalBuyWisemen.Symbol, &altIntervalBuyWisemen.IsAltIntervalOperation); err2 != nil {
// 			fmt.Println("err2")
// 		}
// 		altIntervalBuyWisemenList = append(altIntervalBuyWisemenList, altIntervalBuyWisemen)
// 	}
// 	return altIntervalBuyWisemenList
// }

// func postNeoBuyOrderResponse(holdingWisemen HoldingWisemen) string {

// 	// type HoldingWisemen struct {
// 	// 	CreatedAt   string
// 	// 	Symbol      string
// 	// 	Price       string
// 	// 	Qty         string
// 	// 	QtyBought   string
// 	// 	OrderStatus string
// 	// }

// 	// json := `{
// 	// 	"requestType": "postNeoBuyOrderResponse",
// 	// 	"data": [
// 	// 	`

// 	// // for indexEvalResult, evalResult := range evalResults {
// 	// json += "\"" + holdingWisemen.Symbol + "\","
// 	// //bool to string. strconv.FormatBool(v)
// 	// json += "\"" + holdingWisemen.Price + "\","
// 	// json += "\"" + holdingWisemen.QtyBought + "\","

// 	// json += "\"" + holdingWisemen.OrderStatus + "\""
// 	// json = json + `]}`

// 	json := `{
// 		"request_type": "postNeoBuyOrderResponse",
// 		"data": {
// 		`
// 	json = json + "\"symbol\":" + "\"" + holdingWisemen.Symbol + "\","
// 	json = json + "\"limit\":" + "\"" + holdingWisemen.Price + "\","
// 	json = json + "\"qty\":" + "\"" + holdingWisemen.OrderStatus + "\""
// 	json = json + `}}`

// 	url := "http://localhost:11000/databaseQuery"
// 	response := post(url, json)

// 	// fmt.Println("json")
// 	// fmt.Println(json)

// 	// response := ""
// 	return response
// }

// type HoldingWisemen struct {
// 	CreatedAt   string
// 	Symbol      string
// 	Price       string
// 	Qty         string
// 	QtyBought   string
// 	OrderStatus string
// }

func postNeoBuyOrderResponse(holdingWisemen HoldingWisemen) string {
	symbol := holdingWisemen.Symbol
	buyPrice := holdingWisemen.Price
	qty := holdingWisemen.Qty
	status := holdingWisemen.OrderStatus
	json := `{
		"requestType": "postNeoBuyOrderResponse",
		"data": [
			`
	json += "\"" + symbol + "\","
	json += "\"" + buyPrice + "\","
	json += "\"" + qty + "\","
	json += "\"" + status + "\""
	json = json + `]}`

	url := "http://localhost:11000/databaseQuery"
	response := post(url, json)
	return response
}

func postNeoHealthCheck() string {
	json := `{
		"requestType": "postNeoHealthCheck",
		"data": [
			]}`

	url := "http://localhost:11000/databaseQuery"
	response := post(url, json)
	return response
}

func postNeoTradeDayResult(symbol string, tradeDayResult string) string {
	json := `{
		"requestType": "tradeDayResult",
		"data": [
			`
	json += "\"" + symbol + "\","
	json += "\"" + tradeDayResult + "\""
	json = json + `]}`

	url := "http://localhost:11000/databaseQuery"
	response := post(url, json)
	return response
}
