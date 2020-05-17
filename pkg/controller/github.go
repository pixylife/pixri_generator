package controller

import (
	"errors"
	"flag"
	"fmt"
	"github.com/google/go-github/github"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"pixri_generator/pixriLogger"
	"strconv"
	"strings"
	"time"
)


var (
	authorName    = flag.String("author-name", "pixylife", "Name of the author of the commit.")
	authorEmail   = flag.String("author-email", "sahanvijaya@gmail.com", "Email of the author of the commit.")
	commitMessage = flag.String("commit-message", "Adding Project", "Content of the commit message.")
	baseBranch    = flag.String("base-branch", "master", "Name of branch to create the `commit-branch` from.")
	sourceOwner   = flag.String("source-owner", "pixylife", "Name of the owner (user or org) of the repo to create the commit in.")
	commitBranch  = flag.String("commit-branch", "dev", "Name of branch to create the commit in. If it does not already exists, it will be created using the `base-branch` parameter")

)

func CreateRepository(appName string) (*github.Repository, *github.Response, error) {
	ctx := context.Background()
	now := time.Now()
	name := appName + "-" + strconv.Itoa(int(now.Unix())) + "-Pixri"
	fmt.Println(name)
	repo := &github.Repository{
		Name:    github.String(name),
		Private: github.Bool(false),
	}

	return getClient().Repositories.Create(ctx, "", repo)

}

func getClient() *github.Client {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: "4126f1e5b8e5cf4ce71aaf9f3bec4e8cdad1b051"},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	return client
}

func exe(targetDir string, command string, args []string) bool {
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

func GitClone(url string, localDir string) error {
	pixriLogger.Log.Debug(" cloning started : ", url)
	arguments := []string{"clone", url}
	val := exe(localDir, "git", arguments)
	pixriLogger.Log.Debug("command status : ", val)
	if !val {
		return errors.New("Cloning  Failed")
	}
	return nil
}

func CreateCommitPush(ref *github.Reference, tree *github.Tree, repoName string) error {
	ctx := context.Background()

	parent, _, err := getClient().Repositories.GetCommit(ctx,*sourceOwner, repoName, *ref.Object.SHA)
	if err != nil {
		return err
	}
	parent.Commit.SHA = parent.SHA

	date := time.Now()
	author := &github.CommitAuthor{Date: &date, Name: authorName, Email: authorEmail}
	commit := &github.Commit{Author: author, Message: commitMessage, Tree: tree, Parents: []github.Commit{*parent.Commit}}

	newCommit, _, err := getClient().Git.CreateCommit(ctx, *sourceOwner, repoName, commit)
	if err != nil {
		return err
	}
	ref.Object.SHA = newCommit.SHA
	_, _, err = getClient().Git.UpdateRef(ctx, *sourceOwner, repoName, ref, false)

	return err
}

func GetTree(ref *github.Reference, rootpath string, repoName string) (tree *github.Tree, err error) {
	ctx := context.Background()

	var entries []github.TreeEntry

	fileList := ListAllFiles(rootpath)
	// Load each file into the tree.
	for _, fileArg := range *fileList {
		file, content, err := GetFileContent(fileArg)
		if err != nil {
			return nil, err
		}
		entries = append(entries, github.TreeEntry{Path: github.String(file), Type: github.String("blob"), Content: github.String(string(content)), Mode: github.String("100644")})
	}
	tree, _, err = getClient().Git.CreateTree(ctx,*sourceOwner, repoName, *ref.Object.SHA, entries)
	return tree, err
}

func GetFileContent(fileArg string) (targetName string, b []byte, err error) {
	var localFile string
	files := strings.Split(fileArg, ":")
	switch {
	case len(files) < 1:
		return "", nil, errors.New("empty `-files` parameter")
	case len(files) == 1:
		localFile = files[0]
		targetName = files[0]
	default:
		localFile = files[0]
		targetName = files[1]
	}

	b, err = ioutil.ReadFile(localFile)
	return targetName, b, err
}

func ListAllFiles(root string) *[]string {

	var files []string

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {

		if !info.IsDir() {
			files = append(files, path)
		}

		return nil
	})
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		fmt.Println(file)
	}
	return &files
}

