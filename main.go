package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host   = "localhost"
	port   = 5432
	user   = "postgres"
	dbname = "money_machine"
)

type User struct {
	ID        int
	Age       int
	FirstName string
	LastName  string
	Email     string
}

func main() {
	id := createEntry(25, "trucker17@gmail.com", "Joe", "Tammison")
	fmt.Println(readEntry(id))
}

func createEntry(age int, email string, firstName string, lastName string) int {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("Error 1")
		panic(err)
	}
	defer db.Close()

	sqlStatement := `
		INSERT INTO users (age, email, first_name, last_name)
		VALUES ($1, $2, $3, $4)
		RETURNING id, age, email, first_name, last_name
		`
	var user User
	row := db.QueryRow(sqlStatement, age, email, firstName, lastName)
	err1 := row.Scan(&user.ID, &user.Age, &user.Email, &user.FirstName, &user.LastName)
	if err1 != nil {
		fmt.Println("Error 2")
		panic(err1)
	}
	return user.ID
}

func readEntry(idToSearch int) (int, int, string, string, string) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("Error 1")
		panic(err)
	}
	defer db.Close()

	sqlStatement := `SELECT * FROM users WHERE id=$1;`
	var user User
	row := db.QueryRow(sqlStatement, idToSearch)
	err1 := row.Scan(&user.ID, &user.Age, &user.FirstName,
		&user.LastName, &user.Email)
	if err1 != nil {
		fmt.Println("Error 2")
		panic(err1)
	}
	fmt.Println(user.ID, user.Age, user.Email, user.FirstName, user.LastName)
	return user.ID, user.Age, user.Email, user.FirstName, user.LastName
}

func updateEntry(idToSearch int) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	sqlStatement := `SELECT * FROM users WHERE id=$1;`
	var user User
	row := db.QueryRow(sqlStatement, idToSearch)
	err1 := row.Scan(&user.ID, &user.Age, &user.FirstName,
		&user.LastName, &user.Email)
	switch err1 {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return
	case nil:
		fmt.Println(user)
	default:
		panic(err)
	}
}

func deleteEntry(idToSearch int) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	sqlStatement := `SELECT * FROM users WHERE id=$1;`
	var user User
	row := db.QueryRow(sqlStatement, idToSearch)
	err1 := row.Scan(&user.ID, &user.Age, &user.FirstName,
		&user.LastName, &user.Email)
	switch err1 {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return
	case nil:
		fmt.Println(user)
	default:
		panic(err)
	}
}
