package common

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var onlyOneSignalHandler = make(chan struct{})
var shutdownSignals = []os.Signal{os.Interrupt, syscall.SIGTERM}

func SetupSignalHandler() (stopChan <-chan struct{}) {
	close(onlyOneSignalHandler)

	stop := make(chan struct{})
	c := make(chan os.Signal, 2)
	signal.Notify(c, shutdownSignals...)
	go func() {
		<-c
		close(stop)
		<-c
		os.Exit(1)
	}()

	return stop
}

func TestChan() {
	stopChan := SetupSignalHandler()
	fmt.Println("TestChan start.")
	<-stopChan
	time.Sleep(time.Second * 2)
	fmt.Println("TestChan end.")
}
