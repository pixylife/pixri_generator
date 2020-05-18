package app

import (
	"path/filepath"
	"pixri_generator/pkg/controller"
	"pixri_generator/pkg/env"
	"text/template"
)


func CreateAppClass(project Project)  {
	fileRoot := project.Root+ filepath.FromSlash(env.Root)
	controller.GenerateDir(fileRoot)
	tmpl := template.Must(template.ParseFiles("./templates/app/app.tp"))


	var imports []string

	imports = append(imports,  project.Name+env.Src+env.UI_PATH+"home"+env.DartExtension)

	data := make(map[string]interface{})
	data["imports"] = imports
	data["title"] = project.Name
	data["body"] = "HomePage"

	filePath :=fileRoot+"app.dart"
	controller.TemplateFileWriter(data, filePath, tmpl)
}