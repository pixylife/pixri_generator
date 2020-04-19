package app

import (
	"pixri_generator/pkg/controller"
	"text/template"
)

func createMain(generatedRoot string,projectName string)  {
	modelRoot := generatedRoot+"/lib/"
	controller.GenerateDir(modelRoot)
	tmpl := template.Must(template.ParseFiles("./templates/app/main.tp"))
	filePath :=modelRoot+"main.dart"
	controller.TemplateFileWriter(projectName, filePath, tmpl)
}