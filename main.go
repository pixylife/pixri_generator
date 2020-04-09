package main

import "pixri_generator/pkg/generator"

func main()  {
	project := generator.GenerateInit("sample")
	generator.GenerateControllers("sample",project.Name,project.Root)
}
