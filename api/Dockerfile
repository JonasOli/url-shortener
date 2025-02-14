# Use Golang base image
FROM golang:1.24 AS builder

WORKDIR /app

# Copy everything
COPY . .

# Download dependencies
RUN go mod tidy

# Build the Go application
RUN go build -o url-shortener main.go

# Use a minimal base image
FROM alpine:latest

WORKDIR /root/

# Copy built binary from builder stage
COPY --from=builder /app/url-shortener .

# Expose port 3000
EXPOSE 3000

# Run the application
CMD ["./url-shortener"]
