package main

import (
	"github.com/sirupsen/logrus"
	"os"
	"periph.io/x/conn/v3/gpio"
	"tic-linky/internal/config"
	"tic-linky/internal/influxdb"
	"tic-linky/internal/led"
	"tic-linky/internal/processing"
	"tic-linky/internal/serialport"
)

const (
	undefinedValue = "N/A"
	debugMode      = "DEBUG"
)

var (
	version   = undefinedValue
	commit    = undefinedValue
	buildDate = undefinedValue
)

var logger = logrus.WithField("logger", "main")

func main() {
	setLoggerConf()
	logger.Infof("tic-linky v%s - commit %s - build date : %s", version, commit, buildDate)
	config.ParseEnvVars()

	// start all the stuff
	led.InitLedGpio()
	influxdb.NewInfluxdbClient()
	serialPort := serialport.CreateSerialPort()
	serialPort.ExitWG.Add(1)
	processing.StartProcessing(serialPort.StreamingChan, serialPort.ExitChan, serialPort.ExitWG)
	led.TurnLEDOnOff(gpio.High)

	// wait for exit signal
	waitForSignal(serialPort)
	serialPort.ExitWG.Wait()
	influxdb.CloseInfluxDBConnection()
	led.TurnLEDOnOff(gpio.Low)
	logger.Info("Final exit")
}

// setLoggerConf set the global logger configuration
func setLoggerConf() {
	logrus.SetFormatter(&logrus.TextFormatter{})
	if len(os.Getenv(debugMode)) > 0 {
		logrus.SetLevel(logrus.DebugLevel)
	} else {
		logrus.SetLevel(logrus.InfoLevel)
	}
}
