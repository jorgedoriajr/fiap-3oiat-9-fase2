# Hamburgueria

Projeto FIAP 3SOAT - Grupo 9

## Responsability

Responsável por gerenciar um sistema de totem para pedidos

## Requirements

- Golang 1.21+
- Docker

## Folders Structure

This repository contains five main folders: `cmd`, `build`, `internal` , `pkg` and `configs`.

The `internal` folder contains all the go code, modules and tests that compose the
service.

The `build` folder contains the Dockerfile used for building the container,
the docker-compose file with the description of the local development environment.

The `cmd` folder stores the `main.go` file and the application **handlers**

The `pkg` folder stores the external implementations that could be used by other projects(ex: Kafka Client)

The `configs` folder contains the configuration, loading environment variables, loading feature flags, etc.

## Open Api Doc

    http://ifoodcard-redeem-handler.aws.cluster-01.k8s.enterprise-sandbox.dc-ifood.com/swagger/index.html

    Documentacao gerada pelo swaggo. (http://github.com/swaggo/echo-swagger)
    Acessar doc: ${path}/swagger/index.html

    Após documentar novos endpoints é necessario gerar novamente os arquivos através do comando:
    swag init -g cmd/main.go

## Executing

### **Run Project**

Install dependencies. 

`go get -u ./...`

Starter database

`docker compose -f ./deployments/compose/docker-compose.yml up -d`

If have problems with permission with the database startup, maybe you will need to give permission to this file

`chmod +x ./deployments/compose/database/setup-postgresql-databases.sh`, them starter the docker compose again.

Run application

`CONF=./config/config.yml go run -race cmd/main.go`

Run unit test

`go test -v -race ./...`
