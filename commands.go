package main

import (
	"fmt"

	"github.com/codegangsta/cli"
	"github.com/rchrd/xke-cli/xke"
)

// Commands is an array containing the available commands.
var Commands = []cli.Command{
	commandList,
	commandSessions,
}

var commandList = cli.Command{
	Name:      "list",
	ShortName: "l",
	Usage:     "List all XKE's and Innovation days",
	Action:    list,
	Flags: []cli.Flag{
		cli.BoolFlag{
			Name:  "all, a",
			Usage: "List all events, including past ones",
		},
	},
}

var commandSessions = cli.Command{
	Name:      "sessions",
	ShortName: "s",
	Usage:     "List sessions of a XKE or Innovation Day",
	ArgsUsage: "<yyyy-mm-dd>",
	Action:    sessions,
}

func list(c *cli.Context) {
	client := xke.NewClient(c.GlobalString("token"))
	var events []xke.Event
	if c.IsSet("all") {
		events, _ = client.AllEvents()
	} else {
		events, _ = client.FutureEvents()
	}

	for _, e := range events {
		fmt.Println(e)
	}
}

func sessions(c *cli.Context) {
	client := xke.NewClient(c.GlobalString("token"))
	var event xke.Event
	event, _ = client.NextEvent()
	fmt.Println("Next event: ", event)
}
