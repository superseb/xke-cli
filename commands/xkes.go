package commands

import (
	"fmt"

	"github.com/codegangsta/cli"
	"github.com/rchrd/xke-cli/xke"
)

var XKEs = cli.Command{
	Name:      "xkes",
	ShortName: "x",
	Usage:     "List all future XKE's (and Innovation Days)",
	Action:    xkes,
	Flags: []cli.Flag{
		cli.BoolFlag{
			Name:  "all, a",
			Usage: "List all XKE's, including past ones",
		},
	},
}

func xkes(c *cli.Context) {
	client := getClient(c)
	var xkes []xke.XKE
	if c.IsSet("all") {
		xkes, _ = client.AllXKEs()
	} else {
		xkes, _ = client.FutureXKEs()
	}

	for _, e := range xkes {
		fmt.Println(e)
	}
}
