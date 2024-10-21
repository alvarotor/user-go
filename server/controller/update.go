package controller

import (
	"context"
	"errors"
	"net/http"

	"github.com/alvarotor/user-go/server/dto"
	"github.com/alvarotor/user-go/server/model"
)

func (u *controllerUser) Update(c context.Context, id uint, userToUpdate dto.UserRequest) (int, error) {
	user, err := u.svc.Get(c, id, "")
	if errors.Is(err, model.ErrUserNotFound) {
		u.log.Info(model.ErrUserNotFound.Error())
		return http.StatusNotFound, err
	}
	if err != nil {
		u.log.Error(err.Error())
		return http.StatusInternalServerError, err
	}

	user.Name = userToUpdate.Name
	user.Age = uint32(userToUpdate.Age)
	user.Gender = uint32(userToUpdate.Gender)
	user.CountryOrigin = userToUpdate.CountryOrigin

	err = u.svc.Update(c, user.ID, *user)
	if err != nil {
		errMsg := "user not updated"
		u.log.Error(errMsg)
		return http.StatusInternalServerError, errors.New(errMsg)
	}

	return http.StatusOK, nil
}
