package commands

import (
	"fmt"
	"os"
	"strconv"

	"github.com/codegangsta/cli"
	"github.com/rchrd/xke-cli/xke"
)

var Details = cli.Command{
	Name:      "details",
	ShortName: "d",
	Usage:     "Show the details of a session",
	ArgsUsage: "<id>",
	Action:    details,
}

func details(c *cli.Context) {
	client := xke.NewClient(c.GlobalString("token"))
	if c.NArg() == 0 {
		fmt.Println("ERROR - Please specify the id of the session")
		os.Exit(1)
	}
	id, _ := strconv.Atoi(c.Args().First())
	session, _ := client.Session(id)
	fmt.Println(session.PrintDetails())
}
