# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Common Development Commands

### Running the Application
```bash
# Run the server
go run server/*.go

# Run the client
go run client/client.go
```

### Testing
```bash
# Run tests
go test ./...

# Run specific test
go test ./server -v
```

### Protocol Buffers
```bash
# Generate protobuf files (when .proto files change)
make protoc

# Or manually:
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative server/user-pb/user.proto
```

### Database
```bash
# Ensure PostgreSQL is running with proper environment variables
# Check server/config.go for required environment variables
```

## Architecture Overview

This is a **gRPC-based user management system** with the following key components:

### Core Structure
- **Server (`/server`)**: Main gRPC server implementation
  - `main.go`: Entry point and server setup
  - `config.go`: Environment variable configuration and validation
  - `user-pb/`: Protocol buffer definitions and generated Go code
  
- **Client (`/client`)**: Example gRPC client implementation
  - `client.go`: Main client with usage examples
  - `health/`: Health check client utilities

### Service Layers
1. **Controllers (`/server/controllers`)**: Handle gRPC requests and responses
2. **Server (`/server/server`)**: Core business logic implementation
3. **Services (`/server/services`)**: Utility services (validation, token handling, etc.)
4. **Models (`/server/models`)**: Data structures and database models
5. **Database (`/server/db`)**: Database connection and operations using GORM

### Key Features
- **Authentication**: JWT-based with access and refresh tokens
- **User Management**: Full CRUD operations with role-based access (Admin/SuperAdmin)
- **Email Validation**: Code-based email verification system
- **Session Management**: Browser fingerprinting and token management
- **Health Checks**: gRPC health checking implementation
- **Database**: PostgreSQL with GORM ORM

### Environment Configuration
The application requires extensive environment configuration (see `server/config.go:11-26`). All variables are validated at startup and will cause the application to fail if missing.

### Protocol Buffers
- Service definition: `server/user-pb/user.proto`
- Generated files: `user.pb.go` and `user_grpc.pb.go`
- Service provides 13 main endpoints including CRUD operations, authentication, and health checks

### Database Schema
Uses GORM with PostgreSQL. Models are defined in `/server/models/` and include:
- User model with authentication fields
- Token management structures
- Configuration models
- Custom error types

## Development Notes

### Dependencies
- Go 1.23.2+
- PostgreSQL database
- Protocol Buffers compiler (protoc)
- Key libraries: gRPC, GORM, JWT, validator

### Testing
- Tests are located in `server/main_test.go`
- Run with standard Go testing tools
- Currently minimal test coverage (single example test)

### Security
- JWT tokens with configurable expiration times
- Environment-based configuration for secrets
- Role-based access control (Admin/SuperAdmin)
- Browser fingerprinting for session management