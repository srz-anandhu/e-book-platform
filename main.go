package main

import (
	"log"

	gdb "ebook/pkg/database"
	"ebook/pkg/repo"

	_ "github.com/lib/pq"
)

func main() {

	db, err := gdb.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	// migrating models(User)
	if err := gdb.AutoMigrateModels(db); err != nil {
		log.Fatal(err)
	}
	// User creation
	// u := repo.User{

	// 	Username: "test5user",
	// 	Mail:     "test5@gmail.com",
	// 	Password: "testpassword",
	// 	Salt:     "testsalt",
	// }

	// userID, err := u.CreateUser(db)
	// if err != nil {
	// 	log.Printf("user creation failed due to : %v", err)
	// 	return
	// }
	// log.Printf("user created with ID: %d", userID)

	// Get One user
	// user, err := repo.GetOneUser(db, 13)
	// if err != nil {
	// 	log.Printf("cant get user due to : %v", err)
	// 	return
	// }
	//fmt.Printf(" ID: %d\n Username: %s\n Email: %s\n Password: %s\n CreatedAt: %s\n UpdatedAt: %s", user.ID, user.Username, user.Mail, user.Password, user.CreatedAt, user.UpdatedAt)

	// User Deletion (Soft delete)
	if err := repo.DeleteUser(db, 18); err != nil {
		log.Printf("cant delete user due to : %v", err)
	}

}
