package controller

import (
	"context"

	"github.com/alvarotor/user-go/server/dto"
	"github.com/alvarotor/user-go/server/model"
	"github.com/alvarotor/user-go/server/service"
)

type IControllerUser interface {
	service.IUserService
	Login(context.Context, dto.UserLogin) (int, string, error)
	LogOut(context.Context, string) (int, error)
	Validate(context.Context, string) (int, model.Token, error)
	TokenToUser(context.Context, string) (*model.User, error)
	Health(context.Context, uint32) int
	UpdateUserAdminStatus(context.Context, string, bool) error
}
