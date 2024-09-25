package main

import (
	"emailcheck/service"
	"log"
	"net"

	"github.com/justIGreK/emailcheck/go/emailcheck"
	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen on port 50051: %v", err)
	}

	grpcServer := grpc.NewServer()
	emailServiceServer := &service.EmailServiceServer{}

	emailcheck.RegisterEmailServiceServer(grpcServer, emailServiceServer)

	log.Println("Starting gRPC server on port 50051...")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}