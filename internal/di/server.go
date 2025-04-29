//go:build wireinject
// +build wireinject

package di

import (
	"github.com/berkaydedeoglu/zenenv/internal/api"
	"github.com/berkaydedeoglu/zenenv/internal/api/handlers"
	"github.com/berkaydedeoglu/zenenv/internal/config"
	"github.com/berkaydedeoglu/zenenv/internal/repository"
	"github.com/berkaydedeoglu/zenenv/internal/service"
	"github.com/berkaydedeoglu/zenenv/pkg/printer"
	"github.com/google/wire"
)

func InitializeServer() (*api.Server, error) {
	wire.Build(
		api.NewServer,
		api.NewApiRouter,
		handlers.NewHealthHandler,
		handlers.NewZenenvHandler,
		service.NewZenService,
		repository.NewFactory,
		wire.Bind(
			new(repository.RepositoryFactory), // hedef: interface
			new(*repository.Factory),          // kaynak: struct
		),
		printer.NewPrinter,
		config.NewConfig)
	return &api.Server{}, nil
}
