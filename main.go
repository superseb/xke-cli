package main

import (
	"os"

	"github.com/codegangsta/cli"
)

type Session struct {
	ID          int
	SessionType string `json:"session_type"`
	Title       string
	Presenter   string
	XKE         string
	StartTime   string `json:"start_time"`
	EndTime     string `json:"end_time"`
	Unit        string
	Location    string
	Goal        string
	Summary     string
	Preparation string
}

func main() {
	app := cli.NewApp()
	app.Name = "xke-cli"
	app.Version = Version
	app.Usage = "Get information about XKE's and Innovation Days"
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "Richard Woudenberg",
			Email: "rwoudenberg@xebia.com",
		},
	}
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "token",
			Usage: "Authorization token",
		},
	}
	app.Commands = Commands

	app.Run(os.Args)
}
