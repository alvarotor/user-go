package services

import "context"

func (s *userService) ValidateSvc(ctx context.Context, email string) error {
	user, err := s.GetByEmail(ctx, email)
	if err != nil {
		return err
	}

	user.Validated = true

	return s.Update(ctx, user.ID, *user)
}
