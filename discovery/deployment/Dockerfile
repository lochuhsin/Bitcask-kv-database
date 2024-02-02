FROM golang:1.21

WORKDIR /app

COPY . .

RUN go mod tidy

RUN make build

CMD ["./app"]

EXPOSE 8765
