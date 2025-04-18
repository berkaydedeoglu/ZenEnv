package cli

import (
	urfave "github.com/urfave/cli/v3"
)

type Cli struct {
	cli *urfave.Command
}

func NewCli() *Cli {
	return &Cli{
		cli: &urfave.Command{
			Name:     "ZenEnv",
			Usage:    "a pure environment variable tool",
			Commands: []*urfave.Command{},
		},
	}
}

func (c *Cli) registerServerCommands() {
}
