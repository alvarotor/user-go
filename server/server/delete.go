package server

import (
	"context"
	"errors"

	pb "github.com/alvarotor/user-go/server/user-pb"
)

func (s *UserServer) Delete(ctx context.Context, req *pb.UserDeleteRequest) (*pb.UserStatusResponse, error) {

	if req.GetEmail() == "" {
		s.Log.Error("email is required")
		return &pb.UserStatusResponse{}, errors.New("email is required")
	}

	userMailRequest := pb.UserMailRequest{
		Email: req.GetEmail(),
	}
	user, err := s.GetIDByEmail(ctx, &userMailRequest)
	if err != nil {
		s.Log.Error(err.Error())
		return &pb.UserStatusResponse{}, err
	}

	err = s.UserController.Delete(ctx, uint(user.GetId()), req.Permanently)
	if err != nil {
		s.Log.Error(err.Error())
		return &pb.UserStatusResponse{}, err
	}

	return &pb.UserStatusResponse{
		Status: 1,
	}, nil
}
