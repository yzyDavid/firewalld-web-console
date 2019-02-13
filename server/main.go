package main

import (
        "net/http"
	"context"
	"time"
	"log"

	service "github.com/yzyDavid/firewalld-web-console/proto"
	client "github.com/yzyDavid/firewalld-web-console/server/agent_client"
        echo "github.com/labstack/echo"
)

const (
        WebPort = ":32126"
)

var clients map[client.AgentSpec]service.AgentClient

func GetOrMakeClient(spec client.AgentSpec) service.AgentClient {
        if clients[spec] == nil {
	        clients[spec] = client.New(spec.Address, spec.Port)
	}
	return clients[spec]
}

type State struct{
        state string
}

func main() {
     	e := echo.New()
	clients = make(map[client.AgentSpec]service.AgentClient)
	
	e.GET("/echo", func(c echo.Context) error {
		return c.String(http.StatusOK, "running")
	})
	e.GET("/state", func(c echo.Context) error {
		address := c.QueryParam("address")
		port := c.QueryParam("port")
		spec := client.NewSpec(address, port)
		rpcClient := GetOrMakeClient(*spec)
		ctx, cancel := context.WithTimeout(context.Background(),
		time.Second)
		defer cancel()
		r, err := rpcClient.GetStatus(ctx, &service.StatusRequest{Token:
		""})
		if err != nil {
		        log.Printf("get state error: %v", err)
		}
		state := new(State)
		state.state = r.BasicInfo
		return c.JSON(http.StatusOK, state)
	})
	
	e.Logger.Fatal(e.Start(WebPort))
}
