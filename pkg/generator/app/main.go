package app

import (
	"path/filepath"
	"pixri_generator/pkg/controller"
	"pixri_generator/pkg/env"
	"text/template"
)

func CreateMain(generatedRoot string,projectName string)  {
	root := generatedRoot+filepath.FromSlash(env.Lib)
	controller.GenerateDir(root)
	tmpl := template.Must(template.ParseFiles("./templates/app/main.tp"))
	filePath :=root+"/main.dart"
	controller.TemplateFileWriter(projectName, filePath, tmpl)
}