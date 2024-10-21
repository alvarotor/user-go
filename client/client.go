package main

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/alvarotor/user-go/server/user-pb"
)

func main() {
	addr := "localhost:50051"

	// Set up a connection to the server.
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewUserClient(conn)

	// Create a new user
	user := &pb.UserRequest{
		Email:    "test@example.com",
		Password: "password123",
		Name:     "Test User",
		// Set other fields as needed
	}
	r, err := c.Create(context.Background(), user)
	if err != nil {
		log.Fatalf("could not create user: %v", err)
	}
	log.Printf("Created user with ID: %d", r.Id)

	// Implement other operations (Get, Update, Delete, List) as needed
}
