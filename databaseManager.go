package main

import (
	"database/sql"
	"fmt"
)

const (
	host   = "localhost"
	port   = 5432
	user   = "postgres"
	dbname = "money_machine"
)

func insertDay(dayOfWeek string) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("Create Error 1")
	}
	defer db.Close()

	sqlStatement := `
		INSERT INTO day (day_of_Week)
		VALUES ($1)
		RETURNING id, created_at, day_of_Week
		`
	var day Day
	row := db.QueryRow(sqlStatement, dayOfWeek)
	err1 := row.Scan(&day.ID, &day.CreatedAt, &day.DayOfWeek)
	if err1 != nil {
		fmt.Println("Create Error 2")
	}
	fmt.Println(day.ID, day.DayOfWeek, day.CreatedAt)
}
func setDay() {
}
func selectDay() {
}
func deleteDay() {
}

func insertNews(newsEntry News) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("Create Error 1")
	}
	defer db.Close()

	sqlStatement := `
		INSERT INTO news (day_id, news_info)
		VALUES ($1, $2)
		RETURNING id, day_id, created_at, news_info
		`
	var news News
	row := db.QueryRow(sqlStatement, newsEntry.DayID, newsEntry.NewsInfo)
	err1 := row.Scan(&news.ID, &news.DayID, &news.CreatedAt, &news.NewsInfo)
	if err1 != nil {
		fmt.Println("Create Error 2")
	}
	fmt.Println(news.ID, news.DayID, news.CreatedAt, news.NewsInfo)
}
func setNews() {
}
func selectNews() {
}
func deleteNews() {
}

func insertDow(currentDowValue string, pointsChanged string, percentageChange string) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("Create Error 1")
	}
	defer db.Close()

	sqlStatement := `
		INSERT INTO dow (current_dow_value, points_changed, percentage_change)
		VALUES ($1, $2, $3)
		RETURNING id
		`
	var dow Dow
	row := db.QueryRow(sqlStatement, currentDowValue, pointsChanged, percentageChange)
	err1 := row.Scan(&dow.ID)
	if err1 != nil {
		fmt.Println("Create Error 2")
	}
	fmt.Println(dow.ID)
}
func setDow() {
}
func selectDow() []Dow {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("Read Error 1")
		panic(err)
	}
	defer db.Close()

	rows, err1 := db.Query("SELECT id, created_at, current_dow_value, points_changed, percentage_change FROM dow")
	if err1 != nil {
		fmt.Println(err1)
	}
	defer rows.Close()
	dowList := make([]Dow, 0)

	for rows.Next() {
		// var symbol string
		var dowInstance Dow
		if err2 := rows.Scan(&dowInstance.ID, &dowInstance.CreatedAt, &dowInstance.CurrentDowValue, &dowInstance.PointsChanged, &dowInstance.PercentageChange); err2 != nil {
			fmt.Println("err2")
		}
		dowList = append(dowList, dowInstance)
	}
	return dowList
}

func deleteDow() {
}

