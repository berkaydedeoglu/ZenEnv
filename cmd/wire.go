//go:build wireinject

package main

import (
	"github.com/berkaydedeoglu/zenenv/internal/api"
	"github.com/berkaydedeoglu/zenenv/internal/api/handlers"

	"github.com/google/wire"
)

func InitializeServer() (*api.Server, error) {
	wire.Build(
		handlers.NewHealthHandler,
		api.NewApiRouter,
		api.NewServer,
	)
	return &api.Server{}, nil
}
