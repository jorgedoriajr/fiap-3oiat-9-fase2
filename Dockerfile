FROM golang:1.21-alpine as base

WORKDIR /hamburgueria

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# BUILDER
FROM base as builder

COPY . /hamburgueria

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64  go build -ldflags "${LDFLAGS} -s -w" -a -o hamburgueria-app /hamburgueria/cmd/main.go

FROM scratch

WORKDIR /hamburgueria

COPY --from=builder /hamburgueria/config /hamburgueria/config
COPY --from=builder /hamburgueria/hamburgueria-app /hamburgueria


EXPOSE 8080 8081


CMD ["./hamburgueria-app"]
