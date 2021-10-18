package main

import (
	"github.com/Anuj123Verma/Rest_Api_Golang_2/controller"
	"github.com/labstack/echo/v4"
)

const (
	port = "8000"
)

func main() {
	e := echo.New()
	e.GET("/state/:data", controller.Getstate)
	e.Start(":" + port)
}
