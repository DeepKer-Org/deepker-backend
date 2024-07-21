# Biometric data backend

This is a Go project that uses the Gin framework for routing and GORM for database interaction. Swagger is used for API documentation.

## Requirements

- Go 1.18 or higher
- Docker
- MySQL

## Project Setup

### Step 1: Clone the Repository

```sh
git clone https://github.com/Kudas-AI/biometric-data-backend.git
cd biometric-data-backend
```

### Step 2: Configure the `.env` File

Create a `.env` file in the root of the project with the following content:

```env
DB_USER=root
DB_PASSWORD=root
DB_NAME=go_db
DB_HOST=localhost
DB_PORT=3306
```

### Step 3: Initialize the Database

Use Docker to start a MySQL container with the following command:

```sh
docker run -d --name mysql-container -e MYSQL_ROOT_PASSWORD=root -p 3306:3306 mysql:8.0.37 --default-authentication-plugin=mysql_native_password --bind-address=0.0.0.0
```

Connect to the MySQL container and create the `go_db` database:

```sh
docker exec -it mysql-container mysql -u root -proot -e "CREATE DATABASE go_db;"
```

### Step 4: Install Dependencies

Run the following command to install project dependencies:

```sh
go mod tidy
```

### Step 5: Generate Swagger Documentation (optional)

Install Swagger if you haven't already:

```sh
go install github.com/swaggo/swag/cmd/swag@latest
```

Generate the Swagger documentation:

```sh
swag init
```

### Step 6: Run the Project

Run the project:

```sh
go run main.go
```

### Step 7: View the API Documentation (if the documentation was generated)

Open your browser and visit the following link to view the API documentation:

[http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)

## Project Structure

The project follows a standard structure for Go applications, with the following main directories:

- `controllers`: Contains the controllers to handle HTTP requests.
- `models`: Contains the data models.
- `repositories`: Contains the repositories for database interactions.
- `services`: Contains the business logic.
- `routes`: Contains the route configuration.
- `config`: Contains the database configuration and environment variable loading.
