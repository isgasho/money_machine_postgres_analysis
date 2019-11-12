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

func insertDow(currentDowValue string) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("Create Error 1")
	}
	defer db.Close()

	sqlStatement := `
		INSERT INTO dow (current_dow_value)
		VALUES ($1)
		RETURNING id
		`
	var dow Dow
	row := db.QueryRow(sqlStatement, currentDowValue)
	err1 := row.Scan(&dow.ID)
	if err1 != nil {
		fmt.Println("Create Error 2")
	}
	fmt.Println(dow.ID)
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

	rows, err1 := db.Query("SELECT id, created_at, current_dow_value FROM dow")
	if err1 != nil {
		fmt.Println(err1)
	}
	defer rows.Close()
	dowList := make([]Dow, 0)

	for rows.Next() {
		// var symbol string
		var dowInstance Dow
		if err2 := rows.Scan(&dowInstance.ID, &dowInstance.CreatedAt, &dowInstance.CurrentDowValue); err2 != nil {
			fmt.Println("err2")
		}
		dowList = append(dowList, dowInstance)
	}
	return dowList
}

func dropDow() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("Read Error 1")
		panic(err)
	}
	defer db.Close()

	res, err1 := db.Exec("drop table dow")
	if err1 != nil {
		fmt.Println("Delete Error 2")
	}
	count, err2 := res.RowsAffected()
	if err2 != nil {
		fmt.Println("Delete Error 3")
	}
	fmt.Println(count)
}

func createDow() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("Read Error 1")
		panic(err)
	}
	defer db.Close()

	res, err1 := db.Exec(`CREATE TABLE dow
	( 
	   id SERIAL PRIMARY KEY,
	   created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
	   current_dow_value VARCHAR
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

//Stock Whale high
func insertStockWhaleHigh(stockEntry Stock) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("Create Error 1")
	}
	defer db.Close()

	sqlStatement := `
		INSERT INTO stock_whale_high (symbol, bid, ask, last, pchg, pcls, opn, vl, pvol, volatility12, wk52hi, wk52hidate, wk52lo, wk52lodate, hi, lo, pr_adp_50, pr_adp_100, prchg, adp_50, adp_100, adv_30, adv_90)
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
func selectAllStockWhaleHigh() []Stock {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("Read Error 1")
		panic(err)
	}
	defer db.Close()

	rows, err1 := db.Query("SELECT id, created_at, symbol, bid, ask, last, pchg, pcls, opn, vl, pvol, volatility12, wk52hi, wk52hidate, wk52lo, wk52lodate, hi, lo, pr_adp_50, pr_adp_100, prchg, adp_50, adp_100, adv_30, adv_90 FROM stock_whale_high")
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

func selectStockWhaleHigh(symbol string) []Stock {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("Read Error 1")
		panic(err)
	}
	defer db.Close()

	rows, err1 := db.Query("SELECT id, created_at, symbol, bid, ask, last, pchg, pcls, opn, vl, pvol, volatility12, wk52hi, wk52hidate, wk52lo, wk52lodate, hi, lo, pr_adp_50, pr_adp_100, prchg, adp_50, adp_100, adv_30, adv_90 FROM stock_whale_high WHERE symbol=$1", symbol)
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

func deleteAllStockOfSymbolInWhaleHigh(symbol string) []Stock {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("Read Error 1")
		panic(err)
	}
	defer db.Close()

	rows, err1 := db.Query("DELETE FROM stock_whale_high WHERE symbol=$1;", symbol)

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

//Stock Whale low
func insertStockWhaleLow(stockEntry Stock) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("Create Error 1")
	}
	defer db.Close()

	sqlStatement := `
		INSERT INTO stock_whale_low (symbol, bid, ask, last, pchg, pcls, opn, vl, pvol, volatility12, wk52hi, wk52hidate, wk52lo, wk52lodate, hi, lo, pr_adp_50, pr_adp_100, prchg, adp_50, adp_100, adv_30, adv_90)
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
func selectAllStockWhaleLow() []Stock {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("Read Error 1")
		panic(err)
	}
	defer db.Close()

	rows, err1 := db.Query("SELECT id, created_at, symbol, bid, ask, last, pchg, pcls, opn, vl, pvol, volatility12, wk52hi, wk52hidate, wk52lo, wk52lodate, hi, lo, pr_adp_50, pr_adp_100, prchg, adp_50, adp_100, adv_30, adv_90 FROM stock_whale_low")
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

func selectStockWhaleLow(symbol string) []Stock {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("Read Error 1")
		panic(err)
	}
	defer db.Close()

	rows, err1 := db.Query("SELECT id, created_at, symbol, bid, ask, last, pchg, pcls, opn, vl, pvol, volatility12, wk52hi, wk52hidate, wk52lo, wk52lodate, hi, lo, pr_adp_50, pr_adp_100, prchg, adp_50, adp_100, adv_30, adv_90 FROM stock_whale_low WHERE symbol=$1", symbol)
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

func deleteAllStockOfSymbolInWhaleLow(symbol string) []Stock {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("Read Error 1")
		panic(err)
	}
	defer db.Close()

	rows, err1 := db.Query("DELETE FROM stock_whale_low WHERE symbol=$1;", symbol)

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

//Temp symbol high
func insertTempSymbolHoldHigh(symbol string, userInput bool) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("Create Error 1")
	}
	defer db.Close()

	sqlStatement := `
		INSERT INTO temp_symbol_hold_high (symbol, user_inputed)
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

