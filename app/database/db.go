package database

import (
	"database/sql"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	user     = "postgres"
	password = "password"
	host     = "localhost"
	port     = 5432
	dbname   = "ebook"
)

func ConnectDB() (*gorm.DB, *sql.DB, error) {
	// postgres://postgres:password@localhost:5432/ebook?sslmode=disable
	dsn := fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s sslmode=disable", user, password, host, port, dbname)
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
