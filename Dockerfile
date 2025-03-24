# Usa uma imagem oficial do Golang para build
FROM golang:1.23 AS builder

WORKDIR /app

# Copia os arquivos do projeto
COPY . .

# Baixa as dependências do Go
RUN go mod tidy

# Compila a aplicação
RUN go build -o /app/parseflow .

# Usa uma imagem minimalista para rodar a aplicação
FROM debian:bookworm-slim

WORKDIR /app

# Copia o binário compilado da etapa anterior
COPY --from=builder /app/parseflow .

# Criar o diretório de uploads e garantir permissões
RUN mkdir -p /app/uploads && chmod -R 777 /app/uploads

# Garante permissão de execução
RUN chmod +x /app/parseflow

# Expor a porta da aplicação
EXPOSE 8080

# Executa a aplicação
CMD ["/app/parseflow"]

