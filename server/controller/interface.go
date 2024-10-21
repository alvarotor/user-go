package controller

import (
	"context"

	"github.com/alvarotor/user-go/server/dto"
	"github.com/alvarotor/user-go/server/model"
)

type IControllerUser interface {
	Create(context.Context, model.User) (int, error)
	Login(context.Context, dto.UserLogin) (int, error)
	LogOut(context.Context, string) (int, error)
	Get(context.Context, uint) (int, dto.UserResponse, error)
	Update(context.Context, uint, dto.UserRequest) (int, error)
	Delete(context.Context, uint, bool) (int, error)
	Validate(context.Context, string, string, string) (int, model.Token, error)
}
