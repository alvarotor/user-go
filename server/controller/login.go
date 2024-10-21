package controller

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/alvarotor/user-go/server/dto"
	"github.com/alvarotor/user-go/server/model"
	"golang.org/x/exp/rand"
)

func (u *controllerUser) Login(c context.Context, userLogin dto.UserLogin) (int, error) {
	user, err := u.svc.GetByEmail(c, userLogin.Email)
	if err != nil && err.Error() != "user not found" {
		u.log.Error(err.Error())
		return http.StatusInternalServerError, err
	}

	tenMinutes := time.Now().UTC().Add(10 * time.Minute)

	if user == nil {

		user = new(model.User)
		user.Email = userLogin.Email
		user.LoginLengthTime = uint32(userLogin.Time)
		user.Code = u.generateRandomString(u.conf.SizeRandomStringValidation)
		user.CodeExpire = tenMinutes
		user.Validated = false

		user, err := u.svc.Create(c, *user)
		if err != nil {
			u.log.Error(err.Error())
			return http.StatusInternalServerError, err
		}

		if user.ID == 0 {
			return http.StatusInternalServerError, errors.New("user not created")
		}

	} else {

		user.Code = u.generateRandomString(u.conf.SizeRandomStringValidation)
		user.CodeExpire = tenMinutes
		user.LoginLengthTime = uint32(userLogin.Time)

		err := u.svc.Update(c, user.ID, *user)
		if err != nil {
			u.log.Error(err.Error())
			return http.StatusInternalServerError, err
		}
	}

	u.log.Info("user code: " + user.Code)

	return http.StatusOK, nil
}

func (u *controllerUser) generateRandomString(length int) string {
	r := rand.New(rand.NewSource(uint64(time.Now().UnixNano())))
	var letters = []rune(u.conf.RandomStringValidation)
	b := make([]rune, length)
	for i := range b {
		b[i] = letters[r.Intn(len(letters))]
	}
	return string(b)
}
