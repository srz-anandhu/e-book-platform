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

	userRepo := repo.NewUserRepo(db)
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)

	r.Route("/users", func(r chi.Router) {
		r.Get("/", userController.GetAllUsers)
		r.Post("/create", userController.CreateUser)
	})

	return r
}
