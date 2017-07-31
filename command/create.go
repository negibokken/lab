package command

import (
	"fmt"
	"io/ioutil"
	"net/http"
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
	err := createRepository()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	}
}

func createRepository() error {
	req, err := http.NewRequest(
		"POST", GitLabApiUrl, nil,
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
