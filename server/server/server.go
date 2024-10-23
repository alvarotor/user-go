package server

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/alvarotor/user-go/server/controller"
	"github.com/alvarotor/user-go/server/dto"
	"github.com/alvarotor/user-go/server/model"
	pb "github.com/alvarotor/user-go/server/user-pb"
	"github.com/go-playground/validator/v10"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type UserServer struct {
	pb.UnimplementedUserServer
	// users map[uint32]*pb.UserResponse
	Controller controller.IControllerUser
	Log        *slog.Logger
}

func NewServer(
	controller controller.IControllerUser,
	log *slog.Logger,
) *UserServer {
	return &UserServer{
		Controller: controller,
		Log:        log,
	}
}

func (s *UserServer) Create(ctx context.Context, req *pb.UserRequest) (*pb.UserIDRequest, error) {
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
		return &pb.UserIDRequest{}, err
	}

	userCreated, err := s.Controller.Create(ctx, user)
	if err != nil {
		s.Log.Error(err.Error())
		return &pb.UserIDRequest{}, err
	}

	return &pb.UserIDRequest{
		Id: uint32(userCreated.ID),
	}, nil
}

func (s *UserServer) Get(ctx context.Context, req *pb.UserIDRequest) (*pb.UserResponse, error) {

	user, err := s.Controller.Get(ctx, uint(req.Id), "")
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

func (s *UserServer) Delete(ctx context.Context, req *pb.UserIDRequest) (*pb.UserIDResponse, error) {

	err := s.Controller.Delete(ctx, uint(req.Id), false)
	if err != nil {
		s.Log.Error(err.Error())
		return &pb.UserIDResponse{}, err
	}

	return &pb.UserIDResponse{
		Id: uint32(req.Id),
	}, nil
}

func (s *UserServer) List(ctx context.Context, _ *emptypb.Empty) (*pb.ListUsersResponse, error) {

	users, err := s.Controller.GetAll(ctx)
	if err != nil {
		s.Log.Error(err.Error())
		return &pb.ListUsersResponse{}, err
	}

	pbUsers := []*pb.UserResponse{}

	for _, user := range users {
		pbUser := pb.UserResponse{
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
		}
		pbUsers = append(pbUsers, &pbUser)
	}

	return &pb.ListUsersResponse{
		Users: pbUsers,
	}, nil
}

func (s *UserServer) Login(ctx context.Context, req *pb.UserLoginRequest) (*pb.UserIDResponse, error) {
	userLogin := dto.UserLogin{
		Email: req.Email,
		Time:  uint(req.LoginLengthTime),
	}

	validator := validator.New(validator.WithRequiredStructEnabled())
	err := validator.Struct(userLogin)
	if err != nil {
		s.Log.Error(err.Error())
		return &pb.UserIDResponse{}, err
	}

	status, id, err := s.Controller.Login(ctx, userLogin)
	if err != nil {
		s.Log.Error(err.Error())
		return &pb.UserIDResponse{}, err
	}

	return &pb.UserIDResponse{
		Id:     uint32(id),
		Status: uint32(status),
	}, nil
}

func (s *UserServer) LogOut(ctx context.Context, req *pb.UserIDRequest) (*pb.UserIDResponse, error) {
	user, err := s.Controller.Get(ctx, uint(req.Id), "")
	if err != nil {
		s.Log.Error(err.Error())
		return &pb.UserIDResponse{}, err
	}
	email := user.Email
	status, err := s.Controller.LogOut(ctx, email)
	if err != nil {
		s.Log.Error(err.Error())
		return &pb.UserIDResponse{}, err
	}

	return &pb.UserIDResponse{
		Id:     uint32(req.Id),
		Status: uint32(status),
	}, nil
}

func (s *UserServer) Validate(ctx context.Context, req *pb.UserValidateRequest) (*pb.UserTokenResponse, error) {

	status, token, err := s.Controller.Validate(ctx, req.Code)
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
