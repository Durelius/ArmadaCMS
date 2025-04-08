package db

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	_ "github.com/lib/pq"
)

var DB *gorm.DB

func ConnectDB() {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	sslmode := os.Getenv("DB_SSLMODE")

	dsn := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=%s",
		user, password, host, port, dbname, sslmode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	DB = db
}

// func ConnectDB() {
// 	host := os.Getenv("DB_HOST")
// 	port := os.Getenv("DB_PORT")
// 	user := os.Getenv("DB_USER")
// 	password := os.Getenv("DB_PASSWORD")
// 	dbname := os.Getenv("DB_NAME")
// 	sslmode := os.Getenv("DB_SSLMODE")

// 	dsn := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=%s",
// 		user, password, host, port, dbname, sslmode)

// 	db, err := sqlx.Open("postgres", dsn)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	err = db.Ping()
// 	if err != nil {
// 		log.Fatalf("Error connecting to DB: %v", err)
// 	}
// 	DB = db
// }
