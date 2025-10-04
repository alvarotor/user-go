package server

import (
	"context"

	"github.com/alvarotor/user-go/server/dto"
	pb "github.com/alvarotor/user-go/server/user-pb"
	"github.com/golang-jwt/jwt/v5"

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

	// Parse token to extract expiration times for security fix
	var accessTokenExpiresAt int64
	claims := &dto.ClaimsResponse{}
	// Parse without validation since token was already validated by controller
	if _, _, err := jwt.NewParser().ParseUnverified(req.GetToken(), claims); err == nil {
		if claims.ExpiresAt != nil {
			accessTokenExpiresAt = claims.ExpiresAt.Unix()
		}
	} else {
		s.Log.Warn("Failed to parse token claims for expiration", "error", err.Error())
	}

	return &pb.UserResponse{
		Email:                 user.Email,
		Name:                  user.Name,
		ProfilePic:            user.ProfilePic,
		Validated:             user.Validated,
		Admin:                 user.Admin,
		SuperAdmin:            user.SuperAdmin,
		Code:                  user.Code,
		CodeExpire:            timestamppb.New(user.CodeExpire),
		Bucket:                user.Bucket,
		AccessTokenExpiresAt:  accessTokenExpiresAt,
		RefreshTokenExpiresAt: 0, // Not implemented yet, set to 0
	}, nil
}
