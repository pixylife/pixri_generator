package main

import (
	"pixri_generator/pkg/controller"
	"pixri_generator/pkg/generator"
)

func main()  {
	var projectDir = "sample";
	project := generator.GenerateInit(projectDir)
	_ = controller.GitInit(project.Root)
	_ = controller.SetRemote("https://github.com/pixylife/Test-1589634460-Pixri", project.Root)

	//generator.GenerateModelFunctions(projectDir,project.Root)
	//generator.ModifyProjectFiles(project)


/*	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	r := e.Group("/")
	controller.GenerateController(r, "api")

	e.Logger.Fatal(e.Start(":5003"))*/
}

