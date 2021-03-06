package command

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/codegangsta/cli"
)

func executeCommand(str []string) (out string, err error) {
	var buf []byte
	buf, err = exec.Command(str[0], str[1:]...).Output()
	out = string(buf)
	return
}

func CmdPush(c *cli.Context) {
	pushCodes()
}

func pushCodes() {
	commands := [][]string{
		[]string{"git", "add", "."},
		[]string{"git", "commit", "-m", "Initial commit"},
		[]string{"git", "push", "origin", "master"},
	}
	for _, command := range commands {
		fmt.Println(command)
		out, err := executeCommand(command)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v", err)
			// Not return because it's not critical
		}
		fmt.Println(out)
	}
}
