package dto

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
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

// ToDo : validation

type UserCreateRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password"`
	Mail     string `json:"mail" validate:"required,mail"`
}

func (u *UserCreateRequest) Parse(r *http.Request) error {
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		return err
	}
	return nil
}

// Todo : validation

type UserUpdateRequest struct {
	ID       int    `validate:"required"`
	NewUsername string `json:"username" validate:"required"`
	NewMail     string `json:"mail" validate:"required,mail"`
	NewPassword string `json:"password"`
}
