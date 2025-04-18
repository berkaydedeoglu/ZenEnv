package cli

import (
	"context"
	"log"

	"github.com/berkaydedeoglu/zenenv/internal/di"
	"github.com/urfave/cli/v3"
)

func StartServer(ctx context.Context, cmd *cli.Command) error {
	server, err := di.InitializeServer()
	if err != nil {
		log.Fatalf("failed to initialize server: %v", err)
		return err
	}

	server.StartServer()
	return nil
}
