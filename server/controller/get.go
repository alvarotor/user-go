package controller

import (
	"context"
	"errors"
	"net/http"

	"github.com/alvarotor/user-go/server/dto"
	"github.com/alvarotor/user-go/server/model"
)

func (u *controllerUser) Get(c context.Context, id uint) (int, dto.UserResponse, error) {
	user, err := u.svc.Get(c, id, "")
	if errors.Is(err, model.ErrUserNotFound) {
		u.log.Info(model.ErrUserNotFound.Error())
		return http.StatusNotFound, dto.UserResponse{}, err
	}
	if err != nil {
		u.log.Error(err.Error())
		return http.StatusInternalServerError, dto.UserResponse{}, err
	}

	userRes := dto.UserResponse{
		Name:          user.Name,
		Email:         user.Email,
		ProfilePic:    user.ProfilePic,
		Validated:     user.Validated,
		Age:           uint(user.Age),
		Gender:        uint(user.Gender),
		CountryOrigin: user.CountryOrigin,
	}

	return http.StatusOK, userRes, nil
}
