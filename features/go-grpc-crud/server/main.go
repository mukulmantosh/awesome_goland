package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	pb "github.com/mukulmantosh/go-grpc-crud/proto"
	"github.com/mukulmantosh/go-grpc-crud/server/service"
	"github.com/mukulmantosh/go-grpc-crud/server/store"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

func main() {
	// Create a TCP listener
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Create a new gRPC server
	grpcServer := grpc.NewServer()

	// Create a new user store
	userStore := store.NewUserStore()

	// Create a new user service
	userService := service.NewUserService(userStore)

	// Register the user service with the gRPC server
	pb.RegisterUserServiceServer(grpcServer, userService)

	// Register reflection service on gRPC server
	// This allows clients to query the server for its supported methods
	reflection.Register(grpcServer)

	// Handle graceful shutdown
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)
		<-c
		fmt.Println("\nShutting down gRPC server...")
		grpcServer.GracefulStop()
	}()

	// Start the gRPC server
	fmt.Printf("Starting gRPC server on %s...\n", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
