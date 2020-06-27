package monitor

import (
	"encoding/json"
	"envmon/sensor"
	"log"
	"time"
)

func Start() {
	t := time.NewTicker(5 * time.Second)
	for {
		r := sensor.Read()
		log.Printf("%s", r)

		_, err := json.Marshal(&r)
		if err != nil {
			log.Panic(err)
		}
		//log.Printf("json: %s", string(jsonBytes))
		<-t.C
	}
}
