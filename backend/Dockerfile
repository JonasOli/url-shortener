# Build stage
FROM golang:1.24 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o url-shortener .

# Production stage
FROM alpine:latest

WORKDIR /app

COPY .env.prod /app/.env

COPY --from=builder /app/url-shortener .

EXPOSE 8080

CMD ["./url-shortener"]
