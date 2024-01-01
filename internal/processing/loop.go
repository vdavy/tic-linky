package processing

import (
	"github.com/sirupsen/logrus"
	"sync"
)

var logger = logrus.WithField("logger", "processing")

func StartProcessing(streamingChan <-chan string, extChan <-chan bool, exitWG *sync.WaitGroup) {
	go func(streamingChan <-chan string, extChan <-chan bool, exitWG *sync.WaitGroup) {
		keepGoing := true
		for keepGoing {
			select {
			case line := <-streamingChan:
				logger.Debugf("Received line : %s", line)
			case <-extChan:
				keepGoing = false
			}
		}
		logger.Debug("Exiting processing")
		exitWG.Done()
	}(streamingChan, extChan, exitWG)
}