func selectTempSymbolHoldHigh() []string {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("Read Error 1")
		panic(err)
	}
	defer db.Close()

	rows, err1 := db.Query("SELECT symbol FROM temp_symbol_hold_high")
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

func deleteTempSymbolHoldHigh(symbolToDel string) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("Read Error 1")
		panic(err)
	}
	defer db.Close()

	res, err1 := db.Exec("DELETE FROM temp_symbol_hold_high WHERE symbol=$1", symbolToDel)
	if err1 != nil {
		fmt.Println("Delete Error 2")
	}
	count, err2 := res.RowsAffected()
	if err2 != nil {
		fmt.Println("Delete Error 3")
	}
	fmt.Println(count)
}

func dropTempSymbolHoldHigh() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("Read Error 1")
		panic(err)
	}
	defer db.Close()

	res, err1 := db.Exec("drop table temp_symbol_hold_high")
	if err1 != nil {
		fmt.Println("Delete Error 2")
	}
	count, err2 := res.RowsAffected()
	if err2 != nil {
		fmt.Println("Delete Error 3")
	}
	fmt.Println(count)
}

func createTempSymbolHoldHigh() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("Read Error 1")
		panic(err)
	}
	defer db.Close()

	res, err1 := db.Exec(`CREATE TABLE temp_symbol_hold_high
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

func insertTempSymbolHoldLow(symbol string, userInput bool) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("Create Error 1")
	}
	defer db.Close()

	sqlStatement := `
		INSERT INTO temp_symbol_hold_low (symbol, user_inputed)
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

func selectTempSymbolHoldLow() []string {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("Read Error 1")
		panic(err)
	}
	defer db.Close()

	rows, err1 := db.Query("SELECT symbol FROM temp_symbol_hold_low")
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

func deleteTempSymbolHoldLow(symbolToDel string) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("Read Error 1")
		panic(err)
	}
	defer db.Close()

	res, err1 := db.Exec("DELETE FROM temp_symbol_hold_low WHERE symbol=$1", symbolToDel)
	if err1 != nil {
		fmt.Println("Delete Error 2")
	}
	count, err2 := res.RowsAffected()
	if err2 != nil {
		fmt.Println("Delete Error 3")
	}
	fmt.Println(count)
}

