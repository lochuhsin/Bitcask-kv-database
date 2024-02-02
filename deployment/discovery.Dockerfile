FROM golang:1.21

WORKDIR /app

COPY . .

RUN make discovery-init

RUN make discovery-build

CMD ["./app.discovery"]

EXPOSE 8765
