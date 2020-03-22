package database

import (
	"database/sql"
	"errors"
	"fmt"
)

func OpenDB(user string, password string) (db *sql.DB, err error) {
	db, err = sql.Open("mysql", user+":"+password+"@/test")

	return db, err
}

func InsQuery(db *sql.DB, table string, columns []string, values []string) error {
	if len(values) == 0 {
		return errors.New("No values passed")
	}

	query := "INSERT INTO " + table

	if len(columns) != 0 {
		query += " ("
		for i, column := range columns {
			if i != len(columns)-1 {
				query += column + ", "
			} else {
				query += column + ")"
			}
		}
	}

	query += " VALUES"

	for i, value := range values {
		if i != len(values)-1 {
			query += " (\"" + value + "\"),"
		} else {
			query += " (\"" + value + "\");"
		}
	}

	fmt.Println(query)

	insert, err := db.Query(query)

	if err != nil {
		return err
	}

	defer insert.Close()
	return nil
}
