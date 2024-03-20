Claro, aqui está um exemplo de arquivo `README.md` para a solução apresentada:

```markdown
# IoT Simulator

Este é um projeto de simulação de dispositivos IoT utilizando a linguagem Go. Ele inclui testes unitários para garantir a funcionalidade correta dos componentes.

## Pré-requisitos

Certifique-se de ter o Go instalado em seu sistema. Você pode encontrar instruções de instalação em [golang.org](https://golang.org/doc/install).

## Instalação

Para instalar as dependências do projeto, você pode usar o comando `go get`:

```bash
go get -u github.com/eclipse/paho.mqtt.golang
```

Isso garantirá que o cliente MQTT necessário seja baixado e instalado corretamente.

## Executando os Testes

Você pode executar os testes do projeto utilizando o comando `go test`:

```bash
go test
```

Isso executará todos os testes presentes no projeto e fornecerá informações sobre sua execução.

## Arquivos do Projeto

- `main.go`: Contém a implementação principal do simulador de IoT.
- `main_test.go`: Contém os testes unitários para garantir o funcionamento correto do código.
- `go.mod`: O arquivo de manifesto de módulo Go, especificando as dependências do projeto.
- `go.sum`: O arquivo de somas de verificação das dependências do projeto.

## Contribuindo

Contribuições são bem-vindas! Sinta-se à vontade para abrir uma issue para relatar problemas ou propor novas funcionalidades. Se você quiser contribuir diretamente, por favor, abra um pull request.

## Licença

Este projeto está licenciado sob a Licença MIT - veja o arquivo [LICENSE](LICENSE) para detalhes.

```