func insertStockWisemen(stockEntry Stock) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("Create Error 1")
	}
	defer db.Close()

	sqlStatement := `
		INSERT INTO stock_wisemen (symbol, bid, ask, last, pchg, pcls, opn, vl, pvol, volatility12, wk52hi, wk52hidate, wk52lo, wk52lodate, hi, lo, pr_adp_50, pr_adp_100, prchg, adp_50, adp_100, adv_30, adv_90)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23)
		RETURNING id, created_at, symbol, bid, ask, last, pchg, pcls, opn, vl, pvol, volatility12, wk52hi, wk52hidate, wk52lo, wk52lodate, hi, lo, pr_adp_50, pr_adp_100, prchg, adp_50, adp_100, adv_30, adv_90
		`
	var stock Stock

	row := db.QueryRow(sqlStatement, stockEntry.Symbol, stockEntry.Bid, stockEntry.Ask, stockEntry.Last, stockEntry.Pchg, stockEntry.Pcls, stockEntry.Opn, stockEntry.Vl, stockEntry.Pvol, stockEntry.Volatility12, stockEntry.Wk52hi, stockEntry.Wk52hidate, stockEntry.Wk52lo, stockEntry.Wk52lodate, stockEntry.Hi, stockEntry.Lo, stockEntry.PrAdp50, stockEntry.PrAdp100, stockEntry.Prchg, stockEntry.Adp50, stockEntry.Adp100, stockEntry.Adv30, stockEntry.Adv90)
	err1 := row.Scan(&stock.ID, &stock.CreatedAt, &stock.Symbol, &stock.Bid, &stock.Ask, &stock.Last, &stock.Pchg, &stock.Pcls, &stock.Opn, &stock.Vl, &stock.Pvol, &stock.Volatility12, &stock.Wk52hi, &stock.Wk52hidate, &stock.Wk52lo, &stock.Wk52lodate, &stock.Hi, &stock.Lo, &stock.PrAdp50, &stock.PrAdp100, &stock.Prchg, &stock.Adp50, &stock.Adp100, &stock.Adv30, &stock.Adv90)
	if err1 != nil {
		fmt.Println("Create Error 2")
	}
}
func selectAllStockWisemen() []Stock {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("Read Error 1")
		panic(err)
	}
	defer db.Close()

	rows, err1 := db.Query("SELECT id, created_at, symbol, bid, ask, last, pchg, pcls, opn, vl, pvol, volatility12, wk52hi, wk52hidate, wk52lo, wk52lodate, hi, lo, pr_adp_50, pr_adp_100, prchg, adp_50, adp_100, adv_30, adv_90 FROM stock_wisemen")
	if err1 != nil {
		fmt.Println(err1)
	}
	defer rows.Close()
	stockList := make([]Stock, 0)

	for rows.Next() {
		var stock Stock
		if err2 := rows.Scan(&stock.ID, &stock.CreatedAt, &stock.Symbol, &stock.Bid, &stock.Ask, &stock.Last, &stock.Pchg, &stock.Pcls, &stock.Opn, &stock.Vl, &stock.Pvol, &stock.Volatility12, &stock.Wk52hi, &stock.Wk52hidate, &stock.Wk52lo, &stock.Wk52lodate, &stock.Hi, &stock.Lo, &stock.PrAdp50, &stock.PrAdp100, &stock.Prchg, &stock.Adp50, &stock.Adp100, &stock.Adv30, &stock.Adv90); err2 != nil {
			fmt.Println("err2")
		}
		stockList = append(stockList, stock)
	}
	return stockList
}

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

func insertStockWhale(stockEntry Stock) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("Create Error 1")
	}
	defer db.Close()

	sqlStatement := `
		INSERT INTO stock_whale (symbol, bid, ask, last, pchg, pcls, opn, vl, pvol, volatility12, wk52hi, wk52hidate, wk52lo, wk52lodate, hi, lo, pr_adp_50, pr_adp_100, prchg, adp_50, adp_100, adv_30, adv_90)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23)
		RETURNING id, created_at, symbol, bid, ask, last, pchg, pcls, opn, vl, pvol, volatility12, wk52hi, wk52hidate, wk52lo, wk52lodate, hi, lo, pr_adp_50, pr_adp_100, prchg, adp_50, adp_100, adv_30, adv_90
		`
	var stock Stock

	row := db.QueryRow(sqlStatement, stockEntry.Symbol, stockEntry.Bid, stockEntry.Ask, stockEntry.Last, stockEntry.Pchg, stockEntry.Pcls, stockEntry.Opn, stockEntry.Vl, stockEntry.Pvol, stockEntry.Volatility12, stockEntry.Wk52hi, stockEntry.Wk52hidate, stockEntry.Wk52lo, stockEntry.Wk52lodate, stockEntry.Hi, stockEntry.Lo, stockEntry.PrAdp50, stockEntry.PrAdp100, stockEntry.Prchg, stockEntry.Adp50, stockEntry.Adp100, stockEntry.Adv30, stockEntry.Adv90)
	err1 := row.Scan(&stock.ID, &stock.CreatedAt, &stock.Symbol, &stock.Bid, &stock.Ask, &stock.Last, &stock.Pchg, &stock.Pcls, &stock.Opn, &stock.Vl, &stock.Pvol, &stock.Volatility12, &stock.Wk52hi, &stock.Wk52hidate, &stock.Wk52lo, &stock.Wk52lodate, &stock.Hi, &stock.Lo, &stock.PrAdp50, &stock.PrAdp100, &stock.Prchg, &stock.Adp50, &stock.Adp100, &stock.Adv30, &stock.Adv90)
	if err1 != nil {
		fmt.Println("Create Error 2")
	}
}
func selectAllStockWhale() []Stock {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("Read Error 1")
		panic(err)
	}
	defer db.Close()

	rows, err1 := db.Query("SELECT id, created_at, symbol, bid, ask, last, pchg, pcls, opn, vl, pvol, volatility12, wk52hi, wk52hidate, wk52lo, wk52lodate, hi, lo, pr_adp_50, pr_adp_100, prchg, adp_50, adp_100, adv_30, adv_90 FROM stock_whale")
	if err1 != nil {
		fmt.Println(err1)
	}
	defer rows.Close()
	stockList := make([]Stock, 0)

	for rows.Next() {
		var stock Stock
		if err2 := rows.Scan(&stock.ID, &stock.CreatedAt, &stock.Symbol, &stock.Bid, &stock.Ask, &stock.Last, &stock.Pchg, &stock.Pcls, &stock.Opn, &stock.Vl, &stock.Pvol, &stock.Volatility12, &stock.Wk52hi, &stock.Wk52hidate, &stock.Wk52lo, &stock.Wk52lodate, &stock.Hi, &stock.Lo, &stock.PrAdp50, &stock.PrAdp100, &stock.Prchg, &stock.Adp50, &stock.Adp100, &stock.Adv30, &stock.Adv90); err2 != nil {
			fmt.Println("err2")
		}
		stockList = append(stockList, stock)
	}
	return stockList
}

