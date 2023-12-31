package main

import (
	"os"
	"os/signal"
	"syscall"
	"tic-linky/internal/serialport"
)

// waitForSignal block and wait for the exit signal
func waitForSignal(serialPort *serialport.SerialPort) {
	quitCh := make(chan os.Signal, 1)
	signal.Notify(quitCh, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)
	//lint:ignore S1000 wait for close signal
	for {
		select {
		case c := <-quitCh:
			logger.Infof("Exiting - signal %s", c.String())
			serialPort.Close()
			return
		}
	}
}
