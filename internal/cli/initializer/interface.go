package initializer

import "github.com/berkaydedeoglu/zenenv/internal/api"

type Initializer interface {
	Server() *api.Server
}
