FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o client-exe ./client/main.go

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/client-exe .
CMD ["./client-exe"]