func selectStockWhale(symbol string) []Stock {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("Read Error 1")
		panic(err)
	}
	defer db.Close()

	rows, err1 := db.Query("SELECT id, created_at, symbol, bid, ask, last, pchg, pcls, opn, vl, pvol, volatility12, wk52hi, wk52hidate, wk52lo, wk52lodate, hi, lo, pr_adp_50, pr_adp_100, prchg, adp_50, adp_100, adv_30, adv_90 FROM stock_whale WHERE symbol=$1", symbol)
	if err1 != nil {
		fmt.Println(err1)
	}
	defer rows.Close()
	stockList := make([]Stock, 0)

	for rows.Next() {
		var stock Stock
		if err2 := rows.Scan(&stock.ID, &stock.CreatedAt, &stock.Symbol, &stock.Bid, &stock.Ask, &stock.Last, &stock.Pchg, &stock.Pcls, &stock.Opn, &stock.Vl, &stock.Pvol, &stock.Volatility12, &stock.Wk52hi, &stock.Wk52hidate, &stock.Wk52lo, &stock.Wk52lodate, &stock.Hi, &stock.Lo, &stock.PrAdp50, &stock.PrAdp100, &stock.Prchg, &stock.Adp50, &stock.Adp100, &stock.Adv30, &stock.Adv90); err2 != nil {
			fmt.Println("err2")
		}
		stockList = append(stockList, stock)
	}
	return stockList
}

func deleteAllStockOfSymbolInWhale(symbol string) []Stock {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("Read Error 1")
		panic(err)
	}
	defer db.Close()

	rows, err1 := db.Query("DELETE FROM stock_whale WHERE symbol=$1;", symbol)

	if err1 != nil {
		fmt.Println(err1)
	}
	defer rows.Close()
	stockList := make([]Stock, 0)

	for rows.Next() {
		var stock Stock
		if err2 := rows.Scan(&stock.ID, &stock.CreatedAt, &stock.Symbol, &stock.Bid, &stock.Ask, &stock.Last, &stock.Pchg, &stock.Pcls, &stock.Opn, &stock.Vl, &stock.Pvol, &stock.Volatility12, &stock.Wk52hi, &stock.Wk52hidate, &stock.Wk52lo, &stock.Wk52lodate, &stock.Hi, &stock.Lo, &stock.PrAdp50, &stock.PrAdp100, &stock.Prchg, &stock.Adp50, &stock.Adp100, &stock.Adv30, &stock.Adv90); err2 != nil {
			fmt.Println("err2")
		}
		stockList = append(stockList, stock)
	}
	return stockList
}

func setStock() {
}

func selectAllStockOfSymbol(symbolToSearch string) []Stock {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("Read Error 1")
		panic(err)
	}
	defer db.Close()

	rows, err1 := db.Query("SELECT id, created_at, monitoring, symbol, bid, ask, last, pchg, pcls, opn, vl, pvol, volatility12, wk52hi, wk52hidate, wk52lo, wk52lodate, hi, lo, pr_adp_50, pr_adp_100, prchg, adp_50, adp_100, adv_30, adv_90 FROM stock WHERE symbol=$1", symbolToSearch)
	if err1 != nil {
		fmt.Println(err1)
	}
	defer rows.Close()
	stockList := make([]Stock, 0)

	for rows.Next() {
		var stock Stock
		if err2 := rows.Scan(&stock.ID, &stock.CreatedAt, &stock.Monitoring, &stock.Symbol, &stock.Bid, &stock.Ask, &stock.Last, &stock.Pchg, &stock.Pcls, &stock.Opn, &stock.Vl, &stock.Pvol, &stock.Volatility12, &stock.Wk52hi, &stock.Wk52hidate, &stock.Wk52lo, &stock.Wk52lodate, &stock.Hi, &stock.Lo, &stock.PrAdp50, &stock.PrAdp100, &stock.Prchg, &stock.Adp50, &stock.Adp100, &stock.Adv30, &stock.Adv90); err2 != nil {
			fmt.Println("err2")
		}
		stockList = append(stockList, stock)
	}
	return stockList
}

