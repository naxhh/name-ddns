package name

import (
	"log"
	"time"
)

type poller struct {
	ticker *time.Ticker
	stop   chan struct{}
	api    *api
}

func New(config *Config) *poller {
	p := &poller{
		ticker: time.NewTicker(config.UpdateEvery),
		stop:   config.StopChannel,
		api:    newApi(config),
	}

	return p
}

func (p *poller) Run() {
	p.run()

	for {
		select {
		case <-p.ticker.C:
			p.run()
		case <-p.stop:
			p.ticker.Stop()
			return
		}
	}
}

func (p *poller) run() {
	ip, err := p.api.getIp()

	if err != nil {
		log.Println("Failed retrieving ip:", err, "no changes will be performed this execution")
		return
	}

	if err = p.api.update(ip); err != nil {
		log.Println("Failed updating record", err, "no changes will be performed this execution")
		return
	}
}
