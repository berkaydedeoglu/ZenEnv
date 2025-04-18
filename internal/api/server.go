package api

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type Server struct {
	Api    *fiber.App
	routes *ApiRouter
	host   string
	port   string
}

func NewServer(
	routes *ApiRouter,
) *Server {
	return &Server{
		host:   "0.0.0.0",
		port:   "6567",
		routes: routes,
	}
}

func (s *Server) StartServer() *fiber.App {
	s.bootstrapServer()
	s.routes.RegisterHealthRoutes(s.Api)

	u := fmt.Sprintf("%s:%s", s.host, s.port)
	s.Api.Listen(u)
	return s.Api
}

func (s *Server) bootstrapServer() {
	if s.Api == nil {
		s.Api = fiber.New()
	}
}
