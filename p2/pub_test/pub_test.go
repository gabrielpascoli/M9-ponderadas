package main

import (
	"fmt"
	"math"
	"math/rand"
	"testing"
	"time"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

func testeAbobrinha(t *testing.T) string {
	tempoInicio := time.Now()

	medida := rand.Intn(1281)
	taxaExtravagante := 60.0 * 1e9
	taxaExtravagante = taxaExtravagante / 100.0

	time.Sleep(time.Duration(taxaExtravagante))
	texto := fmt.Sprintf("Medição de Radiação Solar: %d W/m2", medida)

	tempoDecorrido := time.Since(tempoInicio).Seconds()
	tempoEsperado := time.Duration(taxaExtravagante).Seconds()

	if math.Abs(tempoDecorrido-tempoEsperado) > tempoEsperado {
		t.Errorf("O Sensor de Radiação Solar levou %f segundos para executar, esperado %f segundos", tempoDecorrido, tempoEsperado)
	}

	return texto
}

func testePublicador(t *testing.T, texto string) {
	opcoes := MQTT.NewClientOptions().AddBroker("tcp://localhost:1891")
	opcoes.SetClientID("go_teste_publicador")

	cliente := MQTT.NewClient(opcoes)
	if token := cliente.Connect(); token.Wait() && token.Error() != nil {
		t.Error(token.Error())
	}

	token := cliente.Publish("teste/tópico", 0, false, texto)
	token.Wait()

	t.Log("Publicado: ", texto)
}

func TestEngraçadinho(t *testing.T) {
	texto := testeAbobrinha(t)

	testePublicador(t, texto)
}