func dropTempSymbolHoldLow() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("Read Error 1")
		panic(err)
	}
	defer db.Close()

	res, err1 := db.Exec("drop table temp_symbol_hold_low")
	if err1 != nil {
		fmt.Println("Delete Error 2")
	}
	count, err2 := res.RowsAffected()
	if err2 != nil {
		fmt.Println("Delete Error 3")
	}
	fmt.Println(count)
}

func createTempSymbolHoldLow() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("Read Error 1")
		panic(err)
	}
	defer db.Close()

	res, err1 := db.Exec(`CREATE TABLE temp_symbol_hold_low
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
func insertMetricsWisemen(desired_price_range_high string, desired_price_range_low string, price_high_pchg string, price_low_pchg string, desired_pchg_variance_value string, desired_volatility_variance_value string, trade_buy_monitor_delay_seconds string, trade_buy_monitor_delay_query_seconds string, trade_buy_monitor_delay_iteration_count string) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("Create Error 1")
	}
	defer db.Close()

	// 	CREATE TABLE metrics_wisemen
	// (
	//    id SERIAL PRIMARY KEY,
	//    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
	//    desired_price_range_high VARCHAR,
	//    desired_price_range_low VARCHAR,
	//    price_high_pchg VARCHAR,
	//    price_low_pchg VARCHAR,
	//    desired_pchg_variance_value VARCHAR,
	//    desired_volatility_variance_value VARCHAR,
	//    trade_buy_monitor_delay_seconds VARCHAR,
	//    trade_buy_monitor_delay_query_seconds VARCHAR,
	//    trade_buy_monitor_delay_iteration_count VARCHAR
	// );
	sqlStatement := `
		INSERT INTO metrics_wisemen (desired_price_range_high, desired_price_range_low, price_high_pchg, price_low_pchg, desired_pchg_variance_value, desired_volatility_variance_value, trade_buy_monitor_delay_seconds, trade_buy_monitor_delay_query_seconds, trade_buy_monitor_delay_iteration_count, end_trade_time)
			VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)
			RETURNING created_at
		`
	var metricsWisemen MetricsWisemen

	row := db.QueryRow(sqlStatement, desired_price_range_high, desired_price_range_low, price_high_pchg, price_low_pchg, desired_pchg_variance_value, desired_volatility_variance_value, trade_buy_monitor_delay_seconds, trade_buy_monitor_delay_query_seconds, trade_buy_monitor_delay_iteration_count, end_trade_time)
	err1 := row.Scan(&metricsWisemen.CreatedAt)
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

	// 	CREATE TABLE metrics_wisemen
	// (
	//    id SERIAL PRIMARY KEY,
	//    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
	//    desired_price_range_high VARCHAR,
	//    desired_price_range_low VARCHAR,
	//    price_high_pchg VARCHAR,
	//    price_low_pchg VARCHAR,
	//    desired_pchg_variance_value VARCHAR,
	//    desired_volatility_variance_value VARCHAR,
	//    trade_buy_monitor_delay_seconds VARCHAR,
	//    trade_buy_monitor_delay_query_seconds VARCHAR,
	//    trade_buy_monitor_delay_iteration_count VARCHAR
	// );

	rows, err1 := db.Query("SELECT created_at, desired_price_range_high, desired_price_range_low, price_high_pchg, price_low_pchg, desired_pchg_variance_value, desired_volatility_variance_value, trade_buy_monitor_delay_seconds, trade_buy_monitor_delay_query_seconds, trade_buy_monitor_delay_iteration_count FROM metrics_wisemen")
	if err1 != nil {
		// log.Fatal(err)
		fmt.Println(err1)
	}
	defer rows.Close()
	metricsList := make([]MetricsWisemen, 0)

	// DesiredPriceRangeHigh
	// DesiredPriceRangeLow
	for rows.Next() {
		var metricsWisemen MetricsWisemen
		if err2 := rows.Scan(&metricsWisemen.CreatedAt, &metricsWisemen.DesiredPriceRangeHigh, &metricsWisemen.DesiredPriceRangeLow, &metricsWisemen.PriceHighPchg, &metricsWisemen.PriceLowPchg, &metricsWisemen.DesiredPchgVarianceValue, &metricsWisemen.DesiredVolatilityVarianceValue, &metricsWisemen.TradeBuyMonitorDelaySeconds, &metricsWisemen.TradeBuyMonitorDelayQuerySeconds, &metricsWisemen.TradeBuyMonitorDelayIterationCount); err2 != nil {
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

	res, err1 := db.Exec("drop table metrics_wisemen")
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

	res, err1 := db.Exec(`CREATE TABLE metrics_wisemen
	( 
	   id SERIAL PRIMARY KEY,
	   created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
	   desired_price_range_high VARCHAR, 
	   desired_price_range_low VARCHAR,
	   price_high_pchg VARCHAR,
	   price_low_pchg VARCHAR,
	   desired_pchg_variance_value VARCHAR,
	   desired_volatility_variance_value VARCHAR,
	   trade_buy_monitor_delay_seconds VARCHAR,
	   trade_buy_monitor_delay_query_seconds VARCHAR,
	   trade_buy_monitor_delay_iteration_count VARCHAR
	);
	`)

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

//Whale high
func insertWhaleSymbolHoldHigh(symbol string, userInput bool) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("Create Error 1")
	}
	defer db.Close()

	sqlStatement := `
		INSERT INTO whale_symbol_hold_high (symbol, user_inputed)
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

