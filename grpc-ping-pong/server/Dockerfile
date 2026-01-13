FROM golang:1.24-alpine AS builder
WORKDIR /app
# Copiamos todo el proyecto (incluyendo la carpeta proto y go.mod)
COPY . .
RUN go mod download
# Entramos a la carpeta del servidor para compilar
WORKDIR /app/server
RUN go build -o /server-exe main.go

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /server-exe .
EXPOSE 50051
CMD ["./server-exe"]