func deleteStock(symbolToDel string) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("Read Error 1")
		panic(err)
	}
	defer db.Close()

	res, err1 := db.Exec("DELETE FROM stock WHERE symbol=$1", symbolToDel)
	if err1 != nil {
		fmt.Println("Delete Error 2")
	}
	count, err2 := res.RowsAffected()
	if err2 != nil {
		fmt.Println("Delete Error 3")
	}
	fmt.Println(count)
}

// func selectMonitoringStock() []string {
// 	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
// 		"dbname=%s sslmode=disable",
// 		host, port, user, dbname)
// 	db, err := sql.Open("postgres", psqlInfo)
// 	if err != nil {
// 		fmt.Println("Read Error 1")
// 		panic(err)
// 	}
// 	defer db.Close()

// 	rows, err1 := db.Query("SELECT symbol FROM stock WHERE monitoring=true")
// 	if err1 != nil {
// 		fmt.Println(err1)
// 	}
// 	defer rows.Close()
// 	resultList := make([]string, 0)

// 	for rows.Next() {
// 		var symbol string
// 		if err2 := rows.Scan(&symbol); err2 != nil {
// 			fmt.Println("err2")
// 		}

// 		resultList = append(resultList, symbol)
// 	}
// 	return resultList
// }

// func selectAllMonitoringStock() []string {
// 	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
// 		"dbname=%s sslmode=disable",
// 		host, port, user, dbname)
// 	db, err := sql.Open("postgres", psqlInfo)
// 	if err != nil {
// 		fmt.Println("Read Error 1")
// 		panic(err)
// 	}
// 	defer db.Close()

// 	// sqlStatement := `SELECT symbol FROM stock WHERE monitoring =$1;`
// 	// var stock Stock
// 	rows, err1 := db.Query("SELECT symbol FROM stock WHERE monitoring=true LIMIT 1")
// 	if err1 != nil {
// 		fmt.Println(err1)
// 	}
// 	defer rows.Close()
// 	// idList := make([]string, 0)
// 	// resultList := make([]string, 0)
// 	resultList := []string{}

// 	for rows.Next() {
// 		var symbol string
// 		if err2 := rows.Scan(&symbol); err2 != nil {
// 			fmt.Println("err2")
// 		}

// 		resultList = append(resultList, symbol)
// 		// fmt.Println(symbol, last)
// 	}
// 	return resultList
// }

//

func insertTempSymbolHold(symbol string, userInput bool) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("Create Error 1")
	}
	defer db.Close()

	sqlStatement := `
		INSERT INTO temp_symbol_hold (symbol, user_inputed)
			VALUES ($1,$2)
			RETURNING symbol
		`
	var stock Stock

	row := db.QueryRow(sqlStatement, symbol, userInput)
	err1 := row.Scan(&stock.Symbol)
	if err1 != nil {
		fmt.Println("Create Error 2")
	}
}

func selectTempSymbolHold() []string {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("Read Error 1")
		panic(err)
	}
	defer db.Close()

	rows, err1 := db.Query("SELECT symbol FROM temp_symbol_hold")
	if err1 != nil {
		// log.Fatal(err)
		fmt.Println(err1)
	}
	defer rows.Close()
	symbolList := make([]string, 0)

	for rows.Next() {
		var symbol string
		if err2 := rows.Scan(&symbol); err2 != nil {
			fmt.Println("err2")
		}
		symbolList = append(symbolList, symbol)
	}
	return symbolList
}

func deleteTempSymbolHold(symbolToDel string) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("Read Error 1")
		panic(err)
	}
	defer db.Close()

	res, err1 := db.Exec("DELETE FROM temp_symbol_hold WHERE symbol=$1", symbolToDel)
	if err1 != nil {
		fmt.Println("Delete Error 2")
	}
	count, err2 := res.RowsAffected()
	if err2 != nil {
		fmt.Println("Delete Error 3")
	}
	fmt.Println(count)
}

func dropTempSymbolHold() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("Read Error 1")
		panic(err)
	}
	defer db.Close()

	res, err1 := db.Exec("drop table temp_symbol_hold")
	if err1 != nil {
		fmt.Println("Delete Error 2")
	}
	count, err2 := res.RowsAffected()
	if err2 != nil {
		fmt.Println("Delete Error 3")
	}
	fmt.Println(count)
}

