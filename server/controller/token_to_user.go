package controller

import (
	"context"
	"errors"

	"github.com/alvarotor/user-go/server/dto"
	"github.com/alvarotor/user-go/server/model"
	"github.com/golang-jwt/jwt/v5"
)

func (u *controllerUser) TokenToUser(c context.Context, token string) (*model.User, error) {

	claims := &dto.ClaimsResponse{}

	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (any, error) {
		return u.conf.JWTKey, nil
	})
	if errors.Is(err, jwt.ErrSignatureInvalid) {
		errMsg := "invalid signature token"
		u.log.Error(errMsg)
		return &model.User{}, errors.New(errMsg)
	}
	if errors.Is(err, jwt.ErrTokenExpired) {
		errMsg := "invalid signature token"
		u.log.Error(errMsg)
		return &model.User{}, errors.New(errMsg)
	}
	if err != nil {
		errMsg := "error parsing token"
		u.log.Error(errMsg)
		return &model.User{}, errors.New(errMsg)
	}
	if !tkn.Valid {
		errMsg := "invalid token"
		u.log.Error(errMsg)
		return &model.User{}, errors.New(errMsg)
	}
	user, err := u.GetByEmail(c, claims.Email)
	if err != nil {
		return &model.User{}, err
	}
	if user.Code == "" {
		errMsg := "user not logged"
		u.log.Error(errMsg)
		return &model.User{}, err

	}
	if len(user.Code) != u.conf.SizeRandomStringValidation {
		errMsg := "invalid user"
		u.log.Error(errMsg)
		return &model.User{}, err
	}

	return user, nil
}
