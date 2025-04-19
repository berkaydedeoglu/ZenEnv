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
			Name:  "zenenv",
			Usage: "a zen way of managing environment variables",
			Commands: []*urfave.Command{
				serviceCommands,
			},
		},
	}
}

func registerServerCommands() *urfave.Command {
	return &urfave.Command{
		Name:  "server",
		Usage: "to manage zenenv service (server)",
		Commands: []*urfave.Command{
			{
				Name:   "up",
				Usage:  "starts zenenv server",
				Action: StartServer,
				Flags: []urfave.Flag{
					&urfave.StringFlag{
						Name:  "host",
						Usage: "host ip address of server",
					},
					&urfave.StringFlag{
						Name:  "port",
						Usage: "port of given server",
					},
				},
			},
		},
	}
}
