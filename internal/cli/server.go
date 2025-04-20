package cli

import (
	"context"

	"github.com/urfave/cli/v3"
)

func (c *Cli) StartServer(ctx context.Context, cmd *cli.Command) error {
	c.config.Server.SetBindPoints(cmd.String("host"), cmd.String("port"))

	server := c.initializer.Server()
	server.StartServer()
	return nil
}
