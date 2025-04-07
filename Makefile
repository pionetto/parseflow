APP_NAME=parseflow
DOCKER_COMPOSE=docker-compose
DOCKERFILE=Dockerfile
DOCKERFILE_TEST=Dockerfile.test

# Sobe a aplicação
up:
	$(DOCKER_COMPOSE) up app

# Derruba os containers
down:
	$(DOCKER_COMPOSE) down

# Sobe apenas o banco de dados
db-up:
	$(DOCKER_COMPOSE) up db

# Roda os testes
test:
	$(DOCKER_COMPOSE) up test

# Limpa containers e arquivos
clean:
	$(DOCKER_COMPOSE) down -v --remove-orphans
	rm -f cover.out cover.html || true

# Build da aplicação
build:
	docker build -t $(APP_NAME) -f $(DOCKERFILE) .

# Build da imagem de testes
build-test:
	docker build -t $(APP_NAME)-test -f $(DOCKERFILE_TEST) .

# Build da imagem de teste e roda os teste
rebuild-test: clean
	$(DOCKER_COMPOSE) build test
	$(DOCKER_COMPOSE) up --abort-on-container-exit --exit-code-from test test

.PHONY: up down db-up test clean build build-test rebuild-test