package test

import (
	"path/filepath"
	"pixri_generator/pkg/controller"
	"text/template"
)

func CreateTestClass(generatedRoot string,projectName string)  {
	root := generatedRoot+filepath.FromSlash("/test")
	controller.GenerateDir(root)
	tmpl := template.Must(template.ParseFiles("./templates/test/test.tp"))
	filePath :=root+"/widget_test.dart"
	controller.TemplateFileWriter(projectName, filePath, tmpl)
}