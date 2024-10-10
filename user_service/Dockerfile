# Use the official Golang image as the base image
FROM golang:1.23

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire project directory into the container
COPY . .

# Set the working directory to the user_service directory
WORKDIR /app/user_service

# Build the Go application
RUN go build -o user_service user_handler.go

# Command to run the application
CMD ["./user_service"]