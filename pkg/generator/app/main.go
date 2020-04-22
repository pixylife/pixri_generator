package app

import (
	"path/filepath"
	"pixri_generator/pkg/controller"
	"pixri_generator/pkg/env"
	"text/template"
)

func CreateMain(generatedRoot string,projectName string)  {
	modelRoot := generatedRoot+filepath.FromSlash(env.Lib)
	controller.GenerateDir(modelRoot)
	tmpl := template.Must(template.ParseFiles("./templates/app/main.tp"))
	filePath :=modelRoot+"main.dart"
	controller.TemplateFileWriter(projectName, filePath, tmpl)
}