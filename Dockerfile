# syntax=docker/dockerfile:1

FROM golang:latest

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download
RUN go mod tidy

COPY *.go ./

RUN go get github.com/go-sql-driver/mysql
RUN go get github.com/gorilla/mux

RUN go build -o /docker-fun

EXPOSE 80

CMD [ "/docker-fun" ]