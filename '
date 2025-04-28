package repository

import (
	"github.com/berkaydedeoglu/zenenv/internal/config"
	"github.com/berkaydedeoglu/zenenv/pkg/db"
)

const (
	BBoltRepository       string = "bbolt"
	SecureBBoltRepository string = "secure-bbolt"
)

type Factory struct {
	config *config.Config
}

func NewFactory(c *config.Config) *Factory {
	return &Factory{config: c}
}

func (f *Factory) CreateRepository(repoType string) Repository {
	switch repoType {
	case BBoltRepository:
		return f.CreateBBolt()
	default:
		return nil
	}
}

func (f *Factory) CreateBBolt() *Bbolt {
	db := db.NewDb(f.config)
	return NewBbolt(db, f.config)
}
