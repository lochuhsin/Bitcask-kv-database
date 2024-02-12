FROM golang:1.21

WORKDIR /app

COPY . .

RUN make init-network && make init

RUN make build

CMD ["./app"]

EXPOSE 8080

EXPOSE 9090