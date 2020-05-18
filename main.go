package main

import (
	"pixri_generator/pkg/controller"
	"pixri_generator/pkg/generator"
)

func main()  {
	/*e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	r := e.Group("/")
	controller.GenerateController(r, "api")

	e.Logger.Fatal(e.Start(":5003"))*/

	var request =  controller.GenRequest{}
	generator.GenerateApplication(request)

}

