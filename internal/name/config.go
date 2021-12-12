package name

import (
	"os"
	"strconv"
	"time"
)

type Config struct {
	User  string
	Token string

	UpdateEvery time.Duration
	StopChannel chan struct{}

	Domain string
	Host   string
}

func NewConfig(stopChannel chan struct{}) *Config {
	updateMinutes, err := strconv.ParseInt(os.Getenv("NAME_DDNS_UPDATE_EVERY_MINUTES"), 10, 0)

	if err != nil {
		updateMinutes = 30
	}

	return &Config{
		User:  os.Getenv("NAME_DDNS_USER"),
		Token: os.Getenv("NAME_DDNS_TOKEN"),

		UpdateEvery: time.Duration(updateMinutes) * time.Minute,
		StopChannel: stopChannel,

		Domain: os.Getenv("NAME_DDNS_DOMAIN"),
		Host:   os.Getenv("NAME_DDNS_HOST"),
	}
}

func (c *Config) Validate() *error {
	// TODO: validate config
	return nil
}