func selectWhaleSymbolHoldHigh() []string {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("Read Error 1")
		panic(err)
	}
	defer db.Close()

	rows, err1 := db.Query("SELECT symbol FROM whale_symbol_hold_high")
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

func deleteWhaleSymbolHoldHigh(symbolToDel string) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("Read Error 1")
		panic(err)
	}
	defer db.Close()

	res, err1 := db.Exec("DELETE FROM whale_symbol_hold_high WHERE symbol=$1", symbolToDel)
	if err1 != nil {
		fmt.Println("Delete Error 2")
	}
	count, err2 := res.RowsAffected()
	if err2 != nil {
		fmt.Println("Delete Error 3")
	}
	fmt.Println(count)
}

//Whale low
func insertWhaleSymbolHoldLow(symbol string, userInput bool) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("Create Error 1")
	}
	defer db.Close()

	sqlStatement := `
		INSERT INTO whale_symbol_hold_low (symbol, user_inputed)
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

func selectWhaleSymbolHoldLow() []string {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("Read Error 1")
		panic(err)
	}
	defer db.Close()

	rows, err1 := db.Query("SELECT symbol FROM whale_symbol_hold_low")
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

func deleteWhaleSymbolHoldLow(symbolToDel string) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("Read Error 1")
		panic(err)
	}
	defer db.Close()

	res, err1 := db.Exec("DELETE FROM whale_symbol_hold_low WHERE symbol=$1", symbolToDel)
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

// func insertTradeBoughtEvaluation(tradeBoughtEvaluation TradeBoughtEvaluation) {
// 	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
// 		"dbname=%s sslmode=disable",
// 		host, port, user, dbname)
// 	db, err := sql.Open("postgres", psqlInfo)
// 	if err != nil {
// 		fmt.Println("Create Error 1")
// 	}
// 	defer db.Close()

// 	sqlStatement := `
// 		INSERT INTO trade_bought_evaluation (symbol, is_bought)
// 		VALUES ($1, $2)
// 		RETURNING id
// 		`
// 	var id int
// 	row := db.QueryRow(sqlStatement, tradeBoughtEvaluation.HoldingList.Symbol, tradeBoughtEvaluation.IsBought)
// 	err1 := row.Scan(&id)
// 	if err1 != nil {
// 		fmt.Println("Create Error 2")
// 	}
// }

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

