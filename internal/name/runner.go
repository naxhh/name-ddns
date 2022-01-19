package name

import (
	"log"
	"fmt"
)

type runner struct {
	api    *api
}

func New(config *Config) *runner {
	p := &runner{
		api:    newApi(config),
	}

	return p
}

func (p *runner) Run() {
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
