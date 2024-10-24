package server

import (
	"context"

	"github.com/alvarotor/user-go/server/model"
	pb "github.com/alvarotor/user-go/server/user-pb"
	"github.com/go-playground/validator/v10"
)

func (s *UserServer) Update(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UserIDResponse, error) {

	user := model.User{
		Email:           req.User.Email,
		Name:            req.User.Name,
		Password:        req.User.Password,
		ProfilePic:      req.User.ProfilePic,
		LoginLengthTime: req.User.LoginLengthTime,
		Admin:           req.User.Admin,
		SuperAdmin:      req.User.SuperAdmin,
		ValidationCode:  req.User.ValidationCode,
		Code:            req.User.Code,
		CodeExpire:      req.User.CodeExpire.AsTime(),
	}

	validate := validator.New(validator.WithRequiredStructEnabled())
	err := validate.Struct(user)
	if err != nil {
		s.Log.Error(err.Error())
		return &pb.UserIDResponse{}, err
	}

	err = s.Controller.Update(ctx, uint(req.Id), user)
	if err != nil {
		s.Log.Error(err.Error())
		return &pb.UserIDResponse{}, err

	}

	return &pb.UserIDResponse{
		Id: uint32(req.Id),
	}, nil
}
