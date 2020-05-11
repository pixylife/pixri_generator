package entity

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"pixri_generator/functions"
	"pixri_generator/pixriLogger"
	"pixri_generator/pkg/controller"
	"pixri_generator/pkg/env"
	"pixri_generator/pkg/generator/app"
	"strings"
	"text/template"
)

type API struct {
	Name       string `json:"name"`
	MethodName string `json:"methodName"`
	Type       string `json:"type"`
	Operation  string `json:"operation"`
	Resource   string `json:"resource"`
	Target     Target `json:"target"`
	Params     []struct {
		Type      string `json:"type"`
		InputType string `json:"inputType"`
		InputName string `json:"inputName"`
	} `json:"params"`
	Ruleid string `json:"ruleid"`
	Return Return `json:"return"`
}

type Target struct {
	Type string `json:"type"`
	Name string `json:"name"`
}

type Return struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	Record string `json:"record"`
}

func readAPIJson(projectDir string) []API {

	var files []string
	var inputs []API

	root := projectDir + "/api/"
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {

		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		pixriLogger.Log.Error(err)
	}
	for _, file := range files {
		pixriLogger.Log.Debug(file)
	}
	for _, file := range files {
		pixriLogger.Log.Debug(file)
		jsonFile, err := os.Open(file)
		if err != nil {
			pixriLogger.Log.Error(err)
		}
		defer jsonFile.Close()
		byteValue, _ := ioutil.ReadAll(jsonFile)
		var input API

		err = json.Unmarshal(byteValue, &input)
		if err != nil {
			pixriLogger.Log.Error(err)
		}

		i := append(inputs, input)
		inputs = i
	}
	return inputs

}

/*
func GenerateAPI(projectDir string,generatedRoot string){
	apis :=readAPIJson(projectDir)
	for _,api:= range apis{
		createApi(generatedRoot,api)
	}
}
*/

func GenerateApi(generatedRoot string, model Model) {
	apis := []API{}
	for _, apiType := range env.GetSupportedCruds() {
		api := API{}
		api.Name = apiType + model.Name
		api.Type = apiType
		api.MethodName = apiType + model.Name
		api.Resource = "/"+strings.ToLower(controller.ToPlural(model.Name))
		apis = append(apis, api)
	}
	createModelApi(generatedRoot, model, apis)

}

func createModelApi(generatedRoot string, model Model, api []API)  {

	apiRoot := generatedRoot + filepath.FromSlash(env.Root + env.API_PATH)
	controller.GenerateDir(apiRoot)

	tmpl := template.New("api")
	funcMap := template.FuncMap{}
	funcMap["dict"] = functions.Dict
	tmpl.Funcs(funcMap)

	data := make(map[string]interface{})

	tmpl, _ = tmpl.ParseFiles("./templates/api/api_class.tp",
		"./templates/api/api_create.tp",
		"./templates/api/api_get_list.tp",
		"./templates/api/api_update.tp",
		"./templates/api/api_delete.tp",
		"./templates/api/api_get.tp")

	var imports []string

	imports = append(imports, app.ProjectData.Name+env.Src+env.MODEL_PATH+model.Name+env.DartExtension)

	data["api"] = api
	data["model"] = model
	data["imports"] = imports
	data["class"] = model.Name+env.API_Class

	filePath := apiRoot + model.Name + env.API_SUFFIX
	controller.TemplateFileWriterByName(data, filePath, tmpl, "api")
}


