package service

import (
	"ebook/app/dto"
	"ebook/app/repo"
	"net/http"
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
		return nil, err
	}
	if err := req.Validate(); err != nil {
		return nil, err
	}
	result, err := s.bookRepo.GetBook(req.ID)
	if err != nil {
		return nil, err
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
		return 0, err
	}
	if err := body.Validate(); err != nil {
		return 0, err
	}
	bookID, err := s.bookRepo.CreateBook(body)
	if err != nil {
		return 0, err
	}

	return bookID, nil
}

func (s *BookServiceImpl) UpdateBook(r *http.Request) error {
	body := &dto.BookUpdateRequest{}
	if err := body.Parse(r); err != nil {
		return err
	}
	if err := body.Validate(); err != nil {
		return err
	}
	if err := s.bookRepo.UpdateBook(body); err != nil {
		return err
	}

	return nil
}

func (s *BookServiceImpl) DeleteBook(r *http.Request) error {
	body := &dto.BookDeleteRequest{}
	if err := body.Parse(r); err != nil {
		return err
	}
	if err := body.Validate(); err != nil {
		return err
	}
	if err := s.bookRepo.DeleteBook(body); err != nil {
		return err
	}

	return nil
}

func (s *BookServiceImpl) GetAllBooks() ([]*dto.BookResponse, error) {
	results, err := s.bookRepo.GetAllBooks()
	if err != nil {
		return nil, err
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
