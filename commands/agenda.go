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
	xke, _ = client.NextXKE()
	fmt.Println(xke)
	sessions, _ := client.Sessions(xke.Date)
	for _, s := range sessions {
		fmt.Println(s.PrintSummary())
	}
}