func insertDayTrackingRecord(dayTrackingRecord DayTrackingRecord) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("Create Error 1")
	}
	defer db.Close()
	sqlStatement := `
		INSERT INTO day_tracking_record (symbol, day_of_week_created, day_of_week_day_iteration, last_day_of_week_day_update, amount_of_trades, is_week_passed)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id
		`
	var id int
	row := db.QueryRow(sqlStatement, dayTrackingRecord.Symbol, dayTrackingRecord.DayOfWeekCreated, dayTrackingRecord.DayOfWeekDayIteration, dayTrackingRecord.LastDayOfWeekDayUpdate, dayTrackingRecord.AmountOfTrades, dayTrackingRecord.IsWeekPassed)
	err1 := row.Scan(&id)
	if err1 != nil {
		fmt.Println("Create Error 2")
	}
}

func selectDayTrackingRecord() []DayTrackingRecord {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("Read Error 1")
		panic(err)
	}
	defer db.Close()

	rows, err1 := db.Query("SELECT created_at, symbol, day_of_week_created, day_of_week_day_iteration, last_day_of_week_day_update, amount_of_trades, is_week_passed FROM day_tracking_record")
	if err1 != nil {
		fmt.Println(err1)
	}
	defer rows.Close()
	dayTrackingRecordList := make([]DayTrackingRecord, 0)

	for rows.Next() {
		var dayTrackingRecord DayTrackingRecord
		if err2 := rows.Scan(&dayTrackingRecord.CreatedAt, &dayTrackingRecord.Symbol, &dayTrackingRecord.DayOfWeekCreated, &dayTrackingRecord.DayOfWeekDayIteration, &dayTrackingRecord.LastDayOfWeekDayUpdate, &dayTrackingRecord.AmountOfTrades, &dayTrackingRecord.IsWeekPassed); err2 != nil {
			fmt.Println("err2")
		}
		dayTrackingRecordList = append(dayTrackingRecordList, dayTrackingRecord)
	}
	return dayTrackingRecordList
}

func dropDayTrackingRecord() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("Read Error 1")
		panic(err)
	}
	defer db.Close()

	res, err1 := db.Exec("drop table day_tracking_record")
	if err1 != nil {
		fmt.Println("Delete Error 2")
	}
	count, err2 := res.RowsAffected()
	if err2 != nil {
		fmt.Println("Delete Error 3")
	}
	fmt.Println(count)
}

