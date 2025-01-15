package controller

import (
	"ebook/app/service"
	"ebook/pkg/api"
	"ebook/pkg/e"
	"net/http"
)

type BookController interface {
	GetBook(w http.ResponseWriter, r *http.Request)
	CreateBook(w http.ResponseWriter, r *http.Request)
	UpdateBook(w http.ResponseWriter, r *http.Request)
	DeleteBook(w http.ResponseWriter, r *http.Request)
	GetAllBooks(w http.ResponseWriter, r *http.Request)
}

type BookControllerImpl struct {
	bookService service.BookService
}

// For checking implementation of BookController interface
var _ BookController = (*BookControllerImpl)(nil)

func NewBookController(bookService service.BookService) BookController {
	return &BookControllerImpl{
		bookService: bookService,
	}
}

func (c *BookControllerImpl) GetBook(w http.ResponseWriter, r *http.Request) {
	result, err := c.bookService.GetBook(r)
	if err != nil {
		httpErr := e.NewAPIError(err, "can't get book")
		api.Fail(w, httpErr.StatusCode, httpErr.Code, httpErr.Message, err.Error())
		return
	}
	api.Success(w, http.StatusOK, result)
}

func (c *BookControllerImpl) CreateBook(w http.ResponseWriter, r *http.Request) {
	bookID, err := c.bookService.CreateBook(r)
	if err != nil {
		httpErr := e.NewAPIError(err, "can't create book")
		api.Fail(w, httpErr.StatusCode, httpErr.Code, httpErr.Message, err.Error())
		return
	}
	api.Success(w, http.StatusCreated, bookID)
}

func (c *BookControllerImpl) UpdateBook(w http.ResponseWriter, r *http.Request) {
	if err := c.bookService.UpdateBook(r); err != nil {
		httpErr := e.NewAPIError(err, "can't update book")
		api.Fail(w, httpErr.StatusCode, httpErr.Code, httpErr.Message, err.Error())
		return
	}
	api.Success(w, http.StatusOK, "book updated successfully")
}

func (c *BookControllerImpl) DeleteBook(w http.ResponseWriter, r *http.Request) {
	if err := c.bookService.DeleteBook(r); err != nil {
		httpErr := e.NewAPIError(err, "can't delete book")
		api.Fail(w, httpErr.StatusCode, httpErr.Code, httpErr.Message, err.Error())
		return
	}
	api.Success(w, http.StatusOK, "book deleted successfully")
}

func (c *BookControllerImpl) GetAllBooks(w http.ResponseWriter, r *http.Request) {
	results, err := c.bookService.GetAllBooks()
	if err != nil {
		httpErr := e.NewAPIError(err, "can't get all books")
		api.Fail(w, httpErr.StatusCode, httpErr.Code, httpErr.Message, err.Error())
		return
	}
	api.Success(w, http.StatusOK, results)
}
