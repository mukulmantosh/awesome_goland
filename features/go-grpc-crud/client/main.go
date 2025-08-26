package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/mukulmantosh/go-grpc-crud/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	address = "localhost:50051"
)

func main() {
	// Set up a connection to the server
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// Create a client
	client := pb.NewUserServiceClient(conn)

	// Set timeout for API calls
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	fmt.Println("=== Creating Users ===")
	// Create users
	user1, err := client.CreateUser(ctx, &pb.CreateUserRequest{
		Name:  "John Doe",
		Email: "john@example.com",
		Age:   30,
	})
	if err != nil {
		log.Fatalf("Could not create user: %v", err)
	}
	fmt.Printf("Created User: %v\n", user1.User)

	user2, err := client.CreateUser(ctx, &pb.CreateUserRequest{
		Name:  "Jane Smith",
		Email: "jane@example.com",
		Age:   25,
	})
	if err != nil {
		log.Fatalf("Could not create user: %v", err)
	}
	fmt.Printf("Created User: %v\n", user2.User)

	fmt.Println("\n=== Listing Users ===")
	// List users
	listResp, err := client.ListUsers(ctx, &pb.ListUsersRequest{})
	if err != nil {
		log.Fatalf("Could not list users: %v", err)
	}
	for i, user := range listResp.Users {
		fmt.Printf("%d. %s (%s) - Age: %d\n", i+1, user.Name, user.Email, user.Age)
	}

	fmt.Println("\n=== Getting User ===")
	// Get a user
	getResp, err := client.GetUser(ctx, &pb.GetUserRequest{Id: user1.User.Id})
	if err != nil {
		log.Fatalf("Could not get user: %v", err)
	}
	fmt.Printf("Got User: %v\n", getResp.User)

	fmt.Println("\n=== Updating User ===")
	// Update a user
	updateResp, err := client.UpdateUser(ctx, &pb.UpdateUserRequest{
		Id:    user1.User.Id,
		Name:  "John Updated",
		Email: "john.updated@example.com",
		Age:   31,
	})
	if err != nil {
		log.Fatalf("Could not update user: %v", err)
	}
	fmt.Printf("Updated User: %v\n", updateResp.User)

	fmt.Println("\n=== Listing Users After Update ===")
	// List users again to see the update
	listResp, err = client.ListUsers(ctx, &pb.ListUsersRequest{})
	if err != nil {
		log.Fatalf("Could not list users: %v", err)
	}
	for i, user := range listResp.Users {
		fmt.Printf("%d. %s (%s) - Age: %d\n", i+1, user.Name, user.Email, user.Age)
	}

	fmt.Println("\n=== Deleting User ===")
	// Delete a user
	deleteResp, err := client.DeleteUser(ctx, &pb.DeleteUserRequest{Id: user2.User.Id})
	if err != nil {
		log.Fatalf("Could not delete user: %v", err)
	}
	fmt.Printf("Delete success: %v\n", deleteResp.Success)

	fmt.Println("\n=== Listing Users After Delete ===")
	// List users again to confirm deletion
	listResp, err = client.ListUsers(ctx, &pb.ListUsersRequest{})
	if err != nil {
		log.Fatalf("Could not list users: %v", err)
	}
	for i, user := range listResp.Users {
		fmt.Printf("%d. %s (%s) - Age: %d\n", i+1, user.Name, user.Email, user.Age)
	}
}
