// Copyright (c) 2020  Shuo Ding
// Open Source: MIT License
// Author: SDing <deen.job@qq.com>
// License that can be found in the LICENSE file.

package main

import (
	"os"

	"github.com/fast-mq/server/internal/app"
	"github.com/fast-mq/server/internal/cmd"
	"github.com/higker/logker"
	"github.com/urfave/cli"
)

func main() {
	appCli := cli.NewApp()

	appCli.Name = "Fast Message Queue"

	appCli.Usage = "is light weight message queue middleware 🚀."

	appCli.Version = app.Verison

	appCli.Commands = []cli.Command{
		cmd.Start,
		cmd.Stop,
	}
	err := appCli.Run(os.Args)
	if err != nil {
		logker.Error("application init error:%v", err)
	}
}
