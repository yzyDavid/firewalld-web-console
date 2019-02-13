package main

import (
        "net/http"
	
        echo "github.com/labstack/echo"
)

const (
        port = ":32126"
)

func main() {
     	e := echo.New()
	e.GET("/echo", func(c echo.Context) error {
		return c.String(http.StatusOK, "running")
	})
	e.Logger.Fatal(e.Start(port))
}
