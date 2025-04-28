package service

import (
	"github.com/berkaydedeoglu/zenenv/internal/config"
	"github.com/berkaydedeoglu/zenenv/internal/repository"
	"github.com/berkaydedeoglu/zenenv/pkg/zenvar"
)

type ZenService struct {
	config     *config.Config
	repository repository.Repository
}

func NewZenService(
	c *config.Config,
	rf repository.RepositoryFactory,
) *ZenService {
	return &ZenService{
		config:     c,
		repository: rf.CreateRepository(c.Server.DBStrategy),
	}
}

func (zs *ZenService) WriteZenVar(zv *zenvar.ZenVar) error {
	return nil
}

func (zs *ZenService) ReadZenVar(key string) (*zenvar.ZenVar, error) {
	return nil, nil
}

func (zs *ZenService) GetProjects() []string {
	return []string{}
}

func (zs *ZenService) GetPlatforms(project string) []string {
	return []string{}
}

func (zs *ZenService) GetTags(project, platform string) []string {
	return []string{}
}

func (zs *ZenService) Close() error {
	return nil
}

func (zs *ZenService) Clear() error {
	return nil
}
