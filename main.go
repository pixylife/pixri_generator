package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"pixri_generator/pkg/generator"
)

func main()  {

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	r := e.Group("/")

	generator.GenerateController(r, "api")


	e.Logger.Fatal(e.Start(":5003"))
}

