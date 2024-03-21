package main

import (
	"fmt"
	"os"
	"regexp"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	godotenv "github.com/joho/godotenv"
)

func loadEnv() {
	projectDirName := "pond4"
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
	opts.SetClientID("GoSubscriber")
	opts.SetUsername(os.Getenv("HIVE_USER"))
	opts.SetPassword(os.Getenv("HIVE_PSWD"))

	var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
		fmt.Printf("Recebido: %s\n", msg.Payload())
	}

	var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
		fmt.Println("Connected")
	}

	var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
		fmt.Printf("Connection lost: %v", err)
	}

	opts.SetDefaultPublishHandler(messagePubHandler)
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	// Defina as funções de manipulador de mensagem para os tópicos aos quais você deseja se inscrever
	topicos := map[string]byte{
		"sensor/topic": 1, // 1 para QoS 1
		// Adicione outros tópicos aqui, se necessário
	}

	if token := client.SubscribeMultiple(topicos, nil); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		return
	}

	fmt.Println("Subscribed running...")
	select {} // Bloqueia indefinidamente
	//client.Disconnect(250)
}
