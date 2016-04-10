package commands

import (
	"fmt"

	"github.com/codegangsta/cli"
	"github.com/rchrd/xke-cli/xke"
)

var AgendaCommand = cli.Command{
	Name:      "agenda",
	ShortName: "a",
	Usage:     "List sessions of a XKE or Innovation Day",
	ArgsUsage: "<yyyy-mm-dd>",
	Action:    Sessions,
}

func Sessions(c *cli.Context) {
	client := getClient(c)
	var xke xke.XKE
	var d string
	if c.NArg() == 1 {
		d = c.Args().First()
	} else {
		xke, _ = client.NextXKE()
		d = xke.Date
	}
	fmt.Printf("Agenda for the XKE of %s\n\n", d)
	sessions, _ := client.Sessions(d)
	for i := range sessions {
		if i > 0 && sessions[i-1].StartTime != sessions[i].StartTime {
			fmt.Println("")
		}
		fmt.Println(sessions[i].PrintSummary())
	}
}
