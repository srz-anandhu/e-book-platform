package controller

import (
	"ebook/app/service"
	"ebook/pkg/api"
	"ebook/pkg/e"
	"net/http"
)

type AuthorController interface {
	GetAuthor(w http.ResponseWriter, r *http.Request)
	CreateAuthor(w http.ResponseWriter, r *http.Request)
	UpdateAuthor(w http.ResponseWriter, r *http.Request)
	DeleteAuthor(w http.ResponseWriter, r *http.Request)
	GetAllAuthors(w http.ResponseWriter, r *http.Request)
}

type AuthorControllerImpl struct {
	authorService service.AuthorService
}

// To check implementation of AuthorController interface
var _ AuthorController = (*AuthorControllerImpl)(nil)

// Constructor
func NewAuthorController(authorService service.AuthorService) AuthorController {
	return &AuthorControllerImpl{
		authorService: authorService,
	}
}

func (c *AuthorControllerImpl) GetAuthor(w http.ResponseWriter, r *http.Request) {
	result, err := c.authorService.GetAuthor(r)
	if err != nil {
		httpErr := e.NewAPIError(err, "can't get author")
		api.Fail(w, httpErr.StatusCode, httpErr.Code, httpErr.Message, err.Error())
		return
	}
	api.Success(w, http.StatusOK, result)
}

func (c *AuthorControllerImpl) CreateAuthor(w http.ResponseWriter, r *http.Request) {
	result, err := c.authorService.CreateAuthor(r)
	if err != nil {
		httpErr := e.NewAPIError(err, "can't create author")
		api.Fail(w, httpErr.StatusCode, httpErr.Code, httpErr.Message, err.Error())
		return
	}
	api.Success(w, http.StatusCreated, result)
}

func (c *AuthorControllerImpl) UpdateAuthor(w http.ResponseWriter, r *http.Request) {
	if err := c.authorService.UpdateAuthor(r); err != nil {
		httpErr := e.NewAPIError(err, "can't update author")
		api.Fail(w, httpErr.StatusCode, httpErr.Code, httpErr.Message, err.Error())
		return
	}
	api.Success(w, http.StatusOK, "author updated successfully")
}

func (c *AuthorControllerImpl) DeleteAuthor(w http.ResponseWriter, r *http.Request) {
	if err := c.authorService.DeleteAuthor(r); err != nil {
		httpErr := e.NewAPIError(err, "can't delete author")
		api.Fail(w, httpErr.StatusCode, httpErr.Code, httpErr.Message, err.Error())
		return
	}
	api.Success(w, http.StatusOK, "deleted author successfully")
}

func (c *AuthorControllerImpl) GetAllAuthors(w http.ResponseWriter, r *http.Request) {
	result, err := c.authorService.GetAllAuthors()
	if err != nil {
		httpErr := e.NewAPIError(err, "can't get all authors")
		api.Fail(w, httpErr.StatusCode, httpErr.Code, httpErr.Message, err.Error())
		return
	}
	api.Success(w, http.StatusOK, result)
}
