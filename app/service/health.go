package service

import "ebook/app/repo"

type HealthService interface {
	CheckHealth() (string, string, error)
}

type HealthServiceImpl struct {
	healthRepo repo.HealthRepo
}

func NewHealthService(healthRepo repo.HealthRepo) HealthService {
	return &HealthServiceImpl{
		healthRepo: healthRepo,
	}
}

func (s *HealthServiceImpl) CheckHealth() (string, string, error) {
	if err := s.healthRepo.CheckPing(); err != nil {
		return "notok", "DB connection issue", err
	}
	return "ok", "e-book is running smoothly", nil
}