package generator

import (
	"pixri_generator/pkg/controller"
	"text/template"
)

type Model struct {
	Name          string `json:"name"`
	Fields    []struct {
		Name  string `json:"name"`
		Type string `json:"value"`
	} `json:"fields"`
	API []struct{
		Name  string `json:"name"`
		Type string `json:"value"`
		URL string `json:"url"`
	}
}

func CreateModel(projectName string, generatedRoot string,model Model) {
	modelRoot := generatedRoot+"/"+projectName+"/lib/model"
	GenerateDir(modelRoot)
	tmpl := template.Must(template.ParseFiles("pixri_generator/templates/model.tp"))
	filePath :=modelRoot+model.Name+".dart"
	controller.TemplateFileWriter(model, filePath, tmpl)


}