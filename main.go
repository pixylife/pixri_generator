package main

import (
	"fmt"
	"pixri_generator/pkg/generator"
)

func main()  {
	var projectDir = "sample";
	project := generator.GenerateInit(projectDir)
	generator.GenerateModelFunctions(projectDir,project.Root)
	generator.ModifyProjectFiles(project)
	fmt.Println("XXXXXXXXXXXXXXXXXXXXx")
	fmt.Println(project.Root)
	//controller.GitAddAll(project.Root)
	//controller.GitCommit(project.Root,"Initial Commit")
	//controller.GitPush(project.Root,"master")



/*	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	r := e.Group("/")
	controller.GenerateController(r, "api")

	e.Logger.Fatal(e.Start(":5003"))*/
}

