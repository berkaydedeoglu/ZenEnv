package main

import (
	"context"
	"log"
	"os"

	"github.com/berkaydedeoglu/zenenv/internal/cli"
	"github.com/berkaydedeoglu/zenenv/internal/config"
)

func main() {
	cfg := config.NewConfig()
	cli := cli.NewCli(cfg)
	if err := cli.Cli.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
