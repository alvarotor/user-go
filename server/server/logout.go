package server

import (
	"context"

	pb "github.com/alvarotor/user-go/server/user-pb"
)

func (s *UserServer) LogOut(ctx context.Context, req *pb.UserIDRequest) (*pb.UserIDResponse, error) {
	user, err := s.Controller.Get(ctx, uint(req.Id), "")
	if err != nil {
		s.Log.Error(err.Error())
		return &pb.UserIDResponse{}, err
	}
	email := user.Email
	status, err := s.Controller.LogOut(ctx, email)
	if err != nil {
		s.Log.Error(err.Error())
		return &pb.UserIDResponse{}, err
	}

	return &pb.UserIDResponse{
		Id:     uint32(req.Id),
		Status: uint32(status),
	}, nil
}
