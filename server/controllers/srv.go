package controllers

import (
	"log/slog"

	"github.com/alvarotor/user-go/server/models"
	"github.com/alvarotor/user-go/server/services"
)

type controllerUser struct {
	services.IUserService
	log  *slog.Logger
	conf *models.Config
}

func NewUserController(log *slog.Logger, svc services.IUserService, conf *models.Config) IControllerUser {
	return &controllerUser{
		IUserService: svc,
		log:          log,
		conf:         conf,
	}
}
