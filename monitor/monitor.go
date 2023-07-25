package monitor

import (
	"context"
	"envmon/sensor"
	"log"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go"
)

func Start() {
	client := influxdb2.NewClient("http://192.168.88.21:8086", "admin:admin")
	writeApi := client.WriteApiBlocking("", "db0")
	t := time.NewTicker(15 * time.Second)
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