func createDayTrackingRecord() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("Read Error 1")
		panic(err)
	}
	defer db.Close()

	res, err1 := db.Exec(`CREATE TABLE same_day_trade_tracking_record
	( 
	   id SERIAL PRIMARY KEY,
	   created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
	   symbol VARCHAR,
	   day_of_week_created VARCHAR,
	   day_of_week_day_iteration VARCHAR,
	   last_day_of_week_day_update VARCHAR,
	   amount_of_trades VARCHAR,
	   is_week_passed VARCHAR
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

//information_at_trade
func insertInformationAtTrade(informationAtTrade InformationAtTrade) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("Create Error 1")
	}
	defer db.Close()
	sqlStatement := `
		INSERT INTO information_at_trade (symbol, hour, minute, dow, bid, ask, last)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id
		`
	var id int
	row := db.QueryRow(sqlStatement, informationAtTrade.Symbol, informationAtTrade.Hour, informationAtTrade.Minute, informationAtTrade.Dow, informationAtTrade.Bid, informationAtTrade.Ask, informationAtTrade.Last)
	err1 := row.Scan(&id)
	if err1 != nil {
		fmt.Println("Create Error 2")
	}
}

func selectInformationAtTrade() []InformationAtTrade {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("Read Error 1")
		panic(err)
	}
	defer db.Close()

	rows, err1 := db.Query("SELECT created_at, symbol, hour, minute, dow, bid, ask, last FROM information_at_trade")
	if err1 != nil {
		fmt.Println(err1)
	}
	defer rows.Close()
	informationAtTradeList := make([]InformationAtTrade, 0)

	for rows.Next() {
		var informationAtTrade InformationAtTrade
		if err2 := rows.Scan(&informationAtTrade.CreatedAt, &informationAtTrade.Symbol, &informationAtTrade.Hour, &informationAtTrade.Minute, &informationAtTrade.Dow, &informationAtTrade.Bid, &informationAtTrade.Ask, &informationAtTrade.Last); err2 != nil {
			fmt.Println("err2")
		}
		informationAtTradeList = append(informationAtTradeList, informationAtTrade)
	}
	return informationAtTradeList
}

func dropInformationAtTrade() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("Read Error 1")
		panic(err)
	}
	defer db.Close()

	res, err1 := db.Exec("drop table information_at_trade")
	if err1 != nil {
		fmt.Println("Delete Error 2")
	}
	count, err2 := res.RowsAffected()
	if err2 != nil {
		fmt.Println("Delete Error 3")
	}
	fmt.Println(count)
}

func createInformationAtTrade() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("Read Error 1")
		panic(err)
	}
	defer db.Close()
	res, err1 := db.Exec(`CREATE TABLE information_at_trade
	(
	   id SERIAL PRIMARY KEY,
	   created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
	   symbol VARCHAR,
	   hour VARCHAR,
	   minute VARCHAR,
	   dow VARCHAR,
	   bid VARCHAR,
	   ask VARCHAR,
   	   last VARCHAR
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
//trade_conditional_metrics
func inserTradeConditionalMetrics(tradeConditionalMetrics TradeConditionalMetrics) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("Create Error 1")
	}
	// 	CREATE TABLE trade_conditional_metrics
	// (
	//    id SERIAL PRIMARY KEY,
	//    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
	//    symbol VARCHAR,
	//    time_start VARCHAR,
	//    time_end VARCHAR,
	//    price_dropout VARCHAR
	// );
	defer db.Close()
	sqlStatement := `
		INSERT INTO trade_conditional_metrics (symbol, time_start, time_end, price_dropout)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
		`
	var id int
	row := db.QueryRow(sqlStatement, tradeConditionalMetrics.Symbol, tradeConditionalMetrics.TimeStart, tradeConditionalMetrics.TimeEnd, tradeConditionalMetrics.PriceDropout)
	err1 := row.Scan(&id)
	if err1 != nil {
		fmt.Println("Create Error 2")
	}
}

func selectTradeConditionalMetrics() []TradeConditionalMetrics {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("Read Error 1")
		panic(err)
	}
	defer db.Close()

	// 	CREATE TABLE trade_conditional_metrics
	// (
	//    id SERIAL PRIMARY KEY,
	//    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
	//    symbol VARCHAR,
	//    time_start VARCHAR,
	//    time_end VARCHAR,
	//    price_dropout VARCHAR
	// );

	rows, err1 := db.Query("SELECT created_at, symbol, time_start, time_end, price_dropout FROM trade_conditional_metrics")
	if err1 != nil {
		fmt.Println(err1)
	}
	defer rows.Close()
	tradeConditionalMetricsList := make([]TradeConditionalMetrics, 0)

	for rows.Next() {
		var tradeConditionalMetrics TradeConditionalMetrics
		if err2 := rows.Scan(&tradeConditionalMetrics.CreatedAt, &tradeConditionalMetrics.Symbol, &tradeConditionalMetrics.TimeStart, &tradeConditionalMetrics.TimeEnd, &tradeConditionalMetrics.PriceDropout); err2 != nil {
			fmt.Println("err2")
		}
		tradeConditionalMetricsList = append(tradeConditionalMetricsList, tradeConditionalMetrics)
	}
	return tradeConditionalMetricsList
}

func dropTradeConditionalMetrics() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("Read Error 1")
		panic(err)
	}
	defer db.Close()

	res, err1 := db.Exec("drop table trade_conditional_metrics")
	if err1 != nil {
		fmt.Println("Delete Error 2")
	}
	count, err2 := res.RowsAffected()
	if err2 != nil {
		fmt.Println("Delete Error 3")
	}
	fmt.Println(count)
}

