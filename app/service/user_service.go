package service

import (
	"ebook/app/dto"
	"ebook/app/repo"
	"net/http"
)

type UserService interface {
	GetUser(r *http.Request) (*dto.UserResponse, error)
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
