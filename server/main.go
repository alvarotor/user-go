package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/alvarotor/user-go/user"
)

type server struct {
	pb.UnimplementedUserServer
}

func (s *server) GetUser(_ context.Context, in *pb.MailUserRequest) (*pb.UserResponse, error) {
	log.Printf("Received: %v", in.GetEmail())
	return &pb.UserResponse{Email: "User " + in.GetEmail()}, nil
}

func main() {
	port := 50051

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterUserServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
