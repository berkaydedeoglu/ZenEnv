package main

import (
	"context"
	"log"
	"os"

	"github.com/berkaydedeoglu/zenenv/internal/cli"
	"github.com/berkaydedeoglu/zenenv/internal/cli/initializer"
	"github.com/berkaydedeoglu/zenenv/internal/config"
	"github.com/berkaydedeoglu/zenenv/pkg/printer"
)

func main() {
	cfg := config.NewConfig()
	initializer := initializer.NewDiInitializer()
	printer := printer.NewPrinter()

	cli := cli.NewCli(cfg, initializer, printer)
	if err := cli.Cli.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
