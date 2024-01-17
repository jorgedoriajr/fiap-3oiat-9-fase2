# Hamburgueria

Projeto FIAP 3SOAT - Grupo 9

<img src="/docs/fiap-3soat-9.drawio.png" alt="desenho arquitetura">

## Responsability

Responsável por gerenciar um sistema de totem para pedidos

## Requirements

- Golang 1.21+
- Docker

## Folders Structure

This repository contains five main folders: `cmd`, `build`, `internal` , `pkg` and `config`.

The `internal` folder contains all the go code, modules and tests that compose the
service.

The `build` folder contains the Dockerfile used for building the container,
the docker-compose file with the description of the local development environment.

The `cmd` folder stores the `main.go` file and the application **handlers**

The `pkg` folder stores the external implementations that could be used by other projects(ex: Kafka Client)

The `config` folder contains the configuration, loading environment variables, loading feature flags, etc.

## Open Api Doc

    http://localhost:8080/swagger/index.html

    Documentacao gerada pelo swaggo. (http://github.com/swaggo/echo-swagger)
    Acessar doc: ${path}/swagger/index.html

    Após documentar novos endpoints é necessario gerar novamente os arquivos através do comando:
    swag init -g cmd/main.go

## Executing

Install dependencies. 

`go get -u ./...`

Run unit test

`go test -v -race ./...`

Run application - *this need the database up*

`CONF=./config/config.yml go run -race cmd/main.go`

### Start application using docker

`docker-compose -f ./deployments/compose/docker-compose.yml up -d`

If have problems with permission with the database startup, maybe you will need to give permission to this file and them starter the docker compose again.

`chmod +x ./deployments/compose/database/setup-postgresql-databases.sh`

### Start application with minukube

Primeiramente verifique se você já tem o minikube instalado.

Passo a passo para instalação conforme sistema operacional: https://minikube.sigs.k8s.io/docs/start/

Também é necessário ter o kubectl instalado: https://kubernetes.io/docs/tasks/tools/

Após a instalação, rode o seguinte comando para habilitar métricas:

`kubectl apply -f https://github.com/kubernetes-sigs/metrics-server/releases/download/v0.4.1/components.yaml`

`minikube addons enable metrics-server`

Agora já é possível iniciar o cluster

`minikube start`

Para subir nosso deployment `hamburgueria-app` primeiramente você precisa ter o banco de dados up.

`kubectl apply -f deployments/k8s/database.yml`

E rodar os scripts para criar as tabelas do banco

`kubectl apply -f deployments/k8s/job-flyway.yml`

Agora é só rodar o seguinte comando para subir a aplicação:

`kubectl apply -f deployments/k8s/application.yml`

Para acessar através da porta 8080, é necessário dar um port forward

`kubectl port-forward service/hamburgueria-app-service 8080:80`

### Access
Application will be available at http://localhost:8080

To access database directly:

    connectionUrl: jdbc:postgresql://localhost:5432/hamburgueria
    user: hamburgueria-app-read
    password: hamburgueria-app-read
