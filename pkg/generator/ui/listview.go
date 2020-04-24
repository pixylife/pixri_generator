package ui

import (
	"pixri_generator/functions"
	"pixri_generator/pkg/controller"
	"pixri_generator/pkg/env"
	"pixri_generator/pkg/generator/entity"
	"text/template"
)

func CreateListViewUI(generatedRoot string, projectName string, model entity.Model) entity.Model {
	uiRoot := generatedRoot + "/lib/ui/"
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
	filePath := uiRoot + model.Name + env.FormViewSuffix
	controller.TemplateFileWriterByName(model, filePath, tmpl, "UI-List-View")
	return model
}
