package controller

import (
	"errors"
	"fmt"
	"github.com/google/go-github/github"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"os"
	"os/exec"
	"pixri_generator/pixriLogger"
	"strconv"
	"time"
)

func CreateRepository(appName string) (*github.Repository, *github.Response, error)  {
	ctx := context.Background()
	now := time.Now()
	name := appName+"-"+strconv.Itoa(int(now.Unix()))+"-Pixri"
	fmt.Println(name)
	repo := &github.Repository{
		Name:    github.String(name),
		Private: github.Bool(true),
	}

	return getClient().Repositories.Create(ctx, "", repo)


}

func getClient() *github.Client{
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: "4126f1e5b8e5cf4ce71aaf9f3bec4e8cdad1b051"},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	return client
}

func exe(targetDir string,command string,args []string) bool{
	cmd := exec.Command(command, args...)
	cmd.Dir = targetDir
	cmd.Stderr = os.Stderr
	if pixriLogger.IsDebugEnabled() {
		pixriLogger.Log.Debug("Location : @  ", targetDir)
		pixriLogger.Log.Debug(cmd.Args)
		cmd.Stderr = os.Stdout
	}
	err := cmd.Run()
	if err != nil {
		pixriLogger.Logger().Error(err)
		return false
	}
	return true
}

func exeClone(url string, localDir string) error {
	pixriLogger.Log.Debug(" cloning started : ",url)
	arguments :=[]string{"clone", url}
	val :=exe(localDir,"git",arguments)
	pixriLogger.Log.Debug("command status : ",val)
	if !val {
		return errors.New("Cloning  Failed")
	}
	return nil
}