func createTempSymbolHold() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("Read Error 1")
		panic(err)
	}
	defer db.Close()

	res, err1 := db.Exec(`CREATE TABLE temp_symbol_hold
	( 
	   id SERIAL PRIMARY KEY,
	   created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
	   symbol VARCHAR,
	   user_inputed boolean
	);`)

	if err1 != nil {
		fmt.Println("Delete Error 2")
	}
	count, err2 := res.RowsAffected()
	if err2 != nil {
		fmt.Println("Delete Error 3")
	}
	fmt.Println(count)
}

//Whale metrics
func insertMetricsWhale(desired_price_range_high string, desired_price_range_low string, desired_pchg, desired_pchg_variance_value string, desired_volatility_variance_value string) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("Create Error 1")
	}
	defer db.Close()

	sqlStatement := `
		INSERT INTO metrics_whale (desired_price_range_high, desired_price_range_low, desired_pchg, desired_pchg_variance_value, desired_volatility_variance_value)
			VALUES ($1,$2,$3,$4,$5)
			RETURNING created_at
		`
	var metricsWhale MetricsWhale

	row := db.QueryRow(sqlStatement, desired_price_range_high, desired_price_range_low, desired_pchg, desired_pchg_variance_value, desired_volatility_variance_value)
	err1 := row.Scan(&metricsWhale.CreatedAt)
	if err1 != nil {
		fmt.Println("Create Error 2")
	}
}

func selectMetricsWhale() []MetricsWhale {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("Read Error 1")
		panic(err)
	}
	defer db.Close()

	rows, err1 := db.Query("SELECT created_at, desired_price_range_high, desired_price_range_low, desired_pchg, desired_pchg_variance_value, desired_volatility_variance_value FROM metrics_whale")
	if err1 != nil {
		// log.Fatal(err)
		fmt.Println(err1)
	}
	defer rows.Close()
	metricsList := make([]MetricsWhale, 0)

	for rows.Next() {
		var metricsWhale MetricsWhale
		// DesiredPchg                    string
		if err2 := rows.Scan(&metricsWhale.CreatedAt, &metricsWhale.DesiredPriceRangeHigh, &metricsWhale.DesiredPriceRangeLow, &metricsWhale.DesiredPchg, &metricsWhale.DesiredPchgVarianceValue, &metricsWhale.DesiredVolatilityVarianceValue); err2 != nil {
			fmt.Println("err2")
		}
		metricsList = append(metricsList, metricsWhale)
	}
	return metricsList
}

func dropMetricsWhale() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("Read Error 1")
		panic(err)
	}
	defer db.Close()

	res, err1 := db.Exec("drop table metrics_whale")
	if err1 != nil {
		fmt.Println("Delete Error 2")
	}
	count, err2 := res.RowsAffected()
	if err2 != nil {
		fmt.Println("Delete Error 3")
	}
	fmt.Println(count)
}

func createMetricsWhale() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("Read Error 1")
		panic(err)
	}
	defer db.Close()

	res, err1 := db.Exec(`CREATE TABLE metrics_whale
	( 
		id SERIAL PRIMARY KEY,
		created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
		desired_price_range_high VARCHAR,
		desired_price_range_low VARCHAR,
		desired_pchg VARCHAR,
		desired_pchg_variance_value VARCHAR,
		desired_volatility_variance_value VARCHAR
	);`)

	if err1 != nil {
		fmt.Println("Delete Error 2")
	}
	count, err2 := res.RowsAffected()
	if err2 != nil {
		fmt.Println("Delete Error 3")
	}
	fmt.Println(count)
}

//Wisemen metrics
func insertMetricsWisemen(desired_price_range_high string, desired_price_range_low string, desired_pchg, desired_pchg_variance_value string, desired_volatility_variance_value string) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("Create Error 1")
	}
	defer db.Close()

	sqlStatement := `
		INSERT INTO metrics_wisemen (desired_price_range_high, desired_price_range_low, desired_pchg, desired_pchg_variance_value, desired_volatility_variance_value)
			VALUES ($1,$2,$3,$4,$5)
			RETURNING created_at
		`
	var metricsWhale MetricsWhale

	row := db.QueryRow(sqlStatement, desired_price_range_high, desired_price_range_low, desired_pchg, desired_pchg_variance_value, desired_volatility_variance_value)
	err1 := row.Scan(&metricsWhale.CreatedAt)
	if err1 != nil {
		fmt.Println("Create Error 2")
	}
}

