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

var modelMap = make(map[string]Model)

type Model struct {
	Name                  string         `json:"name"`
	UIName                string         `json:"uiName"`
	Fields                []Field        `json:"fields"`
	Relationships         []Relationship `json:"relationships"`
	ChangelogDate         string         `json:"changelogDate"`
	EntityTableName       string         `json:"entityTableName"`
	Dto                   string         `json:"dto"`
	Pagination            string         `json:"pagination"`
	Service               string         `json:"service"`
	JpaMetamodelFiltering bool           `json:"jpaMetamodelFiltering"`
	FluentMethods         bool           `json:"fluentMethods"`
	ClientRootFolder      string         `json:"clientRootFolder"`
	Applications          string         `json:"applications"`
	Generate              bool           `json:"generate"`
}

type Field struct {
	FieldName          string   `json:"fieldName"`
	FieldType          string   `json:"fieldType"`
	FieldUIName        string   `json:"fieldUIName"`
	FieldValidateRules []string `json:"fieldValidateRules,omitempty"`
	FieldValues        string   `json:"fieldValues,omitempty"`
}

type Relationship struct {
	RelationshipName            string `json:"relationshipName"`
	OtherEntityName             string `json:"otherEntityName"`
	RelationshipType            string `json:"relationshipType"`
	OtherEntityRelationshipName string `json:"otherEntityRelationshipName"`
	OtherEntityField            string `json:"otherEntityField"`
	OwnerSide                   bool   `json:"ownerSide,omitempty"`
}

func (u *Model) Modify() {
	for i, model := range u.Fields {
		u.Fields[i].FieldUIName = strings.Title(u.Fields[i].FieldUIName)
		switch model.FieldType {
		case env.Integer:
			u.Fields[i].FieldType = "int"
		case env.String:
			u.Fields[i].FieldType = "String"
		}
	}
	createEntityRelationshipStatements(u)
}

func readEntityJson(projectDir string) []Model {

	var files []string
	var inputs []Model

	root := projectDir + "/entity/"
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

func createEntityRelationshipStatements(u *Model) {
	pixriLogger.Log.Debug(" 🔸 Generating : Entity Relationship")
	for _, field := range u.Relationships {
		otherEntity := field.OtherEntityName
		relationship := field.RelationshipType

		if relationship == "one-to-one" {
			pixriLogger.Log.Debug(" 🔸 Generating : Entity Relationship : one-to-one ")
			newField := Field{}
			newField.FieldName = functions.MakeFirstLowerCase(otherEntity)
			newField.FieldType = strings.Title(otherEntity)
			newField.FieldUIName = field.OtherEntityField

			u.Fields = append(u.Fields, newField)

		} else if relationship == "one-to-many" {
			pixriLogger.Log.Debug(" 🔸 Generating : Entity Relationship : one-to-many ")
			newField := Field{}
			newField.FieldName = functions.MakeFirstLowerCase(otherEntity)
			newField.FieldType = "List<" + otherEntity + ">"
			newField.FieldValues = field.OtherEntityField
			newField.FieldUIName = field.OtherEntityField
			u.Fields = append(u.Fields, newField)

		} else if relationship == "many-to-one" {
			pixriLogger.Log.Debug(" 🔸 Generating : Entity Relationship : many-to-one ")
			newField := Field{}
			newField.FieldName = functions.MakeFirstLowerCase(otherEntity)
			newField.FieldType = strings.Title(otherEntity)
			newField.FieldValues = field.OtherEntityField
			newField.FieldUIName = field.OtherEntityField
			u.Fields = append(u.Fields, newField)

		} else if relationship == "many-to-many" {
			pixriLogger.Log.Debug(" 🔸 Generating : Entity Relationship : many-to-many ")
			newField := Field{}
			newField.FieldName = functions.MakeFirstLowerCase(otherEntity)
			newField.FieldType = "List<" + otherEntity + ">"
			newField.FieldValues = field.OtherEntityField
			newField.FieldUIName = field.OtherEntityField
			u.Fields = append(u.Fields, newField)
		}

	}
}

func GenerateModel(projectDir string, generatedRoot string) []Model {
	var modelList []Model
	models := readEntityJson(projectDir)
	for _, model := range models {
		model.Modify()
		modelMap[model.Name] = model
		createModel(generatedRoot, model)
		GenerateApi(generatedRoot, model)
		modelList = append(modelList, model)
	}
	return modelList
}

func createModel(generatedRoot string, model Model) {
	modelRoot := generatedRoot + filepath.FromSlash( env.Root + env.MODEL_PATH)

	controller.GenerateDir(modelRoot)
	tmpl := template.Must(template.ParseFiles("./templates/controller/model.tp"))

	var imports []string

	if model.Relationships != nil{
		for _, relationship := range model.Relationships{
			imports = append(imports, app.ProjectData.Name+env.Src+env.MODEL_PATH+strings.Title(relationship.OtherEntityName)+env.DartExtension)
		}
	}

	data := make(map[string]interface{})

	data["model"] = model
	data["imports"] = imports

	filePath := modelRoot + model.Name + ".dart"
	controller.TemplateFileWriter(data, filePath, tmpl)
}
