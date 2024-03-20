package main

import (
	"testing"
	"time"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

func TestSubscriber(t *testing.T) {
	messageReceived := make(chan struct{})

	messagePubHandler := func(client MQTT.Client, msg MQTT.Message) {
		t.Logf("Received: %s\n", msg.Payload())
		close(messageReceived)
	}

	opts := MQTT.NewClientOptions().AddBroker("tcp://localhost:1891")
	opts.SetClientID("go_test_subscriber")
	opts.SetDefaultPublishHandler(messagePubHandler)

	client := MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		t.Error(token.Error())
	}

	if token := client.Subscribe("test/topic", 1, nil); token.Wait() && token.Error() != nil {
		t.Error(token.Error())
		return
	}

	t.Log("Subscriber running.")

	select {
	case <-messageReceived:
		t.Log("Received a message. Test successful!")
	case <-time.After(30 * time.Second):
		t.Error("Timeout reached. No message received.")
	}
}
