package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/Devrar/DbTest/database"
	"github.com/joho/godotenv"

	_ "github.com/go-sql-driver/mysql"
)

func init() {
	err := godotenv.Load()

	if err != nil {
		log.Print("No .env file found!")
	}
}

func main() {
	fmt.Println("Go MySQL Tutorial")

	DbUser, exists := os.LookupEnv("DB_USER")

	if !exists {
		panic(errors.New("No database user found"))
	}

	DbPassword, exists := os.LookupEnv("DB_PASSWORD")

	if !exists {
		panic(errors.New("No database password find"))
	}

	db, err := database.OpenDB(DbUser, DbPassword)

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
