package server

import (
	"context"

	pb "github.com/alvarotor/user-go/server/user-pb"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *UserServer) TokenToUser(ctx context.Context, req *pb.UserTokenRequest) (*pb.UserResponse, error) {

	user, err := s.Controller.TokenToUser(ctx, req.GetToken())
	if err != nil {
		s.Log.Error(err.Error())
		return &pb.UserResponse{}, err
	}

	return &pb.UserResponse{
		Email:           user.Email,
		Name:            user.Name,
		ProfilePic:      user.ProfilePic,
		Validated:       user.Validated,
		Admin:           user.Admin,
		SuperAdmin:      user.SuperAdmin,
		LoginLengthTime: user.LoginLengthTime,
		ValidationCode:  user.ValidationCode,
		Code:            user.Code,
		CodeExpire:      timestamppb.New(user.CodeExpire),
	}, nil
}
