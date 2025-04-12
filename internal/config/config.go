package config

import "flag"

type Config struct {
	ServerAddress string `json:"server_address"`
	BaseURL       string `json:"base_url"`
}

func NewConfig() Config {
	address := flag.String("a", "localhost:8080", "input address")
	url := flag.String("b", "http://localhost:8080", "input URL")

	flag.Parse()

	return Config{
		ServerAddress: *address,
		BaseURL:       *url}

}
