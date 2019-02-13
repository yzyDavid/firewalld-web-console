package main

import (
        "context"
	service "github.com/yzyDavid/firewalld-web-console/proto"
)

type server struct{}

func (s *server) GetStatus(ctx context.Context, in *service.StatusRequest) (*service.StatusResponse, error) {
	status := State()
	return &service.StatusResponse{Version: 1, BasicInfo: status}, nil
}
