package name

import (
	"errors"
	"fmt"
	"log"
	"os"
)

type Task struct {
	CronExpression string

	User  string
	Token string

	Domain string
	Host   string
}

type Config struct {
	StopChannel chan struct{}
	Tasks       []Task
}

func NewConfig(stopChannel chan struct{}) *Config {
	defaultUser := os.Getenv("NAME_DDNS_USER")
	defaultToken := os.Getenv("NAME_DDNS_TOKEN")

	var tasks []Task

	i := 0
	for {
		domain := os.Getenv(fmt.Sprintf("NAME_DDNS_DOMAIN_%d", i))

		if domain == "" {
			break
		}

		cron := os.Getenv(fmt.Sprintf("NAME_DDNS_CRON_%d", i))
		host := os.Getenv(fmt.Sprintf("NAME_DDNS_HOST_%d", i))

		user := os.Getenv(fmt.Sprintf("NAME_DDNS_USER_%d", i))
		token := os.Getenv(fmt.Sprintf("NAME_DDNS_TOKEN_%d", i))

		if user == "" {
			user = defaultUser
		}

		if token == "" {
			token = defaultToken
		}

		tasks = append(tasks, Task{
			CronExpression: cron,
			User:           user,
			Token:          token,
			Domain:         domain,
			Host:           host,
		})

		i++
	}

	return &Config{
		StopChannel: stopChannel,
		Tasks:       tasks,
	}
}

func (c *Config) Validate() error {
	if len(c.Tasks) == 0 {
		return errors.New("No entries configured, closing")
	} else {
		log.Printf("Validating %d entries", len(c.Tasks))
	}

	for key, task := range c.Tasks {
		if task.CronExpression == "" {
			return fmt.Errorf("Cron expression was empty on entry %d", key)
		}

		if task.User == "" {
			return fmt.Errorf("User was empty on entry %d", key)
		}

		if task.Token == "" {
			return fmt.Errorf("Token was empty on entry %d", key)
		}

		if task.Domain == "" {
			return fmt.Errorf("Domain was empty on entry %d", key)
		}

		if task.Host == "" {
			return fmt.Errorf("Host was empty on entry %d", key)
		}
	}
	return nil
}
