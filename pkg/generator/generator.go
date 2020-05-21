package generator

import (
	"fmt"
	"github.com/google/go-github/github"
	"pixri_generator/pixriLogger"
	"pixri_generator/pkg/controller"
	"pixri_generator/pkg/generator/app"
	"pixri_generator/pkg/generator/entity"
	"pixri_generator/pkg/generator/test"
	"pixri_generator/pkg/generator/ui"
	"pixri_generator/pkg/generator/ui/crud"
	"pixri_generator/pkg/generator/util"
	"pixri_generator/pkg/model"
)

var project app.Project



func GenerateInit(projectDir string,request model.GenRequest,repo *github.Repository) app.Project {
	pixriLogger.Log.Debug("Generating : Init")
	project = app.GetProject(projectDir,request,repo)
	project.Packgeroot = project.Name
	return project
}

func GenerateModelFunctions(project app.Project,request model.GenRequest){
	generatedRoot := project.Root
	models := entity.GenerateModel(generatedRoot,request)
	fmt.Println(models)
	for _,model := range models {
		crud.CreateFormUI(generatedRoot, model)
		crud.CreateListViewUI(generatedRoot, model)
		crud.CreateCRUDPageUI(generatedRoot, model)
	}
	ui.CreateHomeClass(generatedRoot,models)
}

func ModifyProjectFiles(project app.Project,request model.GenRequest){
	pixriLogger.Log.Debug("Modifying Project files :")
	app.UpdatePubspec(project)
	app.CreateAppClass(project,request.Theme)
	app.CreateMain(project.Root,project.Name)
	test.CreateTestClass(project.Root,project.Name)

}


func GenerateApplication(request model.GenRequest,repo *github.Repository){
	var projectDir = "sample"
	project := GenerateInit(projectDir,request,repo)
	GenerateModelFunctions(project,request)
	util.CreateHexColorUtil(project.Root)
	ModifyProjectFiles(project,request)
	controller.GitAddAll(project.Root)
	controller.GitCommit(project.Root,"Initial Commit")
	controller.GitPush(project.Root,"master")
}

