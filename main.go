package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	user     = "postgres"
	password = "password"
	host     = "localhost"
	port     = 5432
	dbname   = "ebook"
)

func main() {
	// postgres://postgres:password@localhost:5432/ebook?sslmode=disable
	connectionStr := fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s sslmode=disable", user, password, host, port, dbname)

	db, err := sql.Open("postgres", connectionStr)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	// connection checking
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to db successfully...")
}
