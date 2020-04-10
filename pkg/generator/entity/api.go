package entity

import (
	"errors"
	"pixri_generator/pkg/controller"
	"text/template"
)

type API struct {
	Name string `json:"name"`
	Type string `json:"type"`
	URL  string `json:"url"`
}

type ApiData struct {
	Name      string `json:"name"`
	BaseURL   string `json:"base_url"`
	ModelData ModelData
	API       []API
	PackageName string
	Path string
}

type ModelData struct {
	Name     string `json:"name"`
	Package  string `json:"package"`
	Path     string
	PrimaryField PrimaryField
}

type PrimaryField struct {
	Name string `json:"name"`
	Type string `json:"type"`
}


func GenerateApi(generatedRoot string, api ApiData) ApiData{
	apiRoot := generatedRoot + "/lib/api/"
	controller.GenerateDir(apiRoot)

	tmpl := template.New("api")
	funcMap := template.FuncMap{}
	funcMap["dict"] = dict
	tmpl.Funcs(funcMap)

	tmpl, _ = tmpl.ParseFiles("./templates/api/api_class.tp",
		"./templates/api/api_create.tp",
		"./templates/api/api_get_list.tp",
		"./templates/api/api_update.tp",
		"./templates/api/api_delete.tp",
		"./templates/api/api_get.tp")

	filePath := apiRoot + api.Name + "_api_service.dart"
	controller.TemplateFileWriterByName(api, filePath, tmpl, "api")

	api.Path = filePath
	return api
}

func dict(values ...interface{}) (map[string]interface{}, error) {
	if len(values)%2 != 0 {
		return nil, errors.New("invalid dict call")
	}
	dict := make(map[string]interface{}, len(values)/2)
	for i := 0; i < len(values); i += 2 {
		key, ok := values[i].(string)
		if !ok {
			return nil, errors.New("dict keys must be strings")
		}
		dict[key] = values[i+1]
	}
	return dict, nil
}
