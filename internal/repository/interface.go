package repository

import "github.com/berkaydedeoglu/zenenv/pkg/zenvar"

type RepositoryFactory interface {
	CreateRepository(string) Repository
}

type Repository interface {
	WriteZenVar(*zenvar.ZenVar) error
	ReadZenVar(string) (*zenvar.ZenVar, error)
	GetProjects() []string
	GetPlatforms(string) []string
	GetTags(string, string, string) []string

	Connect() error
	Close() error
	Clear() error
}
