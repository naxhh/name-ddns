package name

import (
	"os"
)

type Config struct {
	User  string
	Token string

	UpdateEveryCronFormat string
	StopChannel chan struct{}

	Domain string
	Host   string
}

func NewConfig(stopChannel chan struct{}) *Config {
	cronUpdateTime := os.Getenv("NAME_DDNS_UPDATE_EVERY")

	return &Config{
		User:  os.Getenv("NAME_DDNS_USER"),
		Token: os.Getenv("NAME_DDNS_TOKEN"),

		UpdateEveryCronFormat: cronUpdateTime,
		StopChannel: stopChannel,

		Domain: os.Getenv("NAME_DDNS_DOMAIN"),
		Host:   os.Getenv("NAME_DDNS_HOST"),
	}
}

func (c *Config) Validate() *error {
	// TODO: validate config
	return nil
}
