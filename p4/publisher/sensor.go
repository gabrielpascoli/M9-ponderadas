package main

import (
	"fmt"
	"math/rand"
	"time"
)

func SolarRadiationSensor() string {
	// Gerar um valor aleatório entre 0 e 1280
	medicao := rand.Intn(1281)

	// Calcular a taxa de disparo do sensor
	taxa_disparo := 60.0 * 1e9

	// Converta a taxa de disparo para fins educacionais
	taxa_disparo = taxa_disparo / 100.0

	// Dormir por uma curta duração para simular medidas em tempo real
	time.Sleep(time.Duration(taxa_disparo))

	// Criar uma mensagem de texto com a medição
	texto := fmt.Sprintf("Medição de Radiação Solar: %d W/m2", medicao)

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
