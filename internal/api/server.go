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

func (s *Server) StartServer(host string, port string) *fiber.App {
	// TODO: This appends 7 lines of code to this method. It
	// will be better if we can fetch from config. Hence config can be changed
	// on cli method
	if host != "" {
		s.host = host
	}

	if port != "" {
		s.port = port
	}

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
