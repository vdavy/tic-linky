package serialport

import (
	"github.com/sirupsen/logrus"
	"github.com/tarm/serial"
	"sync"
	"tic-linky/internal/config"
	"time"
)

const (
	baudRate      = 9600
	frameSize     = 7
	readTimeout   = 5 * time.Second
	lineSeparator = '\n'
)

var logger = logrus.WithField("logger", "serial")

type SerialPort struct {
	serialPort    *serial.Port
	StreamingChan chan string
	ExitChan      chan bool
	ExitWG        *sync.WaitGroup
}

// CreateSerialPort open the serial port
func CreateSerialPort() *SerialPort {
	port, err := serial.OpenPort(&serial.Config{
		Name:        config.SerialPort,
		Baud:        baudRate,
		Parity:      serial.ParityEven,
		Size:        frameSize,
		StopBits:    serial.Stop1,
		ReadTimeout: readTimeout,
	})

	if err != nil {
		logger.WithError(err).Fatal("Error opening serial port")
	}

	serialPort := createBufferReader(port)
	logger.Infof("Serial port opened at %s", config.SerialPort)

	return serialPort
}

// Close close the serial port and manage error
func (serialPort *SerialPort) Close() {
	if err := serialPort.serialPort.Close(); err != nil {
		logger.WithError(err).Error("Error closing serial port")
	} else {
		logger.Info("Serial port closed")
	}
}
