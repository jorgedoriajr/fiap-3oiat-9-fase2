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
	docker compose -f ./deployments/compose/docker-compose.yml up postgres flyway  -d

run-on-k8s: compose-db minikube-start apply-k8s-config port-forward

apply-k8s-config:
	kubectl apply -f ./deployments/k8s/configmap.yaml ./deployments/k8s/deployment.yaml ./deployments/k8s/service.yaml ./deployments/k8s/hpa.yaml

port-forward:
	kubectl port-forward service/hamburgueria-app-service 8080:80

minikube-start:
	minikube start

validate: vet tests

tests:
	go test -v -race ./...

vet:
	go vet -v ./...

fmt:
	go fmt ./...