package main

import (
	"context"
	"log"
	"net"
	
	grpc "google.golang.org/grpc"
	service "github.com/yzyDavid/firewalld-web-console/proto"
)

const (
	port = ":32116"
)

var _ service.StatusResponse

type server struct {}

func (s *server) GetStatus(ctx context.Context, in *service.StatusRequest) (*service.StatusResponse, error) {
	return &service.StatusResponse{}, nil
}

func main() {
	log.Print("agent started.")
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("listening %s", port)
	s := grpc.NewServer()
	service.RegisterAgentServer(s, &server{})
	if err := s.Serve(lis); err != nil {
	        log.Fatalf("failed to serve: %v", err)
	}
}
