package config

var DEFAULT_CONFIG_FILE_PATH = [7]string{
	"/etc/zenenv/config.toml",
	"/etc/zenenv/zenenv.toml",
	"~/.zenenv.toml",
	"~/.config/zenenv/config.toml",
	"~/.config/zenenv.toml",
	"./config.toml",
	"./zenenv.toml",
}
