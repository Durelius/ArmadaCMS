package db

import "log"

func TestQuery() string {
	var test string
	testSelectSQL := "SELECT test FROM testtable WHERE id = $1"
	err := DB.Get(&test, testSelectSQL, 1)
	if err != nil {
		log.Fatal(err)
	}
	return test

}
