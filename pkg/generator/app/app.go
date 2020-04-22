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
	filePath :=fileRoot+"app.dart"
	controller.TemplateFileWriter(project, filePath, tmpl)
}