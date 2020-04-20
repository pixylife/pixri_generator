package app

import (
	"pixri_generator/pkg/controller"
	"text/template"
)

func UpdatePubspec(project Project)  {
	fileRoot := project.Root+"/"
	controller.GenerateDir(fileRoot)
	tmpl := template.Must(template.ParseFiles("./templates/app/pubspec.tp"))
	filePath :=fileRoot+"pubspec.yaml"
	controller.TemplateFileWriter(project, filePath, tmpl)
}