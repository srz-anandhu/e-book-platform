package controller

import (
	"ebook/app/service"
	"ebook/pkg/api"
	"net/http"
)

type HealthController interface {
	CheckHealth(w http.ResponseWriter, r *http.Request)
}

type HealthControllerImpl struct {
	healthService service.HealthService
}

func NewHealthController(healthService service.HealthService) HealthController {
	return &HealthControllerImpl{
		healthService: healthService,
	}
}

func (c *HealthControllerImpl) CheckHealth(w http.ResponseWriter, r *http.Request) {
	status, result, err := c.healthService.CheckHealth()
	if err != nil {
		api.Fail(w, 00, 00, "bad health", status, result)
	}
	api.Success(w, http.StatusOK, result)
}