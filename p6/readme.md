otsimMeta
IotsimMeta é uma atualização de IotsimHive, adicionando armazenamento em um banco de dados SQlite e visualização em uma dashboard pelo Metabase.

Visualizando e armazenando com IotsimMeta!
Para armazenar os dados do simulador de sensor de radiação solar e visualizá-los em uma dashboard basta seguir o tutorial a partir do diretório pond6.

Configuração das variaveis de ambiente
Primeiramente, para acessar o binário do Go, rode o seguinte comando:

source .bashrc# IotsimMeta
IotsimMeta é uma atualização de
[IotsimHive](https://github.com/IgorSFG/Modulo9/tree/main/src/pond4),
adicionando armazenamento em um banco de dados SQlite e visualização em uma dashboard pelo Metabase.

## Visualizando e armazenando com IotsimMeta!
Para armazenar os dados do simulador de sensor de radiação solar e visualizá-los em uma dashboard basta seguir o tutorial a partir do diretório `pond6`.

### Configuração das variaveis de ambiente
Primeiramente, para acessar o binário do Go, rode o seguinte comando:
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


### Visualizando a Dashboard
Para a visualização dos dados, rode o seguinte comando:
sudo docker run -d -p 3000:3000
-v ~/Modulo9/src/pond6/metabase.db:/metabase.db
-v ~/Modulo9/src/pond6/database:/database
--name metabase metabase/metabase


No metabase, insira suas informção e conecte-se ao banco de dados com o seguinte comando:
database/data.db


## IotsimMeta em Ação!
Você pode conferir o funcionamento de IotsimMeta no vídeo a seguir:

https://drive.google.com/file/d/1j1XDhm1z2UYCpQuphqfuQU5veI6NnfUh/view?usp=sharing
Em seguida, crie um arquivo .env com as suas credenciais:

BROKER_ADDR = "379d67d20bd940f2921461046040735b.s1.eu.hivemq.cloud"
HIVE_USER = "<seu-nome-de-usuario-aqui>"
HIVE_PSWD = "<sua-senha-cadastrada-aqui>"
API para gerenciamento de dados
Para que os dados sejam armazenados no banco de dados, é necessário ativar um api, então rode o seguinte comando no diretório /api:

go run .
```

8. No terminal restante, execute o seguinte comando para iniciar o subscriber:

```bash
cd subscriber
go run .
Visualizando a Dashboard
Para a visualização dos dados, rode o seguinte comando:

sudo docker run -d -p 3000:3000 \
-v ~<caminho-absoluto>/Modulo9/src/pond6/metabase.db:/metabase.db \
-v ~<caminho-absoluto>/Modulo9/src/pond6/database:/database \
--name metabase metabase/metabase
No metabase, insira suas informção e conecte-se ao banco de dados com o seguinte comando:

database/data.db
IotsimMeta em Ação!
Você pode conferir o funcionamento de IotsimMeta no vídeo a seguir:
