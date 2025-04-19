package config

import "reflect"

type Config struct {
	Server *ServerConfig
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) ReadConfig(path string) error {
	return nil
}

func (c *Config) SetDefaults(configPart ConfigPart) error {
	if reflect.ValueOf(configPart).IsNil() {
		return configPart.SetDefaults()
	}
	return c.setAllDefaults()
}

func (c *Config) setAllDefaults() error {
	return c.Server.SetDefaults()
}
