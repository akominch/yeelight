package main

import (
	"github.com/akominch/yeelight"
	"time"
)

func main() {
	config := yeelight.YeelightConfig{
		Ip: "192.168.100.24",
		Effect: yeelight.Smooth,
	}
	bulb := yeelight.New(config)

	_, _ = bulb.TurnOnWithParams(yeelight.Last, 200)

	time.Sleep(5 * time.Second)

	_, _ = bulb.TurnOff()
}
