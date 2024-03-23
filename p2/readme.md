

---

# Sistema de Sensores Engraçados

Este é um sistema de simulação de sensores engraçados que gera dados de medição fictícios para um sensor de radiação solar. Os códigos fornecidos incluem um publicador (publisher) e um assinante (subscriber) para simular a geração e recebimento desses dados.

## Requisitos

- Go 1.13 ou superior
- Pacote MQTT (github.com/eclipse/paho.mqtt.golang)

## Como Usar

1. Certifique-se de ter Go instalado em seu sistema.
2. Instale o pacote MQTT executando o seguinte comando:
   ```
   go get github.com/eclipse/paho.mqtt.golang
   ```
3. Clone ou baixe este repositório para o seu sistema.
4. Navegue até o diretório onde os códigos estão localizados.
5. Execute o programa do publicador e do assinante em terminais separados.

### Publicador (Publisher)

O publicador é responsável por gerar dados de medição engraçados e enviá-los para um tópico MQTT.

```bash
go run publisher.go
```

### Assinante (Subscriber)

O assinante se conecta ao tópico MQTT e recebe os dados de medição engraçados enviados pelo publicador.

```bash
go run subscriber.go
```


## Link do vídeo 

https://youtu.be/fjSUbT1inVg?si=fLzLzp9Y-myTqxJf
