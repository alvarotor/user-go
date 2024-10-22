package server

import (
	"context"
	"log/slog"
	"time"

	"github.com/alvarotor/user-go/server/controller"
	"github.com/alvarotor/user-go/server/model"
	"github.com/alvarotor/user-go/server/service"
	pb "github.com/alvarotor/user-go/server/user-pb"
	"github.com/go-playground/validator/v10"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type UserServer struct {
	pb.UnimplementedUserServer
	// users map[uint32]*pb.UserResponse
	Controller controller.IControllerUser
	Svc        service.IUserService
	Log        *slog.Logger
}

func NewServer(
	controller controller.IControllerUser,
	svc service.IUserService,
	log *slog.Logger,
) *UserServer {
	return &UserServer{
		Controller: controller,
		Svc:        svc,
		Log:        log,
	}
}

func (s *UserServer) Create(ctx context.Context, req *pb.UserRequest) (*pb.UserIDRequest, error) {
	user := model.User{
		Email:           req.Email,
		Name:            req.Name,
		Password:        req.Password,
		Age:             req.Age,
		Gender:          req.Gender,
		CountryOrigin:   req.CountryOrigin,
		ProfilePic:      req.ProfilePic,
		LoginLengthTime: req.LoginLengthTime,
		Validated:       false,
		Admin:           false,
		SuperAdmin:      false,
		ValidationCode:  "1",
		Code:            "1",
		CodeExpire:      time.Time{},
	}

	// s.Log.Info(fmt.Sprintf("%v\n", user))

	validate := validator.New(validator.WithRequiredStructEnabled())
	err := validate.Struct(user)
	if err != nil {
		s.Log.Error(err.Error())
		return &pb.UserIDRequest{}, err
	}

	userCreated, err := s.Svc.Create(ctx, user)
	if err != nil {
		s.Log.Error(err.Error())
		return &pb.UserIDRequest{}, err
	}

	return &pb.UserIDRequest{
		Id: uint32(userCreated.ID),
	}, nil
}

func (s *UserServer) Get(ctx context.Context, req *pb.UserIDRequest) (*pb.UserResponse, error) {

	user, err := s.Svc.Get(ctx, uint(req.Id), "")
	if err != nil {
		s.Log.Error(err.Error())
		return &pb.UserResponse{}, err
	}

	return &pb.UserResponse{
		Email:           user.Email,
		Name:            user.Name,
		ProfilePic:      user.ProfilePic,
		Validated:       user.Validated,
		Age:             uint32(user.Age),
		Gender:          uint32(user.Gender),
		CountryOrigin:   user.CountryOrigin,
		Admin:           user.Admin,
		SuperAdmin:      user.SuperAdmin,
		LoginLengthTime: user.LoginLengthTime,
		ValidationCode:  user.ValidationCode,
		Code:            user.Code,
		CodeExpire:      timestamppb.New(user.CodeExpire),
	}, nil
}

func (s *UserServer) Update(ctx context.Context, req *pb.UserRequest) (*pb.UserIDResponse, error) {

	user := model.User{
		Email:           req.Email,
		Name:            req.Name,
		Password:        req.Password,
		Age:             req.Age,
		Gender:          req.Gender,
		CountryOrigin:   req.CountryOrigin,
		ProfilePic:      req.ProfilePic,
		LoginLengthTime: req.LoginLengthTime,
		Validated:       false,
		Admin:           false,
		SuperAdmin:      false,
		ValidationCode:  "1",
		Code:            "1",
		CodeExpire:      time.Time{},
	}

	// s.Log.Info(fmt.Sprintf("%v\n", user))

	validate := validator.New(validator.WithRequiredStructEnabled())
	err := validate.Struct(user)
	if err != nil {
		s.Log.Error(err.Error())
		return &pb.UserIDResponse{}, err
	}

	err = s.Svc.Update(ctx, uint(req.Id), user)
	if err != nil {
		s.Log.Error(err.Error())
		return &pb.UserIDResponse{}, err

	}

	return &pb.UserIDResponse{
		Id: uint32(userUpdated.ID),
	}, nil
}

func (s *UserServer) Delete(ctx context.Context, req *pb.UserIDRequest) (*pb.UserIDResponse, error) {
	// Implementation
	return &pb.UserIDResponse{}, nil
}

func (s *UserServer) List(ctx context.Context, _ *emptypb.Empty) (*pb.ListUsersResponse, error) {
	// Implementation
	return &pb.ListUsersResponse{}, nil
}
