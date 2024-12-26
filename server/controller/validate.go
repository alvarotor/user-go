package controller

import (
	"context"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/alvarotor/user-go/server/dto"
	"github.com/alvarotor/user-go/server/models"
	"github.com/golang-jwt/jwt/v5"
)

func (u *controllerUser) Validate(c context.Context, code string) (int, models.Token, error) {
	user, err := u.GetByCode(c, code)
	if err != nil {
		u.log.Error(err.Error())
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

	if user.CodeExpire.Before(time.Now().UTC()) {
		errMsg := "code is expired"
		u.log.Error(errMsg)
		return http.StatusBadRequest, models.Token{}, errors.New(errMsg)
	}

	if !user.Validated {
		err = u.ValidateSvc(c, user.Email)
		if err != nil {
			u.log.Error(err.Error())
			return http.StatusInternalServerError, models.Token{}, err
		}
	}

	expirationTime := getExpirationTime(uint(user.LoginLengthTime))

	claims := &dto.ClaimsResponse{
		Email:      user.Email,
		Admin:      user.Admin,
		SuperAdmin: user.SuperAdmin,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			Issuer:    u.conf.Issuer,
		},
		DeviceInfo: models.DeviceInfo{
			Browser:                user.Browser,
			BrowserVersion:         user.BrowserVersion,
			OperatingSystem:        user.OperatingSystem,
			OperatingSystemVersion: user.OperatingSystemVersion,
			Cpu:                    user.Cpu,
			Language:               user.Language,
			Timezone:               user.Timezone,
			CookiesEnabled:         user.CookiesEnabled,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(u.conf.JWTKey)
	if err != nil {
		u.log.Error(err.Error())
		return http.StatusInternalServerError, models.Token{}, err
	}

	model := models.Token{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
		Email:   user.Email,
	}

	return http.StatusOK, model, nil
}

func getExpirationTime(timeUser uint) time.Time {

	var expirationTime time.Time
	now := time.Now().UTC()
	switch timeUser {
	case 1:
		expirationTime = now.Add(time.Minute)
	case 5:
		expirationTime = now.Add(time.Duration(timeUser) * time.Minute)
	case 60:
		expirationTime = now.Add(time.Hour)
	case 12:
		expirationTime = now.Add(time.Duration(timeUser) * time.Hour)
	case 24:
		expirationTime = now.Add(time.Duration(timeUser) * time.Hour)
	case 30:
		expirationTime = now.Add(time.Duration(timeUser) * 24 * time.Hour)
	case 365:
		expirationTime = now.Add(time.Duration(timeUser) * 24 * time.Hour)
	default:
		expirationTime = now.Add(24 * time.Hour)
	}

	return expirationTime
}
