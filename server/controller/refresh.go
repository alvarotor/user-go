package controller

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/alvarotor/user-go/server/models"
)

func (u *controllerUser) Refresh(ctx context.Context, refreshToken string) (int, models.Token, error) {
	user, err := u.GetByCodeRefresh(ctx, refreshToken)
	if err != nil {
		return http.StatusNotFound, models.Token{}, err
	}

	if user == nil {
		errMsg := "code is invalid"
		u.log.Error(errMsg)
		return http.StatusBadRequest, models.Token{}, errors.New(errMsg)
	}

	if user.Code == "OUT" || strings.TrimSpace(user.Code) == "" {
		errMsg := "code is invalid"
		u.log.Error(errMsg)
		return http.StatusBadRequest, models.Token{}, errors.New(errMsg)
	}

	status, token, err := u.Validate(ctx, user.Code)
	if err != nil {
		return http.StatusBadRequest, models.Token{}, err
	}

	user.CodeRefresh = u.GenerateRandomString(u.conf.SizeRandomStringValidationRefresh)
	err = u.UpdateField(ctx, user.ID, "code_refresh", user.CodeRefresh)
	if err != nil {
		u.log.Error(err.Error())
		return http.StatusInternalServerError, models.Token{}, err
	}

	return status, token, nil
}
