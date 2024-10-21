package service

import (
	"context"
	"errors"
	"time"

	"github.com/alvarotor/entitier-go/repository"
	"github.com/alvarotor/user-go/server/model"
	"gorm.io/gorm"
)

type IUserService interface {
	repository.IGenericRepo[model.User, uint]
	GetByEmail(context.Context, string) (*model.User, error)
	GetByCode(context.Context, string) (*model.User, error)
	Validate(context.Context, string) error
	LogOut(context.Context, string) error
}

type userService struct {
	repository.IGenericRepo[model.User, uint]
}

func NewUserService(
	db *gorm.DB,
) IUserService {
	repo := repository.NewGenericRepository[model.User, uint](db)
	return &userService{
		IGenericRepo: repo,
	}
}

func (s userService) GetByEmail(ctx context.Context, email string) (*model.User, error) {
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

func (s *userService) Validate(ctx context.Context, email string) error {
	user, err := s.GetByEmail(ctx, email)
	if err != nil {
		return err
	}

	user.Validated = true

	return s.Update(ctx, user.ID, *user)
}

func (s *userService) LogOut(ctx context.Context, email string) error {
	user, err := s.GetByEmail(ctx, email)
	if err != nil {
		return err
	}

	user.Code = ""
	user.LoginLengthTime = 0
	user.CodeExpire = time.Time{}

	return s.Update(ctx, user.ID, *user)
}
