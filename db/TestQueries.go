package db

import (
	"log"
)

func TestQuery() string {
	var test string
	row := DB.QueryRow("SELECT test FROM testtable  WHERE id = $1 ", 1)
	if err := row.Scan(&test); err != nil {
		log.Fatalf("Error scanning row: %v", err)
	}
	return test

}
