package main

import (
	"fmt"
	"math/rand"
	"time"
)

// simularLeituraSensor simula uma leitura do sensor SPS30, retornando dois valores, PM2.5 (µg/m³) e PM10 (µg/m³).
func simularLeituraSensor() (float64, float64) {
	var1 := rand.Float64() * 10
	var2 := rand.Float64()*10 + 10
	return var1, var2
}

func main() {
	// Canal para sinalizar a conclusão da simulação
	done := make(chan struct{})

	// Captura de sinal para lidar com o encerramento
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt)

	// Simular leituras do sensor continuamente
	go func() {
		for {
			select {
			case <-done:
				return
			default:
				var1, var2 := simularLeituraSensor()
				fmt.Printf("Leitura do sensor: %.2f, %.2f\n", var1, var2)
				time.Sleep(time.Duration(rand.Intn(2000)+1000) * time.Millisecond) // Espera aleatória entre 1 e 3 segundos
			}
		}
	}()

	// Aguardar um sinal de interrupção para terminar a simulação
	<-signalChannel
	close(done)
	fmt.Println("Simulação encerrada")
}
