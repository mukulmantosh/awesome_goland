package service

import (
	"context"

	pb "github.com/mukulmantosh/go-grpc-crud/proto"
	"github.com/mukulmantosh/go-grpc-crud/server/store"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// UserService implements the gRPC UserService
type UserService struct {
	pb.UnimplementedUserServiceServer
	store *store.UserStore
}

// NewUserService creates a new UserService
func NewUserService(store *store.UserStore) *UserService {
	return &UserService{
		store: store,
	}
}

// CreateUser implements the CreateUser RPC method
func (s *UserService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.UserResponse, error) {
	// Validate request
	if req.Name == "" {
		return nil, status.Error(codes.InvalidArgument, "name is required")
	}
	if req.Email == "" {
		return nil, status.Error(codes.InvalidArgument, "email is required")
	}

	// Create user in store
	user, err := s.store.Create(req.Name, req.Email, req.Age)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.UserResponse{User: user}, nil
}

// GetUser implements the GetUser RPC method
func (s *UserService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.UserResponse, error) {
	// Validate request
	if req.Id == "" {
		return nil, status.Error(codes.InvalidArgument, "user id is required")
	}

	// Get user from store
	user, err := s.store.Get(req.Id)
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return &pb.UserResponse{User: user}, nil
}

// ListUsers implements the ListUsers RPC method
func (s *UserService) ListUsers(ctx context.Context, req *pb.ListUsersRequest) (*pb.ListUsersResponse, error) {
	// Get all users from store
	users := s.store.List()
	return &pb.ListUsersResponse{Users: users}, nil
}

// UpdateUser implements the UpdateUser RPC method
func (s *UserService) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UserResponse, error) {
	// Validate request
	if req.Id == "" {
		return nil, status.Error(codes.InvalidArgument, "user id is required")
	}
	if req.Name == "" {
		return nil, status.Error(codes.InvalidArgument, "name is required")
	}
	if req.Email == "" {
		return nil, status.Error(codes.InvalidArgument, "email is required")
	}

	// Update user in store
	user, err := s.store.Update(req.Id, req.Name, req.Email, req.Age)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.UserResponse{User: user}, nil
}

// DeleteUser implements the DeleteUser RPC method
func (s *UserService) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	// Validate request
	if req.Id == "" {
		return nil, status.Error(codes.InvalidArgument, "user id is required")
	}

	// Delete user from store
	err := s.store.Delete(req.Id)
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return &pb.DeleteUserResponse{Success: true}, nil
}
