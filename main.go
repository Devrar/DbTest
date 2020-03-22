package main

import (
	"fmt"

	"github.com/Devrar/DbTest/database"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("Go MySQL Tutorial")

	err := GetConst()

	if err != nil {
		panic(err)
	}

	db, err := database.OpenDB(DbUser, DbPassword, DbName)

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	columns := []string{"username"}
	values := []string{"Io", "Tizio"}

	err = database.InsQuery(db, "Persone", columns, values)

	if err != nil {
		panic(err.Error())
	}
}
