package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDB() {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	sslmode := os.Getenv("DB_SSLMODE")

	dsn := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=%s",
		user, password, host, port, dbname, sslmode)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
		return
	}
	err = db.Ping()
	if err != nil {
		log.Fatalf("Error connecting to DB: %v", err)
	}
	DB = db
}
