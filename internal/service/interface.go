package service

import "github.com/berkaydedeoglu/zenenv/pkg/zenvar"

type Service interface {
	WriteZenVar(*zenvar.ZenVar) error
	ReadZenVar(string) (*zenvar.ZenVar, error)
	GetProjects() []string
	GetPlatforms(string) []string
	GetTags(string, string) []string

	Close() error
	Clear() error
}