func createTradeConditionalMetrics() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("Read Error 1")
		panic(err)
	}
	defer db.Close()

	res, err1 := db.Exec(`CREATE TABLE trade_conditional_metrics
	(
	   id SERIAL PRIMARY KEY,
	   created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
	   symbol VARCHAR,
	   time_start VARCHAR,
	   time_end VARCHAR,
	   price_dropout VARCHAR
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

func deleteTradeConditionalMetrics(symbolToDel string) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("Read Error 1")
		panic(err)
	}
	defer db.Close()

	res, err1 := db.Exec("DELETE FROM trade_conditional_metrics WHERE symbol=$1", symbolToDel)
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

//
//holdingsWisemen
func insertHoldingWisemen(holdingWisemen HoldingWisemen) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("Create Error 1")
	}
	// 	CREATE TABLE holding_wisemen
	// (
	//    id SERIAL PRIMARY KEY,
	//    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
	//    symbol VARCHAR,
	//    price VARCHAR,
	//    qty VARCHAR,
	//    qty_bought VARCHAR,
	//    order_status VARCHAR
	// );

	defer db.Close()
	sqlStatement := `
		INSERT INTO holding_wisemen (symbol, price, qty, order_status)
		VALUES ($1, $2, $3, $4)
		RETURNING id
		`
	var id int

	row := db.QueryRow(sqlStatement, holdingWisemen.Symbol, holdingWisemen.Price, holdingWisemen.Qty, holdingWisemen.OrderStatus)
	err1 := row.Scan(&id)
	if err1 != nil {
		fmt.Println("Create Error 2")
	}
}

func selectHoldingWisemen() []HoldingWisemen {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("Read Error 1")
		panic(err)
	}
	defer db.Close()

	// 	CREATE TABLE holding_wisemen
	// (
	//    id SERIAL PRIMARY KEY,
	//    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
	//    symbol VARCHAR,
	//    price VARCHAR,
	//    qty VARCHAR,
	//    qty_bought VARCHAR,
	//    order_status VARCHAR
	// );

	rows, err1 := db.Query("SELECT created_at, symbol, price, qty, order_status FROM holding_wisemen")
	if err1 != nil {
		fmt.Println(err1)
	}
	defer rows.Close()
	holdingWisemenList := make([]HoldingWisemen, 0)

	for rows.Next() {
		var holdingWisemen HoldingWisemen
		if err2 := rows.Scan(&holdingWisemen.CreatedAt, &holdingWisemen.Symbol, &holdingWisemen.Price, &holdingWisemen.Qty, &holdingWisemen.OrderStatus); err2 != nil {
			fmt.Println("err2")
		}
		holdingWisemenList = append(holdingWisemenList, holdingWisemen)
	}
	return holdingWisemenList
}

func deleteHoldingWisemen(symbolToDel string) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("Read Error 1")
		panic(err)
	}
	defer db.Close()

	res, err1 := db.Exec("DELETE FROM holding_wisemen WHERE symbol=$1", symbolToDel)
	if err1 != nil {
		fmt.Println("Delete Error 2")
	}
	count, err2 := res.RowsAffected()
	if err2 != nil {
		fmt.Println("Delete Error 3")
	}
	fmt.Println(count)
}

func dropHoldingWisemen() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("Read Error 1")
		panic(err)
	}
	defer db.Close()

	res, err1 := db.Exec("drop table holding_wisemen")
	if err1 != nil {
		fmt.Println("Delete Error 2")
	}
	count, err2 := res.RowsAffected()
	if err2 != nil {
		fmt.Println("Delete Error 3")
	}
	fmt.Println(count)
}

func createHoldingWisemen() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("Read Error 1")
		panic(err)
	}
	defer db.Close()

	res, err1 := db.Exec(`CREATE TABLE holding_wisemen
	( 
	   id SERIAL PRIMARY KEY,
	   created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
	   symbol VARCHAR,
	   price VARCHAR,
	   qty VARCHAR,
	   order_status VARCHAR
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

//
//TradeResultStore
func insertTradeResultStore(tradeResultStore TradeResultStore) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("Create Error 1")
	}
	// 	CREATE TABLE trade_result_store
	// (
	//    id SERIAL PRIMARY KEY,
	//    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
	//    algorithm_used VARCHAR,
	//    result VARCHAR,
	//    change_amount VARCHAR,
	//    stock_symbol VARCHAR,
	//    time_start VARCHAR,
	//    time_end VARCHAR,
	//    time_trade_buy VARCHAR,
	//    time_trade_sell VARCHAR,
	//    dow1 VARCHAR,
	//    dow2 VARCHAR,
	//    dow3 VARCHAR,
	//    dow4 VARCHAR
	// );
	defer db.Close()
	sqlStatement := `
		INSERT INTO trade_result_store (algorithm_used, result, change_amount, stock_symbol, time_start, time_end, time_trade_buy, time_trade_sell, dow1, dow2, dow3, dow4)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
		RETURNING id
		`
	var id int
	row := db.QueryRow(sqlStatement, tradeResultStore.AlgorithmUsed, tradeResultStore.Result, tradeResultStore.ChangeAmount, tradeResultStore.StockSymbol, tradeResultStore.TimeStart, tradeResultStore.TimeEnd, tradeResultStore.TimeTradeBuy, tradeResultStore.TimeTradeSell, tradeResultStore.Dow1, tradeResultStore.Dow2, tradeResultStore.Dow3, tradeResultStore.Dow4)
	err1 := row.Scan(&id)
	if err1 != nil {
		fmt.Println("Create Error 2")
	}
}

