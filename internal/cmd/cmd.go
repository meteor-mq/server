// Copyright (c) 2020  Shuo Ding
// Open Source: MIT License
// Author: SDing <deen.job@qq.com>
// License that can be found in the LICENSE file.

package cmd

import (
	"time"

	"github.com/fast-mq/server/internal/app"
	"github.com/higker/logker"
	"github.com/urfave/cli"
)

var (

	// whether enable  backgrund running
	isBackgrund bool

	// Start program start command.
	Start = cli.Command{
		// Command name
		Name: "start",
		// Command abbreviation
		Aliases: []string{"up"},
		// Command usage notes, here will show how to use the command when you enter the program name -help
		Usage: "FastMQ program start command✅.\n",
		// Command processing function
		Action: func(c *cli.Context) error {
			//c.Args().First()
			app.EchoInfo()

			if isBackgrund {
				logker.Info("%s", "FastMQ serve start up.")
				logker.Info("%s", "FastMQ is enable daemon mode running.")
			} else {
				logker.Info("%s", "FastMQ serve start up.")
			}
			time.Sleep(time.Second * 5)
			return nil
		},
	}
	// Stop program stop command.
	Stop = cli.Command{

		Name: "stop",

		Aliases: []string{"down"},

		Usage: "FastMQ program stop command❎.\n",

		Action: func(c *cli.Context) error {
			// STOP
			return nil
		},
	}
)

func init() {
	Start.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:        "daemon, d",
			Usage:       "Whether the program is running in the background.",
			Destination: &isBackgrund,
		},
	}
}
