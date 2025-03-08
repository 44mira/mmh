FROM golang:1.24.1-alpine

WORKDIR /usr/src/app

COPY go.* ./

COPY . .
RUN go build -v -o ./app .
