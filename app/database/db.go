package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	_"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// const (
// 	user     = "postgres"
// 	password = "password"
// 	host     = "localhost"
// 	port     = 5432
// 	dbname   = "ebook"
// )

func ConnectDB() (*gorm.DB, *sql.DB, error) {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_NAME")

	// postgres://postgres:password@localhost:5432/ebook?sslmode=disable
	dsn := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable", user, password, host, port, dbname)
	gDb, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database : %v", err)
	}

	// Getting SQL DB object
	sqlDB, err := gDb.DB()
	if err != nil {
		log.Fatal(err)
	}

	// Checking the connection is alive
	if err := sqlDB.Ping(); err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to db successfully...")
	return gDb, sqlDB, nil
}
