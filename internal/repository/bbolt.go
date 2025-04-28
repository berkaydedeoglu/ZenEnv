package repository

import (
	"github.com/berkaydedeoglu/zenenv/internal/config"
	"github.com/berkaydedeoglu/zenenv/pkg/db"
	"github.com/berkaydedeoglu/zenenv/pkg/zenvar"
)

type Bbolt struct {
	db     *db.Db
	config *config.Config
}

func NewBbolt(db *db.Db, c *config.Config) *Bbolt {
	return &Bbolt{
		db:     db,
		config: c,
	}
}

func (b *Bbolt) WriteZenVar(zv *zenvar.ZenVar) error {
	return nil
}

func (b *Bbolt) ReadZenVar(v string) (*zenvar.ZenVar, error) {
	return nil, nil
}

func (b *Bbolt) GetProjects() []string {
	return []string{}
}

func (b *Bbolt) GetPlatforms(project string) []string {
	return []string{}
}

func (b *Bbolt) GetTags(project string, platform string, v string) []string {
	return []string{}
}

func (b *Bbolt) Connect() error {
	return nil
}

func (b *Bbolt) Close() error {
	return nil
}

func (b *Bbolt) Clear() error {
	return nil
}
