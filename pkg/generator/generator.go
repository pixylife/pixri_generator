package generator

import (
	"pixri_generator/pixriLogger"
	"pixri_generator/pkg/generator/app"
	"pixri_generator/pkg/generator/entity"
	"pixri_generator/pkg/generator/ui"
)

var project *app.Project

func GenerateInit(projectDir string) *app.Project {
	pixriLogger.Log.Debug("Generating : Init")
	project = app.GetProject(projectDir)
	project.Packgeroot = project.Name
	return project
}

func GenerateControllers(projectDir string,projectName string, generatedRoot string){
	models := entity.GenerateModel(projectDir,generatedRoot,projectName)
	for _,model := range models {
		ui.CreateFormUI(generatedRoot, projectName, model)
		ui.CreateListViewUI(generatedRoot, projectName, model)
	}
}

func ModifyProjectFiles(projectDir string,project app.Project){
	app.UpdatePubspec(project)
	app.CreateAppClass(project)
	app.CreateMain(project.Root,project.Name)
}

