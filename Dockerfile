FROM golang:alpine3.16 AS initial
WORKDIR /keeper
ENV CGO_ENABLED=0
EXPOSE 8080
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .

FROM initial AS builder
WORKDIR /keeper/cmd
RUN go build -o /keeper/keeper

FROM builder
WORKDIR /keeper
CMD ["/bin/sh", "-c", "./keeper"]
