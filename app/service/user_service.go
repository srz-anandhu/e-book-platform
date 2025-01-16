package service

import (
	"ebook/app/dto"
	"ebook/app/repo"
	"ebook/pkg/e"
	"net/http"
)

type UserService interface {
	GetUser(r *http.Request) (*dto.UserResponse, error)
	CreateUser(r *http.Request) (int64, error)
	UpdateUser(r *http.Request) error
	DeleteUser(r *http.Request) error
	GetAllUsers() ([]*dto.UserResponse, error)
}

type UserServiceImpl struct {
	userRepo repo.UserRepo
}

func NewUserService(userRepo repo.UserRepo) UserService {
	return &UserServiceImpl{
		userRepo: userRepo,
	}
}

var _ UserService = (*UserServiceImpl)(nil)

func (s *UserServiceImpl) GetUser(r *http.Request) (*dto.UserResponse, error) {
	req := &dto.UserRequest{}
	if err := req.Parse(r); err != nil {
		return nil, e.NewError(e.ErrInvalidRequest, "user id parse error", err)
	}
	if err := req.Validate(); err != nil {
		return nil, e.NewError(e.ErrValidateRequest, "user request validate error", err)
	}
	userResp, err := s.userRepo.GetUser(req.ID)
	if err != nil {
		return nil, e.NewError(e.ErrResourceNotFound, "can't get user", err)
	}
	var user dto.UserResponse

	user.ID = userResp.ID
	user.Username = userResp.Username
	user.Mail = userResp.Mail
	user.Password = userResp.Password
	user.CreatedAt = userResp.CreatedAt
	user.UpdatedAt = userResp.UpdatedAt
	user.DeletedAt = userResp.DeletedAt

	return &user, nil
}

func (s *UserServiceImpl) CreateUser(r *http.Request) (int64, error) {
	body := &dto.UserCreateRequest{}
	if err := body.Parse(r); err != nil {
		return 0, e.NewError(e.ErrDecodeRequestBody, "can't decode user create request", err)
	}
	if err := body.Validate(); err != nil {
		return 0, e.NewError(e.ErrValidateRequest, "can't validate user create request", err)
	}
	userID, err := s.userRepo.CreateUser(body)
	if err != nil {
		return 0, e.NewError(e.ErrInternalServer, "can't create user", err)
	}
	return userID, nil
}

func (s *UserServiceImpl) UpdateUser(r *http.Request) error {
	body := &dto.UserUpdateRequest{}
	if err := body.Parse(r); err != nil {
		return e.NewError(e.ErrDecodeRequestBody, "can't decode user update request", err)
	}
	if err := body.Validate(); err != nil {
		return e.NewError(e.ErrValidateRequest, "can't validate user create request", err)
	}
	if err := s.userRepo.UpdateUser(body); err != nil {
		return e.NewError(e.ErrInternalServer, "can't update user", err)
	}
	return nil
}

func (s *UserServiceImpl) DeleteUser(r *http.Request) error {
	req := &dto.UserRequest{}
	if err := req.Parse(r); err != nil {
		return e.NewError(e.ErrInvalidRequest, "user request parese error", err)
	}
	if err := req.Validate(); err != nil {
		return e.NewError(e.ErrValidateRequest, "user request validate error", err)
	}
	if err := s.userRepo.DeleteUser(req.ID); err != nil {
		return e.NewError(e.ErrDecodeRequestBody, "can't delete user", err)
	}
	return nil
}

func (s *UserServiceImpl) GetAllUsers() ([]*dto.UserResponse, error) {
	result, err := s.userRepo.GetAllUsers()
	if err != nil {
		return nil, e.NewError(e.ErrInternalServer, "can't get all users", err)
	}

	var users []*dto.UserResponse

	for _, val := range result {

		var user dto.UserResponse

		user.ID = val.ID
		user.Username = val.Username
		user.Password = val.Password
		user.Mail = val.Mail
		user.CreatedAt = val.CreatedAt
		user.UpdatedAt = val.UpdatedAt
		user.IsDeleted = val.IsDeleted
		user.DeletedAt = val.DeletedAt

		users = append(users, &user)
	}
	return users, nil
}
