package service

import (
	"ebook/app/dto"
	"ebook/app/repo"
	"ebook/pkg/e"
	"net/http"
)

type AuthorService interface {
	GetAuthor(r *http.Request) (*dto.AuthorResponse, error)
	CreateAuthor(r *http.Request) (int64, error)
	UpdateAuthor(r *http.Request) error
	DeleteAuthor(r *http.Request) error
	GetAllAuthors() ([]*dto.AuthorResponse, error)
}

type AuthorServiceImpl struct {
	authorRepo repo.AuthorRepo
}

// For checking implementation of Service interface
var _ AuthorService = (*AuthorServiceImpl)(nil)

func NewAuthorService(authorRepo repo.AuthorRepo) AuthorService {
	return &AuthorServiceImpl{
		authorRepo: authorRepo,
	}
}

func (s *AuthorServiceImpl) GetAuthor(r *http.Request) (*dto.AuthorResponse, error) {
	req := &dto.AuthorRequest{}
	if err := req.Parse(r); err != nil {
		return nil, err
	}
	if err := req.Validate(); err != nil {
		return nil, err
	}
	resp, err := s.authorRepo.GetAuthor(req.ID)
	if err != nil {
		return nil, err
	}
	var author dto.AuthorResponse

	author.ID = resp.ID
	author.Name = resp.Name
	author.Status = resp.Status
	author.CreatedAt = resp.CreatedAt
	author.CreatedBy = resp.CreatedBy
	author.UpdatedAt = resp.UpdatedAt
	author.UpdatedBy = resp.UpdatedBy
	author.DeletedAt = resp.DeletedAt
	author.DeletedBy = resp.DeletedBy

	return &author, nil
}

func (s *AuthorServiceImpl) CreateAuthor(r *http.Request) (int64, error) {
	body := &dto.AuthorCreateRequest{}

	if err := body.Parse(r); err != nil {
		return 0, e.NewError(e.ErrDecodeRequestBody, "can't decode author create request", err)
	}
	if err := body.Validate(); err != nil {
		return 0, e.NewError(e.ErrValidateRequest, "can't validate author create request", err)
	}
	authorID, err := s.authorRepo.CreateAuthor(body)
	if err != nil {
		return 0, e.NewError(e.ErrInternalServer, "can't create author", err)
	}
	return authorID, nil
}

func (s *AuthorServiceImpl) UpdateAuthor(r *http.Request) error {
	body := &dto.AuthorUpdateRequest{}

	if err := body.Parse(r); err != nil {
		return e.NewError(e.ErrDecodeRequestBody, "can't decode author update request", err)
	}
	if err := body.Validate(); err != nil {
		return e.NewError(e.ErrValidateRequest, "can't validate author update request", err)
	}
	if err := s.authorRepo.UpdateAuthor(body); err != nil {
		return e.NewError(e.ErrInternalServer, "can't update author", err)
	}
	return nil
}

func (s *AuthorServiceImpl) DeleteAuthor(r *http.Request) error {
	body := &dto.AuthorDeleteReq{}

	if err := body.Parse(r); err != nil {
		return e.NewError(e.ErrInvalidRequest, "author delete request parse error", err)
	}
	if err := body.Validate(); err != nil {
		return e.NewError(e.ErrValidateRequest, "author delete request validate error", err)
	}
	if err := s.authorRepo.DeleteAuthor(body); err != nil {
		return e.NewError(e.ErrInternalServer, "can't delete author", err)
	}
	return nil
}

func (s *AuthorServiceImpl) GetAllAuthors() ([]*dto.AuthorResponse, error) {
	results, err := s.authorRepo.GetAllAuthors()
	if err != nil {
		return nil, e.NewError(e.ErrInternalServer, "can't get all authors", err)
	}
	var authors []*dto.AuthorResponse

	for _, val := range results {

		var author dto.AuthorResponse

		author.ID = val.ID
		author.Name = val.Name
		author.Status = val.Status
		author.CreatedBy = val.CreatedBy
		author.CreatedAt = val.CreatedAt
		author.UpdatedBy = val.UpdatedBy
		author.UpdatedAt = val.UpdatedAt
		author.DeletedBy = val.DeletedBy
		author.DeletedAt = val.DeletedAt

		authors = append(authors, &author)
	}

	return authors, nil
}
