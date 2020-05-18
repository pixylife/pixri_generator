package generator

import (
	"pixri_generator/pixriLogger"
	"pixri_generator/pkg/controller"
	"pixri_generator/pkg/generator/app"
	"pixri_generator/pkg/generator/entity"
	"pixri_generator/pkg/generator/test"
	"pixri_generator/pkg/generator/ui"
	"pixri_generator/pkg/generator/ui/crud"
)

var project app.Project

func GenerateInit(projectDir string,request controller.GenRequest) app.Project {
	pixriLogger.Log.Debug("Generating : Init")
	project = app.GetProject(projectDir,request)
	project.Packgeroot = project.Name
	return project
}

func GenerateModelFunctions(projectDir string, generatedRoot string){
	models := entity.GenerateModel(projectDir,generatedRoot)
	for _,model := range models {
		crud.CreateFormUI(generatedRoot, model)
		crud.CreateListViewUI(generatedRoot, model)
		crud.CreateCRUDPageUI(generatedRoot, model)
	}
	ui.CreateHomeClass(generatedRoot,models)
}

func ModifyProjectFiles(project app.Project){
	pixriLogger.Log.Debug("Modifying Project files :")
	app.UpdatePubspec(project)
	app.CreateAppClass(project)
	app.CreateMain(project.Root,project.Name)
	test.CreateTestClass(project.Root,project.Name)

}


func GenerateApplication(request controller.GenRequest){
	var projectDir = "sample"
	project := GenerateInit(projectDir,request)
	GenerateModelFunctions(projectDir,project.Root)
	ModifyProjectFiles(project)
	/*controller.GitAddAll(project.Root)
	controller.GitCommit(project.Root,"Initial Commit")
	controller.GitPush(project.Root,"master")*/
}

