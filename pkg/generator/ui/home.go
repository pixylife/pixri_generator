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
	tmpl := template.New("UI-Home-Page")
	//adding template functions
	funcMap := template.FuncMap{}
	funcMap["dict"] = functions.Dict
	funcMap["plus1"] = functions.Plus1
	funcMap["first_letter_to_upper"] = functions.FirstLetterUpper
	funcMap["first_letter_to_lower"] = functions.MakeFirstLowerCase
	tmpl.Funcs(funcMap)



	controller.GenerateDir(fileRoot)
	tmpl,_ = tmpl.ParseFiles(
		"./templates/ui/homepage.tp",
		"./templates/ui/widget/home_page_button_card.tp")


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
	controller.TemplateFileWriterByName(data, filePath, tmpl, "UI-Home-Page")
}