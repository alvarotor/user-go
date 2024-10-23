package service

import (
	"context"
	"errors"
	"time"

	"github.com/alvarotor/user-go/server/model"
)

func (s userService) GetByCode(ctx context.Context, code string) (*model.User, error) {
	users, err := s.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	for _, user := range users {
		if user.Code == code {
			if time.Now().UTC().After(user.CodeExpire) {
				return nil, errors.New("code is expired")
			}
			return user, nil
		}
	}

	return nil, model.ErrUserNotFound
}
