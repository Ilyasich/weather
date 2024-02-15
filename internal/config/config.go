package config

import (
	"os"

	"github.com/caarlos0/env"
	"github.com/pelletier/go-toml/v2"
)

const (
	Apikey string = "3caf85347f7e49e481d110120241401"
	City   string = "Batumi"
	Lang   string = "RU"
)


//конфигурационный параметр
type Config struct {
	Server ServerConfig
}


//конфигурационный параметр
type ServerConfig struct {
	DebugMode  bool   `env:"DEBUG" toml:"debug_mode"`
	ServerHost string `env:"SERVER_HOST" envDefault:":8080" toml:"server_host"`
}


// считывает значения конфигурации из переменных окружения. Значения из переменных окружения загружаются в структуру `ServerConfig`, которая затем используется для создания объекта `Config`.
func Read() (Config, error) {
	cfg := ServerConfig{}
	if err := env.Parse(&cfg); err != nil {
		return Config{}, err
	}

	return Config{
		Server: cfg,
	}, nil
}

//считывает значения конфигурации из файла по указанному пути с использованием пакета `toml`
func ReadFile(path string) (Config, error) {
	configFile, err := os.ReadFile(path)
	if err != nil {
		return Config{}, err
	}
//Обе функции возвращают объект `Config` с заполненными значениями конфигурации или ошибку, если чтение конфигурации не удалось.
	srvCfg := ServerConfig{}
	err = toml.Unmarshal(configFile, &srvCfg)
	return Config{
		srvCfg,
	}, err
}
