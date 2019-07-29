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

func insertDow(dowEntry Dow) {
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
	var dow Dow
	row := db.QueryRow(sqlStatement, dowEntry.DayID, dowEntry.DowInfo)
	err1 := row.Scan(&dow.ID, &dow.DayID, &dow.CreatedAt, &dow.DowInfo)
	if err1 != nil {
		fmt.Println("Create Error 2")
	}
	fmt.Println(dow.ID, dow.DayID, dow.CreatedAt, dow.DowInfo)
}
func setDow() {
}
func selectDow() {
}
func deleteDow() {
}

func insertStock(stockEntry Stock) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("Create Error 1")
	}
	defer db.Close()

	sqlStatement := `
		INSERT INTO stock (day_id, monitoring, symbol, bid, ask, last, pchg, pcls, opn, vl, pvol, volatility12, wk52hi, wk52hidate, wk52lo, wk52lodate, hi, low, pr_adp_50, pr_adp_100, prchg, adp_50, adp_100, adv_30, adv_90)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25)
		RETURNING id, day_id, created_at, monitoring, symbol, bid, ask, last, pchg, pcls, opn, vl, pvol, volatility12, wk52hi, wk52hidate, wk52lo, wk52lodate, hi, low, pr_adp_50, pr_adp_100, prchg, adp_50, adp_100, adv_30, adv_90
		`
	// ,
	var stock Stock
	row := db.QueryRow(sqlStatement, stockEntry.DayID, stockEntry.Monitoring, stockEntry.Symbol, stockEntry.Bid, stockEntry.Ask, stockEntry.Last, stockEntry.Pchg, stockEntry.Pcls, stockEntry.Opn, stockEntry.Vl, stockEntry.Pvol, stockEntry.Volatility12, stockEntry.Wk52hi, stockEntry.Wk52hidate, stockEntry.Wk52lo, stockEntry.Wk52lodate, stockEntry.Hi, stockEntry.Low, stockEntry.PrAdp50, stockEntry.PrAdp100, stockEntry.Prchg, stockEntry.Adp50, stockEntry.Adp100, stockEntry.Adv30, stockEntry.Adv90)
	err1 := row.Scan(&stock.ID, &stock.DayID, &stock.CreatedAt, &stock.Monitoring, &stock.Symbol, &stock.Bid, &stock.Ask, &stock.Last, &stock.Pchg, &stock.Pcls, &stock.Opn, &stock.Vl, &stock.Pvol, &stock.Volatility12, &stock.Wk52hi, &stock.Wk52hidate, &stock.Wk52lo, &stock.Wk52lodate, &stock.Hi, &stock.Low, &stock.PrAdp50, &stock.PrAdp100, &stock.Prchg, &stock.Adp50, &stock.Adp100, &stock.Adv30, &stock.Adv90)
	if err1 != nil {
		fmt.Println("Create Error 2")
	}
	fmt.Println(stock.ID, stock.DayID, stock.CreatedAt, stock.Monitoring, stock.Symbol, stock.Bid, stock.Ask) //, stock.Symbol, stock.Bid)
}
func setStock() {
}
func selectAllStock(symbolToSearch string) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("Read Error 1")
		panic(err)
	}
	defer db.Close()

	rows, err1 := db.Query("SELECT id, symbol, last FROM stock WHERE symbol=$1", symbolToSearch)
	if err1 != nil {
		// log.Fatal(err)
		fmt.Println(err1)
	}
	defer rows.Close()
	// idList := make([]string, 0)
	// symbolList := make([]string, 0)

	for rows.Next() {
		var id string
		var symbol string
		var last string
		if err2 := rows.Scan(&id, &symbol, &last); err2 != nil {
			// Check for a scan error.
			// Query rows will be closed with defer.
			// log.Fatal(err)
			fmt.Println("err2")
		}
		fmt.Println(id, symbol, last)
		// idList = append(idList, id)
		// symbolList = append(symbolList, id)
		// fmt.Println(idList)
	}

	// sqlStatement := `SELECT id FROM stock WHERE symbol=$1;`

	// // var stock Stock
	// rows, err1 := db.Query(sqlStatement, symbolToSearch)
	// if err1 != nil {
	// 	// handle this error better than this
	// 	// panic(err)
	// 	fmt.Println("Select all error1")
	// }
	// for rows.Next() {
	// 	var id int
	// 	// var Symbol string
	// 	err2 := rows.Scan(id)
	// 	if err2 != nil {
	// 		// handle this error
	// 		// panic(err)
	// 		print("Error 2")
	// 	}
	// 	fmt.Println(id)
	// }

	// var stock Stock
	// row := db.QueryRow(sqlStatement, symbolToSearch)
	// err1 := row.Scan(&stock.ID, &stock.DayID)
	// if err1 != nil {
	// 	fmt.Println("Read Error 2")
	// }
	// fmt.Println(stock.ID, stock.DayID)
}

// func selectStockCondition(condition int) {
// 	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
// 		"dbname=%s sslmode=disable",
// 		host, port, user, dbname)
// 	db, err := sql.Open("postgres", psqlInfo)
// 	if err != nil {
// 		fmt.Println("Read Error 1")
// 		panic(err)
// 	}
// 	defer db.Close()

// 	sqlStatement := `SELECT * FROM users WHERE id=$1;`
// 	var user User
// 	row := db.QueryRow(sqlStatement, idToSearch)
// 	err1 := row.Scan(&user.ID, &user.Age, &user.FirstName,
// 		&user.LastName, &user.Email)
// 	if err1 != nil {
// 		fmt.Println("Read Error 2")
// 		return 0, 0, "null", "null", "null"
// 	}
// 	return user.ID, user.Age, user.Email, user.FirstName, user.LastName
// }
func deleteStock() {
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
