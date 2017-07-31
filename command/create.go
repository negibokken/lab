package command

import (
	"fmt"
	"os"
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
		flags["GL_ENDPOINT"] = false
		fmt.Fprintf(os.Stderr, "Environement variable is not set")
	}
	if UserID == "" {
		flags["GL_USER"] = false
		fmt.Fprintf(os.Stderr, "Environement variable is not set")
	}
	if PrivateToken == "" {
		flags["GL_PRIVATE"] = false
		fmt.Fprintf(os.Stderr, "Environement variable is not set")
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
	fmt.Println(UserID)
	fmt.Println(PrivateToken)
	fmt.Println(GitLabApiUrl)
}
