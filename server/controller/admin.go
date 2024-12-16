package controller

import (
	"context"
	"errors"
)

func (s *controllerUser) UpdateUserAdminStatus(ctx context.Context, email string, admin bool) error {
	if email == "" {
		return errors.New("user email is required")
	}

	user, err := s.GetByEmail(ctx, email)
	if err != nil {
		return errors.New("failed to get user")
	}
	if user == nil {
		return errors.New("user not found")
	}

	user.Admin = admin

	err = s.UpdateField(ctx, user.ID, "admin", admin)
	if err != nil {
		return errors.New("failed to update user")
	}

	return nil
}
