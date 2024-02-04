FROM golang:1.21

WORKDIR /app

COPY . .

RUN make init-discovery

RUN make build-discovery

CMD ["./app.discovery"]

EXPOSE 8765
