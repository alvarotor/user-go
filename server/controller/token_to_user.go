package controller

import (
	"context"
	"errors"

	"github.com/alvarotor/user-go/server/dto"
	"github.com/alvarotor/user-go/server/models"
	"github.com/golang-jwt/jwt/v5"
)

func (u *controllerUser) TokenToUser(
	c context.Context,
	token string,
	browser string,
	browserVersion string,
	operatingSystem string,
	operatingSystemVersion string,
	cpu string,
	language string,
	timezone string,
	cookiesEnabled bool,
) (*models.User, error) {
	claims := &dto.ClaimsResponse{}

	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (any, error) {
		return u.conf.JWTKey, nil
	})
	if errors.Is(err, jwt.ErrSignatureInvalid) {
		errMsg := "invalid signature token"
		u.log.Error(errMsg)
		return &models.User{}, errors.New(errMsg)
	}
	if errors.Is(err, jwt.ErrTokenExpired) {
		errMsg := "invalid signature token"
		u.log.Error(errMsg)
		return &models.User{}, errors.New(errMsg)
	}
	if err != nil {
		errMsg := "error parsing token"
		u.log.Error(errMsg)
		return &models.User{}, errors.New(errMsg)
	}
	if !tkn.Valid {
		errMsg := "invalid token"
		u.log.Error(errMsg)
		return &models.User{}, errors.New(errMsg)
	}
	user, err := u.GetByEmail(c, claims.Email)
	if err != nil {
		return &models.User{}, err
	}
	if user.Code == "OUT" {
		errMsg := "user not logged"
		u.log.Error(errMsg)
		return &models.User{}, errors.New(errMsg)
	}
	if len(user.Code) != u.conf.SizeRandomStringValidation {
		errMsg := "invalid user"
		u.log.Error(errMsg)
		return &models.User{}, errors.New(errMsg)
	}

	secs := models.BaseSecurityLogin{
		Browser:                browser,
		BrowserVersion:         browserVersion,
		OperatingSystem:        operatingSystem,
		OperatingSystemVersion: operatingSystemVersion,
		Cpu:                    cpu,
		Language:               language,
		Timezone:               timezone,
		CookiesEnabled:         cookiesEnabled,
	}

	if claims.BaseSecurityLogin != secs {
		errMsg := "security questions user not match"
		u.log.Error(errMsg)
		return &models.User{}, errors.New(errMsg)
	}

	return user, nil
}
