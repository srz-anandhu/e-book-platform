package dto

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	validator "github.com/go-playground/validator/v10"
)

type BookResponse struct {
	ID       int64  `json:"id"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	AuthorID int64  `json:"author_id"`
	Status   int    `json:"status"` // 1 - Drafted, 2 - Published, 3 - Deleted
	CreateUpdateResponse
	DeleteInfoResponse
}

// To get a Book
type BookRequest struct {
	ID int `validate:"required"`
}

func (b *BookRequest) Parse(r *http.Request) error {
	strID := chi.URLParam(r, "id")
	intID, err := strconv.Atoi(strID)
	if err != nil {
		return err
	}
	b.ID = intID
	return nil
}

func (b *BookRequest) Validate() error {
	validate := validator.New()
	if err := validate.Struct(b); err != nil {
		return err
	}
	return nil
}

type BookCreateRequest struct {
	Title     string `json:"title"`
	Content   string `json:"content"`
	AuthorID  int    `json:"author_id"`
	Status    int    `json:"status"`
	CreatedBy int    `json:"created_by"`
}

func (b *BookCreateRequest) Parse(r *http.Request) error {
	if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
		return err
	}
	return nil
}

func (b *BookCreateRequest) Validate() error {
	validate := validator.New()
	if err := validate.Struct(b); err != nil {
		return err
	}
	return nil
}

type BookUpdateRequest struct {
	ID        int    `validate:"required"`
	Status    int    `validate:"required"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	UpdatedBy int    `json:"updated_by" validate:"required"` // User.ID
}

func (b *BookUpdateRequest) Parse(r *http.Request) error {
	// Get ID from request
	strID := chi.URLParam(r, "id")
	intID, err := strconv.Atoi(strID)
	if err != nil {
		return err
	}
	b.ID = intID
	// Decode to BookUpdateRequest
	if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
		return err
	}
	return nil
}

func (b *BookUpdateRequest) Validate() error {
	validate := validator.New()
	if err := validate.Struct(b); err != nil {
		return err
	}
	return nil
}

type BookDeleteRequest struct {
	ID int `validate:"required"`
	DeletedBy int `validate:"required"` // User.ID
}

func (b *BookDeleteRequest) Parse(r *http.Request) error {
	strID := chi.URLParam(r, "id")
	intID, err := strconv.Atoi(strID)
	if err != nil {
		return err
	}
	b.ID = intID

	// Decode to BookDeleteReq
	if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
		return err
	}
	return nil
}

func (b *BookDeleteRequest) Validate() error {
	validate := validator.New()
	if err := validate.Struct(b); err != nil {
		return err
	}
	return nil
}