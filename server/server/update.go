package server

import (
	"context"

	"github.com/alvarotor/user-go/server/dto"
	pb "github.com/alvarotor/user-go/server/user-pb"
	"github.com/go-playground/validator/v10"
)

func (s *UserServer) Update(ctx context.Context, req *pb.UserUpdateRequest) (*pb.UserStatusResponse, error) {

	user := dto.UserUpdate{
		Name:       req.Name,
		ProfilePic: req.ProfilePic,
	}

	validate := validator.New(validator.WithRequiredStructEnabled())
	err := validate.Struct(user)
	if err != nil {
		s.Log.Error(err.Error())
		return &pb.UserStatusResponse{}, err
	}

	userID, err := s.Controller.GetByEmail(ctx, req.Email)
	if err != nil {
		s.Log.Error(err.Error())
		return &pb.UserStatusResponse{}, err
	}

	userID.Name = user.Name
	userID.ProfilePic = user.ProfilePic
	err = s.Controller.Update(ctx, uint(userID.ID), *userID)
	if err != nil {
		s.Log.Error(err.Error())
		return &pb.UserStatusResponse{}, err
	}

	return &pb.UserStatusResponse{
		Status: 1,
	}, nil
}
