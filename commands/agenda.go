package commands

import (
	"fmt"

	"github.com/codegangsta/cli"
	"github.com/rchrd/xke-cli/xke"
)

var Agenda = cli.Command{
	Name:      "agenda",
	ShortName: "a",
	Usage:     "List sessions of a XKE or Innovation Day",
	ArgsUsage: "<yyyy-mm-dd>",
	Action:    sessions,
}

func sessions(c *cli.Context) {
	client := getClient(c)
	var xke xke.XKE
	var d string
	if c.NArg() == 1 {
		d = c.Args().First()
	} else {
		xke, _ = client.NextXKE()
		d = xke.Date
	}
	fmt.Println(d)
	sessions, _ := client.Sessions(d)
	for _, s := range sessions {
		fmt.Println(s.PrintSummary())
	}
}
