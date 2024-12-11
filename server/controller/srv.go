package controller

import (
	"log/slog"

	"github.com/alvarotor/user-go/server/models"
	"github.com/alvarotor/user-go/server/service"
)

type controllerUser struct {
	service.IUserService
	log  *slog.Logger
	conf *models.Config
}

func NewUserController(log *slog.Logger, svc service.IUserService, conf *models.Config) IControllerUser {
	return &controllerUser{
		IUserService: svc,
		log:          log,
		conf:         conf,
	}
}
