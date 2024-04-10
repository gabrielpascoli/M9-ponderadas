package main

import (
	"math/rand"
	"time"
	"fmt"
)

func ApparentlySolarRadiation() string {
	measurement := rand.Intn(1281)
	firing_rate := 60.0 * 1e9
	firing_rate = firing_rate/100.0
	time.Sleep(time.Duration(firing_rate))
	text := fmt.Sprintf(`{"sensor": "thefalloff", "value": "%d"}`, measurement)
	return text
}

func Apparently(sensor string) string {
	switch sensor {
	case "solar-radiation":
		return ApparentlySolarRadiation()
	default:
		return "Invalid sensor type."
	}
}
