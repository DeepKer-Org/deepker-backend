# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Development Commands

### Running the Application
- **Local development**: `go run cmd/server/main.go`
- **Build application**: `go build -o deepker-app cmd/server/main.go`
- **Docker build**: `docker build -t deepker-backend .`
- **Docker Compose**: `docker-compose up` (includes PostgreSQL database)

### Dependencies
- **Install/update dependencies**: `go mod tidy`
- **Download dependencies**: `go mod download`

### Database Setup
```bash
# PostgreSQL container
docker run --name postgres-deepker -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=root -e POSTGRES_DB=deepker -p 5432:5432 -d postgres

# Redis container (optional - can be disabled)
docker run -d --name redis-local -p 6379:6379 -e REDIS_PASSWORD=your_secure_password redis redis-server --requirepass your_secure_password
```

### Cache Configuration
```bash
# Enable Redis cache (default)
CACHE_ENABLED=true
REDIS_ENABLED=true

# Disable cache for local development
CACHE_ENABLED=false
REDIS_ENABLED=false

# Configure cache TTL
CACHE_DEFAULT_TTL=5m
```

### Documentation
- **Swagger API docs**: Available at `http://localhost:8080/swagger/index.html` when server is running
- **Generate Swagger docs**: `swag init` (requires `go install github.com/swaggo/swag/cmd/swag@latest`)

### Testing
**Important**: This project currently has NO testing infrastructure. There are no test files, testing dependencies, or testing commands available.

## Architecture

### High-Level Structure
This is a Go REST API backend using clean architecture patterns with the following layers:

- **Controllers** (`controller/`): HTTP request handlers and response formatting
- **Services** (`service/`): Business logic layer
- **Repositories** (`repository/`): Data access layer (GORM-based)
- **Models** (`models/`): Database entities and DTOs
- **Routes** (`routes/`): Route definitions and middleware configuration

### Key Frameworks & Libraries
- **Web Framework**: Gin (github.com/gin-gonic/gin)
- **ORM**: GORM with PostgreSQL driver
- **Authentication**: JWT tokens (golang-jwt/jwt/v5)
- **Caching**: Redis (go-redis/redis/v8)
- **Documentation**: Swagger/OpenAPI
- **Password Hashing**: golang.org/x/crypto

### Database
- **Primary Database**: PostgreSQL
- **Cache**: Redis
- **Migrations**: Located in `migrations/postgres/` (manual SQL files)
- **Connection**: Configured via environment variables in `.env` file

### Environment Configuration
Required `.env` file variables:
```env
DB_USER=postgres
DB_PASSWORD=root
DB_HOST=localhost
DB_PORT=5432
DB_NAME=deepker
JWT_SECRET_KEY=mySecretKey
ALLOWED_ORIGIN=http://localhost:3000
```

### API Resources
The API provides CRUD operations for:
- **Authorization**: User authentication and management
- **Doctors**: Doctor profiles and specializations
- **Patients**: Patient records and data
- **Biometrics**: Biometric data collection
- **Computer Diagnostics**: AI-generated diagnostic results
- **Monitoring Devices**: IoT device management
- **Alerts**: Notification system
- **Comorbidities**: Medical condition tracking
- **Medications**: Drug and prescription management
- **Phones**: Device registration for notifications

### Authentication & Authorization
- JWT-based authentication
- Role-based access control (Admin, Doctor, Tester roles)
- Middleware-protected endpoints with role validation
- Password change functionality for doctors

### Caching Strategy
- Redis-based caching with 5-minute default TTL
- Cache manager wrapper for consistent cache operations
- Applied to service layer for database query optimization

### CORS Configuration
- Configurable allowed origins via `ALLOWED_ORIGIN` environment variable
- Special handling for React Native apps via `X-App-Origin` header
- Supports credentials and common HTTP methods

## Development Workflow

### Adding New Features
1. Create model in `models/` (include DTO if needed)
2. Implement repository in `repository/`
3. Add business logic in `service/`
4. Create controller in `controller/`
5. Register routes in `routes/routes.go`
6. Update database migrations if schema changes required

### Code Patterns
- All entities extend `BaseModel` with common fields (ID, CreatedAt, UpdatedAt, DeletedAt)
- DTOs use consistent naming: `{Resource}DTO`, `{Resource}CreateDTO`, `{Resource}UpdateDTO`
- Controllers follow REST conventions with standard CRUD methods
- Services handle business logic and cache operations
- Repositories use GORM for database operations with soft deletes

### Code Improvements (feat/code-refactoring-improvements branch)
This branch contains significant improvements for maintainability, scalability, and reduced duplication:

- **Centralized Configuration**: All app config in `config/app_config.go` with environment variable validation
- **Configurable Redis Cache**: Easy enable/disable via `CACHE_ENABLED` and `REDIS_ENABLED`
- **Generic Base Services**: `service/base_service.go` eliminates ~75% of CRUD duplication
- **Generic Controllers**: `controller/base_controller.go` reduces controller code by ~80%
- **Enhanced Repositories**: Additional methods for pagination, batch operations, field queries
- **Standardized Responses**: `utils/response_utils.go` and `utils/validation_utils.go`
- **Better Error Handling**: Consistent error responses and logging

#### Usage Examples
```go
// Using improved patterns
type MyService struct {
    *BaseService[Model, CreateDTO, UpdateDTO]
}

type MyController struct {
    *EnhancedController[CreateDTO, UpdateDTO]
}
```

### Default Credentials
Available user accounts for testing:
- **Admin users**: 44556677, 55667788, 66778899 (password: hashed_password1!)
- **Doctor users**: doctor1@example.com, doctor2@example.com, doctor3@example.com (password: hashed_password1!)