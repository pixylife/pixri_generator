package ui

import (
	"path/filepath"
	"pixri_generator/functions"
	"pixri_generator/pkg/controller"
	"pixri_generator/pkg/env"
	"pixri_generator/pkg/generator/app"
	"pixri_generator/pkg/generator/entity"
	"text/template"

)

type PageData struct {
	Name string
	ConstructorName string
}


func CreateHomeClass(generatedRoot string,models []entity.Model)  {
	fileRoot := generatedRoot + filepath.FromSlash(env.Root+env.UI_PATH)
	controller.GenerateDir(fileRoot)
	tmpl := template.Must(template.ParseFiles(
		"./templates/ui/homepage.tp",
		"./templates/ui/widget/home_page_button_card.tp"))


	var imports []string
	var pages []PageData


	for _,model := range models {
		imports = append(imports, app.ProjectData.Name+env.Src+env.UI_PATH+functions.TOLower(model.Name)+"/"+model.Name+env.PageSuffix)
		pages = append(pages,PageData{model.Name,model.Name+env.Page})
	}

	data := make(map[string]interface{})
	data["imports"] = imports
	data["pages"] = pages

	filePath :=fileRoot+"home.dart"
	controller.TemplateFileWriter(data, filePath, tmpl)
}