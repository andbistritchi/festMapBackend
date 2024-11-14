# Use the official Golang image for ARM64 architecture
FROM golang:1.23 AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN go build -o backend .

# Start a new stage from scratch
FROM golang:1.23

# Set the Current Working Directory inside the container
WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY homePage.json /app/homePage.json
COPY --from=builder /app/backend .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./backend"]
