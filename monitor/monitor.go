package monitor

import (
	"context"
	"envmon/sensor"
	"log"
	"os"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api/write"
)

func Start() {
	client := influxdb2.NewClient("http://192.168.88.21:8086", os.Getenv("INFLUXDB_TOKEN"))
	writeAPI := client.WriteAPIBlocking("Home", "Home")
	t := time.NewTicker(15 * time.Second)
	for {
		r := sensor.Read()
		log.Printf("%s", r)

		p := write.NewPoint(
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

		err := writeAPI.WritePoint(context.Background(), p)
		if err != nil {
			log.Println(err)
		}
		<-t.C
	}
}
