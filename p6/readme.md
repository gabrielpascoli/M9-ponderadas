
---

# iotsimMeta

iotsimMeta é uma atualização do iotsimHive, adicionando armazenamento em um banco de dados SQLite e visualização em um painel através do Metabase.

## Visualizando e armazenando com iotsimMeta!

Para armazenar os dados do simulador de sensor de radiação solar e visualizá-los em um painel, siga o tutorial do diretório `pond6`.

### Configurando as variáveis de ambiente

Primeiramente, para acessar o binário do Go, execute o seguinte comando:
```
source .bashrc
```

Em seguida, crie um arquivo `.env` com suas credenciais:
```
BROKER_ADDR="379d67d20bd940f2921461046040735b.s1.eu.hivemq.cloud"
HIVE_USER=""
HIVE_PSWD=""
```

### API para gerenciamento de dados

Para armazenar os dados no banco de dados, é necessário ativar uma API, então execute o seguinte comando no diretório `/api`:
```
go run .
```

### Publicando e Subscrevendo no HiveMQ

Para enviar e receber dados para o cluster do HiveMQ, execute o seguinte comando nos diretórios `/publisher` e `/subscriber`:
```
go run .
```

### Visualizando o Painel

Para a visualização dos dados, execute o seguinte comando:
```
sudo docker run -d -p 3000:3000 \
-v ~/Modulo9/src/pond6/metabase.db:/metabase.db \
-v ~/Modulo9/src/pond6/database:/database \
--name metabase metabase/metabase
```

No Metabase, insira suas informações e conecte-se ao banco de dados com o seguinte comando:
```
database/data.db
```

## iotsimMeta em Ação!

Você pode ver o iotsimMeta em ação no seguinte vídeo:

[Assistir ao Vídeo](https://drive.google.com/file/d/1j1XDhm1z2UYCpQuphqfuQU5veI6NnfUh/view?usp=sharing)

---
