package service

import (
	"ebook/app/dto"
	"ebook/app/repo"
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
		return nil, err
	}
	if err := req.Validate(); err != nil {
		return nil, err
	}
	userResp, err := s.userRepo.GetUser(req.ID)
	if err != nil {
		return nil, err
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
		return 0, err
	}
	if err := body.Validate(); err != nil {
		return 0, err
	}
	userID, err := s.userRepo.CreateUser(body)
	if err != nil {
		return 0, err
	}
	return userID, nil
}

func (s *UserServiceImpl) UpdateUser(r *http.Request) error {
	body := &dto.UserUpdateRequest{}
	if err := body.Parse(r); err != nil {
		return err
	}
	if err := body.Validate(); err != nil {
		return err
	}
	if err := s.userRepo.UpdateUser(body); err != nil {
		return err
	}
	return nil
}

func (s *UserServiceImpl) DeleteUser(r *http.Request) error {
	req := &dto.UserRequest{}
	if err := req.Parse(r); err != nil {
		return err
	}
	if err := req.Validate(); err != nil {
		return err
	}
	if err := s.userRepo.DeleteUser(req.ID); err != nil {
		return err
	}
	return nil
}

func (s *UserServiceImpl) GetAllUsers() ([]*dto.UserResponse, error) {
	result, err := s.userRepo.GetAllUsers()
	if err != nil {
		return nil, err
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
