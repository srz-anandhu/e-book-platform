package main

import (
	"log"

	gdb "ebook/app/database"

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

	// 	Username: "test444user",
	// 	Mail:     "test555@gmail.com",
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
	// user, err := repo.GetOneUser(db, 5)
	// if err != nil {
	// 	log.Printf("cant get user due to : %v", err)
	// 	return
	// }
	//fmt.Printf(" ID: %d\n Username: %s\n Email: %s\n Password: %s\n CreatedAt: %s\n UpdatedAt: %s", user.ID, user.Username, user.Mail, user.Password, user.CreatedAt, user.UpdatedAt)

	// User Deletion (Soft delete)
	// if err := repo.DeleteUser(db, 18); err != nil {
	// 	log.Printf("cant delete user due to : %v", err)
	// }

	// User Updtion
	// if err := repo.UpdateUser(db, 5, "updatedPassword"); err != nil {

	// 	log.Printf("updation failed due to : %v", err)
	// }

	// Get all users
	// users, err := repo.GetAllUsers(db)
	// if err != nil {
	// 	log.Printf("can't get users due to : %v", err)
	// }

	// for _, user := range users {
	// 	fmt.Printf("ID: %d\n UserName: %s\n Email: %s\n CreatedAt: %s\n UpdatedAt: %s\n", user.ID, user.Username, user.Mail, user.CreatedAt, user.UpdatedAt)
	// }

}
