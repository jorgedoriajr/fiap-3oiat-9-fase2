# Hamburgueria

Projeto FIAP 3SOAT - Grupo 9

## Responsability

Responsável por gerenciar um sistema de totem para pedidos

## Requirements

- Golang 1.20+
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