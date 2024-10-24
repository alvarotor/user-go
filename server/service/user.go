package service

import (
	"github.com/alvarotor/entitier-go/repository"
	"github.com/alvarotor/user-go/server/model"
	"gorm.io/gorm"
)

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
