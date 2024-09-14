FROM golang:1.23.1-alpine AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go run github.com/swaggo/swag/cmd/swag@v1.16.3 init -g cmd/api/main.go -o ./docs

WORKDIR /app/cmd/api

RUN go build -o /main .

FROM alpine:latest

RUN apk add --no-cache netcat-openbsd

WORKDIR /root/

COPY --from=builder /main .
COPY --from=builder /app/docs ./docs
COPY --from=builder /app/migrations ./migrations

EXPOSE 8080

CMD ["./main"]