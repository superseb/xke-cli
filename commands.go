package main

import (
	"github.com/codegangsta/cli"
	"github.com/rchrd/xke-cli/commands"
)

// Commands is an array containing the available commands.
var Commands = []cli.Command{
	commands.Events,
	commands.Sessions,
	commands.Details,
	commands.Locations,
}
