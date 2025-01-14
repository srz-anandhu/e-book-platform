package dto

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	validator "github.com/go-playground/validator/v10"
)

type AuthorResponse struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Status bool   `json:"status"`
	CreateUpdateResponse
	DeleteInfoResponse
}

// To get and delete an Author
type AuthorRequest struct {
	ID int `validate:"required"`
}

func (a *AuthorRequest) Parse(r *http.Request) error {
	strID := chi.URLParam(r, "id")
	intID, err := strconv.Atoi(strID)
	if err != nil {
		return err
	}
	a.ID = intID
	return nil
}

func (a *AuthorRequest) Validate() error {
	validate := validator.New()
	if err := validate.Struct(a); err != nil {
		return err
	}
	return nil
}

type AuthorCreateRequest struct {
	Name string `json:"name"`
}

func (a *AuthorCreateRequest) Parse(r *http.Request) error {
	if err := json.NewDecoder(r.Body).Decode(&a); err != nil {
		return err
	}
	return nil
}

func (a *AuthorCreateRequest) Validate() error {
	validate := validator.New()
	if err := validate.Struct(a); err != nil {
		return err
	}
	return nil
}

type AuthorUpdateRequest struct {
	ID        int    `validate:"required"`
	NewName   string `validate:"required"`
	UpdatedBy int    `validate:"required"`
}

func (a *AuthorUpdateRequest) Parse(r *http.Request) error {
	// Get ID from request
	strID := chi.URLParam(r, "id")
	intID, err := strconv.Atoi(strID)
	if err != nil {
		return err
	}
	a.ID = intID
	// Decode to AuthorUpdateRequest
	if err := json.NewDecoder(r.Body).Decode(&a); err != nil {
		return err
	}
	return nil
}

func (a *AuthorUpdateRequest) Validate() error {
	validate := validator.New()
	if err := validate.Struct(a); err != nil {
		return err
	}
	return nil
}

type AuthorDeleteReq struct {
	ID        int `validate:"required"`
	DeletedBy int `validate:"required"` // userID
}

func (a *AuthorDeleteReq) Parse(r *http.Request) error {
	strID := chi.URLParam(r, "id")
	intID, err := strconv.Atoi(strID)
	if err != nil {
		return err
	}
	a.ID = intID

	// Decode to AuthorDeleteReq
	if err := json.NewDecoder(r.Body).Decode(&a); err != nil {
		return err
	}
	return nil
}

func (a *AuthorDeleteReq) Validate() error {
	validate := validator.New()
	if err := validate.Struct(a); err != nil {
		return err
	}
	return nil
}
