package main

import "pixri_generator/pkg/generator"



func main()  {

}

func GenerateFromFile(projectDir string)  {
	generator.GenerateInit(projectDir)
//	generator.GenerateBackend()
	//generator.GenerateFrontend(projectDir)
}