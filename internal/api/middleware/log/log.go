package log

import (
	"fmt"
	"time"

	"github.com/berkaydedeoglu/zenenv/internal/config"
	"github.com/berkaydedeoglu/zenenv/pkg/printer"
	"github.com/gofiber/fiber/v2"
)

func New(c *config.ServerConfig, p *printer.Printer) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		url := ctx.OriginalURL()
		ip := ctx.IP()
		t := time.Now().Local().Format("yyyy-MM-dd HH:mm:ss Z")

		info := fmt.Sprintf("%s | %s | %s", t, ip, url)
		p.Info(info)

		return ctx.Next()
	}
}
