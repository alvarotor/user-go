package controller

import (
	"context"
	"errors"
	"net/http"

	"github.com/alvarotor/user-go/server/dto"
	"github.com/alvarotor/user-go/server/model"
	"github.com/golang-jwt/jwt/v5"
)

func (u *controllerUser) TokenToUser(c context.Context, token string) (int, model.User, error) {

	claims := &dto.ClaimsResponse{}

	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (any, error) {
		return []byte{}, nil
	})
	if errors.Is(err, jwt.ErrSignatureInvalid) {
		errMsg := "invalid signature token"
		u.log.Error(errMsg)
		return http.StatusUnauthorized, model.User{}, errors.New(errMsg)
	}
	if errors.Is(err, jwt.ErrTokenExpired) {
		errMsg := "invalid signature token"
		u.log.Error(errMsg)
		return http.StatusUnauthorized, model.User{}, errors.New(errMsg)
	}
	if err != nil {
		errMsg := "parsing token"
		u.log.Error(errMsg)
		return http.StatusUnauthorized, model.User{}, errors.New(errMsg)
	}
	if !tkn.Valid {
		errMsg := "invalid token"
		u.log.Error(errMsg)
		return http.StatusUnauthorized, model.User{}, errors.New(errMsg)
	}
	user, err := u.GetByEmail(c, claims.Email)
	if err != nil {
		return http.StatusInternalServerError, model.User{}, err
	}
	if user.Code == "" {
		errMsg := "user not logged"
		u.log.Error(errMsg)
		return http.StatusUnauthorized, model.User{}, err

	}
	if len(user.Code) != u.conf.SizeRandomStringValidation {
		errMsg := "invalid user"
		u.log.Error(errMsg)
		return http.StatusUnauthorized, model.User{}, err
	}

	return http.StatusOK, *user, nil
}
