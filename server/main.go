package main

import (
	"fmt"
	"log"
	"log/slog"
	"net"
	"os"

	"google.golang.org/grpc"

	"github.com/alvarotor/user-go/server/controller"
	"github.com/alvarotor/user-go/server/db"
	"github.com/alvarotor/user-go/server/model"
	"github.com/alvarotor/user-go/server/server"
	"github.com/alvarotor/user-go/server/service"
	pb "github.com/alvarotor/user-go/server/user-pb"
	"github.com/joho/godotenv"
)

func main() {
	var conf model.Config

	if conf.IsLocalENV() {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("error loading .env file")
		}
	}
	checkEnvVarsConf(&conf)

	opts := &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}
	l := slog.New(slog.NewJSONHandler(os.Stdout, opts))

	dbUser := db.GetDB_PG(&conf, l)

	svc := service.NewUserService(dbUser)
	con := controller.NewUserController(l, svc, &conf)
	userServer := server.UserServer{
		Controller: con,
		Log:        l,
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", conf.PROJECT_PORT))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterUserServer(s, &userServer)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
