package commands

import (
	"fmt"

	"github.com/codegangsta/cli"
	"github.com/rchrd/xke-cli/xke"
)

var Sessions = cli.Command{
	Name:      "sessions",
	ShortName: "s",
	Usage:     "List sessions of a XKE or Innovation Day",
	ArgsUsage: "<yyyy-mm-dd>",
	Action:    sessions,
}

func sessions(c *cli.Context) {
	client := xke.NewClient(c.GlobalString("token"))
	var event xke.Event
	event, _ = client.NextEvent()
	fmt.Println(event)
	sessions, _ := client.Sessions(event.Date)
	for _, s := range sessions {
		fmt.Println(s.PrintSummary())
	}
}
