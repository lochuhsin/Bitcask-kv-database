FROM golang:1.21

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o app cmd/main.go

CMD ["./app"]

EXPOSE 8080

EXPOSE 9090