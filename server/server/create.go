package server

import (
	"context"
	"fmt"
	"time"

	"github.com/alvarotor/user-go/server/model"
	pb "github.com/alvarotor/user-go/server/user-pb"
	"github.com/go-playground/validator/v10"
)

func (s *UserServer) Create(ctx context.Context, req *pb.UserRequest) (*pb.UserIDResponse, error) {
	user := model.User{
		Email:           req.Email,
		Name:            req.Name,
		Password:        req.Password,
		ProfilePic:      req.ProfilePic,
		LoginLengthTime: req.LoginLengthTime,
		Validated:       false,
		Admin:           false,
		SuperAdmin:      false,
		ValidationCode:  "1",
		Code:            "1",
		CodeExpire:      time.Time{},
	}

	s.Log.Info(fmt.Sprintf("%v\n", user))

	validate := validator.New(validator.WithRequiredStructEnabled())
	err := validate.Struct(user)
	if err != nil {
		s.Log.Error(err.Error())
		return &pb.UserIDResponse{}, err
	}

	userCreated, err := s.Controller.Create(ctx, user)
	if err != nil {
		s.Log.Error(err.Error())
		return &pb.UserIDResponse{}, err
	}

	return &pb.UserIDResponse{
		Id:     uint32(userCreated.ID),
		Status: 1,
	}, nil
}
