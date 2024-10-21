package server

import (
	"context"

	"github.com/alvarotor/user-go/server/controller"
	"github.com/alvarotor/user-go/server/model"
	"github.com/alvarotor/user-go/server/service"
	pb "github.com/alvarotor/user-go/server/user-pb"
	"google.golang.org/protobuf/types/known/emptypb"
)

type UserServer struct {
	pb.UnimplementedUserServer
	// users map[uint32]*pb.UserResponse
	Controller controller.IControllerUser
	Svc        service.IUserService
}

func NewServer(
	controller controller.IControllerUser,
	svc service.IUserService,
) *UserServer {
	return &UserServer{
		Controller: controller,
		Svc:        svc,
	}
}

func (s *UserServer) Create(ctx context.Context, req *pb.UserRequest) (*pb.UserIDRequest, error) {
	user := model.User{
		Email:           req.Email,
		Name:            req.Name,
		Age:             req.Age,
		Gender:          req.Gender,
		CountryOrigin:   req.CountryOrigin,
		ProfilePic:      req.ProfilePic,
		LoginLengthTime: req.LoginLengthTime,
	}

	userCreated, err := s.Svc.Create(ctx, user)
	if err != nil {
		return &pb.UserIDRequest{}, err
	}
	return &pb.UserIDRequest{
		Id: uint32(userCreated.ID),
	}, nil
}

func (s *UserServer) Get(ctx context.Context, req *pb.UserIDRequest) (*pb.UserResponse, error) {
	// Implementation
	return &pb.UserResponse{}, nil
}

func (s *UserServer) Update(ctx context.Context, req *pb.UserRequest) (*pb.UserIDResponse, error) {
	// Implementation
	return &pb.UserIDResponse{}, nil
}

func (s *UserServer) Delete(ctx context.Context, req *pb.UserIDRequest) (*pb.UserIDResponse, error) {
	// Implementation
	return &pb.UserIDResponse{}, nil
}

func (s *UserServer) List(ctx context.Context, _ *emptypb.Empty) (*pb.ListUsersResponse, error) {
	// Implementation
	return &pb.ListUsersResponse{}, nil
}