func selectMetricsWisemen() []MetricsWisemen {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("Read Error 1")
		panic(err)
	}
	defer db.Close()

	rows, err1 := db.Query("SELECT created_at, desired_price_range_high, desired_price_range_low, desired_pchg, desired_pchg_variance_value, desired_volatility_variance_value FROM metrics_wisemen")
	if err1 != nil {
		// log.Fatal(err)
		fmt.Println(err1)
	}
	defer rows.Close()
	metricsList := make([]MetricsWisemen, 0)

	for rows.Next() {
		var metricsWisemen MetricsWisemen
		// DesiredPchg                    string
		if err2 := rows.Scan(&metricsWisemen.CreatedAt, &metricsWisemen.DesiredPriceRangeHigh, &metricsWisemen.DesiredPriceRangeLow, &metricsWisemen.DesiredPchg, &metricsWisemen.DesiredPchgVarianceValue, &metricsWisemen.DesiredVolatilityVarianceValue); err2 != nil {
			fmt.Println("err2")
		}
		metricsList = append(metricsList, metricsWisemen)
	}
	return metricsList
}

func dropMetricsWisemen() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("Read Error 1")
		panic(err)
	}
	defer db.Close()

	res, err1 := db.Exec("drop table metrics_whale")
	if err1 != nil {
		fmt.Println("Delete Error 2")
	}
	count, err2 := res.RowsAffected()
	if err2 != nil {
		fmt.Println("Delete Error 3")
	}
	fmt.Println(count)
}

func createMetricsWisemen() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("Read Error 1")
		panic(err)
	}
	defer db.Close()

	res, err1 := db.Exec(`CREATE TABLE metrics_whale
	( 
		id SERIAL PRIMARY KEY,
		created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
		desired_price_range_high VARCHAR,
		desired_price_range_low VARCHAR,
		desired_pchg VARCHAR,
		desired_pchg_variance_value VARCHAR,
		desired_volatility_variance_value VARCHAR
	);`)

	if err1 != nil {
		fmt.Println("Delete Error 2")
	}
	count, err2 := res.RowsAffected()
	if err2 != nil {
		fmt.Println("Delete Error 3")
	}
	fmt.Println(count)
}

//
func insertWisemenSymbolHold(symbol string, userInput bool) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("Create Error 1")
	}
	defer db.Close()

	sqlStatement := `
		INSERT INTO wisemen_symbol_hold (symbol, user_inputed)
			VALUES ($1,$2)
			RETURNING symbol
		`
	var stock Stock

	row := db.QueryRow(sqlStatement, symbol, userInput)
	err1 := row.Scan(&stock.Symbol)
	if err1 != nil {
		fmt.Println("Create Error 2")
	}
}

func selectWisemenSymbolHold() []string {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("Read Error 1")
		panic(err)
	}
	defer db.Close()

	rows, err1 := db.Query("SELECT symbol FROM wisemen_symbol_hold")
	if err1 != nil {
		// log.Fatal(err)
		fmt.Println(err1)
	}
	defer rows.Close()
	symbolList := make([]string, 0)

	for rows.Next() {
		var symbol string
		if err2 := rows.Scan(&symbol); err2 != nil {
			fmt.Println("err2")
		}
		symbolList = append(symbolList, symbol)
	}
	return symbolList
}

func deleteWisemenSymbolHold(symbolToDel string) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("Read Error 1")
		panic(err)
	}
	defer db.Close()

	res, err1 := db.Exec("DELETE FROM wisemen_symbol_hold WHERE symbol=$1", symbolToDel)
	if err1 != nil {
		fmt.Println("Delete Error 2")
	}
	count, err2 := res.RowsAffected()
	if err2 != nil {
		fmt.Println("Delete Error 3")
	}
	fmt.Println(count)
}

func dropWisemenSymbolHold() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("Read Error 1")
		panic(err)
	}
	defer db.Close()

	res, err1 := db.Exec("drop table wisemen_symbol_hold")
	if err1 != nil {
		fmt.Println("Delete Error 2")
	}
	count, err2 := res.RowsAffected()
	if err2 != nil {
		fmt.Println("Delete Error 3")
	}
	fmt.Println(count)
}

func createWisemenSymbolHold() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("Read Error 1")
		panic(err)
	}
	defer db.Close()

	res, err1 := db.Exec(`CREATE TABLE wisemen_symbol_hold
	( 
	   id SERIAL PRIMARY KEY,
	   created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
	   symbol VARCHAR,
	   user_inputed boolean
	);`)

	if err1 != nil {
		fmt.Println("Delete Error 2")
	}
	count, err2 := res.RowsAffected()
	if err2 != nil {
		fmt.Println("Delete Error 3")
	}
	fmt.Println(count)
}

