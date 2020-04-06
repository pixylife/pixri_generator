package main

import (
	"fmt"
	"github.com/labstack/echo"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
)

type GenJob struct {

	SolutionId     int    `json:"solutionId"`
	SolutionName   string `json:"solutionName"`
	ProjectUUId    string  `json:"project_uu_id"`
	JobDescription string `json:"job_description"`

}


var generateJobs = make(chan GenJob,10)
var finishedJobs = make (chan GenResult,15)
//worker
func webAppGenerator(wg *sync.WaitGroup) {
	for job := range generateJobs {
		xiLogger.Log.Debug("Web App generator pic the job ")
		Generate(&job)
		genResult := pub.GenResult{SolutionId: int32(job.SolutionId),ProjectUUId:job.ProjectUUId,Status:pub.JobStatus_COMPLETED}
		output := GenResult{CallbackUrl:job.CallbackUrl,result: genResult}
		select {
		case finishedJobs <- output:
			xiLogger.Log.Debug("Job is finished and insert the result into finish job channel ")
		default:
			xiLogger.Log.Error("Job is finished and but cannot insert the result into finish job channel ")
		}
	}
	wg.Done()
}

func createWebAppGeneratorsPool(noOfWorkers int) {
	xiLogger.Log.Info("Creating generator worker pool..." )
	var wg sync.WaitGroup
	for i := 0; i < noOfWorkers; i++ {
		wg.Add(1)
		go webAppGenerator(&wg)
	}
	wg.Wait()
}

func notifier(wg *sync.WaitGroup)  {

	for result := range finishedJobs {
		xiLogger.Log.Debug(" Finished Job result , going to notify of job ID",strconv.Itoa(int(result.result.SolutionId)))
		notify(&result)
	}
	wg.Done()
}

func createNotifierPool(noOfWorkers int) {
	xiLogger.Log.Info("Creating Notifier worker pool..." )
	var wg sync.WaitGroup
	for i := 0; i < noOfWorkers; i++ {
		wg.Add(1)
		go notifier(&wg)
	}
	wg.Wait()
}

func notify(result *GenResult)  {

	if err := result.CallbackUrl.Send(&result.result); err != nil {
		xiLogger.Log.Error("send error %v", err)
	}
}

/*func restNotify(result *GenResult)  {

	client := resty.New()
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(result).Post(result.CallbackUrl)

	if err != nil {
		 xiLogger.Log.Error("Error while nonflying to remote web-studio",err)
	 //	finishedJobs <- *result
	}

	if resp != nil {
		if resp.IsSuccess() {
			xiLogger.Log.Info("Successfully Delivered :",resp.Status())
		}
	}
}
*/

func generateJobRequest(job *GenJob) bool {
	xiLogger.Log.Debug("New Generate request")
	status := make(chan bool)
	go addToJobQue(job,status)
	result:= <- status
	if result {
		return true
	}else {
		return false
	}
}

func generateJobRequestFromREST(c echo.Context) error {

	xiLogger.Log.Debug("New Generate request")
	job := new(GenJob)
	er1 := c.Bind(job)
	if er1 != nil {
		return c.JSON(http.StatusBadRequest,job)
	}
	result:=  generateJobRequest(job)
	if result {
		return c.JSON(http.StatusOK,job)
	}else {
		return c.JSON(http.StatusServiceUnavailable,job)
	}
}

func generateJobRequestFromRPC(job *pub.GenJob, srv pub.Generate_GenerateServer) bool  {

	genjob :=new(GenJob)
	genjob.SolutionId = int(job.GetSolutionId())
	genjob.SolutionName = job.GetSolutionName()
	genjob.ProjectUUId = job.GetProjectUUId()
	genjob.JobDescription = job.GetDescription()
	genjob.CallbackUrl = srv
	result:=  generateJobRequest(genjob)

	return result
}

