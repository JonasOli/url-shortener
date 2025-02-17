# Use Golang base image
FROM golang:1.24 AS builder

WORKDIR /app

# Copy everything
COPY . .

# Download dependencies
RUN go mod tidy

# Build the Go application
RUN go build -o /url-shortener main.go

# Expose port 8080
EXPOSE 8080

# Run the application
CMD ["/url-shortener"]
