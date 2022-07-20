package main

import (
	"github.com/naxhh/name-ddns/internal/name"
	"github.com/naxhh/name-ddns/internal/system"
	"log"
	"os"
)

func main() {
	stopChannel := system.GetSignalNotifier()
	config := name.NewConfig(stopChannel)

	if err := config.Validate(); err != nil {
		log.Println(err)
		os.Exit(1)
	}

	app := name.New(config)

	app.Run()

	log.Println("Shuting down")
}
