FROM golang:1.21-alpine as base

WORKDIR /hamburgueria

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# BUILDER
FROM base as builder

COPY . /hamburgueria

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "${LDFLAGS} -s -w" -a -o hamburgueria-app /hamburgueria

FROM alpine:latest

RUN apk update && apk add --no-cache ca-certificates

WORKDIR /hamburgueria

COPY --from=builder /hamburgueria/hamburgueria-app /hamburgueria

# Optional:
# To bind to a TCP port, runtime parameters must be supplied to the docker command.
# But we can document in the Dockerfile what ports
# the application is going to listen on by default.
# https://docs.docker.com/engine/reference/builder/#expose
EXPOSE 8080


CMD ["./hamburgueria-app"]