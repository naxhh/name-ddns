package system

import (
	"os"
	"os/signal"
	"syscall"
)

func GetSignalNotifier() chan struct{} {
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(
		signalChannel,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)

	stopChannel := make(chan struct{})

	go func() {
		<-signalChannel
		stopChannel <- struct{}{}
	}()

	return stopChannel
}
