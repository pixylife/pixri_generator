package crud

import (
	"path/filepath"
	"pixri_generator/functions"
	"pixri_generator/pkg/controller"
	"pixri_generator/pkg/env"
	"pixri_generator/pkg/generator/app"
	"pixri_generator/pkg/generator/entity"
	"strings"
	"text/template"
)



func CreateFormUI(generatedRoot string, projectName string, model entity.Model) entity.Model {
	uiRoot := generatedRoot + filepath.FromSlash(env.Root+env.UI_PATH+strings.ToLower(model.Name+"/"))
	controller.GenerateDir(uiRoot)
	tmpl := template.New("UI-Basic-Form")
	funcMap := template.FuncMap{}
	funcMap["dict"] = functions.Dict
	funcMap["plus1"] = functions.Plus1
	funcMap["first_letter_to_upper"] = functions.FirstLetterUpper
	tmpl.Funcs(funcMap)



	tmpl, _ = tmpl.ParseFiles("./templates/ui/form/basic_input_form.tp",
										"./templates/ui/widget/raised_button_widget.tp",
										"./templates/ui/widget/text_field_widget.tp")

	data := make(map[string]interface{})
	var imports []string

	imports = append(imports, app.ProjectData.Name+env.Src+env.MODEL_PATH+model.Name+env.DartExtension)
	imports = append(imports, app.ProjectData.Name+env.Src+env.API_PATH+model.Name+env.API_SUFFIX)

	data["model"] = model
	data["imports"] = imports
	data["class"] = model.Name+env.UI_Form
	data["class_api"] = model.Name+env.API_Class


	filePath := uiRoot + model.Name + env.FormViewSuffix
	controller.TemplateFileWriterByName(data, filePath, tmpl, "UI-Basic-Form")
	return model
}
