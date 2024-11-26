package server

import (
	"context"
	"errors"
	"net/http"

	pb "github.com/alvarotor/user-go/server/user-pb"
)

func (s *UserServer) Health(ctx context.Context, req *pb.UserIDRequest) (*pb.UserStatusResponse, error) {

	status := s.Controller.Health(ctx, req.GetId())
	if status != http.StatusOK {
		errMsg := "Not Healthy"
		s.Log.Error(errMsg)
		return &pb.UserStatusResponse{}, errors.New(errMsg)
	}

	s.Log.Info("Healthy")

	return &pb.UserStatusResponse{
		Status: http.StatusOK,
	}, nil
}
