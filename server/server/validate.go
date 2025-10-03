package server

import (
	"context"

	pb "github.com/alvarotor/user-go/server/user-pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *UserServer) Validate(ctx context.Context, req *pb.UserValidateRequest) (*pb.UserTokenResponse, error) {
	s.Log.Info("Validate gRPC request received", "code_length", len(req.GetCode()))

	status, token, err := s.UserController.Validate(ctx, req.GetCode())
	if err != nil {
		s.Log.Error("Validate controller failed", "error", err.Error())
		return &pb.UserTokenResponse{}, err
	}

	response := &pb.UserTokenResponse{
		Token:               token.Token,
		TokenRefresh:        token.TokenRefresh,
		TokenRefreshExpires: timestamppb.New(token.TokenRefreshExpires),
		TokenExpires:        timestamppb.New(token.TokenExpires),
		Email:               token.Email,
		Status:              uint32(status),
	}

	s.Log.Info("Validate gRPC response prepared",
		"email", token.Email,
		"token_length", len(token.Token),
		"refresh_token_length", len(token.TokenRefresh),
		"expires_at", token.TokenExpires,
		"refresh_expires_at", token.TokenRefreshExpires)

	return response, nil
}
