package processing

import (
	"github.com/sirupsen/logrus"
	"sync"
)

var (
	logger = logrus.WithField("logger", "processing")

	// currentFrameData the current frame data for holding context
	currentFrameData *frameData
)

// StartProcessing start the processing for the TIC data
func StartProcessing(streamingChan <-chan string, extChan <-chan bool, exitWG *sync.WaitGroup) {
	initCurrentFrameData() // init data struct before starting parser

	go internalProcessing(streamingChan, extChan, exitWG)
}

// internalProcessing internal processing function running in another thread
func internalProcessing(streamingChan <-chan string, extChan <-chan bool, exitWG *sync.WaitGroup) {
	func(streamingChan <-chan string, extChan <-chan bool, exitWG *sync.WaitGroup) {
		keepGoing := true
		for keepGoing {
			select {
			case line := <-streamingChan:
				currentFrameData.processLine(line)
				if detectEndOfFrame(line) { // init current frame data for the next new line
					initCurrentFrameData()
				}
			case <-extChan:
				keepGoing = false
			}
		}
		logger.Debug("Exiting processing")
		exitWG.Done()
	}(streamingChan, extChan, exitWG)
}

// initCurrentFrameData init current frame data with non nil value
func initCurrentFrameData() {
	currentFrameData = &frameData{
		indexMap: make(map[string]uint64),
		powerMap: make(map[string]uint64),
	}
}
