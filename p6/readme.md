# Ponderada 6 - Metabase

A ponderada opera da seguinte maneira: o HiveMQ atua como broker, facilitando a conexão entre o subscriber e o publisher. O publisher envia as informações e o subscriber as recebe. A API captura os dados do subscriber e os encaminha para o banco de dados. O Metabase está conectado diretamente ao banco, permitindo atualizações em tempo real.

## Execução do Projeto

Para executar o projeto, siga estes passos:

1. Clone o projeto:

```bash
git clone https://github.com/gabrielpascoli/M9-ponderadas.git
```

2. Abra 4 terminais dentro do projeto. Em cada terminal, execute o seguinte comando para acessar o ambiente:

```bash
cd ambiente_clonado/M9-ponderadas/p6
```

3. Execute o seguinte comando em 3 terminais para acessar o binário do Go:

```bash
source .bashrc
```

4. Para visualizar o banco de dados em funcionamento, execute o seguinte comando em um terminal que não executou o comando anterior, para acessar o Metabase através de um container Docker. Siga os passos fornecidos pelo Metabase.

```bash
cd database
sudo docker run -d -p 3000:3000 -v ~/caminho-para-o-repositorio/M9-ponderadas/p6/database:/database --name metabase metabase/metabase
```

5. Ao conectar o banco de dados no Metabase, utilize o seguinte caminho:

```plaintext
database/database.db
```

6. Em um dos terminais, execute os seguintes comandos para iniciar a API:

```bash
cd api
go run .
```

7. Em outro terminal, execute os seguintes comandos para iniciar o publisher:

```bash
cd publisher
go run .
```

8. No terminal restante, execute o seguinte comando para iniciar o subscriber:

```bash
cd subscriber
go run .
```

## Funcionamento Comprovado

Veja a ponderada em funcionamento corretamente através [deste link](POR O LINK AQUI).
