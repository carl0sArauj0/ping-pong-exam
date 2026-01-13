FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
WORKDIR /app/server
RUN go build -o /server-exe main.go

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /server-exe .
EXPOSE 50051
CMD ["./server-exe"]