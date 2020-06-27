package sensor

import (
	"fmt"
	"github.com/d2r2/go-bsbmp"
	"github.com/d2r2/go-i2c"
	"log"
	"strings"
)

var sensor *bsbmp.BMP

func Init(addr uint8, bus int) {
	// Create new connection to i2c-bus on 1 line with address 0x76.
	// Use i2cdetect utility to find device address over the i2c-bus
	i2cConn, err := i2c.NewI2C(addr, bus)
	if err != nil {
		log.Fatal(err)
	}

	sensor, err = bsbmp.NewBMP(bsbmp.BME280, i2cConn) // signature=0x60
	if err != nil {
		log.Fatal(err)
	}

	err = sensor.IsValidCoefficients()
	if err != nil {
		log.Fatal(err)
	}
}

type Reading struct {
	Temperature float32
	Pressure    float32
	Humidity    float32
	Altitude    float32
}

func (r Reading) String() string {
	return strings.TrimSpace(
		fmt.Sprintf(`
Reading {
	Temperature:	%.2f °C,
	Pressure:	%.2f Pa,
	Humidity:	%.2f%%,
	Altitude:	%.2f m,
}
		`,
			r.Temperature,
			r.Pressure,
			r.Humidity,
			r.Altitude,
		),
	)
}

func Read() Reading {
	// Read temperature in celsius degree
	t, err := sensor.ReadTemperatureC(bsbmp.ACCURACY_ULTRA_HIGH)
	if err != nil {
		log.Panic(err)
	}

	// Read atmospheric pressure in pascal
	p, err := sensor.ReadPressurePa(bsbmp.ACCURACY_ULTRA_HIGH)
	if err != nil {
		log.Panic(err)
	}

	// Read atmospheric pressure in mmHg
	supported, rh, err := sensor.ReadHumidityRH(bsbmp.ACCURACY_ULTRA_HIGH)
	if !supported {
		log.Panic("BME280 should support humidity")
	}
	if err != nil {
		log.Panic(err)
	}

	// Read atmospheric altitude in meters above sea level, if we assume
	// that pressure at see level is equal to 101325 Pa.
	a, err := sensor.ReadAltitude(bsbmp.ACCURACY_ULTRA_HIGH)
	if err != nil {
		log.Panic(err)
	}

	return Reading{
		Temperature: t,
		Pressure:    p,
		Humidity:    rh,
		Altitude:    a,
	}
}
