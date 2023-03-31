package common

import "os"

var onlyOneSignalHandler = make(chan struct{})

func SetupSignalHandler() (stopChan <-chan struct{}) {
	close(onlyOneSignalHandler)

	stop := make(chan struct{})
	c := make(chan os.Signal, 2)

}
