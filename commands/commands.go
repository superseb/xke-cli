package commands

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/rchrd/xke-cli/xke"
)

// Commands is an array containing the available commands.
var Commands = []cli.Command{
	XKEs,
	Agenda,
	Session,
}

func getClient(c *cli.Context) *xke.Client {
	client, err := xke.NewClient(c.GlobalString("token"))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return client
}
