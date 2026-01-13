package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "grpc-ping-pong/proto"
)

func main() {
	
	conn, err := grpc.Dial("server:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("No se pudo conectar: %v", err)
	}
	defer conn.Close()
	c := pb.NewPingServiceClient(conn)

	
	for {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		r, err := c.Ping(ctx, &pb.PingRequest{Message: "Ping"})
		cancel()
		if err != nil {
			log.Printf("Esperando al servidor... error: %v", err)
			time.Sleep(2 * time.Second)
			continue
		}
		log.Printf("Respuesta del Servidor: %s", r.GetMessage())
		time.Sleep(5 * time.Second)
	}
}