package commands

import (
	"fmt"
	"os"
	"strconv"

	"github.com/codegangsta/cli"
)

var SessionCommand = cli.Command{
	Name:      "session",
	ShortName: "s",
	Usage:     "Show the details of a session",
	ArgsUsage: "<id>",
	Action:    Session,
}

func Session(c *cli.Context) {
	client := getClient(c)
	if c.NArg() == 0 {
		fmt.Println("ERROR - Please specify the id of the session")
		os.Exit(1)
	}
	id, _ := strconv.Atoi(c.Args().First())
	session, _ := client.Session(id)
	fmt.Println(session.PrintDetails())
}
