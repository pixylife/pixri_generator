package ui

import (
	"pixri_generator/pkg/controller"
	"pixri_generator/pkg/generator/entity"
	"text/template"
	"pixri_generator/pkg/env"
)

var fns = template.FuncMap{
	"plus1": func(x int) int {
		return x + 1
	},
}

func CreateFormUI(generatedRoot string, projectName string, model entity.Model) entity.Model {
	uiRoot := generatedRoot + "/lib/ui/"
	controller.GenerateDir(uiRoot)
	tmpl := template.New("UI-Basic-Form")
	tmpl.Funcs(fns)

	tmpl, _ = tmpl.ParseFiles("./templates/ui/form/basic_input_form.tp",
										"./templates/ui/widget/raised_button_widget.tp",
										"./templates/ui/widget/text_field_widget.tp")
	filePath := uiRoot + model.Name + env.FormViewSuffix
	controller.TemplateFileWriterByName(model, filePath, tmpl, "UI-Basic-Form")
	return model
}
