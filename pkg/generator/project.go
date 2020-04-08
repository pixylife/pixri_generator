package generator

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"sync"
	"time"
	"pixri_generator/pixriLogger"
)

type Project struct {
	Name          string `json:"name"`
	Status        string `json:"status"`
	Properties    []struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	} `json:"properties"`
	root        string
	entityNames []string
	packgeroot  string
}


func getProject(projectDir string) *Project {

	pixriLogger.Log.Debug("Project Directory : ", projectDir)
	pj, er := ioutil.ReadFile(projectDir + "/project.json")
	if er != nil {
		pixriLogger.Log.Error("Error while reading project json", er)
	}
	project := new(Project)
	if er := json.Unmarshal(pj, &project); er != nil {
		pixriLogger.Log.Error("Error while Unmarshal project json", er)
	}
	rootLocation := projectDir + "/generated/" + project.Name
		if _, err := os.Stat(filepath.FromSlash(rootLocation)); os.IsNotExist(err) {
			pixriLogger.Log.Debug( "Project root is not exist , creating",rootLocation)
			projectInit(project.Name,projectDir)
		}else{
			pixriLogger.Log.Info("Project root is exist , ignore project Init step")
		}
	project.root = filepath.FromSlash(rootLocation)
	projectInit(project.Name, projectDir)
	pixriLogger.Log.Info("Project root for generated codes :", project.root)
	return project
}

func projectInit(projectName string, projectDir string){
	pixriLogger.Log.Info("Initialization of the project :", projectName)
	generatedRoot := projectDir + "/generated"
	createProject(projectName, generatedRoot)
	}

func createProject(projectName string, generatedRoot string) {
	now := time.Now()      // current local time
	cmd := exec.Command("flutter", "create", "--org", "io.prixi."+projectName+""+string(now.Unix()), "-i", "swift", "-a", "kotlin", "--description", "'"+projectName +" mobile app'", projectName)
	displayOutput(*cmd)
}

func displayOutput(cmd exec.Cmd) {
	var stdoutBuf, stderrBuf bytes.Buffer
	stdoutIn, _ := cmd.StdoutPipe()
	stderrIn, _ := cmd.StderrPipe()

	var errStdout, errStderr error
	stdout := io.MultiWriter(os.Stdout, &stdoutBuf)
	stderr := io.MultiWriter(os.Stderr, &stderrBuf)
	err := cmd.Start()
	if err != nil {
		pixriLogger.Log.Error("cmd.Start() failed with '%s'\n", err)
	}

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		_, errStdout = io.Copy(stdout, stdoutIn)
		wg.Done()
	}()

	_, errStderr = io.Copy(stderr, stderrIn)
	wg.Wait()

	err = cmd.Wait()
	if err != nil {
		pixriLogger.Log.Error("cmd.Run() failed with %s\n", err)
	}
	if errStdout != nil || errStderr != nil {
		pixriLogger.Log.Error("failed to capture stdout or stderr\n")
	}
	outStr, errStr := string(stdoutBuf.Bytes()), string(stderrBuf.Bytes())
	pixriLogger.Log.Error("\nout:\n%s\nerr:\n%s\n", outStr, errStr)
}
