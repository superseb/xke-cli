package commands

import (
	"fmt"
	"strconv"

	"github.com/codegangsta/cli"
	"github.com/rchrd/xke-cli/xke"
)

var Locations = cli.Command{
	Name:      "locations",
	ShortName: "l",
	Usage:     "List locations",
	ArgsUsage: "<id>",
	Action:    locations,
}

func locations(c *cli.Context) {
	client := xke.NewClient(c.GlobalString("token"))
	if c.NArg() > 0 {
		id, _ := strconv.Atoi(c.Args().First())
		location, _ := client.Location(id)
		fmt.Println(location)
	}
	locations, _ := client.Locations()
	for _, l := range locations {
		fmt.Println(l)
	}
}
