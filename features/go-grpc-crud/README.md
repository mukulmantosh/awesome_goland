# Go gRPC CRUD Application

A simple Go gRPC application that demonstrates CRUD (Create, Read, Update, Delete) operations for a User service.

## Project Structure

```
go-grpc-crud/
├── proto/              # Protocol Buffers definitions
│   ├── user.proto      # User service definition
│   ├── user.pb.go      # Generated Go code for messages
│   └── user_grpc.pb.go # Generated Go code for service
├── server/             # Server implementation
│   ├── main.go         # Server entry point
│   ├── service/        # Service implementations
│   │   └── user_service.go
│   └── store/          # Data storage
│       └── user_store.go
└── client/             # Client for testing
    └── main.go         # Client implementation
```

## Requirements

- Go 1.25 or later
- Protocol Buffers compiler (protoc)
- gRPC tools for Go

## Installation

1. Clone the repository:
```bash
git clone https://github.com/yourusername/go-grpc-crud.git
cd go-grpc-crud
```

2. Install dependencies:
```bash
go mod tidy
```

## Running the Application

### Start the Server

```bash
go run server/main.go
```

The server will start on port 50051.

### Run the Client

In a separate terminal:

```bash
go run client/main.go
```

The client will:
1. Create two users
2. List all users
3. Get a specific user
4. Update a user
5. List users after update
6. Delete a user
7. List users after deletion

## API Description

The User service provides the following operations:

- **CreateUser**: Create a new user with name, email, and age
- **GetUser**: Retrieve a user by ID
- **ListUsers**: List all users
- **UpdateUser**: Update an existing user's information
- **DeleteUser**: Delete a user by ID

## Generating Protocol Buffers Code

If you make changes to the `.proto` files, regenerate the Go code with:

```bash
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/user.proto
```

## Implementation Details

- The application uses an in-memory store for simplicity
- UUIDs are used for user IDs
- Basic validation is implemented for all operations
- The server includes graceful shutdown handling
