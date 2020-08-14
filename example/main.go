package main

import (
	"github.com/akominch/yeelight"
	"github.com/akominch/yeelight/transitions"
	"time"
)

func main() {
	config := yeelight.BulbConfig{
		Ip: "192.168.100.24",
		Effect: yeelight.Smooth,
	}
	bulb := yeelight.New(config)

	t := transitions.Police2()

	flow := yeelight.NewFlow(3, yeelight.Off, t)

	bulb.StartFlow(flow)

	//_, _ = bulb.TurnOnWithParams(yeelight.Last, 200)

	time.Sleep(5 * time.Second)

	_, _ = bulb.TurnOff()
}
