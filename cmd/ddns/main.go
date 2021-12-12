package main

import (
	"fmt"
	"github.com/naxhh/name-ddns/internal/name"
	"github.com/naxhh/name-ddns/internal/system"
	"os"
)

func main() {
	stopChannel := system.GetSignalNotifier()
	config := name.NewConfig(stopChannel)

	if err := config.Validate(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	app := name.New(config)

	app.Run()

	fmt.Println("Shuting down")
}
