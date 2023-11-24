APP=hamburgueria

build-run: fmt vet lint tests compile
	./${APP}

compile:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "${LDFLAGS} -s -w" -a -o ${APP} cmd/main.go

compose:
	docker compose -f ./deployments/compose/docker-compose.yml up -d

compose-app:
	docker compose -f ./deployments/compose/docker-compose.yml app up -d

compose-down:
	docker compose -f ./deployments/compose/docker-compose.yml down

compose-db:
	docker compose -f ./deployments/compose/docker-compose.yml postgres flyway up -d

validate: vet tests

tests:
	go test -v -race ./...

vet:
	go vet -v ./...

fmt:
	go fmt ./...