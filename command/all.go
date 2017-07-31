package command

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
)

func CmdAll(c *cli.Context) {
	if err := createRepository(); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	}
	pushCodes()
}
