package name

import (
	"github.com/robfig/cron/v3"
	"log"
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
		api:    newApi(),
		cron:   cron.New(),
		config: config,
	}

	return p
}

func (p *poller) Run() {
	for _, task := range p.config.Tasks {
		func(task Task) {
			_, err := p.cron.AddFunc(task.CronExpression, func() { p.run(task) })
			if err != nil {
				panic(err)
			}
		}(task)
	}
	p.cron.Start()

	for {
		<-p.stop
		p.cron.Stop()
		return
	}
}

func (p *poller) run(task Task) {
	ip, err := p.api.getIp(task)

	if err != nil {
		log.Printf("Failed retrieving IP for %s.%s: %s no changes will be performed this execution", task.Host, task.Domain, err)
		return
	}

	if err = p.api.update(task, ip); err != nil {
		log.Printf("Failed updating record for %s.%s: %s no changes will be performed this execution", task.Host, task.Domain, err)
		return
	}
}
