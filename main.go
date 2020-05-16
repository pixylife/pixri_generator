package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"pixri_generator/pkg/controller"
)

func main()  {
	/*var projectDir = "sample";
	project := generator.GenerateInit(projectDir)
	generator.GenerateModelFunctions(projectDir,project.Root)
	generator.ModifyProjectFiles(project)*/


	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	r := e.Group("/")
	controller.GenerateController(r, "api")

	e.Logger.Fatal(e.Start(":5003"))
}

