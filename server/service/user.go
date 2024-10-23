package service

import (
	"context"

	"github.com/alvarotor/entitier-go/repository"
	"github.com/alvarotor/user-go/server/model"
	"gorm.io/gorm"
)

type IUserService interface {
	repository.IGenericRepo[model.User, uint]
	GetByEmail(context.Context, string) (*model.User, error)
	GetByCode(context.Context, string) (*model.User, error)
	ValidateSvc(context.Context, string) error
	LogOutSvc(context.Context, string) error
}

type userService struct {
	repository.IGenericRepo[model.User, uint]
}

func NewUserService(
	db *gorm.DB,
) IUserService {
	repo := repository.NewGenericRepository[model.User, uint](db)
	return &userService{
		IGenericRepo: repo,
	}
}
