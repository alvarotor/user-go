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

	expirationTime := getExpirationTime(uint(u.conf.TokenExpirationTime))

	claims := &dto.ClaimsResponse{
		Email:      user.Email,
		Admin:      user.Admin,
		SuperAdmin: user.SuperAdmin,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			Issuer:    u.conf.Issuer,
		},
		DeviceInfo: createDeviceInfo(user),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(u.conf.JWTKey)
	if err != nil {
		u.log.Error(err.Error())
		return http.StatusInternalServerError, models.Token{}, err
	}

	expirationTimeRefresh := getExpirationTime(uint(u.conf.TokenExpirationTimeRefresh))

	claimsRefresh := &dto.ClaimsRefreshResponse{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTimeRefresh),
			Issuer:    u.conf.Issuer,
		},
		DeviceInfo: createDeviceInfo(user),
		CodeRefresh: user.CodeRefresh,
	}
	tokenRefresh := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsRefresh)
	tokenRefreshString, err := tokenRefresh.SignedString(u.conf.JWTKey)
	if err != nil {
		u.log.Error(err.Error())
		return http.StatusInternalServerError, models.Token{}, err
	}

	model := models.Token{
		Name:    "token",
		Token:   tokenString,
		ExpiresRefresh: expirationTimeRefresh,
		Email:   user.Email,
		RefreshToken: tokenRefreshString,
	}

	return http.StatusOK, model, nil
}

func getExpirationTime(seconds uint) time.Time {
	var expirationTime time.Time
	now := time.Now().UTC()
	expirationTime = now.Add(time.Duration(seconds) * time.Second)
	return expirationTime
}

func createDeviceInfo(user *models.User) models.DeviceInfo {
	return models.DeviceInfo{
		Browser:                user.Browser,
		BrowserVersion:         user.BrowserVersion,
		OperatingSystem:        user.OperatingSystem,
		OperatingSystemVersion: user.OperatingSystemVersion,
		Cpu:                    user.Cpu,
		Language:               user.Language,
		Timezone:               user.Timezone,
		CookiesEnabled:         user.CookiesEnabled,
	}
}
