package controller

import (
	"context"
	"net/http"
)

func (u *controllerUser) Delete(c context.Context, id uint, preload bool) (int, error) {
	err := u.svc.Delete(c, id, preload)
	if err != nil {
		u.log.Error(err.Error())
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}
