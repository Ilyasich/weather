package config

import (
	"os"

	"github.com/caarlos0/env"
	"github.com/pelletier/go-toml/v2"
)

type Config struct {
	Server ServerConfig
}

type ServerConfig struct {
	DebugMode  bool   `env:"DEBUG" toml:"debug_mode"`
	ServerHost string `env:"SERVER_HOST" envDefault:":8080" toml:"server_host"`
}

func Read() (Config, error) {
	cfg := ServerConfig{}
	if err := env.Parse(&cfg); err != nil {
		return Config{}, err
	}

	return Config{
		Server: cfg,
	}, nil
}

func ReadFile(path string) (Config, error) {
	configFile, err := os.ReadFile(path)
	if err != nil {
		return Config{}, err
	}

	srvCfg := ServerConfig{}
	err = toml.Unmarshal(configFile, &srvCfg)
	return Config{
		srvCfg,
	}, err
}
