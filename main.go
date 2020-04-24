package main

import "pixri_generator/pkg/generator"

func main()  {
	var projectDir = "sample";
	project := generator.GenerateInit(projectDir)
	generator.GenerateModelFunctions(projectDir,project.Root)
	generator.ModifyProjectFiles(projectDir,project)
}
