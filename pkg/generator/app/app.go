package app

import (
	"path/filepath"
	"pixri_generator/pkg/controller"
	"pixri_generator/pkg/env"
	"pixri_generator/pkg/model"
	"text/template"
)


func CreateAppClass(project Project,theme model.Theme)  {
	fileRoot := project.Root+ filepath.FromSlash(env.Root)
	controller.GenerateDir(fileRoot)
	tmpl := template.Must(template.ParseFiles("./templates/app/app.tp"))


	var imports []string

	imports = append(imports,  project.Name+env.Src+env.UI_PATH+"home"+env.DartExtension)
	imports = append(imports,  project.Name+env.Src+env.UTIL_PATH+"HexColor"+env.DartExtension)



	data := make(map[string]interface{})
	data["imports"] = imports
	data["title"] = project.Name
	data["body"] = "HomePage"
	data["theme"] = theme

	filePath :=fileRoot+"app.dart"
	controller.TemplateFileWriter(data, filePath, tmpl)
}