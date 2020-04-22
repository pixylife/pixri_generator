package app

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"pixri_generator/pixriLogger"
	"time"
)

type Project struct {
	Name          string `json:"name"`
	Status        string `json:"status"`
	Description string `json:"description"`
	Properties    []struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	} `json:"properties"`
	Root       string
	Packgeroot string
}

var ProjectData = Project{}

func GetProject(projectDir string) Project {
	pixriLogger.Log.Debug("Project Directory : ", projectDir)
	pj, er := ioutil.ReadFile(projectDir + "/project.json")
	if er != nil {
		pixriLogger.Log.Error("Error while reading project json", er)
	}
	if er := json.Unmarshal(pj, &ProjectData); er != nil {
		pixriLogger.Log.Error("Error while Unmarshal project json", er)
	}
	rootLocation := projectDir + "/generated/" + ProjectData.Name
		if _, err := os.Stat(filepath.FromSlash(rootLocation)); os.IsNotExist(err) {
			pixriLogger.Log.Debug( "Project root is not exist , creating",rootLocation)
			projectInit(ProjectData.Name,projectDir)
		}else{
			pixriLogger.Log.Info("Project root is exist , ignore project Init step")
		}
	ProjectData.Root = filepath.FromSlash(rootLocation)
	projectInit(ProjectData.Name, projectDir)
	pixriLogger.Log.Info("Project root for generated codes :", ProjectData.Root)
	return ProjectData
}

func projectInit(projectName string, projectDir string){
	pixriLogger.Log.Info("Initialization of the project :", projectName)
	generatedRoot := projectDir + "/generated"
	createProject(projectName, generatedRoot)
	}

func createProject(projectName string, generatedRoot string) {
	now := time.Now()
	cmd := exec.Command("flutter", "create", "--org", "io.prixi."+projectName+""+string(now.Unix()), "-i", "swift", "-a", "kotlin", "--description", "'"+projectName +" mobile app'", projectName)
	cmd.Dir = generatedRoot
	out, err := cmd.Output()
	pixriLogger.Log.Info("Project init",string(out))
	if err !=nil {
		pixriLogger.Log.Error("cmd.Run() failed with %s\n", err)
	}
}