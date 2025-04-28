package db

import (
	"fmt"
	"time"

	"github.com/berkaydedeoglu/zenenv/internal/config"
	err "github.com/berkaydedeoglu/zenenv/pkg/error"
	bolt "go.etcd.io/bbolt"
)

const FileName = "zenenv.db"

type Db struct {
	Database *bolt.DB
	Config   *config.Config
}

func NewDb(config *config.Config) *Db {
	return &Db{
		Config: config,
	}
}

func (d *Db) Open() error {
	if d.Database != nil {
		return nil
	}

	opt := &bolt.Options{Timeout: time.Duration(d.Config.Server.DBTimeout) * time.Second}
	path := fmt.Sprintf("%s/%s", d.Config.Server.DataPath, FileName)

	db, err := bolt.Open(path, 0600, opt)
	if err != nil {
		return err
	}

	d.Database = db
	return nil
}

func (d *Db) Close() error {
	if d.Database == nil {
		return err.ErrDbIsNotOpened
	}
	return d.Database.Close()
}

func (d *Db) Clear() error {
	return nil
}
