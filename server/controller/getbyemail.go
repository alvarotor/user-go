package controller

import (
	"context"

	"github.com/alvarotor/user-go/server/model"
)

func (c *controllerUser) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	user, err := c.GetByEmailSvc(ctx, email)
	if err != nil {
		return nil, err
	}
	return user, nil
}
