package env


const APP_PORT  = "PORT"
const BASE_PATH  = "BASE_PATH"
const WP_SIZE  = "WP" // worker pool size
const GEN_MODE  = "GM"  // code generation mode  BF , FB , Async
const GIT_SERVER  = "GIT_SERVER"
const YO_UIGEN  =  "YO_UIGEN"
const LOG_LEVEL  =  "LOG_LEVEL"

/*   ############################################### System Constants ###################################### */

const defaultBasePath  = "/tmp/pixri/gen/mobileapp"
const defaultPort  = "5005"
const genDirectory = "/gen"
const defaultGitServerPath = "ssh://git@localhost:2222/git-server/repos"
const defaultGid ="io.pixri.generator.mobileApp"
const defaultPoolSize  = 1
const defaultGenMode  = "BF"

/*   ############################################### System Variables ###################################### */

var Gid string
var BasePath string
var Port string
var GitServerURL string
var WorkerPoolSize int
var GenMode string
var LogLevel string

func SetDefaultBasePath(){
	BasePath = defaultBasePath
}

func SetDefaultGitServer()  {
	GitServerURL = defaultGitServerPath
}

func SetDefaultAppPort()  {
	Port = defaultPort
}

func GetGenDirectory() string {
	return BasePath + genDirectory
}

func SetDefaultGeneratorId() {
	Gid = defaultGid
}

func SetDefaultPoolSize()  {
	WorkerPoolSize = defaultPoolSize
}

func SetDefaultGenMode()  {
	GenMode = defaultGenMode
}

func GetServerRepoUrl(repoName string) string {
	return GitServerURL + "/" + repoName
}


func GetLoglevel()string  {

	if LogLevel == ""{
		return "debug"
	}
	return LogLevel
}