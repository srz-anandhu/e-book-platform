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

var db *sql.DB
var err error

func createUser(mail, username, password, salt string) (userID int, err error) {
	query := ` INSERT INTO users(mail,username,password,salt) VALUES($1,$2,$3,$4) RETURNING id `

	if err := db.QueryRow(query, mail, username, password, salt).Scan(&userID); err != nil {
		return 0, err
	}

	return userID, nil
}

func main() {
	// postgres://postgres:password@localhost:5432/ebook?sslmode=disable
	connectionStr := fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s sslmode=disable", user, password, host, port, dbname)

	db, err = sql.Open("postgres", connectionStr)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	// connection checking
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to db successfully...")

	// Email, Username, Password, Salt
	userID, err := createUser("random@gmail.com", "randomUsername", "randomPassword", "randomSalt")
	if err != nil {
		log.Println(err)
	}

	log.Printf(" User created with ID: %d", userID)

}
