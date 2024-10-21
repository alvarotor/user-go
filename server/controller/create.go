package controller

import (
	"context"
	"errors"
	"net/http"

	"github.com/alvarotor/user-go/server/model"
)

func (u *controllerUser) Create(c context.Context, userCreate model.User) (int, error) {
	user, err := u.svc.Create(c, userCreate)
	if err != nil {
		u.log.Error(err.Error())
		return http.StatusInternalServerError, err
	}

	if user.ID == 0 {
		return http.StatusInternalServerError, errors.New("user not created")
	}

	return http.StatusOK, nil
}
