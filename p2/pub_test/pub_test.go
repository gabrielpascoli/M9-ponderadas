package main

import (
	"fmt"
	"math"
	"math/rand"
	"testing"
	"time"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

func testSensor(t *testing.T) string {
	startTime := time.Now()

	measurement := rand.Intn(1281)
	firingRate := 60.0 * 1e9
	firingRate = firingRate / 100.0

	time.Sleep(time.Duration(firingRate))
	text := fmt.Sprintf("Solar Radiation Measurement: %d W/m2", measurement)

	elapsedTime := time.Since(startTime).Seconds()
	expectedTime := firingRate.Seconds()

	if math.Abs(elapsedTime-expectedTime) > expectedTime {
		t.Errorf("Solar Radiation Sensor took %f seconds to execute, expected %f seconds", elapsedTime, expectedTime)
	}

	return text
}

func testPublisher(t *testing.T, text string) {
	opts := MQTT.NewClientOptions().AddBroker("tcp://localhost:1891")
	opts.SetClientID("go_test_publisher")

	client := MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		t.Error(token.Error())
	}

	token := client.Publish("test/topic", 0, false, text)
	token.Wait()

	t.Log("Published: ", text)
}

func TestIotsim(t *testing.T) {
	text := testSensor(t)

	testPublisher(t, text)
}
