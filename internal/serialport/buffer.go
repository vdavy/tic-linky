package serialport

import (
	"bufio"
	"errors"
	"github.com/tarm/serial"
	"io"
	"sync"
)

// createBufferReader create the buffer reader and send data to channel
func createBufferReader(port *serial.Port) *SerialPort {
	buffer := bufio.NewReader(port)
	streamingChan := make(chan string)
	exitChan := make(chan bool)
	exitWG := &sync.WaitGroup{}
	exitWG.Add(1)

	go handleReader(buffer, streamingChan, exitChan, exitWG)

	return &SerialPort{
		serialPort:    port,
		StreamingChan: streamingChan,
		ExitChan:      exitChan,
		ExitWG:        exitWG,
	}
}

// handleReader handle the reader incoming data
func handleReader(buffer *bufio.Reader, streamingChan chan<- string, exitChan chan<- bool, exitWG *sync.WaitGroup) {
	keepGoing := true
	for keepGoing {
		line, err := buffer.ReadString(lineSeparator)
		if err != nil {
			if errors.Is(err, io.EOF) { // io.EOF means closing serial port
				keepGoing = false
				exitChan <- true
				close(streamingChan)
				close(exitChan)
			} else {
				logger.WithError(err).Warn("Error reading the serial port")
			}
		} else {
			streamingChan <- line // send the received line
		}
	}

	exitWG.Done()
	logger.Debug("Serial port reader closed")
}
