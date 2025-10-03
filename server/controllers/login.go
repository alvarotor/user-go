package controllers

import (
	"context"
	"crypto/rand"
	"errors"
	"net/http"
	"time"

	entModels "github.com/alvarotor/entitier-go/models"
	"github.com/alvarotor/user-go/server/dto"
	"github.com/alvarotor/user-go/server/models"
)

func (u *controllerUser) Login(c context.Context, userLogin dto.UserLogin) (int, string, error) {
	user, err := u.IUserService.GetByEmail(c, userLogin.Email)
	if !errors.Is(err, entModels.ErrNotFound) {
		if err != nil {
			u.log.Error(err.Error())
			return http.StatusNotFound, "", err
		}
	}

	tenMinutes := time.Now().UTC().Add(10 * time.Minute)

	if user == nil {

		user = new(models.User)
		user.Email = userLogin.Email
		user.Code = u.GenerateRandomString(u.conf.SizeRandomStringValidation)
		user.CodeRefresh = u.GenerateRandomString(u.conf.SizeRandomStringValidationRefresh)
		user.CodeExpire = tenMinutes
		user.Validated = false
		user.DeviceInfo = userLogin.DeviceInfo

		user, err := u.Create(c, *user)
		if err != nil {
			u.log.Error(err.Error())
			return http.StatusInternalServerError, "", err
		}

		if user.ID == 0 {
			return http.StatusInternalServerError, "", errors.New("user not created")
		}

	} else {

		user.Code = u.GenerateRandomString(u.conf.SizeRandomStringValidation)
		user.CodeRefresh = u.GenerateRandomString(u.conf.SizeRandomStringValidationRefresh)
		user.CodeExpire = tenMinutes
		user.DeviceInfo = userLogin.DeviceInfo

		err := u.Update(c, user.ID, *user)
		if err != nil {
			u.log.Error(err.Error())
			return http.StatusInternalServerError, "", err
		}
	}

	u.log.Info("user code: " + user.Code)
	u.log.Info("user code refresh: " + user.CodeRefresh)

	return http.StatusOK, user.Code, nil
}

func (u *controllerUser) GenerateRandomString(length int) string {
	var letters = []rune(u.conf.RandomStringValidation)
	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		u.log.Error(err.Error())
		return ""
	}
	for i := range b {
		b[i] = byte(letters[int(b[i])%len(letters)])
	}
	return string(b)
}
