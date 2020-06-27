package sensor

import (
	"envmon/units"
	"fmt"
	"github.com/d2r2/go-bsbmp"
	"github.com/d2r2/go-i2c"
	"log"
	"strings"
	"time"
)

var sensor *bsbmp.BMP
var altOffset units.M

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
	Temperature units.C
	Pressure    units.Pa
	Humidity    units.Percent
	Altitude    units.Ft
	Time        time.Duration
}

func (r Reading) String() string {
	return strings.TrimSpace(
		fmt.Sprintf(`
Reading {
	Temperature:	%s,
	Pressure:	%s,
	Humidity:	%s,
	Altitude:	%s,
	MeasureT:	%s,
}
		`,
			r.Temperature,
			r.Pressure,
			r.Humidity,
			r.Altitude,
			r.Time,
		),
	)
}

var altitudeHistory = []units.M{}

func Read() Reading {
	start := time.Now()

	t, err := sensor.ReadTemperatureC(bsbmp.ACCURACY_ULTRA_HIGH)
	if err != nil {
		log.Panic(err)
	}

	p, err := sensor.ReadPressurePa(bsbmp.ACCURACY_ULTRA_HIGH)
	if err != nil {
		log.Panic(err)
	}

	supported, rh, err := sensor.ReadHumidityRH(bsbmp.ACCURACY_ULTRA_HIGH)
	if !supported {
		log.Panic("BME280 should support humidity")
	}
	if err != nil {
		log.Panic(err)
	}

	aRaw, err := sensor.ReadAltitude(bsbmp.ACCURACY_ULTRA_HIGH)
	if err != nil {
		log.Panic(err)
	}
	altitudeHistory = append(altitudeHistory, units.M(aRaw))

	aLarge := units.M(0)

	for _, alt := range altitudeHistory {
		aLarge += alt
	}

	a := units.M(float32(aLarge) / float32(len(altitudeHistory)))

	if len(altitudeHistory) > 2 {
		// moving average filter with IIR, weighted as part of a regular FIR moving average
		altitudeHistory[1] = (altitudeHistory[1] + altitudeHistory[0]) / 2
		altitudeHistory = altitudeHistory[1:]
	}

	if altOffset == 0 {
		altOffset = a
		log.Printf("altitude offset is now %s", altOffset.ToFt())
	}

	a = a - altOffset

	return Reading{
		Temperature: units.C(t),
		Pressure:    units.Pa(p),
		Humidity:    units.Percent(rh),
		Altitude:    a.ToFt(),
		Time:        time.Since(start),
	}
}
