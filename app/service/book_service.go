package service

import (
	"ebook/app/dto"
	"ebook/app/repo"
	"ebook/pkg/e"
	"errors"
	"net/http"

	"github.com/rs/zerolog/log"
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
	log.Info().Msg("successfully completed validation N parsing")
	result, err := s.bookRepo.GetBook(req.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Warn().Msg("record not found")
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
	log.Info().Msgf("successfully retrieved book with ID : %d", book.ID)
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
	log.Info().Msg("successfully completed validation N parsing")
	bookID, err := s.bookRepo.CreateBook(body)
	if err != nil {
		return 0, e.NewError(e.ErrInternalServer, "can't create book", err)
	}
	log.Info().Msgf("successfully created book with ID: %d", body.ID)
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
	log.Info().Msg("successfully completed validation N parsing")
	if err := s.bookRepo.UpdateBook(body); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return e.NewError(e.ErrResourceNotFound, "book not found to update", err)
		}
		return e.NewError(e.ErrInternalServer, "can't update book", err)
	}
	log.Info().Msgf("successfully updated book with ID : %d", body.ID)
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
	log.Info().Msg("successfully completed validation N parsing")
	if err := s.bookRepo.DeleteBook(body); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return e.NewError(e.ErrResourceNotFound, "book not found to delete", err)
		}
		return e.NewError(e.ErrInternalServer, "can't delete book", err)
	}
	log.Info().Msgf("successfully deleted book with ID: %d", body.ID)
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
	log.Info().Msg("successfully retrieved all books")
	return books, nil
}
