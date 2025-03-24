<h1 align="center">:file_cabinet: Parseflow</h1>

ParseFlow é um sistema para processamento e armazenamento eficiente de arquivos de dados em um banco de dados PostgreSQL. Ele lê arquivos .txt e .csv, trata os dados e os armazena de forma otimizada.

## :books: Requisitos

Antes de rodar o projeto, certifique-se de ter instalado:

* Docker

* Docker Compose

* curl para testes de upload de arquivos

## :rocket: Rodando o projeto

Clone este repositório:

```
git clone https://github.com/seu-usuario/parseflow.git
```

```
cd parseflow
```

Configure as variáveis de ambiente no arquivo .env:

```
cp .env.example .env
```

Edite o arquivo .env e ajuste as credenciais do banco de dados, se necessário.

Construa e suba os containers com Docker Compose:

```
docker-compose up --build
```

Isso iniciará os serviços necessários, incluindo o PostgreSQL e a aplicação Go.

Aguarde o PostgreSQL estar pronto. Caso haja problemas, reinicie os containers:

```
docker-compose down
```

```
docker-compose up
```

## :books: Como Fazer Upload de Arquivos

Para enviar um arquivo para processamento, utilize o comando curl:

```
curl -X POST -F "file=@caminho/do/arquivo.csv" http://localhost:8080/upload
```

Altere caminho/do/arquivo.csv pelo caminho real do seu arquivo.
Na raiz do projeto este aqui já existe, chamado base_teste.txt


## Banco de Dados

O sistema conecta-se ao PostgreSQL e cria automaticamente a tabela clientes caso ela não exista. Para acessar o banco via CLI:

```
docker exec -it parseflow_postgres psql -U postgres -d parserflow_db
```

Logs e Debug

Para visualizar os logs da aplicação, use:

```
docker-compose logs -f go_parserflow
```

Se precisar reiniciar completamente os dados do banco de dados:


```
docker-compose down -v
```

Isso removerá todos os volumes e reiniciará o banco do zero.