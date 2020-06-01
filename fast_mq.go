package main

import (
	"github.com/fast-mq/server/command"
	"go.coder.com/cli"
)

func main() {
	cli.RunRoot(command.GetHelpCmd())
}
