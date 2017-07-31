package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/negibokken/lab/command"
)

var GlobalFlags = []cli.Flag{}

var Commands = []cli.Command{
	{
		Name:   "create",
		Usage:  "",
		Action: command.CmdCreate,
		Flags:  []cli.Flag{},
	},
	{
		Name:   "push",
		Usage:  "",
		Action: command.CmdPush,
		Flags:  []cli.Flag{},
	},
	{
		Name:   "all",
		Usage:  "",
		Action: command.CmdAll,
		Flags:  []cli.Flag{},
	},
}

func CommandNotFound(c *cli.Context, command string) {
	fmt.Fprintf(os.Stderr, "%s: '%s' is not a %s command. See '%s --help'.", c.App.Name, command, c.App.Name, c.App.Name)
	os.Exit(2)
}
