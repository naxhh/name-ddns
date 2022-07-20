package name

import (
	"log"
	"github.com/robfig/cron/v3"
)

type poller struct {
	stop   chan struct{}
	api    *api
	cron   *cron.Cron
	config *Config
}

func New(config *Config) *poller {
	p := &poller{
		stop:   config.StopChannel,
		api:    newApi(config),
		cron:   cron.New(),
		config: config,
	}

	return p
}

func (p *poller) Run() {
	p.cron.AddFunc(p.config.UpdateEveryCronFormat, func() { p.run() })
	p.cron.Start()

	for {
		select {
		case <-p.stop:
			p.cron.Stop()
			return
		}
	}
}

func (p *poller) run() {
	ip, err := p.api.getIp()

	if err != nil {
		log.Println("Failed retrieving IP:", err, "no changes will be performed this execution")
		return
	}

	if err = p.api.update(ip); err != nil {
		log.Println("Failed updating record", err, "no changes will be performed this execution")
		return
	}
}
