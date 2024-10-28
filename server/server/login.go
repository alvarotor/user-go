package server

import (
	"context"

	"github.com/alvarotor/user-go/server/dto"
	pb "github.com/alvarotor/user-go/server/user-pb"
	"github.com/go-playground/validator/v10"
)

func (s *UserServer) Login(ctx context.Context, req *pb.UserLoginRequest) (*pb.UserLoginResponse, error) {
	userLogin := dto.UserLogin{
		Email: req.GetEmail(),
		Time:  uint(req.LoginLengthTime),
	}

	validator := validator.New(validator.WithRequiredStructEnabled())
	err := validator.Struct(userLogin)
	if err != nil {
		s.Log.Error(err.Error())
		return &pb.UserLoginResponse{}, err
	}

	status, code, err := s.Controller.Login(ctx, userLogin)
	if err != nil {
		s.Log.Error(err.Error())
		return &pb.UserLoginResponse{}, err
	}

	return &pb.UserLoginResponse{
		Code:   code,
		Status: uint32(status),
	}, nil
}
