# app golang build
FROM golang:1.13-alpine as builder

RUN apk update && apk add --no-cache git bash wget curl && mkdir app

WORKDIR /app

COPY . .

RUN go mod download && go build -o tager-cli

# directory

## Distribution
FROM alpine

RUN apk update && apk add --no-cache git bash wget curl && mkdir app 
RUN mkdir -p /app/storage && mkdir -p /app/storage/logs

WORKDIR /app

COPY --from=builder /app/tager-cli /app
CMD /app/tager-cli

EXPOSE 6060