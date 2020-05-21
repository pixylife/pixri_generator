package util

import (
	"path/filepath"
	"pixri_generator/pkg/controller"
	"pixri_generator/pkg/env"
	"text/template"
)

func CreateHexColorUtil(root string)  {
	fileRoot := root+filepath.FromSlash(env.Root+"/"+env.UTIL_PATH)
	controller.GenerateDir(fileRoot)
	tmpl := template.Must(template.ParseFiles("./templates/util/hex_color.tp"))
	filePath :=fileRoot+"HexColor.dart"
	controller.TemplateFileWriter(nil, filePath, tmpl)
}