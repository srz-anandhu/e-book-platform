package main

import (
	"log"

	gdb "ebook/pkg/database"
	"ebook/pkg/repo"

	_ "github.com/lib/pq"
)

// User creation
// func createUser(mail, username, password, salt string) (userID int, err error) {
// 	query := ` INSERT INTO users(mail,username,password,salt) VALUES($1,$2,$3,$4) RETURNING id `

// 	if err := db.QueryRow(query, mail, username, password, salt).Scan(&userID); err != nil {
// 		return 0, fmt.Errorf("user creation failed due to : %v", err)
// 	}

// 	return userID, nil
// }

// Get one user
// func getOneUser(id int) (username, mail string, createdAt, updatedAt time.Time, err error) {
// 	query := ` SELECT username,mail,created_at,updated_at FROM users WHERE id=$1`

// 	if err := db.QueryRow(query, id).Scan(&username, &mail, &createdAt, &updatedAt); err != nil {
// 		return "", "", time.Time{}, time.Time{}, fmt.Errorf("get one user failed due to : %v", err)
// 	}

// 	return username, mail, createdAt, updatedAt, nil
// }

// Delete user (Soft Delete)
// func deleteUser(id int) (err error) {
// 	query := ` UPDATE users SET is_deleted=$1,deleted_at=$2 WHERE id=$3`

// 	_, err = db.Exec(query, true, time.Now().UTC(), id)
// 	if err != nil {
// 		return fmt.Errorf("user deletion failed due to : %v", err)
// 	}
// 	log.Println("deleted user successfully")
// 	return nil
// }

// Update user
// func updateUser(id int, mail, password string) (err error) {
// 	query := ` UPDATE users SET mail=$1,password=$2,updated_at=$3 WHERE id=$4`

// 	result, err := db.Exec(query, mail, password, time.Now().UTC(), id)
// 	if err != nil {
// 		return fmt.Errorf("user updation failed due to : %v", err)
// 	}

// 	isAffected, err := result.RowsAffected()
// 	if err != nil {
// 		return fmt.Errorf("no affected rows due to : %v", err)
// 	}
// 	if isAffected == 0 {
// 		return fmt.Errorf("no user with ID: %d", id)
// 	}
// 	log.Println("user updated successfully")
// 	return nil
// }

// Author creation
// func createAuthor(name string, createdBy int) (authorID int, err error) {
// 	query := ` INSERT INTO authors(name,created_by) VALUES($1,$2) RETURNING id`

// 	if err := db.QueryRow(query, name, createdBy).Scan(&authorID); err != nil {
// 		return 0, fmt.Errorf("author creation failed due to : %v", err)
// 	}

// 	return authorID, nil
// }

func main() {

	db, err := gdb.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	u := repo.User{

		Username: "test4user",
		Mail:     "test4@gmail.com",
		Password: "testpassword",
		Salt:     "testsalt",
	}
	if err := db.AutoMigrate(&u); err != nil {
		log.Printf("user migration failed : %v", err)
	}
	userID, err := u.CreateUser(db)
	if err != nil {
		log.Printf("user creation failed due to : %v", err)
		return
	}
	log.Printf("user created with ID: %d", userID)

	// Email, Username, Password, Salt
	// userID, err := createUser("random2@gmail.com", "random2Username", "random2Password", "random2Salt")
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }

	// log.Printf(" User created with ID: %d", userID)

	// Author name, createdBy (user ID)
	// authorID, err := createAuthor("random2author name", 5)
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }

	// log.Printf("Author created with ID : %d ", authorID)

	// Get one user
	// userName, mail, createdAt, updatedAt, err := getOneUser(99) // UserID
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }

	// fmt.Printf("Username: %s\n Mail: %s\n CreatedAt: %s\n, UpdatedAt:%s", userName, mail, createdAt, updatedAt)

	// Delete user
	// if err := deleteUser(1); err != nil {
	// 	log.Println(err)
	// }

	// Update User
	// userID, Email, Password
	// if err := updateUser(99, "updatedmail@gmail.com", "updatedPassword"); err != nil {
	// 	log.Println(err)
	// }

}
