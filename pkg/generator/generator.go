package generator

import (
	"pixri_generator/pixriLogger"
	"pixri_generator/pkg/generator/entity"
	"pixri_generator/pkg/generator/ui"
)

var project *Project

func GenerateInit(projectDir string) (*Project) {
	pixriLogger.Log.Debug("Generating : Init")
	project = getProject(projectDir)
	project.packgeroot = project.Name
	return project
}

func GenerateControllers(projectDir string,projectName string, generatedRoot string){
	models := entity.GenerateModel(projectDir,generatedRoot,projectName)
	for _,model := range models {
		ui.CreateFormUI(generatedRoot, projectName, model)
	}

}

