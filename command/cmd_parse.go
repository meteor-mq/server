package command

import (
	"fmt"

	"github.com/spf13/pflag"
	"go.coder.com/cli"
)

// HelpCmd help command
type HelpCmd struct {
	// This root command has no default action, so print the help.
}

// command param
type argsCmd struct {
	// config file path
	cfgPath string
	// Whether to run in the background
	isBackground bool
}

func (c *argsCmd) Run(fl *pflag.FlagSet) {
	if c.isBackground {
		fmt.Println("application is background.")
	}
}

// GetHelpCmd get help command root
func GetHelpCmd() *HelpCmd {
	return &HelpCmd{}
}

func (c *argsCmd) Spec() cli.CommandSpec {
	return cli.CommandSpec{
		Name:  "-back",
		Usage: "",
		Desc:  `Whether to run in the background.`,
	}
}

type cmd struct {
}

func (c *HelpCmd) Run(fl *pflag.FlagSet) {
	// This root command has no default action, so print the help.
	fl.Usage()
}

func (c *HelpCmd) Spec() cli.CommandSpec {
	return cli.CommandSpec{
		Name:  "FastMQ",
		Usage: "[-help]",
		Desc:  `Use the  command to view the help message.`,
	}
}

func (c *HelpCmd) Subcommands() []cli.Command {
	return []cli.Command{
		&argsCmd{},
	}
}
