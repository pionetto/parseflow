<h1 align="center">:file_cabinet: Parseflow</h1>

ParseFlow é um sistema para processamento e armazenamento eficiente de arquivos de dados em um banco de dados PostgreSQL. Ele lê arquivos .txt e .csv, trata os dados e os armazena de forma otimizada.

## :books: Requisitos

Antes de rodar o projeto, certifique-se de ter instalado:

* Go 1.23+

* Docker

* Docker Compose

* Git

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
curl -X POST -F "file=@base_teste.txt" http://localhost:8080/upload
```

Você pode executar este comando diretamente em um terminal Linux.


## Banco de Dados

O sistema conecta-se ao PostgreSQL e cria automaticamente a tabela clientes caso ela não exista. Para acessar o banco via CLI:

```
docker exec -it parseflow_db psql -U postgres -d parserflow_db
```

Se precisar reiniciar completamente os dados do banco de dados:


```
docker-compose down -v
```

Isso removerá todos os volumes e reiniciará o banco do zero.

## Comandos Disponíveis

O projeto inclui um `Makefile` para facilitar a execução de tarefas comuns:

| Comando | Descrição |
|:--------|:----------|
| `make up` | Sobe o ambiente (aplicação + banco de dados) |
| `make down` | Derruba todos os containers Docker |
| `make db-up` | Sobe apenas o banco de dados PostgreSQL |
| `make build` | Build da imagem Docker da aplicação |
| `make build-test` | Build da imagem Docker para rodar testes |
| `make test` | Roda os testes unitários dentro do container |
| `make clean` | Limpa containers, volumes e arquivos de cobertura (`cover.out`, `cover.html`) |
| `make rebuild-test` | Limpa, faz o build e roda testes automaticamente |

---

### Exemplos de Uso:

- Subir o ambiente local:

```bash
make up
