package commands

import (
	"fmt"

	"github.com/codegangsta/cli"
	"github.com/fatih/color"
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
	bold := color.New(color.Bold).SprintFunc()
	client := getClient(c)
	var xke xke.XKE
	var d string
	if c.NArg() == 1 {
		d = c.Args().First()
	} else {
		xke, _ = client.NextXKE()
		d = xke.Date
	}
	fmt.Println(bold("--------------------------------"))
	fmt.Printf(bold("Agenda for the XKE of %s\n"), d)
	fmt.Println(bold("--------------------------------"))
	fmt.Println("")

	sessions, _ := client.Sessions(d)

	if len(sessions) == 0 {
		fmt.Println("No sessions found")
		return
	}

	for i := range sessions {
		if i > 0 && sessions[i-1].StartTime != sessions[i].StartTime {
			fmt.Println("")
		}
		fmt.Println(sessions[i].PrintSummary())
	}
}
