package generator

import (
	"pixri_generator/pixriLogger"
)

var project *Project
var entityNames []string
var entityFiles []string


func GenerateInit(projectDir string) (*Project) {

	pixriLogger.Log.Debug("Generating : Init")
	project = getProject(projectDir)
	project.packgeroot = project.Name


	return project
}