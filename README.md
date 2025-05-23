# Deepker Backend

The Deepker Backend is a high-performance server-side application developed in Go (Golang). It serves as the backbone of the Deepker platform, providing robust and scalable APIs for biometric data processing and analysis. Leveraging modern technologies like the Gin web framework and GORM ORM, it ensures efficient handling of HTTP requests and seamless interactions with a PostgreSQL database.

## Requirements

- Go 1.18 or higher
- Docker
- **PostgreSQL**

## Project Setup

### Step 1: Clone the Repository

```sh
git clone https://github.com/DeepKer-Org/deepker-backend.git
cd deepker-backend
```

### Step 2: Configure the `.env` File

Create a `.env` file in the root of the project with the following content:

```env
DB_USER=postgres
DB_PASSWORD=root
DB_HOST=localhost
DB_PORT=5432
DB_NAME=deepker
JWT_SECRET_KEY=mySecretKey
```

### Step 3: Install Dependencies

Run the following command to install project dependencies:

```sh
go mod tidy
```

This command downloads the necessary Go modules specified in the `go.mod` file.

### Step 4: Initialize the database and cache manager

Use Docker to start a PostgreSQL container with the following command:

```sh
docker run --name postgres-deepker -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=root -e POSTGRES_DB=deepker -p 5432:5432 -d postgres
```

Use Docker to start a Redis container with the following command:

```sh
docker run -d --name redis-local \
  -p 6379:6379 \
  -e REDIS_PASSWORD=your_secure_password \
  redis \
  redis-server --requirepass your_secure_password
```

<!-- ### Step 5: Generate Swagger Documentation (Optional)

Install Swagger if you haven't already:

```sh
go install github.com/swaggo/swag/cmd/swag@latest
```

Generate the Swagger documentation:

```sh
swag init
``` -->

### Step 5: Run the Project

Run the project:

```sh
go run cmd/server/main.go
```
<!-- 
### Step 8: View the API Documentation (If Generated)

Open your browser and visit the following link to view the API documentation:

[http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html) -->

## Default Credentials

- `username`: 44556677 | 55667788 | 66778899
- `password`: hashed_password1!
- `role`: admin

- `username`: doctor{1-3}@example.com
- `password`: hashed_password1!
- `role`: doctor

## Project Structure

The project follows a standard structure for Go applications, with the following main directories:

- `controllers`: Contains the controllers to handle HTTP requests.
- `models`: Contains the data models.
- `repositories`: Contains the repositories for database interactions.
- `services`: Contains the business logic.
- `routes`: Contains the route configuration.
- `config`: Contains the database configuration and environment variable loading.
- `migrations`: Contains SQL files for database migrations.
- `cmd/migrate`: Contains the migration script to run the migrations.