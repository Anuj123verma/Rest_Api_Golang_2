// @title Swagger Covid Information About State API
// @version 1.0
// @description This is a server that can give the covid information about the states of India.

// @contact.name Anuj Verma
// @contact.email anujssooni360@gmail.con

// @host localhost:8000
// @BasePath /
// @Schemes http

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
package main

import (
	"github.com/Anuj123Verma/Rest_Api_Golang_2/controller"
	_ "github.com/Anuj123Verma/Rest_Api_Golang_2/docs"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

const (
	port = "8000"
)

func main() {
	// making an instance of echo
	e := echo.New()
	// calling the api
	e.GET("/state/:data", controller.Getstate)
	// calling the swagger api
	e.GET("/swagger/*any", echoSwagger.WrapHandler)
	// starting the port on 8000
	e.Logger.Fatal(e.Start(":" + port))
}
