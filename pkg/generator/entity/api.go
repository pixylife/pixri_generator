package entity

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"pixri_generator/pixriLogger"
	"pixri_generator/pkg/controller"
	"pixri_generator/pkg/env"
	"text/template"
)


type API struct {
	Name       string `json:"name"`
	MethodName string `json:"methodName"`
	Type       string `json:"type"`
	Operation  string `json:"operation"`
	Resource   string `json:"resource"`
	Target     struct {
		Type string `json:"type"`
		Name string `json:"name"`
	} `json:"target"`
	Params []struct {
		Type      string `json:"type"`
		InputType string `json:"inputType"`
		InputName string `json:"inputName"`
	} `json:"params"`
	Ruleid string `json:"ruleid"`
	Return struct {
		Name   string `json:"name"`
		Type   string `json:"type"`
		Record string `json:"record"`
	} `json:"return"`
}


func readAPIJson(projectDir string)[]API {

	var files []string
	var inputs []API

	root := projectDir+"/api/"
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

func GenerateAPI(projectDir string,generatedRoot string){
	apis :=readAPIJson(projectDir)
	for _,api:= range apis{
		createApi(generatedRoot,api)
	}
}


func createApi(generatedRoot string,api API){
	var apiRoot string
	if api.Type == env.ENTITY_API{
		apiRoot = generatedRoot + "/lib/api/"+api.Target.Name+"/"
	}else {
		apiRoot = generatedRoot + "/lib/api/"
	}
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

	filePath := apiRoot + api.Name + env.API_SUFFIX
	controller.TemplateFileWriterByName(api, filePath, tmpl, "api")
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
