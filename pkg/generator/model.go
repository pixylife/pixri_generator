package generator

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"pixri_generator/pixriLogger"
	"pixri_generator/pkg/controller"
	"text/template"
)

type Model struct {
	Name          string `json:"name"`
	Fields    []struct {
		Name  string `json:"name"`
		Type string `json:"type"`
		Key bool `json:"key"`
	} `json:"fields"`
	API []API `json:"apis"`
	PackageName string
	Path string
	BaseURL string `json:"base_url"`
}


func readEntityJson(projectDir string)[]Model {

	var files []string
	var inputs []Model

	root := projectDir+"/entity/"
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
		var input Model

		err = json.Unmarshal(byteValue, &input)
		if err != nil {
			pixriLogger.Log.Error(err)
		}

		i := append(inputs, input)
		inputs = i
	}
	return inputs

}

func generateModel(projectDir string, generatedRoot string,projectName string)  {
	models := readEntityJson(projectDir)
	for _, model := range models{
		model = createModel(generatedRoot,projectName,model)
		var primaryKey = PrimaryField{}
		for _, field:= range model.Fields{
			if field.Key{
				primaryKey.Name = field.Name
				primaryKey.Type = field.Type
				break
			}
		}

		modelData := ModelData{model.Name,model.PackageName,model.Path,primaryKey}
		api := ApiData{model.Name,model.BaseURL,modelData,model.API}
		GenerateApi(generatedRoot,api)
	}
}



func createModel(generatedRoot string,projectName string,model Model) Model {
	modelRoot := generatedRoot+"/lib/model/"
	GenerateDir(modelRoot)
	tmpl := template.Must(template.ParseFiles("./templates/model.tp"))
	filePath :=modelRoot+model.Name+".dart"
	controller.TemplateFileWriter(model, filePath, tmpl)
	model.Path = filePath
	model.PackageName =projectName+"/src/model/"+model.Name+".dart"
	return model
}