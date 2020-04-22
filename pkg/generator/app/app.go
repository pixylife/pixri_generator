package app

import (
	"pixri_generator/pkg/controller"
	"text/template"
)


func CreateAppClass(project Project)  {
	fileRoot := project.Root+"/lib/"
	controller.GenerateDir(fileRoot)
	tmpl := template.Must(template.ParseFiles("./templates/app/app.tp"))
	filePath :=fileRoot+"app.dart"
	controller.TemplateFileWriter(project, filePath, tmpl)
}