package config

import (
	"fmt"
	"os"
	"sync"

	errors "github.com/berkaydedeoglu/zenenv/pkg/error"
	"github.com/pelletier/go-toml/v2"
)

type Config struct {
	Server *ServerConfig
}

var (
	singleton *Config
	once      sync.Once
)

func NewConfig() *Config {
	once.Do(createConfig)
	return singleton
}

func (c *Config) ReadConfig(path string) error {
	data, err := findConfigFile(path)
	if err != nil {
		return err
	}

	if err = toml.Unmarshal(data, &c); err != nil {
		return err
	}

	return nil
}

func (c *Config) SetDefaultsForConfigPart(configPart ConfigPart) error {
	return configPart.SetDefaults()
}

func (c *Config) setAllDefaultsForAllConfigParts() error {
	return c.Server.SetDefaults()
}

func createConfig() {
	c := &Config{
		Server: &ServerConfig{},
	}
	if err := c.ReadConfig(""); err != nil {
		fmt.Println(err)
		c.setAllDefaultsForAllConfigParts()
	}
	singleton = c
}

func findConfigFile(path string) ([]byte, error) {
	if path == "" {
		return findFileInDefaults()
	}

	if data, err := os.ReadFile(path); err == nil {
		return data, nil
	}

	return nil, errors.ErrSpecifiedFileNotFound
}

func findFileInDefaults() ([]byte, error) {
	for _, path := range DEFAULT_CONFIG_FILE_PATH {
		if data, err := os.ReadFile(path); err == nil {
			return data, nil
		}
	}
	return nil, errors.ErrFileNotFoundInDefaults
}
