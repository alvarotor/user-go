package controller

import (
	"log/slog"

	"github.com/alvarotor/user-go/server/model"
	"github.com/alvarotor/user-go/server/service"
)

type controllerUser struct {
	service.IUserService
	log  *slog.Logger
	conf *model.Config
}

func NewUserController(log *slog.Logger, svc service.IUserService, conf *model.Config) IControllerUser {
	return &controllerUser{
		IUserService: svc,
		log:          log,
		conf:         conf,
	}
}
