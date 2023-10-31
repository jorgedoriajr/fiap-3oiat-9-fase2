FROM golang:1.21-alpine as builder

WORKDIR /hamburgueria

COPY ./ ./

RUN go mod tidy && \
    go build -a -installsuffix cgo -ldflags '-extldflags -s -w' -o ./bin/hamburgueria ./cmd/.

# Execution container
FROM golang:1.21-alpine

WORKDIR /app/app

COPY --from=builder /hamburgueria/config /app/app/config
COPY --from=builder /hamburgueria/hamburgueria /app/app

EXPOSE 8080 8081

CMD ["/app/app/hamburgueria" ]