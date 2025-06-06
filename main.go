package main

import (
	"email-forwarder/api"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/test", api.Test)
	e.Start(":8002")
}
