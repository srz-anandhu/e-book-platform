package app

import (
	"ebook/app/controller"
	"ebook/app/repo"
	"ebook/app/service"

	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

func APIRouter(db *gorm.DB) chi.Router {

	r := chi.NewRouter()

	// r.Route("/", func(r chi.Router) {
	// 	r.Get("/hello", api.DemoHandler)
	// })

	// User
	userRepo := repo.NewUserRepo(db)
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)

	// Author
	authorRepo := repo.NewAuthorRepo(db)
	authorService := service.NewAuthorService(authorRepo)
	authorController := controller.NewAuthorController(authorService)

	// Book
	bookRepo := repo.NewBookRepo(db)
	bookService := service.NewBookService(bookRepo)
	bookController := controller.NewBookController(bookService)

	r.Route("/users", func(r chi.Router) {
		r.Get("/", userController.GetAllUsers)
		r.Post("/create", userController.CreateUser)
		r.Put("/{id}", userController.UpdateUser)
		r.Delete("/{id}", userController.DeleteUser)
		r.Get("/{id}", userController.GetOne)
	})

	r.Route("/authors", func(r chi.Router) {
		r.Get("/", authorController.GetAllAuthors)
		r.Post("/create", authorController.CreateAuthor)
		r.Put("/{id}", authorController.UpdateAuthor)
		r.Delete("/{id}", authorController.DeleteAuthor)
		r.Get("/{id}", authorController.GetAuthor)
	})

	r.Route("/books", func(r chi.Router) {
		r.Get("/", bookController.GetAllBooks)
		r.Post("/create", bookController.CreateBook)
		r.Put("/{id}", bookController.UpdateBook)
		r.Delete("/{id}", bookController.DeleteBook)
		r.Get("/{id}", bookController.GetBook)
	})

	return r
}
