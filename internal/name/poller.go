package name

import (
	"log"
	"fmt"
	"github.com/robfig/cron/v3"
)

type poller struct {
	cronValue string
	stop   chan struct{}
	api    *api
}

func New(config *Config) *poller {
	p := &poller{
		cronValue: config.UpdateCron,
		stop: config.StopChannel,
		api:    newApi(config),
	}

	return p
}

func (p *poller) Run() {
	p.run()
}

func updater(p *poller) {
	ip, err := p.api.getIp()

	if err != nil {
		log.Println(fmt.Sprintf("Failed retrieving IP:", err, "no changes will be performed this execution"))
		return
	}

	if err = p.api.update(ip); err != nil {
		log.Println(fmt.Sprintf("Failed updating record", err, "no changes will be performed this execution"))
		return
	}
}

func (p *poller) run() {

	updater(p)

	log.Println(fmt.Sprintf("Updating with cron schedule: %s", p.cronValue))

	c := cron.New()
	c.AddFunc(p.cronValue, func() {
		updater(p)
	})
	c.Start()

	for {
		select {
		case <-p.stop:
			c.Stop()
			return
		}
	}
}
