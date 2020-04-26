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

func CreateCRUDPageUI(generatedRoot string, model entity.Model) {
	uiRoot := generatedRoot + filepath.FromSlash(env.Root+env.UI_PATH+strings.ToLower(model.Name+"/"))
	controller.GenerateDir(uiRoot)
	tmpl := template.New("model-page")
	funcMap := template.FuncMap{}
	funcMap["dict"] = functions.Dict
	funcMap["first_letter_to_upper"] = functions.FirstLetterUpper
	funcMap["first_letter_to_lower"] = functions.MakeFirstLowerCase
	tmpl.Funcs(funcMap)

	tmpl, _ = tmpl.ParseFiles("./templates/ui/page/entity_crud_page.tp")

	var imports []string

	imports = append(imports, app.ProjectData.Name+env.Src+env.UI_PATH+strings.ToLower(model.Name+"/")+model.Name+env.FormViewSuffix)
	imports = append(imports, app.ProjectData.Name+env.Src+env.UI_PATH+strings.ToLower(model.Name+"/")+model.Name+env.ListViewSuffix)

	data := make(map[string]interface{})
	data["model"] = model
	data["imports"] = imports
	data["class"] = model.Name+env.Page
	data["list_ui"] = model.Name+env.List_View
	data["form_ui"] = model.Name+env.UI_Form

	filePath := uiRoot + model.Name + env.PageSuffix
	controller.TemplateFileWriterByName(data, filePath, tmpl, "model-page")
}
