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

func CreateListViewUI(generatedRoot string, projectName string, model entity.Model) entity.Model {
	uiRoot := generatedRoot + filepath.FromSlash(env.Root+env.UI_PATH+strings.ToLower(model.Name+"/"))
	controller.GenerateDir(uiRoot)
	tmpl := template.New("UI-List-View")
	funcMap := template.FuncMap{}
	funcMap["dict"] = functions.Dict
	funcMap["plus1"] = functions.Plus1
	funcMap["first_letter_to_upper"] = functions.FirstLetterUpper
	tmpl.Funcs(funcMap)

	tmpl, _ = tmpl.ParseFiles("./templates/ui/view/basic_list_view.tp",
		"./templates/ui/dialog/yes_no_confirm_dialog.tp",
		"./templates/ui/widget/text_field_widget.tp")

	var imports []string

	imports = append(imports, app.ProjectData.Name+env.Src+env.MODEL_PATH+model.Name+env.DartExtension)
	imports = append(imports, app.ProjectData.Name+env.Src+env.API_PATH+model.Name+env.API_SUFFIX)
	data := make(map[string]interface{})
	data["model"] = model
	data["imports"] = imports
	data["class"] = model.Name+env.List_View
	data["class_api"] = model.Name+env.API_Class

	filePath := uiRoot + model.Name + env.ListViewSuffix
	controller.TemplateFileWriterByName(data, filePath, tmpl, "UI-List-View")
	return model
}
