package name

import (
	"os"
	"fmt"
	"log"
	"errors"
)

type Task struct {
	CronExpression string

	User string
	Token string

	Domain string
	Host string
}

type Config struct {
	StopChannel chan struct{}
	Tasks []Task
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
			User:  user,
			Token: token,
			Domain: domain,
			Host:   host,
		})


		i++
	}

	return &Config{
		StopChannel: stopChannel,
		Tasks: tasks,
	}
}

func (c *Config) Validate() error {
	if len(c.Tasks) == 0 {
		return errors.New("No entries configured, closing")
	} else {
		log.Println(fmt.Sprintf("Validating %d entries", len(c.Tasks)))
	}

	for key, task := range c.Tasks {
		if task.CronExpression == "" {
			return errors.New(fmt.Sprintf("Cron expression was empty on entry %d", key))
		}

		if task.User == "" {
			return errors.New(fmt.Sprintf("User was empty on entry %d", key))
		}

		if task.Token == "" {
			return errors.New(fmt.Sprintf("Token was empty on entry %d", key))
		}

		if task.Domain == "" {
			return errors.New(fmt.Sprintf("Domain was empty on entry %d", key))
		}

		if task.Host == "" {
			return errors.New(fmt.Sprintf("Host was empty on entry %d", key))
		}
	}
	return nil
}
