package main

import (
	"testing"
	"time"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

func TestEsquisito(t *testing.T) {
	mensagemRecebida := make(chan struct{})

	controladorMensagem := func(cliente MQTT.Client, msg MQTT.Message) {
		t.Logf("Recebido: %s\n", msg.Payload())
		close(mensagemRecebida)
	}

	opcoes := MQTT.NewClientOptions().AddBroker("tcp://localhost:1891")
	opcoes.SetClientID("go_teste_assinante")
	opcoes.SetDefaultPublishHandler(controladorMensagem)

	cliente := MQTT.NewClient(opcoes)
	if token := cliente.Connect(); token.Wait() && token.Error() != nil {
		t.Error(token.Error())
	}

	if token := cliente.Subscribe("teste/tópico", 1, nil); token.Wait() && token.Error() != nil {
		t.Error(token.Error())
		return
	}

	t.Log("Assinante em ação.")

	select {
	case <-mensagemRecebida:
		t.Log("Mensagem recebida. Teste bem sucedido!")
	case <-time.After(30 * time.Second):
		t.Error("Tempo esgotado. Nenhuma mensagem recebida.")
	}
}
