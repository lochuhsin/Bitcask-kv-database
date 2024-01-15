FROM golang:1.21

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o app server/cmd/main.go

CMD ["./app"]

EXPOSE 8000