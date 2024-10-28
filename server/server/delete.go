package server

import (
	"context"

	pb "github.com/alvarotor/user-go/server/user-pb"
)

func (s *UserServer) Delete(ctx context.Context, req *pb.UserDeleteRequest) (*pb.UserIDResponse, error) {

	err := s.Controller.Delete(ctx, uint(req.Id), req.Permanently)
	if err != nil {
		s.Log.Error(err.Error())
		return &pb.UserIDResponse{}, err
	}

	return &pb.UserIDResponse{
		Id:     uint32(req.Id),
		Status: 1,
	}, nil
}
