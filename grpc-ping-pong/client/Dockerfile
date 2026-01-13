FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
# Entramos a la carpeta del cliente para compilar
WORKDIR /app/client
RUN go build -o /client-exe main.go

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /client-exe .
CMD ["./client-exe"]