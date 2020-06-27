package main

import (
	"envmon/monitor"
	"envmon/sensor"
	"github.com/d2r2/go-logger"
)

func main() {
	logger.ChangePackageLogLevel("i2c", logger.InfoLevel)
	logger.ChangePackageLogLevel("bsbmp", logger.InfoLevel)
	sensor.Init(0x77, 1, 90)
	monitor.Start()
}
