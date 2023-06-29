# Use the official Golang image as the base image
FROM golang:1.20.5-alpine AS build

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files to the working directory
COPY go.mod go.sum ./

# Download and cache Go dependencies
RUN go mod download

# Copy the entire project to the working directory
COPY . .

# Build the Go application
RUN go build -o /app/main ./cmd/main.go

# Create a new lightweight image with the Go binary
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the Go binary from the previous build stage
COPY --from=build /app/main .

# Expose the port that the application listens on
EXPOSE 8080

# Set the entry point for the container
ENTRYPOINT ["./main"]
