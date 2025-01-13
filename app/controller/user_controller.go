package controller

import (
	"ebook/app/service"
	"log"
	"net/http"
)

type UserController interface {
	GetOne(w http.ResponseWriter, r *http.Request)
}

type UserControllerImpl struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &UserControllerImpl{
		userService: userService,
	}
}

var _ UserController = (*UserControllerImpl)(nil)

func (c *UserControllerImpl) GetOne(w http.ResponseWriter, r *http.Request) {
	userResp, err := c.userService.GetUser(r)
	if err != nil {
		log.Printf("cant get user due to : %v", err)
		return
	}
	log.Println(userResp)
}