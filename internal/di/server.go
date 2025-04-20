//go:build wireinject
// +build wireinject

package di

import (
	"github.com/berkaydedeoglu/zenenv/internal/api"
	"github.com/berkaydedeoglu/zenenv/internal/api/handlers"
	"github.com/berkaydedeoglu/zenenv/internal/config"

	"github.com/google/wire"
)

func InitializeServer() (*api.Server, error) {
	wire.Build(
		api.NewServer,
		api.NewApiRouter,
		handlers.NewHealthHandler,
		config.NewConfig,
	)
	return &api.Server{}, nil
}
