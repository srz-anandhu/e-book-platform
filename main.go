package main

import (
	"ebook/cmd"

	_ "github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

func main() {
	cmd.Execute()
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }

	// db, sqlDb, err := gdb.ConnectDB()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// migrating models(User, Author, Book)
	// if err := gdb.AutoMigrateModels(db); err != nil {
	// 	log.Fatal(err)
	// }

	// defer func() {
	// 	sqlDb.Close()
	// 	fmt.Println("\n Db connection closed..")
	// }()
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

	// Author Creation
	// a := repo.Author{
	// 	Name: "testauthor777",
	// 	BaseModel: repo.BaseModel{
	// 		CreatedBy: 18,
	// 	},
	// }

	// authorID, err := a.CreateAuthor(db)
	// if err != nil {
	// 	log.Printf("author creation failed due to : %v", err)
	// 	return
	// }
	// log.Printf("author created with ID : %d", authorID)

	// GetOneAuthor
	// author, err := repo.GetOneAuthor(db, 10)
	// if err != nil {
	// 	log.Printf("can't get author due to : %v", err)
	// 	return
	// }
	// fmt.Printf(" ID : %d\n Name : %s\n Status : %t\n CreatedAt : %s\n CreatedBy : %d\n UpdatedAt : %s\n UpdatedBy : %d\n", author.ID, author.Name, author.Status, author.CreatedAt, author.CreatedBy, author.UpdatedAt, author.UpdatedBy)

	// err = repo.DeleteAuthor(db, 6, 18) // authorID, userID

	// if err != nil {
	// 	fmt.Printf("can't delete author due to : %v", err)
	// }

	// Update author
	// err = repo.UpdateAuthor(db, "updatedauthorname", 7, 18) // Authorname, authorID, userID
	// if err != nil {
	// 	fmt.Printf("can't update author due to : %v", err)
	// }

	// authors, err := repo.GetAllAuthors(db)
	// if err != nil {
	// 	log.Printf("can't get authors due to : %v", err)
	// 	return
	// }

	// for _, author := range authors {
	// 	fmt.Printf(" ID : %d\n Author Name : %s\n CreatedBy : %d\n CreatedAt : %s\n UpdatedAt : %s\n", author.ID, author.Name, author.CreatedBy, author.CreatedAt, author.UpdatedAt)
	// }

	// Book Creation
	// b := repo.Book{
	// 	Title:    "testbook",
	// 	Content:  "testbook content",
	// 	AuthorID: 8,
	// 	Status:   2,
	// 	BaseModel: repo.BaseModel{
	// 		CreatedBy: 18,
	// 	},
	// }

	// bookID, err := b.CreateBook(db)
	// if err != nil {
	// 	log.Printf("can't create book due to : %v", err)
	// 	return
	// }
	// fmt.Printf("book created with ID : %d", bookID)

	

	// Get One Book
	// book, err := repo.GetOneBook(db, 3)
	// if err != nil {
	// 	log.Printf("can't find book due to : %v", err)
	// 	return
	// }
	//fmt.Printf(" ID: %d\n Title: %s\n Content: %s\n AuthorID: %d\n CreatedBy: %d\n CreatedAt: %s\n UpdatedAt: %s\n UpdatedBy: %d\n DeletedAt: %v\n DeletedBy: %d", book.ID, book.Title, book.Content, book.AuthorID, book.CreatedBy, book.CreatedAt, book.UpdatedAt, book.UpdatedBy, book.DeletedAt, book.DeletedBy)

	// Update Book
	// err = repo.UpdateBook(db, 3, "updated title", "updated content", 18, 6, 1) // userID, authorID, status
	// if err != nil {
	// 	log.Printf("can't update book due to : %v", err)
	// }

}
