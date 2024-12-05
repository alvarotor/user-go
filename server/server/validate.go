package server

import (
	"context"

	pb "github.com/alvarotor/user-go/server/user-pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *UserServer) Validate(ctx context.Context, req *pb.UserValidateRequest) (*pb.UserTokenResponse, error) {
	status, token, err := s.Controller.Validate(ctx, req.GetCode())
	if err != nil {
		s.Log.Error(err.Error())
		return &pb.UserTokenResponse{}, err
	}

	return &pb.UserTokenResponse{
		Name:    token.Name,
		Value:   token.Value,
		Expires: timestamppb.New(token.Expires),
		Email:   token.Email,
		Status:  uint32(status),
	}, nil
}
