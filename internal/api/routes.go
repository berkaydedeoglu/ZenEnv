package api

import (
	"github.com/berkaydedeoglu/zenenv/internal/api/handlers"
	"github.com/gofiber/fiber/v2"
)

type ApiRouter struct {
	healthHandler *handlers.HealthHandler
}

func NewApiRouter(
	healthHandler *handlers.HealthHandler,
) *ApiRouter {
	return &ApiRouter{
		healthHandler: healthHandler,
	}
}

func (a ApiRouter) RegisterHealthRoutes(app *fiber.App) {
	hr := app.Group(("/health"))
	hr.Get("/", a.healthHandler.HandleHealth)
}
