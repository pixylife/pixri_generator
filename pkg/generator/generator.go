package generator

import (
	"twillo_mobile_generator/xiLogger"
)

var project *Project
var entityNames []string
var entityFiles []string


func GenerateInit(projectDir string) (*Project) {

	xiLogger.Log.Debug("Generating : Init")
	project = getProject(projectDir)
	project.packgeroot = project.Name


	return project
}