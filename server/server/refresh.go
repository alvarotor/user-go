package server

import (
	"context"

	pb "github.com/alvarotor/user-go/server/user-pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *UserServer) Refresh(ctx context.Context, req *pb.UserValidateRefreshRequest) (*pb.UserTokenResponse, error) {
	status, token, err := s.Controller.Refresh(ctx, req.GetRefreshToken())
	if err != nil {
		s.Log.Error(err.Error())
		return &pb.UserTokenResponse{}, err
	}

	return &pb.UserTokenResponse{
		Name:    token.Name,
		Token:   token.Token,
		Expires: timestamppb.New(token.Expires),
		Email:   token.Email,
		Status:  uint32(status),
		RefreshToken: token.RefreshToken,
	}, nil
}
