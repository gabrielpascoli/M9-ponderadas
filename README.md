Claro! Abaixo está um exemplo de README.md para os códigos que você forneceu:

```markdown
# Projeto MQTT com Go

Este é um projeto de exemplo que demonstra como usar o protocolo MQTT (Message Queuing Telemetry Transport) com a linguagem de programação Go. O projeto consiste em um publicador (publisher) e um inscrito (subscriber) que se comunicam através de um servidor MQTT.

## Requisitos

- Go 1.16 ou superior
- Cliente MQTT: Eclipse Paho MQTT Client para Go
- Biblioteca Godotenv para carregar variáveis de ambiente de um arquivo .env

## Configuração

1. Certifique-se de ter um servidor MQTT disponível para comunicação.
2. Crie um arquivo `.env` na raiz do projeto com as seguintes variáveis:

   ```
   BROKER_ADDR=seu_endereco_broker
   HIVE_USER=seu_usuario
   HIVE_PSWD=sua_senha
   ```

   Substitua `seu_endereco_broker`, `seu_usuario` e `sua_senha` com as informações do seu servidor MQTT.

## Instalação

1. Clone este repositório:

   ```bash
   git clone https://github.com/seu-usuario/seu-repositorio.git
   ```

2. Navegue até o diretório do projeto:

   ```bash
   cd seu-repositorio
   ```

3. Instale as dependências do Go:

   ```bash
   go mod tidy
   ```

## Uso

### Publicador (Publisher)

O publicador envia mensagens para um tópico específico no servidor MQTT.

```bash
go run publisher.go
```

### Inscrito (Subscriber)

O inscrito se inscreve em um ou mais tópicos no servidor MQTT para receber mensagens.

```bash
go run subscriber.go
```

## Contribuição

Contribuições são bem-vindas! Para mudanças importantes, abra primeiro uma issue para discutir o que você gostaria de mudar.

## Licença

[MIT](https://choosealicense.com/licenses/mit/)
```

Neste README.md, forneci instruções sobre como configurar e usar o projeto, bem como informações sobre como contribuir e a licença do projeto. Certifique-se de substituir `seu_endereco_broker`, `seu_usuario` e `sua_senha` com as informações reais do seu servidor MQTT.