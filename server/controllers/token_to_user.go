package controllers

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

	if err := u.validateToken(tkn, err); err != nil {
		return &models.User{}, err
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

	secs := models.DeviceInfo{
		Browser:                browser,
		BrowserVersion:         browserVersion,
		OperatingSystem:        operatingSystem,
		OperatingSystemVersion: operatingSystemVersion,
		Cpu:                    cpu,
		Language:               language,
		Timezone:               timezone,
		CookiesEnabled:         cookiesEnabled,
	}

	if claims.DeviceInfo != secs {
		errMsg := "security data user's token don't match"
		u.log.Error(errMsg)
		return &models.User{}, errors.New(errMsg)
	}

	return user, nil
}

func (u *controllerUser) validateToken(tkn *jwt.Token, err error) error {
	if errors.Is(err, jwt.ErrSignatureInvalid) {
		return u.logAndReturnError("invalid signature token")
	}
	if errors.Is(err, jwt.ErrTokenExpired) {
		return u.logAndReturnError("token expired")
	}
	if err != nil {
		return u.logAndReturnError("error parsing token")
	}
	if !tkn.Valid {
		return u.logAndReturnError("invalid token")
	}
	return nil
}

func (u *controllerUser) logAndReturnError(errMsg string) error {
	u.log.Error(errMsg)
	return errors.New(errMsg)
}
