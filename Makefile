APP=hamburgueria

build-run: fmt vet lint tests compile
	./${APP}

compose-up:
	#chmod +x ./deployments/compose/database/setup-postgresql-databases.sh
	docker compose -f ./deployments/compose/docker-compose.yml up -d

compile:
# compile binary
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "${LDFLAGS} -s -w" -a -o ${APP} cmd/app/main.go

validate: vet tests

tests:
	go test -v -race ./...

vet:
	go vet -v ./...

fmt:
	go fmt ./...

run:
	CONF=./config/config.yml go run -race cmd/app/main.go

tools:
	go get github.com/golangci/golangci-lint/cmd/golangci-lint

lint:
	golangci-lint run

tidy:
	go mod tidy -compat=1.19