package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("Go MySQL Tutorial")

	db, err := sql.Open("mysql", "goUser:password@/test")
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	insert, err := db.Query("INSERT INTO Persone (username) VALUES (\"Devrar\")")
	if err != nil {
		panic(err.Error())
	}

	defer insert.Close()
}
