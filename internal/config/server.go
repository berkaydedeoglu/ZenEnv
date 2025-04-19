package config

type ServerConfig struct {
	Host string
	Port string
}

func (s *ServerConfig) SetDefaults() error {
	s.Host = "127.0.0.1"
	s.Port = "6565"
	return nil
}
