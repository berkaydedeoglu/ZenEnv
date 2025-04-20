package cli

import (
	"github.com/berkaydedeoglu/zenenv/internal/cli/initializer"
	"github.com/berkaydedeoglu/zenenv/internal/config"
	"github.com/berkaydedeoglu/zenenv/pkg/printer"
	urfave "github.com/urfave/cli/v3"
)

type Cli struct {
	Cli         *urfave.Command
	config      *config.Config
	initializer initializer.Initializer
	printer     *printer.Printer
}

func NewCli(
	cfg *config.Config,
	i initializer.Initializer,
	p printer.Printer,
) *Cli {
	c := &Cli{
		config:      cfg,
		initializer: i,
		printer:     p,
		Cli: &urfave.Command{
			Name:     "zenenv",
			Usage:    "a zen way of managing environment variables",
			Commands: []*urfave.Command{},
		},
	}

	c.registerServerCommands()

	return c
}

func (c *Cli) registerServerCommands() {
	serverCommand := &urfave.Command{
		Name:  "server",
		Usage: "to manage zenenv service (server)",
		Commands: []*urfave.Command{
			{
				Name:   "up",
				Usage:  "starts zenenv server",
				Action: c.StartServer,
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

	c.Cli.Commands = append(c.Cli.Commands, serverCommand)
}
