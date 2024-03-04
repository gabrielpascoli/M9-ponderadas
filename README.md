# M9-ponderadas
ponderadas do moduo 09

p1 : # MQTT Simulator Readme

Este é um simulador MQTT desenvolvido em Go (Golang) para simular a leitura de dados de um sensor fictício e a comunicação MQTT entre um cliente e um broker. O simulador é dividido em três partes principais:

1. **Publicador MQTT (main.go):**
   - O código simula a leitura de um sensor fictício que mede partículas PM2.5 e PM10.
   - Utiliza a biblioteca Paho MQTT para se conectar a um broker MQTT local (localhost:1883) e publica leituras do sensor no tópico "test/topic".
   - O programa continua a publicação até receber um sinal de interrupção (Ctrl+C).

2. **Simulador de Leitura do Sensor (sensor_simulator.go):**
   - Este programa simula leituras do sensor SPS30, gerando valores aleatórios para PM2.5 e PM10.
   - As leituras são impressas no console, e a simulação ocorre indefinidamente até receber um sinal de interrupção (Ctrl+C).

3. **Assinante MQTT (mqtt_subscriber.go):**
   - O código cria um cliente MQTT que se conecta a um broker MQTT local e assina o tópico "test/topic".
   - Implementa uma função de callback que é chamada sempre que uma mensagem é recebida no tópico assinado.
   - O programa continua aguardando mensagens até receber um sinal de interrupção (Ctrl+C).

## Como Usar:

1. Certifique-se de ter o Go instalado em seu ambiente de desenvolvimento.
2. Instale a biblioteca Paho MQTT com o seguinte comando:
   ```bash
   go get github.com/eclipse/paho.mqtt.golang
   ```
3. Execute os programas individualmente a partir do terminal. Por exemplo, para executar o Publicador MQTT:
   ```bash
   go run main.go
   ```
   Ou para executar o Simulador de Leitura do Sensor:
   ```bash
   go run sensor_simulator.go
   ```
   E para o Assinante MQTT:
   ```bash
   go run mqtt_subscriber.go
   ```

## Observações:
- Certifique-se de ter um broker MQTT em execução localmente no endereço e porta padrão (tcp://localhost:1883) para que os programas se conectem corretamente.
- Os programas são projetados para rodar indefinidamente até receber um sinal de interrupção (Ctrl+C).
- O código utiliza a biblioteca `rand` para gerar valores aleatórios, simbolizando leituras fictícias do sensor.

Esse simulador é uma ferramenta simples para entender e testar a comunicação MQTT em ambientes locais. Sinta-se à vontade para modificar e expandir conforme necessário para atender aos seus requisitos específicos.