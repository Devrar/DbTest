package main

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

var DbUser, DbPassword, DbName string

func GetConst() error {
	err := godotenv.Load()

	if err != nil {
		return err
	}

	var exists bool

	DbUser, exists = os.LookupEnv("DB_USER")

	if !exists {
		return errors.New("No database user found")
	}

	DbPassword, exists = os.LookupEnv("DB_PASSWORD")

	if !exists {
		return errors.New("No database password find")
	}

	DbName, exists = os.LookupEnv("DB_NAME")

	if !exists {
		return errors.New("No database name find")
	}

	return nil
}
