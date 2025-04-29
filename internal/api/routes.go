package api

import (
	"github.com/berkaydedeoglu/zenenv/internal/api/handlers"
	"github.com/gofiber/fiber/v2"
)

type ApiRouter struct {
	healthHandler *handlers.HealthHandler
	zenenvHandler *handlers.ZenenvHandler
}

func NewApiRouter(
	healthHandler *handlers.HealthHandler,
	zenenvHandler *handlers.ZenenvHandler,
) *ApiRouter {
	return &ApiRouter{
		healthHandler: healthHandler,
		zenenvHandler: zenenvHandler,
	}
}

func (a ApiRouter) RegisterHealthRoutes(app *fiber.App) {
	hr := app.Group(("/health"))
	hr.Get("/", a.healthHandler.HandleHealth)
}
