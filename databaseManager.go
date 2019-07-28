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
		INSERT INTO day (age, email, first_name, last_name)
		VALUES ($1, $2, $3, $4)
		RETURNING id, age, email, first_name, last_name
		`
	var day Day
	row := db.QueryRow(sqlStatement, dayOfWeek)
	err1 := row.Scan(&day.ID, &day.Age, &user.Email, &user.FirstName, &user.LastName)
	if err1 != nil {
		fmt.Println("Create Error 2")
	}
	return user.ID
}
func setDay() {
}
func selectDay() {
}
func deleteDay() {
}

func insertNews() {
}
func setNews() {
}
func selectNews() {
}
func deleteNews() {
}

func insertDow() {
}
func setDow() {
}
func selectDow() {
}
func deleteDow() {
}

func insertStock() {
}
func setStock() {
}
func selectStock() {
}
func deleteStock() {
}

func insertTradeInfo() {
}
func setTradeInfo() {
}
func selectTradeInfo() {
}
func deleteTradeInfo() {
}
