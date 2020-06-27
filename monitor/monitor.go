package monitor

import (
	"envmon/sensor"
	"log"
	"time"
)

func Start() {
	t := time.NewTicker(5 * time.Second)
	for {
		r := sensor.Read()
		log.Printf("%s", r)
		<-t.C
	}
}
