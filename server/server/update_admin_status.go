package server

import (
	"context"

	"github.com/alvarotor/user-go/server/models"
	pb "github.com/alvarotor/user-go/server/user-pb"
)

func (s *UserServer) UpdateUserAdminStatus(ctx context.Context, req *pb.UpdateUserAdminRequest) (*pb.UserStatusResponse, error) {
	// Authenticate the requesting user
	authUser, err := s.UserController.TokenToUser(
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
		s.Log.Error("authentication failed for UpdateUserAdminStatus", "error", err)
		return nil, err
	}

	// Check if the authenticated user has SuperAdmin role
	if !authUser.SuperAdmin {
		s.Log.Error("unauthorized access attempt to UpdateUserAdminStatus", "user", authUser.Email, "super_admin", authUser.SuperAdmin)
		return nil, models.ErrUnauthorized
	}

	err = s.UserController.UpdateUserAdminStatus(ctx, req.Email, req.Admin)
	if err != nil {
		return nil, err
	}

	return &pb.UserStatusResponse{
		Status: 1,
	}, nil
}
