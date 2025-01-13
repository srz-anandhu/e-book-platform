package controller

import (
	"ebook/app/service"
	"ebook/pkg/api"
	"ebook/pkg/e"
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
		httpErr := e.NewAPIError(err, "can't get user")
		api.Fail(w, httpErr.StatusCode, httpErr.Code, httpErr.Message, err.Error())
		return
	}
	api.Success(w, http.StatusOK, userResp)
}
