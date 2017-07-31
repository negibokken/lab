package command

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/codegangsta/cli"
)

var GitLabApiUrl string
var UserID string
var PrivateToken string

func checkEnv() {
	flags := map[string]bool{
		"GL_ENDPOINT": true,
		"GL_USER":     true,
		"GL_PRIVATE":  true,
	}
	if GitLabApiUrl == "" {
		GitLabApiUrl = "https://api.gitlab.com/v3"
		fmt.Fprintf(os.Stderr, "Environement variable GL_ENDPOINT is not set. Default endpoint https://api.gitlab.com/v3 is used")
	}
	if UserID == "" {
		flags["GL_USER"] = false
		fmt.Fprintf(os.Stderr, "Environement variable GL_USER is not set")
	}
	if PrivateToken == "" {
		flags["GL_PRIVATE"] = false
		fmt.Fprintf(os.Stderr, "Environement variable GL_PRIVATE is not set")
	}
	keys := []string{}
	for k := range flags {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		v := flags[k]
		if !v {
			fmt.Fprintf(os.Stderr, "%s is not set", k)
		}
	}
}

func setEnvironMent() {
	for _, e := range os.Environ() {
		splited := strings.Split(e, "=")
		if len(splited) <= 1 {
			fmt.Fprintf(os.Stderr, "Environement variable is not set")
			os.Exit(1)
		}
		key, value := splited[0], splited[1]
		switch key {
		case "GL_ENDPOINT":
			GitLabApiUrl = value
		case "GL_USER":
			UserID = value
		case "GL_PRIVATE":
			PrivateToken = value
		}
	}
}

func init() {
	setEnvironMent()
}

func CmdCreate(c *cli.Context) {
	err := createRepository()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	}
}

func currentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return filepath.Base(dir)
}

func createRepository() error {
	projectName := currentDirectory()
	resource := "/projects?name=" + projectName
	req, err := http.NewRequest(
		"POST", GitLabApiUrl+resource, nil,
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", err)
		return err
	}
	req.Header.Set("Content-Type", "application/x-www-urlencoded")
	req.Header.Add("PRIVATE-TOKEN", PrivateToken)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", err)
		return err
	}
	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", err)
		return err
	}
	fmt.Println(string(buf))
	defer resp.Body.Close()
	return nil
}
