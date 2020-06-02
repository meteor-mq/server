// Copyright (c) 2020  Shuo Ding
// Open Source: MIT License
// Author: SDing <deen.job@qq.com>
// License that can be found in the LICENSE file.

package cmd

import (
	"github.com/fast-mq/server/internal/app"
	"github.com/urfave/cli"
)

// Start program start command.
var Start = cli.Command{
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
		return nil
	},
}

// Stop program stop command.
var Stop = cli.Command{

	Name: "stop",

	Aliases: []string{"down"},

	Usage: "FastMQ program stop command❎.\n",

	Action: func(c *cli.Context) error {
		// STOP
		return nil
	},
}
