FROM golang:1.24.3-alpine3.21 AS builder

WORKDIR /app
COPY . .
RUN go mod download
RUN go install github.com/pressly/goose/v3/cmd/goose@v3.24.3
RUN go build -o main ./cmd/server

FROM alpine:3.21
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /go/bin/goose /usr/local/bin/goose
COPY --from=builder /app/migrations ./migrations
# COPY .env .
RUN chmod +x /app/main

EXPOSE 8080
CMD ["/app/main"]