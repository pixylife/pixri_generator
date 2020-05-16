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

var owner = "pixylife"

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


/*
func PushCommit(ref *github.Reference, tree *github.Tree,repoName string) (err error) {
	ctx := context.Background()
	parent, _, err := getClient().Repositories.GetCommit(ctx,owner, repoName, *ref.Object.SHA)
	if err != nil {
		return err
	}
	parent.Commit.SHA = parent.SHA

	date := time.Now()
	author := &github.CommitAuthor{Date: &date, Name: authorName, Email: authorEmail}
	commit := &github.Commit{Author: author, Message: , Tree: tree, Parents: []github.Commit{*parent.Commit}}
	newCommit, _, err := getClient().Git.CreateCommit(ctx, owner, repoName, commit)
	if err != nil {
		return err
	}
	ref.Object.SHA = newCommit.SHA
	_, _, err = getClient().Git.UpdateRef(ctx, owner, repoName, ref, false)

	return err
}*/