func addToJobQue(genJob *GenJob,respond chan<- bool) {

	select {
	case generateJobs <- *genJob: // Put 2 in the channel unless it is full
		log.Println(" <<<<< Inserting a Job = "+strconv.Itoa(genJob.SolutionId))
		xiLogger.Log.Debug(" << Inserting a Job to chanel , jobId : = "+strconv.Itoa(genJob.SolutionId))
		respond <- true
	default:
		fmt.Println("Channel full. Discarding value")
		xiLogger.Log.Warn("Couldn't insert into channel ,Channel may full. Discarding value")
		respond <- false
	}
}


func updateGit(basePath string,solutionId int,solutionName string,branchName string) string {
	xiLogger.Log.Debug(" Getting resources from the Git At :"+branchName)
	solutionDir := basePath+"/"+strconv.Itoa(solutionId)+"_"+strings.TrimSpace(solutionName)
	projectDir := solutionDir +"/"+branchName

	if _, err := os.Stat(solutionDir); os.IsNotExist(err) {
		repoName := strconv.Itoa(solutionId)+"_"+strings.TrimSpace(solutionName)
		er := githubclient.CheckoutRemoteRepository(env.GetServerRepoUrl(repoName),env.GetGenDirectory())
		if er != nil {
			fmt.Println(err)
		}
	}

	githubclient.PullFromRemoteRepository(solutionDir,branchName)
	return filepath.FromSlash(projectDir)
}


func pushUpdateToGit (targetDir string, commitMsg string,branchName string,username string){

	xiLogger.Log.Debug("Pushing update to github >>")
	githubclient.PushDevUpdateToGit(targetDir,commitMsg,branchName,username)

}


func Generate(job *GenJob)   {
	xiLogger.Log.Debug("Web App generator starting the job ")
	solutionDir := env.GetGenDirectory()+"/"+strconv.Itoa(job.SolutionId)+"_"+strings.TrimSpace(job.SolutionName)
	xiLogger.Log.Info(" SolutionDir : ", solutionDir)
	gitBranchName :=job.ProjectUUId
	projectDir := updateGit(env.GetGenDirectory(),job.SolutionId,job.SolutionName,gitBranchName)
	xiLogger.Log.Info(" projectDir : ", projectDir)
	//defer pushUpdateToGit(solutionDir,"")

	generateInit(projectDir)

	if strings.EqualFold("Async",env.GenMode) {
		xiLogger.Log.Debug(" Generating mode :: Async  ")
		var wg sync.WaitGroup

		go generateBackendAsync(&wg)
		go generateFrontendAsync(&wg, projectDir)

		wg.Wait()
	} else if strings.EqualFold("BF",env.GenMode)  {
		xiLogger.Log.Debug(" Generating mode :: BF  ")
		generator.GenerateBackend()
		generator.GenerateFrontend(projectDir)

	} else if strings.EqualFold("FB",env.GenMode) {
		xiLogger.Log.Debug(" Generating mode :: FB  ")
		generator.GenerateFrontend(projectDir)
		generator.GenerateBackend()
	}else {
		xiLogger.Log.Warn(" Generating mode :: Not found ")
		xiLogger.Log.Info(" Generating mode :: FB  ")
		generator.GenerateFrontend(projectDir)
		generator.GenerateBackend()
	}
	// job.description = username
	pushUpdateToGit(solutionDir,"generated",gitBranchName,job.JobDescription)
	//	job.Status = "done"
}

func generateBackendAsync(wg *sync.WaitGroup)  {
	defer wg.Done()
	generator.GenerateBackend()

}

func generateFrontendAsync(wg *sync.WaitGroup ,projectDir string)  {
	defer wg.Done()
	generator.GenerateFrontend(projectDir)
}

func generateInit(projectDir string)  {
	generator.GenerateInit(projectDir)
}

func GenerateFromFile(projectDir string)  {
	generator.GenerateInit(projectDir)
	generator.GenerateBackend()
	generator.GenerateFrontend(projectDir)
}