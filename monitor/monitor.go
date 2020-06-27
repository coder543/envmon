package monitor

import (
	"context"
	"envmon/sensor"
	"github.com/influxdata/influxdb-client-go"
	"log"
	"time"
)

func Start() {
	client := influxdb2.NewClient("http://localhost:8081", "admin:admin")
	writeApi := client.WriteApiBlocking("", "db0")
	t := time.NewTicker(5 * time.Second)
	for {
		r := sensor.Read()
		log.Printf("%s", r)

		p := influxdb2.NewPoint(
			"environment",
			map[string]string{
				"location": "inside",
			},
			map[string]interface{}{
				"temperature": r.Temperature,
				"humidity":    r.Humidity,
				"pressure":    r.Pressure,
				"altitude":    r.Altitude,
			},
			time.Now(),
		)

		err := writeApi.WritePoint(context.Background(), p)
		if err != nil {
			log.Println(err)
		}
		<-t.C
	}
}
