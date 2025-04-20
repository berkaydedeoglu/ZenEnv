package initializer

import (
	"log"

	"github.com/berkaydedeoglu/zenenv/internal/api"
	"github.com/berkaydedeoglu/zenenv/internal/di"
)

type DiInitializer struct{}

func NewDiInitializer() *DiInitializer {
	return &DiInitializer{}
}

func (i *DiInitializer) Server() *api.Server {
	s, err := di.InitializeServer()
	if err != nil {
		log.Fatalf("failed to initialize server: %v", err)
		panic("zenenv exited with error")
	}
	return s
}
