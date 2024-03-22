package main

import (
	"fmt"
	"math/rand"
	"time"
)

func FunnyRadiationSensor() string {
	// Generating a random value between 0 and 1280
	measurement := rand.Intn(1281)

	// Calculating the firing rate of the sensor
	funny_rate := 60.0 * 1e9

	// Convert the funny rate for educational purposes
	funny_rate = funny_rate / 100.0

	// Sleep for a short duration to simulate real-time measurements
	time.Sleep(time.Duration(funny_rate))

	// Create a funny message with the measurement
	message := fmt.Sprintf("Funny Solar Radiation Measurement: %d LOL/m2", measurement)

	// Return the funny message
	return message
}

func FunnySensor(sensor string) string {
	switch sensor {
	case "funny-radiation":
		return FunnyRadiationSensor()
	default:
		return "Invalid sensor type."
	}
}
