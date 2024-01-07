package led

import (
	"github.com/sirupsen/logrus"
	"periph.io/x/conn/v3/gpio"
	"periph.io/x/conn/v3/gpio/gpioreg"
	"periph.io/x/host/v3"
	"tic-linky/internal/config"
)

var (
	logger = logrus.WithField("logger", "gpio")
	// ledGpio the GPIO used to manage LED
	ledGpio gpio.PinIO
)

// InitLedGpio init the LED GPIO on startup
func InitLedGpio() {
	if _, err := host.Init(); err != nil {
		logger.WithError(err).Warnf("Error initializing GPIO")
		return
	}

	ledGpio = gpioreg.ByName("GPIO" + config.LedGPIO)
	if ledGpio == nil {
		logger.Warnf("Led GPIO %s not found", config.LedGPIO)
	}
}

// TurnLEDOnOff turning on or off the LED
func TurnLEDOnOff(level gpio.Level) {
	if ledGpio != nil {
		if err := ledGpio.Out(level); err != nil {
			logger.WithError(err).Warnf("Error turning '%s' led GPIO", level)
			return
		}

		logger.Infof("Turning LED : '%s'", level)
	}
}
