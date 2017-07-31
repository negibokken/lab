package command

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
)

var GitLabApiUrl string
var UserID string
var PrivateToken string

func setEnvironMent() {
	for _, e := range os.Environ() {
		fmt.Println(e)
	}
}

func init() {
	setEnvironMent()
}

func CmdCreate(c *cli.Context) {
	// Write your code here
}
