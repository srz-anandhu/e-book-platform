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
	log.Info().Msg("successfully completed validation N parse")
	resp, err := s.authorRepo.GetAuthor(req.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Warn().Msg("record not found")
			return nil, e.NewError(e.ErrResourceNotFound, "no author found", err)
		}
		return nil, err
	}

	author := &dto.AuthorResponse{}

	author.ID = int(resp.ID)
	author.Name = resp.Name
	author.Status = resp.Status
	author.CreatedAt = resp.CreatedAt
	author.CreatedBy = resp.CreatedBy
	author.UpdatedAt = resp.UpdatedAt
	author.UpdatedBy = resp.UpdatedBy
	author.DeletedAt = resp.DeletedAt.Time
	author.DeletedBy = resp.DeletedBy

	log.Info().Msgf("successfully retrieved author : %v", author)
	return author, nil
}

func (s *AuthorServiceImpl) CreateAuthor(r *http.Request) (int64, error) {
	body := &dto.AuthorCreateRequest{}

	if err := body.Parse(r); err != nil {
		return 0, e.NewError(e.ErrDecodeRequestBody, "can't decode author create request", err)
	}
	if err := body.Validate(); err != nil {
		return 0, e.NewError(e.ErrValidateRequest, "can't validate author create request", err)
	}
	log.Info().Msg("successfully completed validation N parsing")

	authorID, err := s.authorRepo.CreateAuthor(body)
	if err != nil {
		return 0, e.NewError(e.ErrInternalServer, "can't create author", err)
	}
	log.Info().Msgf("successfully created author with id %d", authorID)

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
	log.Info().Msg("successfully completed validation N parsing")
	if err := s.authorRepo.UpdateAuthor(body); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return e.NewError(e.ErrResourceNotFound, "author not found to update", err)
		}
		return e.NewError(e.ErrInternalServer, "can't update author", err)
	}
	log.Info().Msgf("successfully updated author with ID : %d", body.ID)
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
	log.Info().Msg("successfully completed validation N parsing")
	if err := s.authorRepo.DeleteAuthor(body); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return e.NewError(e.ErrResourceNotFound, "author not found to delete", err)
		}
		return e.NewError(e.ErrInternalServer, "can't delete author", err)
	}
	log.Info().Msgf("successfully deleted author with ID : %d", body.ID)
	return nil
}

func (s *AuthorServiceImpl) GetAllAuthors() ([]*dto.AuthorResponse, error) {
	results, err := s.authorRepo.GetAllAuthors()
	if err != nil {
		return nil, e.NewError(e.ErrInternalServer, "can't get all authors", err)
	}
	var authors []*dto.AuthorResponse

	for _, val := range results {

		author := &dto.AuthorResponse{
			ID:     int(val.ID),
			Name:   val.Name,
			Status: val.Status,
			CreateUpdateResponse: dto.CreateUpdateResponse{
				CreatedBy: val.CreatedBy,
				CreatedAt: val.CreatedAt,
				UpdatedAt: val.UpdatedAt,
				UpdatedBy: val.UpdatedBy,
			},
			DeleteInfoResponse: dto.DeleteInfoResponse{
				DeletedAt: val.DeletedAt.Time,
				DeletedBy: val.DeletedBy,
			},
		}

		authors = append(authors, author)
	}
	log.Info().Msg("successfully retrieved all authors")
	return authors, nil
}
