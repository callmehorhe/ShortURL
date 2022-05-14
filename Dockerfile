#build
FROM golang:1.18-alpine3.15 AS builder
WORKDIR /app
COPY . .
RUN go build -o main ./api/server/main.go

#run
FROM alpine:3.15
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/configs/. /app/configs

EXPOSE 8080
CMD ["/app/main"]