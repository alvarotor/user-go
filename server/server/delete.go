package server

import (
	"context"

	pb "github.com/alvarotor/user-go/server/user-pb"
)

func (s *UserServer) Delete(ctx context.Context, req *pb.UserIDRequest) (*pb.UserIDResponse, error) {

	err := s.Controller.Delete(ctx, uint(req.Id), false)
	if err != nil {
		s.Log.Error(err.Error())
		return &pb.UserIDResponse{}, err
	}

	return &pb.UserIDResponse{
		Id: uint32(req.Id),
	}, nil
}
