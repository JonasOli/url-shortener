FROM golang:1.24

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o /url-shortener main.go

EXPOSE 8000

CMD ["/url-shortener"]
