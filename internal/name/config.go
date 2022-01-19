package name

import (
	"os"
)

type Config struct {
	User  string
	Token string
	UpdateCron string
	Domain string
	Host   string
}

func NewConfig() *Config {

	return &Config{
		User:  os.Getenv("NAME_DDNS_USER"),
		Token: os.Getenv("NAME_DDNS_TOKEN"),
		UpdateCron: os.Getenv("NAME_DDNS_UPDATE_CRON"),
		Domain: os.Getenv("NAME_DDNS_DOMAIN"),
		Host:   os.Getenv("NAME_DDNS_HOST"),
	}
}

func (c *Config) Validate() *error {
	// TODO: validate config
	return nil
}
