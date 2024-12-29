package services

import (
	"context"
	"time"
)

func (s *userService) LogOutSvc(ctx context.Context, email string) error {
	user, err := s.GetByEmail(ctx, email)
	if err != nil {
		return err
	}

	user.Code = "OUT"
	user.CodeRefresh = "OUT"
	user.CodeExpire = time.Time{}

	return s.Update(ctx, user.ID, *user)
}
