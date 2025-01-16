package dto

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	validator "github.com/go-playground/validator/v10"
)

type UserResponse struct {
	ID        int    `json:"id"`
	Mail      string `json:"mail"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Salt      string `json:"salt"`
	IsDeleted bool   `json:"is_deleted"`
	CreateUpdateResponse
	DeleteInfoResponse
}

// To get a user and delete a user
type UserRequest struct {
	ID int `validate:"required"`
}

func (u *UserRequest) Parse(r *http.Request) error {
	strID := chi.URLParam(r, "id")
	intID, err := strconv.Atoi(strID)
	if err != nil {
		return err
	}
	u.ID = intID
	return nil
}

func (u *UserRequest) Validate() error {
	validate := validator.New()
	if err := validate.Struct(u); err != nil {
		return err
	}
	return nil
}

type UserCreateRequest struct {
	ID       int    `json:"id"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password"`
	Mail     string `json:"mail" validate:"required"`
}

func (u *UserCreateRequest) Parse(r *http.Request) error {
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		return err
	}
	return nil
}

func (u *UserCreateRequest) Validate() error {
	validate := validator.New()
	if err := validate.Struct(u); err != nil {
		return err
	}
	return nil
}

type UserUpdateRequest struct {
	ID          int    `validate:"required"`
	NewUsername string `json:"username" validate:"required"`
	NewMail     string `json:"mail" validate:"required,mail"`
	NewPassword string `json:"password"`
}

func (u *UserUpdateRequest) Parse(r *http.Request) error {
	// Get ID from request
	strID := chi.URLParam(r, "id")
	intID, err := strconv.Atoi(strID)
	if err != nil {
		return err
	}
	u.ID = intID
	// Decode to UserUpdateRequest
	if err := json.NewDecoder(r.Body).Decode(u); err != nil {
		return err
	}
	return nil
}

func (u *UserUpdateRequest) Validate() error {
	validate := validator.New()
	if err := validate.Struct(u); err != nil {
		return err
	}
	return nil
}
