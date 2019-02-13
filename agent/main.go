package main

import (
	"log"
	"net"

	service "github.com/yzyDavid/firewalld-web-console/proto"
	grpc "google.golang.org/grpc"
)

const (
	port = ":32116"
)

var _ service.StatusResponse

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
