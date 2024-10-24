package service

import (
	"context"

	"github.com/alvarotor/entitier-go/repository"
	"github.com/alvarotor/user-go/server/model"
)

type IUserService interface {
	repository.IGenericRepo[model.User, uint]
	GetByEmail(context.Context, string) (*model.User, error)
	GetByCode(context.Context, string) (*model.User, error)
	ValidateSvc(context.Context, string) error
	LogOutSvc(context.Context, string) error
}
