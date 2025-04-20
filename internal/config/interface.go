package config

type ConfigPart interface {
	SetDefaults() error
}
