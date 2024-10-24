package server

import (
	"log/slog"

	"github.com/alvarotor/user-go/server/controller"
	pb "github.com/alvarotor/user-go/server/user-pb"
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
