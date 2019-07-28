package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func createEntry(age int, email string, firstName string, lastName string) int {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("Create Error 1")
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
		fmt.Println("Create Error 2")
	}
	return user.ID
}

func readEntry(idToSearch int) (int, int, string, string, string) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("Read Error 1")
		panic(err)
	}
	defer db.Close()

	sqlStatement := `SELECT * FROM users WHERE id=$1;`
	var user User
	row := db.QueryRow(sqlStatement, idToSearch)
	err1 := row.Scan(&user.ID, &user.Age, &user.FirstName,
		&user.LastName, &user.Email)
	if err1 != nil {
		fmt.Println("Read Error 2")
		return 0, 0, "null", "null", "null"
	}
	return user.ID, user.Age, user.Email, user.FirstName, user.LastName
}

func updateEntry(idToSearch int, age int, email string, firstName string, lastName string) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("Update Error 1")
	}
	defer db.Close()
	sqlStatement := `
			UPDATE users
			SET age= $2, email = $3, first_name = $4, last_name = $5
			WHERE id = $1
			RETURNING id, age, email, first_name, last_name
			;`
	res, err1 := db.Exec(sqlStatement, idToSearch, age, email, firstName, lastName)
	if err1 != nil {
		fmt.Println("Update Error 2")
	}
	count, err2 := res.RowsAffected()
	if err2 != nil {
		fmt.Println("Update Error 3")
	}
	fmt.Println(count)
}

func deleteEntry(idToSearch int) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("Delete Error 1")
	}
	defer db.Close()

	sqlStatement := `
		DELETE FROM users
		WHERE id = $1;
		`
	res, err1 := db.Exec(sqlStatement, idToSearch)
	if err1 != nil {
		fmt.Println("Delete Error 2")
	}
	count, err2 := res.RowsAffected()
	if err2 != nil {
		fmt.Println("Delete Error 3")
	}
	fmt.Println(count)
}
