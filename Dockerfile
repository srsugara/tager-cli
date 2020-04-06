# app golang build
FROM golang:1.13-alpine as builder

RUN apk update && apk add --no-cache git bash wget curl && mkdir app

WORKDIR /app

COPY . .

RUN go mod download && go build -o tager-cli

ENTRYPOINT ["./tager-cli"]