package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	pb "grpc-ping-pong/proto"
)

type server struct {
	pb.UnimplementedPingServiceServer
}

func (s *server) Ping(ctx context.Context, in *pb.PingRequest) (*pb.PingResponse, error) {
	log.Printf("Recibido: %s", in.GetMessage())
	if in.GetMessage() == "Ping" {
		return &pb.PingResponse{Message: "Pong"}, nil
	}
	return &pb.PingResponse{Message: "Error: Mensaje desconocido"}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Error al escuchar: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterPingServiceServer(s, &server{})
	log.Println("Servidor gRPC iniciado en el puerto :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Error al servir: %v", err)
	}
}