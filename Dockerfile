# Use the official Golang image as a build environment
FROM golang:1.17-alpine AS builder

# Set the current working directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN go build -o main .

# Use a minimal image to run the application
FROM alpine:latest

# Set the current working directory inside the container
WORKDIR /root/

# Copy the built binary from the builder container
COPY --from=builder /app/main .

# Expose port 8000 to the outside world
EXPOSE 8000

# Command to run the executable
CMD ["./main"]
