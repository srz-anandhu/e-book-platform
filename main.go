package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

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

// User creation
func createUser(mail, username, password, salt string) (userID int, err error) {
	query := ` INSERT INTO users(mail,username,password,salt) VALUES($1,$2,$3,$4) RETURNING id `

	if err := db.QueryRow(query, mail, username, password, salt).Scan(&userID); err != nil {
		return 0, fmt.Errorf("user creation failed due to : %v", err)
	}

	return userID, nil
}

// Get one user
func getOneUser(id int) (username, mail string, createdAt, updatedAt time.Time, err error) {
	query := ` SELECT username,mail,created_at,updated_at FROM users WHERE id=$1`

	if err := db.QueryRow(query, id).Scan(&username, &mail, &createdAt, &updatedAt); err != nil {
		return "", "", time.Time{}, time.Time{}, err
	}

	return username, mail, createdAt, updatedAt, nil
}

// Author creation
func createAuthor(name string, createdBy int) (authorID int, err error) {
	query := ` INSERT INTO authors(name,created_by) VALUES($1,$2) RETURNING id`

	if err := db.QueryRow(query, name, createdBy).Scan(&authorID); err != nil {
		return 0, fmt.Errorf("author creation failed due to : %v", err)
	}

	return authorID, nil
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
	userID, err := createUser("random2@gmail.com", "random2Username", "random2Password", "random2Salt")
	if err != nil {
		log.Println(err)
		return
	}

	log.Printf(" User created with ID: %d", userID)

	// Author name, createdBy (user ID)
	authorID, err := createAuthor("random2author name", 5)
	if err != nil {
		log.Println(err)
		return
	}

	log.Printf("Author created with ID : %d ", authorID)

	// Get one user
	userName, mail, createdAt, updatedAt, err := getOneUser(99) // UserID
	if err != nil {
		log.Printf("Get one user failed due to : %v", err)
		return
	}

	fmt.Printf("Username: %s\n Mail: %s\n CreatedAt: %s\n, UpdatedAt:%s", userName, mail, createdAt, updatedAt)

}