func selectTradeResultStore(algorithmUsed string) []TradeResultStore {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("Read Error 1")
		panic(err)
	}
	// 	CREATE TABLE trade_result_store
	// (
	//    id SERIAL PRIMARY KEY,
	//    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
	//    algorithm_used VARCHAR,
	//    result VARCHAR,
	//    change_amount VARCHAR,
	//    stock_symbol VARCHAR,
	//    time_start VARCHAR,
	//    time_end VARCHAR,
	//    time_trade_buy VARCHAR,
	//    time_trade_sell VARCHAR,
	//    dow1 VARCHAR,
	//    dow2 VARCHAR,
	//    dow3 VARCHAR,
	//    dow4 VARCHAR
	// );
	defer db.Close()
	rows, err1 := db.Query("SELECT created_at, algorithm_used, result, change_amount, stock_symbol, time_start, time_end, time_trade_buy, time_trade_sell, dow1, dow2, dow3, dow4 FROM trade_result_store") // WHERE symbol=$1", symbol)
	if err1 != nil {
		fmt.Println(err1)
	}
	defer rows.Close()
	tradeResultStoreList := []TradeResultStore{} //make([]AlgorithmEvaluationForDay, 0)

	for rows.Next() {
		var tradeResultStore TradeResultStore
		if err2 := rows.Scan(&tradeResultStore.CreatedAt, &tradeResultStore.AlgorithmUsed, &tradeResultStore.Result, &tradeResultStore.ChangeAmount, &tradeResultStore.StockSymbol, &tradeResultStore.TimeStart, &tradeResultStore.TimeEnd, &tradeResultStore.TimeTradeBuy, &tradeResultStore.TimeTradeSell, &tradeResultStore.Dow1, &tradeResultStore.Dow2, &tradeResultStore.Dow3, &tradeResultStore.Dow4); err2 != nil {
			fmt.Println("err2")
		}
		tradeResultStoreList = append(tradeResultStoreList, tradeResultStore)
	}
	return tradeResultStoreList
}

func deleteTradeResultStore(createdAt string) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("Read Error 1")
		panic(err)
	}
	defer db.Close()

	res, err1 := db.Exec("DELETE FROM trade_result_store WHERE created_at=$1", createdAt)
	if err1 != nil {
		fmt.Println("Delete Error 2")
	}
	count, err2 := res.RowsAffected()
	if err2 != nil {
		fmt.Println("Delete Error 3")
	}
	fmt.Println(count)
}
