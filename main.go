package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"pixri_generator/pkg/generator"
	"pixri_generator/ws"
)

func main()  {

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	r := e.Group("/")

	generator.GenerateController(r, "api")

	go ws.Manager.Start()
	e.GET("/ws/:userName", ws.WsPage)

	e.Logger.Fatal(e.Start(":5003"))
}

