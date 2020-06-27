package monitor

import (
	"context"
	"envmon/sensor"
	"github.com/influxdata/influxdb-client-go"
	"log"
	"time"
)

func Start() {
	client := influxdb2.NewClient("http://localhost:8086", "admin:admin")
	writeApi := client.WriteApiBlocking("", "db0")
	t := time.NewTicker(1 * time.Minute)
	for {
		r := sensor.Read()
		log.Printf("%s", r)

		p := influxdb2.NewPoint(
			"environment",
			map[string]string{
				"location": "inside",
			},
			map[string]interface{}{
				"temperature": float32(r.Temperature),
				"humidity":    float32(r.Humidity),
				"pressure":    float32(r.Pressure),
				"altitude":    float32(r.Altitude),
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
