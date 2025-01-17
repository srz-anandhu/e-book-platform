package service

import (
	"ebook/app/dto"
	"ebook/app/repo"
	"ebook/pkg/e"
	"errors"
	"net/http"

	"gorm.io/gorm"
)

type BookService interface {
	GetBook(r *http.Request) (*dto.BookResponse, error)
	CreateBook(r *http.Request) (int64, error)
	UpdateBook(r *http.Request) error
	DeleteBook(r *http.Request) error
	GetAllBooks() ([]*dto.BookResponse, error)
}

type BookServiceImpl struct {
	bookRepo repo.BookRepo
}

var _ BookService = (*BookServiceImpl)(nil)

func NewBookService(bookRepo repo.BookRepo) BookService {
	return &BookServiceImpl{
		bookRepo: bookRepo,
	}
}

func (s *BookServiceImpl) GetBook(r *http.Request) (*dto.BookResponse, error) {
	req := &dto.BookRequest{}
	if err := req.Parse(r); err != nil {
		return nil, e.NewError(e.ErrInvalidRequest, "can't parse book request", err)
	}
	if err := req.Validate(); err != nil {
		return nil, e.NewError(e.ErrValidateRequest, "can't validate book request", err)
	}
	result, err := s.bookRepo.GetBook(req.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, e.NewError(e.ErrResourceNotFound, "book not exist", err)
		}
		return nil, e.NewError(e.ErrInternalServer, "can't get book", err)
	}

	var book dto.BookResponse
	book.ID = result.ID
	book.Title = result.Title
	book.Content = result.Content
	book.AuthorID = result.AuthorID
	book.Status = result.Status
	book.CreatedBy = result.CreatedBy
	book.CreatedAt = result.CreatedAt
	book.UpdatedBy = result.UpdatedBy
	book.UpdatedAt = result.UpdatedAt
	book.DeletedBy = result.DeletedBy
	book.DeletedBy = result.DeletedBy

	return &book, nil
}

func (s *BookServiceImpl) CreateBook(r *http.Request) (int64, error) {
	body := &dto.BookCreateRequest{}
	if err := body.Parse(r); err != nil {
		return 0, e.NewError(e.ErrDecodeRequestBody, "can't decode book create request", err)
	}
	if err := body.Validate(); err != nil {
		return 0, e.NewError(e.ErrValidateRequest, "can't validate book request", err)
	}
	bookID, err := s.bookRepo.CreateBook(body)
	if err != nil {
		return 0, e.NewError(e.ErrInternalServer, "can't create book", err)
	}

	return bookID, nil
}

func (s *BookServiceImpl) UpdateBook(r *http.Request) error {
	body := &dto.BookUpdateRequest{}
	if err := body.Parse(r); err != nil {
		return e.NewError(e.ErrDecodeRequestBody, "can't decode book update request", err)
	}
	if err := body.Validate(); err != nil {
		return e.NewError(e.ErrValidateRequest, "can't validate book update request", err)
	}
	if err := s.bookRepo.UpdateBook(body); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return e.NewError(e.ErrResourceNotFound, "book not found to update", err)
		}
		return e.NewError(e.ErrInternalServer, "can't update book", err)
	}

	return nil
}

func (s *BookServiceImpl) DeleteBook(r *http.Request) error {
	body := &dto.BookDeleteRequest{}
	if err := body.Parse(r); err != nil {
		return e.NewError(e.ErrInvalidRequest, "book delete request parse error", err)
	}
	if err := body.Validate(); err != nil {
		return e.NewError(e.ErrInvalidRequest, "book delete request validate error", err)
	}
	if err := s.bookRepo.DeleteBook(body); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return e.NewError(e.ErrResourceNotFound, "book not found to delete", err)
		}
		return e.NewError(e.ErrInternalServer, "can't delete book", err)
	}

	return nil
}

func (s *BookServiceImpl) GetAllBooks() ([]*dto.BookResponse, error) {
	results, err := s.bookRepo.GetAllBooks()
	if err != nil {
		return nil, e.NewError(e.ErrInternalServer, "can't get all books", err)
	}

	var books []*dto.BookResponse

	for _, val := range results {

		var book dto.BookResponse

		book.ID = val.ID
		book.Title = val.Title
		book.Content = val.Content
		book.AuthorID = val.AuthorID
		book.Status = val.Status
		book.CreatedBy = val.CreatedBy
		book.CreatedAt = val.CreatedAt
		book.UpdatedBy = val.UpdatedBy
		book.UpdatedAt = val.UpdatedAt
		book.DeletedBy = val.DeletedBy
		book.DeletedBy = val.DeletedBy

		books = append(books, &book)
	}

	return books, nil
}
