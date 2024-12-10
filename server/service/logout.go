package service

import (
	"context"
	"time"
)

func (s *userService) LogOutSvc(ctx context.Context, email string) error {
	user, err := s.GetByEmail(ctx, email)
	if err != nil {
		return err
	}

	user.Code = " "
	user.LoginLengthTime = 0
	user.CodeExpire = time.Time{}

	return s.Update(ctx, user.ID, *user)
}
