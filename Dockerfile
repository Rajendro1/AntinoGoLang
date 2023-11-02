# Start from the official Golang image to build your application
FROM golang:1.21 as builder

# Set the current working directory inside the container
WORKDIR /app

# Copy go mod and sum files to download dependencies
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Start a new stage from scratch
FROM alpine:latest

# Set the current working directory
WORKDIR /root/

# Copy the output from the builder stage
COPY --from=builder /app/main .

# Expose the application on port 8080
EXPOSE 8080

# Command to run the application
CMD ["./main"]
