# Etapa 1: Compilacion del binario estatico en Go
FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags="-s -w" -o /app/api ./cmd/api

# Etapa 2: Imagen minima de ejecucion (sin compilador ni paquetes innecesarios)
FROM alpine:3.21

RUN apk --no-cache add ca-certificates tzdata && \
    adduser -D -u 10001 appuser

WORKDIR /app
COPY --from=builder /app/api ./api

USER appuser
EXPOSE 8080

CMD ["./api"]
