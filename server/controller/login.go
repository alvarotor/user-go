package controller

import (
	"context"
	"errors"
	"net/http"
	"time"

	entModels "github.com/alvarotor/entitier-go/models"
	"github.com/alvarotor/user-go/server/dto"
	"github.com/alvarotor/user-go/server/models"
	"golang.org/x/exp/rand"
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
		user.LoginLengthTime = uint32(userLogin.Time)
		user.Code = u.generateRandomString(u.conf.SizeRandomStringValidation)
		user.CodeExpire = tenMinutes
		user.Validated = false
		user.BaseSecurityLogin = userLogin.BaseSecurityLogin

		user, err := u.Create(c, *user)
		if err != nil {
			u.log.Error(err.Error())
			return http.StatusInternalServerError, "", err
		}

		if user.ID == 0 {
			return http.StatusInternalServerError, "", errors.New("user not created")
		}

	} else {

		user.Code = u.generateRandomString(u.conf.SizeRandomStringValidation)
		user.CodeExpire = tenMinutes
		user.LoginLengthTime = uint32(userLogin.Time)
		user.BaseSecurityLogin = userLogin.BaseSecurityLogin

		err := u.Update(c, user.ID, *user)
		if err != nil {
			u.log.Error(err.Error())
			return http.StatusInternalServerError, "", err
		}
	}

	u.log.Info("user code: " + user.Code)

	return http.StatusOK, user.Code, nil
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
