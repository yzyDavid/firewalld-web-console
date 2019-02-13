package agent_client

import (
        _ "context"
	"log"

	service "github.com/yzyDavid/firewalld-web-console/proto"
	grpc "google.golang.org/grpc"
)

const (
        defaultPort = "32116"
)

type AgentSpec struct {
        Address string
	Port string
}

func NewSpec(address string, port string) *AgentSpec {
        if port == "" {
	        port = defaultPort
	}
	spec := new(AgentSpec)
	spec.Address = address
	spec.Port = port
	return spec
}

func New(address string, port string) service.AgentClient {
        if port == "" {
	        port = defaultPort
	}
	addr := address + ":" + port
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
	        log.Printf("did not connect: %s %v", addr, err)
		return nil
	}
	return service.NewAgentClient(conn)
}
