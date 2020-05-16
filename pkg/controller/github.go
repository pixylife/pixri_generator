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

var owner = "pixylife"

var (
	authorName    = flag.String("author-name", "pixylife", "Name of the author of the commit.")
	authorEmail   = flag.String("author-email", "sahanvijaya@gmail.com", "Email of the author of the commit.")
	commitMessage = flag.String("commit-message", "Adding Project", "Content of the commit message.")
)

func CreateRepository(appName string) (*github.Repository, *github.Response, error) {
	ctx := context.Background()
	now := time.Now()
	name := appName + "-" + strconv.Itoa(int(now.Unix())) + "-Pixri"
	fmt.Println(name)
	repo := &github.Repository{
		Name:    github.String(name),
		Private: github.Bool(true),
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

func exeClone(url string, localDir string) error {
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

	parent, _, err := getClient().Repositories.GetCommit(ctx, owner, repoName, *ref.Object.SHA)
	if err != nil {
		return err
	}
	parent.Commit.SHA = parent.SHA

	date := time.Now()
	author := &github.CommitAuthor{Date: &date, Name: authorName, Email: authorEmail}
	commit := &github.Commit{Author: author, Message: commitMessage, Tree: tree, Parents: []github.Commit{*parent.Commit}}

	newCommit, _, err := getClient().Git.CreateCommit(ctx, owner, repoName, commit)
	if err != nil {
		return err
	}
	ref.Object.SHA = newCommit.SHA
	_, _, err = getClient().Git.UpdateRef(ctx, owner, repoName, ref, false)

	return err
}

func GetTree(ref *github.Reference, rootpath string, repoName string) (tree *github.Tree, err error) {
	ctx := context.Background()

	entries := []github.TreeEntry{}

	fileList := ListAllFiles(rootpath)
	// Load each file into the tree.
	for _, fileArg := range *fileList {
		file, content, err := GetFileContent(fileArg)
		if err != nil {
			return nil, err
		}
		entries = append(entries, github.TreeEntry{Path: github.String(file), Type: github.String("blob"), Content: github.String(string(content)), Mode: github.String("100644")})
	}

	tree, _, err = getClient().Git.CreateTree(ctx, owner, repoName, *ref.Object.SHA, entries)
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
