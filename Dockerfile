# Stage 1: Build the application
FROM golang:1.23 AS builder

WORKDIR /app

# Copy and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux go build -o /deepker-app cmd/server/main.go

# Stage 2: Create the final image
FROM alpine:3.18

# Install certificates
RUN apk --no-cache add ca-certificates

# Copy the binary
COPY --from=builder /deepker-app /deepker-app

# Copy the migrations
COPY --from=builder /app/migrations/postgres /migrations/postgres

# Expose the port (adjust according to your application)
EXPOSE 8080

# Command to execute
CMD ["/deepker-app"]