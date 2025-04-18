package main

import (
	"context"
	"log"
	"os"

	"github.com/berkaydedeoglu/zenenv/internal/cli"
)

func main() {
	cli := cli.NewCli()
	if err := cli.Cli.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
