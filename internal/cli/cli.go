package cli

import (
	urfave "github.com/urfave/cli/v3"
)

type Cli struct {
	Cli *urfave.Command
}

func NewCli() *Cli {
	serviceCommands := registerServerCommands()
	return &Cli{
		Cli: &urfave.Command{
			Name:  "ZenEnv",
			Usage: "a pure environment variable tool",
			Commands: []*urfave.Command{
				serviceCommands,
			},
		},
	}
}

func registerServerCommands() *urfave.Command {
	return &urfave.Command{
		Name:  "service",
		Usage: "to manage zenenv service (server)",
		Commands: []*urfave.Command{
			{
				Name:   "up",
				Usage:  "starts zenenv server",
				Action: StartServer,
			},
		},
	}
}
