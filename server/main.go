package main

import (
	"fmt"
	"log"
	"log/slog"
	"net"
	"os"
	"strconv"

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
			log.Fatal("Error loading .env file")
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
		Svc:        svc,
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

func checkEnvVarsConf(conf *model.Config) {
	checkEnvVar("PROJECT_PORT")
	checkEnvVar("POSTGRES_HOST")
	checkEnvVar("POSTGRES_DB")
	checkEnvVar("POSTGRES_USER")
	checkEnvVar("POSTGRES_PASSWORD")
	checkEnvVar("POSTGRES_PORT")
	checkEnvVar("RandomStringValidation")
	checkEnvVar("SizeRandomStringValidation")
	checkEnvVar("Issuer")

	PROJECT_PORT, err := strconv.Atoi(os.Getenv("PROJECT_PORT"))
	if err != nil {
		log.Fatalf(`Missing PROJECT_PORT env var`)
	}
	conf.PROJECT_PORT = PROJECT_PORT
	conf.POSTGRES_HOST = os.Getenv("POSTGRES_HOST")
	conf.POSTGRES_DB = os.Getenv("POSTGRES_DB")
	conf.POSTGRES_USER = os.Getenv("POSTGRES_USER")
	conf.POSTGRES_PASSWORD = os.Getenv("POSTGRES_PASSWORD")
	conf.POSTGRES_PORT = os.Getenv("POSTGRES_PORT")
	conf.RandomStringValidation = os.Getenv("RandomStringValidation")
	SizeRandomStringValidation, err := strconv.Atoi(os.Getenv("SizeRandomStringValidation"))
	if err != nil {
		log.Fatalf(`Missing SizeRandomStringValidation env var`)
	}
	conf.SizeRandomStringValidation = SizeRandomStringValidation
	conf.Issuer = os.Getenv("Issuer")
}

func checkEnvVar(envVar string) {
	if len(os.Getenv(envVar)) == 0 {
		log.Fatalf(`Missing %s env var`, envVar)
	}
}
