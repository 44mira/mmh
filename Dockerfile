FROM golang:1.24.1-alpine

WORKDIR /usr/src/app

RUN curl -sL https://github.com/tailwindlabs/tailwindcss/releases/latest/download/tailwindcss-linux-x64 -o tailwindcss && chmod +x tailwindcss

RUN go install github.com/air-verse/air@latest

COPY go.mod go.sum ./
RUN go mod download

COPY . .

EXPOSE 8000
