package generator

import (
	"pixri_generator/pixriLogger"
)

var project *Project

func GenerateInit(projectDir string) (*Project) {
	pixriLogger.Log.Debug("Generating : Init")
	project = getProject(projectDir)
	project.packgeroot = project.Name
	return project
}

func GenerateControllers(projectDir string,projectName string, generatedRoot string){
	generateModel(projectDir,projectName,generatedRoot)
}