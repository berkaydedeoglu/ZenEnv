package config

type ServerConfig struct {
	Enabled bool
	Host    string
	Port    string
}

func (s *ServerConfig) SetDefaults() error {
	s.Enabled = true
	s.Host = "127.0.0.1"
	s.Port = "6565"
	return nil
}

func (s *ServerConfig) SetBindPoints(host string, port string) {
	if host != "" {
		s.Host = host
	}

	if port != "" {
		s.Port = port
	}
}
