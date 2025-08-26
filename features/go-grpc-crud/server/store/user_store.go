package store

import (
	"errors"
	"sync"
	"time"

	"github.com/google/uuid"
	pb "github.com/mukulmantosh/go-grpc-crud/proto"
)

// UserStore provides an in-memory store for users
type UserStore struct {
	mu    sync.RWMutex
	users map[string]*pb.User
}

// NewUserStore creates a new UserStore
func NewUserStore() *UserStore {
	return &UserStore{
		users: make(map[string]*pb.User),
	}
}

// Create adds a new user to the store
func (s *UserStore) Create(name, email string, age int32) (*pb.User, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Check if email already exists
	for _, user := range s.users {
		if user.Email == email {
			return nil, errors.New("user with this email already exists")
		}
	}

	// Generate a new UUID for the user
	id := uuid.New().String()
	now := time.Now().Format(time.RFC3339)

	user := &pb.User{
		Id:        id,
		Name:      name,
		Email:     email,
		Age:       age,
		CreatedAt: now,
		UpdatedAt: now,
	}

	s.users[id] = user
	return user, nil
}

// Get retrieves a user by ID
func (s *UserStore) Get(id string) (*pb.User, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	user, exists := s.users[id]
	if !exists {
		return nil, errors.New("user not found")
	}

	return user, nil
}

// List returns all users
func (s *UserStore) List() []*pb.User {
	s.mu.RLock()
	defer s.mu.RUnlock()

	users := make([]*pb.User, 0, len(s.users))
	for _, user := range s.users {
		users = append(users, user)
	}

	return users
}

// Update modifies an existing user
func (s *UserStore) Update(id, name, email string, age int32) (*pb.User, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	user, exists := s.users[id]
	if !exists {
		return nil, errors.New("user not found")
	}

	// Check if the new email conflicts with another user
	if email != user.Email {
		for _, u := range s.users {
			if u.Id != id && u.Email == email {
				return nil, errors.New("email already in use by another user")
			}
		}
	}

	// Update user fields
	user.Name = name
	user.Email = email
	user.Age = age
	user.UpdatedAt = time.Now().Format(time.RFC3339)

	return user, nil
}

// Delete removes a user from the store
func (s *UserStore) Delete(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.users[id]; !exists {
		return errors.New("user not found")
	}

	delete(s.users, id)
	return nil
}
