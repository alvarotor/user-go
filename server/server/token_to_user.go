package server

import (
	"context"

	pb "github.com/alvarotor/user-go/server/user-pb"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *UserServer) TokenToUser(ctx context.Context, req *pb.UserTokenRequest) (*pb.UserResponse, error) {
	s.Log.Info("TokenToUser gRPC request received",
		"token_length", len(req.GetToken()),
		"browser", req.GetBrowser(),
		"os", req.GetOperatingSystem())

	user, err := s.UserController.TokenToUser(
		ctx,
		req.GetToken(),
		req.GetBrowser(),
		req.GetBrowserVersion(),
		req.GetOperatingSystem(),
		req.GetOperatingSystemVersion(),
		req.GetCpu(),
		req.GetLanguage(),
		req.GetTimezone(),
		req.GetCookiesEnabled(),
	)
	if err != nil {
		s.Log.Error("TokenToUser controller failed", "error", err.Error())
		return &pb.UserResponse{}, err
	}

	s.Log.Info("TokenToUser successful", "email", user.Email, "admin", user.Admin)

	return &pb.UserResponse{
		Email:      user.Email,
		Name:       user.Name,
		ProfilePic: user.ProfilePic,
		Validated:  user.Validated,
		Admin:      user.Admin,
		SuperAdmin: user.SuperAdmin,
		Code:       user.Code,
		CodeExpire: timestamppb.New(user.CodeExpire),
		Bucket:     user.Bucket,
	}, nil
}
