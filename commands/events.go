package commands

import (
	"fmt"

	"github.com/codegangsta/cli"
	"github.com/rchrd/xke-cli/xke"
)

var Events = cli.Command{
	Name:      "events",
	ShortName: "e",
	Usage:     "List all events (XKE's and Innovation Days)",
	Action:    events,
	Flags: []cli.Flag{
		cli.BoolFlag{
			Name:  "all, a",
			Usage: "List all events, including past ones",
		},
	},
}

func events(c *cli.Context) {
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