//

func insertWhaleSymbolHold(symbol string, userInput bool) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("Create Error 1")
	}
	defer db.Close()

	sqlStatement := `
		INSERT INTO whale_symbol_hold (symbol, user_inputed)
			VALUES ($1,$2)
			RETURNING symbol
		`
	var stock Stock

	row := db.QueryRow(sqlStatement, symbol, userInput)
	err1 := row.Scan(&stock.Symbol)
	if err1 != nil {
		fmt.Println("Create Error 2")
	}
}

func selectWhaleSymbolHold() []string {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("Read Error 1")
		panic(err)
	}
	defer db.Close()

	rows, err1 := db.Query("SELECT symbol FROM whale_symbol_hold")
	if err1 != nil {
		// log.Fatal(err)
		fmt.Println(err1)
	}
	defer rows.Close()
	symbolList := make([]string, 0)

	for rows.Next() {
		var symbol string
		if err2 := rows.Scan(&symbol); err2 != nil {
			fmt.Println("err2")
		}
		symbolList = append(symbolList, symbol)
	}
	return symbolList
}

func deleteWhaleSymbolHold(symbolToDel string) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("Read Error 1")
		panic(err)
	}
	defer db.Close()

	res, err1 := db.Exec("DELETE FROM whale_symbol_hold WHERE symbol=$1", symbolToDel)
	if err1 != nil {
		fmt.Println("Delete Error 2")
	}
	count, err2 := res.RowsAffected()
	if err2 != nil {
		fmt.Println("Delete Error 3")
	}
	fmt.Println(count)
}

//
func insertAnalyticsOperations(topStockList []Stock) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("Create Error 1")
	}
	defer db.Close()

	stockRanking := ""
	for indexStock, stock := range topStockList {
		stockRanking += stock.Symbol + ", "
		indexStock++
	}

	sqlStatement := `
		INSERT INTO analytics_operations (stock_ranking)
			VALUES ($1)
			RETURNING stock_ranking
		`
	var ranking string
	row := db.QueryRow(sqlStatement, stockRanking)
	err1 := row.Scan(&ranking)
	if err1 != nil {
		fmt.Println("Create Error 2")
	}
	fmt.Println(ranking)
}

func insertTradeInfo() {
	// psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
	// 	"dbname=%s sslmode=disable",
	// 	host, port, user, dbname)
	// db, err := sql.Open("postgres", psqlInfo)
	// if err != nil {
	// 	fmt.Println("Create Error 1")
	// }
	// defer db.Close()

	// sqlStatement := `
	// 	INSERT INTO dow (day_id, dow_info)
	// 	VALUES ($1, $2)
	// 	RETURNING id, day_id, created_at, dow_info
	// 	`
	// var dow Dow
	// row := db.QueryRow(sqlStatement, dayID, dowInfo)
	// err1 := row.Scan(&dow.ID, &dow.DayID, &dow.CreatedAt, &dow.DowInfo)
	// if err1 != nil {
	// 	fmt.Println("Create Error 2")
	// }
	// fmt.Println(dow.DayID, dow.CreatedAt)
}
func setTradeInfo() {
}
func selectTradeInfo() {
}
func deleteTradeInfo() {
}

func insertEndOfDayAnalyticsOperations(marketClosed bool, day string) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("Create Error 1")
	}
	defer db.Close()

	sqlStatement := `
		INSERT INTO end_of_day_analytics_operations (market_closed, day_of_week)
		VALUES ($1, $2)
		RETURNING id
		`
	var dow Dow
	row := db.QueryRow(sqlStatement, marketClosed, day)
	err1 := row.Scan(&dow.ID)
	if err1 != nil {
		fmt.Println("Create Error 2")
	}
}

func insertEvalResultsWhale(evalResult EvalResultsWhale) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("Create Error 1")
	}
	defer db.Close()

	sqlStatement := `
		INSERT INTO eval_results_whale (symbol, is_breach_worthy, is_pattern_met)
		VALUES ($1, $2, $3)
		RETURNING id
		`
	var evalResultsWhale EvalResultsWhale
	row := db.QueryRow(sqlStatement, evalResult.Symbol, evalResult.IsBreachWorthy, evalResult.IsPatternMet)
	err1 := row.Scan(&evalResultsWhale.ID)
	if err1 != nil {
		fmt.Println("Create Error 2")
	}
}

