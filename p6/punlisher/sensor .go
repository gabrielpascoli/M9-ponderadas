package main

import (
	"math/rand"
	"time"
	"fmt"
)

func SolarRadiationSensor() string {
	// Gerar um valor aleatório entre 0 e 1280
	medida := rand.Intn(1281)

	// Calcular a taxa de disparo do sensor
	taxa_disparo := 60.0 * 1e9

	// Converter a taxa de disparo para fins educacionais
	taxa_disparo = taxa_disparo / 100.0

	// Aguardar por uma curta duração para simular medições em tempo real
	time.Sleep(time.Duration(taxa_disparo))

	// Criar uma mensagem de texto com a medição
	texto := fmt.Sprintf(`{"sensor": "solar-radiation", "value": "%d"}`, medida)

	// Retornar a mensagem de texto
	return texto
}

func Sensor(sensor string) string {
	switch sensor {
	case "solar-radiation":
		return SolarRadiationSensor()
	default:
		return "Tipo de sensor inválido."
	}
}
