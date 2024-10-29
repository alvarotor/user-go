package server

import (
	"context"

	"github.com/alvarotor/user-go/server/model"
	pb "github.com/alvarotor/user-go/server/user-pb"
	"github.com/go-playground/validator/v10"
)

func (s *UserServer) Update(ctx context.Context, req *pb.UserUpdateRequest) (*pb.UserStatusResponse, error) {

	user := model.User{
		Email:      req.Email,
		Name:       req.Name,
		ProfilePic: req.ProfilePic,
	}

	validate := validator.New(validator.WithRequiredStructEnabled())
	err := validate.Struct(user)
	if err != nil {
		s.Log.Error(err.Error())
		return &pb.UserStatusResponse{}, err
	}

	id, err := s.Controller.GetByEmail(ctx, req.Email)
	if err != nil {
		s.Log.Error(err.Error())
		return &pb.UserStatusResponse{}, err
	}

	err = s.Controller.Update(ctx, uint(id.ID), user)
	if err != nil {
		s.Log.Error(err.Error())
		return &pb.UserStatusResponse{}, err

	}

	return &pb.UserStatusResponse{
		Status: 1,
	}, nil
}
