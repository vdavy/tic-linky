package config

import (
	"github.com/sirupsen/logrus"
	"os"
)

const (
	serialPortName        = "SERIAL_PORT"
	ledGPIOName           = "LED_GPIO"
	influxdbServerURLName = "INFLUXDB_URL"
	influxdbUsernameName  = "INFLUXDB_USERNAME"
	influxdbPasswordName  = "INFLUXDB_PASSWORD"
	influxdbDatabaseName  = "INFLUXDB_DATABASE"
)

var (
	// SerialPort serial port env var
	SerialPort string
	// LedGPIO GPIO number for LED
	LedGPIO string
	// InfluxdbServerURL influxdb server URL env var
	InfluxdbServerURL string
	// InfluxdbUsername influxdb username env var
	InfluxdbUsername string
	// InfluxdbPassword influxdb password var
	InfluxdbPassword string
	// InfluxdbDatabase influxdb database var
	InfluxdbDatabase string
	envVarsList      = []*string{&SerialPort, &LedGPIO,
		&InfluxdbServerURL, &InfluxdbUsername, &InfluxdbPassword, &InfluxdbDatabase}
)

// ParseEnvVars parse required env vars, panic if missing
func ParseEnvVars() {
	for idx, envVarName := range []string{serialPortName, ledGPIOName,
		influxdbServerURLName, influxdbUsernameName, influxdbPasswordName, influxdbDatabaseName} {
		envVarValue := os.Getenv(envVarName)
		if len(envVarValue) == 0 {
			logrus.Fatalf("Missing env var %s", envVarName)
		}
		*envVarsList[idx] = envVarValue
	}
}
