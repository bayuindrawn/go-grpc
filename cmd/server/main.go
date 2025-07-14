package main

import (
	"go-grpc/config"
	"go-grpc/internal/employee"
	pb "go-grpc/proto/employee"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	db := config.ConnectDB()
	repo := employee.NewRepository(db)
	service := employee.NewService(repo)
	handler := employee.NewHandler(service)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterEmployeeServiceServer(grpcServer, handler)
	reflection.Register(grpcServer)

	log.Println("gRPC server running on port 50051...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
