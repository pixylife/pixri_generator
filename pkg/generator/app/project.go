package app

import (
	"bytes"
	"github.com/google/go-github/github"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"pixri_generator/functions"
	"pixri_generator/pixriLogger"
	"pixri_generator/pkg/controller"
	"pixri_generator/pkg/model"
	"sync"
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
	GitRepoData GitRepoData
}

type GitRepoData struct {
	Name string
	URL string
	CloneURL string
}

var ProjectData = Project{}
var ProjectResponse = model.GenResponse{}

func GetProject(projectDir string, request model.GenRequest, repo *github.Repository) Project {
	/*pixriLogger.Log.Debug("Project Directory : ", projectDir)
	pj, er := ioutil.ReadFile(projectDir + "/project.json")
	if er != nil {
		pixriLogger.Log.Error("Error while reading project json", er)
	}
	if er := json.Unmarshal(pj, &ProjectData); er != nil {
		pixriLogger.Log.Error("Error while Unmarshal project json", er)
	}
*/
	app := request.Application
	ProjectData.Name = functions.SpaceStringsBuilder(app.Name)
	ProjectData.Description = app.Description


	var git = GitRepoData{}
	git.Name = *repo.Name
	git.URL = *repo.URL
	git.CloneURL = *repo.CloneURL
	ProjectData.GitRepoData = git


	rootLocation := projectDir + "/generated/"+ProjectData.GitRepoData.Name+"/"+ProjectData.Name
	//rootLocation := projectDir + "/generated/"+ProjectData.Name
		if _, err := os.Stat(filepath.FromSlash(rootLocation)); os.IsNotExist(err) {
			pixriLogger.Log.Debug( "Project root is not exist , creating",rootLocation)
		}else{
			pixriLogger.Log.Info("Project root is exist , ignore project Init step")
		}
	ProjectData.Root = filepath.FromSlash(rootLocation)
	projectInit(ProjectData.Name, projectDir)
	pixriLogger.Log.Info("Project root for generated codes :", ProjectData.Root)
	return ProjectData
}

func projectInit(appName string, projectDir string){
	pixriLogger.Log.Info("Initialization of the project :", appName)
	generatedRoot := projectDir + "/generated"
	_ = controller.GitClone(ProjectData.GitRepoData.CloneURL, generatedRoot)
	generatedRoot= generatedRoot+"/"+ProjectData.GitRepoData.Name
	createProject(appName, generatedRoot)
	}

func createProject(projectName string, generatedRoot string) {
	cmd := exec.Command("flutter", "create", "--org", "io.pixri."+projectName, "-i", "swift", "-a", "kotlin", "--description", "'"+projectName +" mobile app'", projectName)
	cmd.Dir = generatedRoot
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