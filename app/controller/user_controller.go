package controller

import (
	"ebook/app/service"
	"ebook/pkg/api"
	"ebook/pkg/e"
	"net/http"
)

type UserController interface {
	GetOne(w http.ResponseWriter, r *http.Request)
	CreateUser(w http.ResponseWriter, r *http.Request)
	UpdateUser(w http.ResponseWriter, r *http.Request)
	DeleteUser(w http.ResponseWriter, r *http.Request)
	GetAllUsers(w http.ResponseWriter, r *http.Request)
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

func (c *UserControllerImpl) CreateUser(w http.ResponseWriter, r *http.Request) {
	result, err := c.userService.CreateUser(r)
	if err != nil {
		httpErr := e.NewAPIError(err, "user creation failed")
		api.Fail(w, httpErr.StatusCode, httpErr.Code, httpErr.Message, err.Error())
		return
	}
	api.Success(w, http.StatusCreated, result)
}

func (c *UserControllerImpl) UpdateUser(w http.ResponseWriter, r *http.Request) {
	if err := c.userService.UpdateUser(r); err != nil {
		httpErr := e.NewAPIError(err, "user updation failed")
		api.Fail(w, httpErr.StatusCode, httpErr.Code, httpErr.Message, err.Error())
		return
	}
	api.Success(w, http.StatusOK, "user updated successfully")
}

func (c *UserControllerImpl) DeleteUser(w http.ResponseWriter, r *http.Request) {
	if err := c.userService.DeleteUser(r); err != nil {
		httpErr := e.NewAPIError(err, "can't delete user")
		api.Fail(w, httpErr.StatusCode, httpErr.Code, httpErr.Message, err.Error())
		return
	}
	api.Success(w, http.StatusOK, "deleted user")
}

func (c *UserControllerImpl) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	result, err := c.userService.GetAllUsers()
	if err != nil {
		httpErr := e.NewAPIError(err, "can't get all users")
		api.Fail(w, httpErr.StatusCode, httpErr.Code, httpErr.Message, err.Error())
		return
	}
	api.Success(w, http.StatusOK, result)
}
