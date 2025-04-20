package api

import (
	"fmt"

	"github.com/berkaydedeoglu/zenenv/internal/config"
	"github.com/gofiber/fiber/v2"
)

type Server struct {
	Api    *fiber.App
	routes *ApiRouter
	config *config.ServerConfig
}

func NewServer(
	routes *ApiRouter,
	config *config.Config,
) *Server {
	return &Server{
		config: config.Server,
		routes: routes,
	}
}

func (s *Server) StartServer() *fiber.App {
	s.bootstrapServer()
	s.routes.RegisterHealthRoutes(s.Api)

	u := fmt.Sprintf("%s:%s", s.config.Host, s.config.Port)
	s.Api.Listen(u)
	return s.Api
}

func (s *Server) bootstrapServer() {
	if s.Api == nil {
		s.Api = fiber.New()
	}
}
