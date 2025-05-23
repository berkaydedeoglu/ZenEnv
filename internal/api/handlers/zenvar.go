package handlers

import (
	"github.com/berkaydedeoglu/zenenv/internal/config"
	"github.com/berkaydedeoglu/zenenv/internal/service"
	"github.com/gofiber/fiber/v2"
)

type zenenvHandler struct {
	config  *config.Config
	service *service.ZenService
}

func NewZenenvHandler(c *config.Config, z *service.ZenService) *zenenvHandler {
	return &zenenvHandler{
		config:  c,
		service: z,
	}
}

func (z *zenenvHandler) HandleHealth(ctx *fiber.Ctx) error {
	ctx.WriteString("OK")
	return nil
}
