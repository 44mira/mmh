FROM golang:1.24.1-alpine

WORKDIR /usr/src/app

RUN go install github.com/air-verse/air@latest

COPY go.mod go.sum ./
RUN go mod download

COPY . .

EXPOSE 8000
