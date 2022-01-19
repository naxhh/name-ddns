package main

import (
	"fmt"
	"log"
	"github.com/naxhh/name-ddns/internal/name"
	"os"
	"flag"
)

func main() {

	config := name.NewConfig()

	var initialRun bool

	flag.BoolVar(&initialRun, "initial-run", false, "Is this initial run")
	flag.Parse()

	if err := config.Validate(); err != nil {
		log.Println(fmt.Println(err))
		os.Exit(1)
	}

	app := name.New(config)

	if initialRun {
		log.Println(fmt.Sprintf("Running with cron expression %s", config.UpdateCron))
	}
	app.Run()
}