func GetRef(repoName string) (ref *github.Reference, err error) {

	ctx := context.Background()

	if ref, _, err = getClient().Git.GetRef(ctx, *sourceOwner, repoName, "refs/heads/"+*commitBranch); err == nil {
		return ref, nil
	}

	// We consider that an error means the branch has not been found and needs to
	// be created.
	if *commitBranch == *baseBranch {
		return nil, errors.New("The commit branch does not exist but `-base-branch` is the same as `-commit-branch`")
	}

	if *baseBranch == "" {
		return nil, errors.New("The `-base-branch` should not be set to an empty string when the branch specified by `-commit-branch` does not exists")
	}

	var baseRef *github.Reference
	if baseRef, _, err = getClient().Git.GetRef(ctx, *sourceOwner, repoName, "refs/heads/"+*baseBranch); err != nil {
		fmt.Println(err)
		return nil, err
	}
	newRef := &github.Reference{Ref: github.String("refs/heads/" + *commitBranch), Object: &github.GitObject{SHA: baseRef.Object.SHA}}
	ref, _, err = getClient().Git.CreateRef(ctx, *sourceOwner, repoName, newRef)
	return ref, err
}


func SetRemote(url string, localDir string) error {
	pixriLogger.Log.Debug(" Setting remote started : ", url)
	arguments := []string{"remote", "add", "origin", url}
	val := exe(localDir, "git", arguments)
	pixriLogger.Log.Debug("command status : ", val)
	if !val {
		return errors.New("failed")
	}
	return nil
}

func GitInit(localDir string) error {
	pixriLogger.Log.Debug(" Git config")
	arguments := []string{"init"}
	val := exe(localDir, "git", arguments)
	pixriLogger.Log.Debug("command status : ", val)
	if !val {
		return errors.New("failed")
	}
	return nil
}

func GitPush(targetDir string,branchName string)  {
	pixriLogger.Log.Debug("exeGitpush for : ",branchName)

	arguments :=[]string{"push", "origin", branchName}
	val :=exe(targetDir,"git",arguments)
	pixriLogger.Log.Debug("command status : ",val)
}


func GitCommit(targetDir string,message string)  {
	pixriLogger.Log.Debug("exe Git commit")
	arguments :=[]string{"commit", "-m", "\""+message+"\""}
	val :=exe(targetDir,"git",arguments)
	pixriLogger.Log.Debug("command status : ",val)
}


func GitAddAll(targetDir string)  {
	pixriLogger.Log.Debug("exe Git add")
	arguments :=[]string{"add", "."}
	val :=exe(targetDir,"git",arguments)
	pixriLogger.Log.Debug("command status : ",val)
}

func exeCheckoutBranch(targetDir string,branchName string)  {
	pixriLogger.Log.Debug(" exeCheckoutBranch ")
	arguments :=[]string{"checkout","-f","-B",branchName}
	val :=exe(targetDir,"git",arguments)
	if  val {
		defer exeClearRootDirectory(targetDir,branchName)
	}
	pixriLogger.Log.Debug("command status : ",val)
}

func exeGitPull(targetDir string, branchName string)  {
	pixriLogger.Log.Debug(" exeGitPull ")
	arguments :=[]string{"pull", "origin",branchName}
	val :=exe(targetDir,"git",arguments)
	pixriLogger.Log.Debug("command status : ",val)
}


func exeClearRootDirectory(targetDir string,branchName string)  {
	pixriLogger.Log.Debug(" clear other resources int the directory : ",targetDir)
	arguments :=[]string{"reset", "--hard","origin/master"}
	val :=exe(targetDir,"git",arguments)
	pixriLogger.Log.Debug("command status : ",val)

	arguments =[]string{"clean", "-f","-d"}
	val =exe(targetDir,"git",arguments)
	pixriLogger.Log.Debug("command status : ",val)
}

