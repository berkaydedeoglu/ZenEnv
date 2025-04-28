package api

import (
	"fmt"

	"github.com/berkaydedeoglu/zenenv/internal/api/middleware/log"
	"github.com/berkaydedeoglu/zenenv/internal/config"
	"github.com/berkaydedeoglu/zenenv/pkg/printer"
	"github.com/gofiber/fiber/v2"
)

type Server struct {
	Api     *fiber.App
	routes  *ApiRouter
	config  *config.ServerConfig
	printer *printer.Printer
}

func NewServer(
	routes *ApiRouter,
	config *config.Config,
	printer *printer.Printer,
) *Server {
	return &Server{
		config:  config.Server,
		routes:  routes,
		printer: printer,
	}
}

func (s *Server) StartServer() *fiber.App {
	s.bootstrapServer()
	s.routes.RegisterHealthRoutes(s.Api)

	u := fmt.Sprintf("%s:%s", s.config.Host, s.config.Port)

	s.Api.Use(log.New(s.config, s.printer))
	s.Api.Listen(u)
	return s.Api
}

func (s *Server) bootstrapServer() {
	if s.Api == nil {
		s.Api = fiber.New(fiber.Config{AppName: "ZenEnv"})
	}
}
