package config

import (
	"flag"
	"github.com/caarlos0/env/v6"
)

type Config struct {
	ServerAddress string `env:"SERVER_ADDRESS"`
	BaseURL       string `env:"BASE_URL"`
}

func NewConfig() Config {

	address := flag.String("a", "localhost:8080", "input address")
	url := flag.String("b", "http://localhost:8080", "input URL")
	flag.Parse()

	cfg := Config{
		ServerAddress: *address,
		BaseURL:       *url,
	}
	_ = env.Parse(&cfg) // перезапишет флаги, если переменные есть

	return cfg
}
