package main

import (
	"fmt"
	"os"
	"regexp"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	godotenv "github.com/joho/godotenv"
)

func loadEnv() {
	p6
	p6
	projectDirName := "p6"
	re := regexp.MustCompile(`^(.*` + projectDirName + `)`)
	cwd, _ := os.Getwd()
	rootPath := re.Find([]byte(cwd))

	err := godotenv.Load(string(rootPath) + `/.env`)

	if err != nil {
		fmt.Printf("Error loading .env file: %s", err)
	}
}

func main() {
	loadEnv()

	broker := os.Getenv("BROKER_ADDR")
	port := 8883

	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tls://%s:%d", broker, port))
	opts.SetClientID("DoIWannaKnowPublisher")
	opts.SetUsername(os.Getenv("HIVE_USER"))
	opts.SetPassword(os.Getenv("HIVE_PSWD"))

	var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
		fmt.Println("Connected")
	}

	var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
		fmt.Printf("Connection lost: %v", err)
	}

	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	for {
		text := Sensor("solar-radiation")
		token := client.Publish("sensor/topic", 0, false, text)
		token.Wait()
		fmt.Println("Published: ", text)
		time.Sleep(5 * time.Second)
	}
}