func selectEvalResultsWhale() []string {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("Read Error 1")
		panic(err)
	}
	defer db.Close()

	rows, err1 := db.Query("SELECT symbol FROM eval_results_whale")
	if err1 != nil {
		fmt.Println(err1)
	}
	defer rows.Close()
	symbolList := make([]string, 0)

	for rows.Next() {
		var symbol string
		if err2 := rows.Scan(&symbol); err2 != nil {
			fmt.Println("err2")
		}
		symbolList = append(symbolList, symbol)
	}
	return symbolList
}

func dropEvalResultsWhale() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("Read Error 1")
		panic(err)
	}
	defer db.Close()

	res, err1 := db.Exec("drop table eval_results_whale")
	if err1 != nil {
		fmt.Println("Delete Error 2")
	}
	count, err2 := res.RowsAffected()
	if err2 != nil {
		fmt.Println("Delete Error 3")
	}
	fmt.Println(count)
}

func createEvalResultsWhale() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("Read Error 1")
		panic(err)
	}
	defer db.Close()

	res, err1 := db.Exec(`CREATE TABLE eval_results_whale
	( 
		id SERIAL PRIMARY KEY,
		created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
		symbol VARCHAR,
		is_breach_worthy VARCHAR,
		is_pattern_met VARCHAR
	 );`)

	if err1 != nil {
		fmt.Println("Delete Error 2")
	}
	count, err2 := res.RowsAffected()
	if err2 != nil {
		fmt.Println("Delete Error 3")
	}
	fmt.Println(count)
}

func insertOrderInformationWisemen() {
	// psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
	// 	"dbname=%s sslmode=disable",
	// 	host, port, user, dbname)
	// db, err := sql.Open("postgres", psqlInfo)
	// if err != nil {
	// 	fmt.Println("Create Error 1")
	// }
	// defer db.Close()

	// sqlStatement := `
	// 	INSERT INTO eval_results_whale (symbol, is_breach_worthy, is_pattern_met)
	// 	VALUES ($1, $2, $3)
	// 	RETURNING id
	// 	`
	// var evalResultsWhale EvalResultsWhale
	// row := db.QueryRow(sqlStatement, evalResult.Symbol, evalResult.IsBreachWorthy, evalResult.IsPatternMet)
	// err1 := row.Scan(&evalResultsWhale.ID)
	// if err1 != nil {
	// 	fmt.Println("Create Error 2")
	// }
}

func selectOrderInformationWisemen() []OrderInformationWisemen {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("Read Error 1")
		panic(err)
	}
	defer db.Close()
	// type OrderInformationWisemen struct {
	// 	CreatedAt string
	// 	IsBought  string
	// 	Symbol    string
	// }
	rows, err1 := db.Query("SELECT symbol, is_bought FROM order_information_wisemen")
	if err1 != nil {
		fmt.Println(err1)
	}
	defer rows.Close()
	orderInformationWisemenList := make([]OrderInformationWisemen, 0)

	for rows.Next() {
		var orderInformationWisemen OrderInformationWisemen
		if err2 := rows.Scan(&orderInformationWisemen.Symbol, orderInformationWisemen.IsBought); err2 != nil {
			fmt.Println("err2")
		}
		orderInformationWisemenList = append(orderInformationWisemenList, orderInformationWisemen)
	}
	return orderInformationWisemenList
}

func insertTradeBoughtEvaluation(tradeBoughtEvaluation TradeBoughtEvaluation) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("Create Error 1")
	}
	defer db.Close()

	sqlStatement := `
		INSERT INTO trade_bought_evaluation (symbol, is_bought)
		VALUES ($1, $2)
		RETURNING id
		`
	var id int
	row := db.QueryRow(sqlStatement, tradeBoughtEvaluation.Holdings.Symbol, tradeBoughtEvaluation.IsBought)
	err1 := row.Scan(&id)
	if err1 != nil {
		fmt.Println("Create Error 2")
	}
}

func selectTradeBoughtEvaluation() []OrderInformationWisemen {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("Read Error 1")
		panic(err)
	}
	defer db.Close()
	// type OrderInformationWisemen struct {
	// 	CreatedAt string
	// 	IsBought  string
	// 	Symbol    string
	// }
	rows, err1 := db.Query("SELECT symbol, is_bought FROM order_information_wisemen")
	if err1 != nil {
		fmt.Println(err1)
	}
	defer rows.Close()
	orderInformationWisemenList := make([]OrderInformationWisemen, 0)

	for rows.Next() {
		var orderInformationWisemen OrderInformationWisemen
		if err2 := rows.Scan(&orderInformationWisemen.Symbol, orderInformationWisemen.IsBought); err2 != nil {
			fmt.Println("err2")
		}
		orderInformationWisemenList = append(orderInformationWisemenList, orderInformationWisemen)
	}
	return orderInformationWisemenList
}
