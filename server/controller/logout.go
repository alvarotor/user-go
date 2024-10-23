package controller

import (
	"context"
	"net/http"
)

func (u *controllerUser) LogOut(c context.Context, email string) (int, error) {
	err := u.LogOutSvc(c, email)
	if err != nil {
		u.log.Error(err.Error())
		return http.StatusBadRequest, err
	}

	return http.StatusOK, nil
}
