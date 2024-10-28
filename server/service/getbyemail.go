package service

import (
	"context"

	"github.com/alvarotor/user-go/server/model"
)

func (s userService) GetByEmailSvc(ctx context.Context, email string) (*model.User, error) {
	users, err := s.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	for _, user := range users {
		if user.Email == email {
			return user, nil
		}
	}

	return nil, model.ErrUserNotFound
}
