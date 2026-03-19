# Encucado Backend (Go)

Estrutura inicial de backend em Go.

## Requisitos

- Go 1.22+

Ou, se preferir, Docker + Docker Compose.

## Executar

```bash
go run ./cmd/api
```

## Executar com Docker

### Subir somente dependencias (PostgreSQL)

```bash
docker compose up -d db
```

### Subir backend + dependencias

```bash
docker compose up --build
```

### Parar os containers

```bash
docker compose down
```

Para remover tambem o volume do banco:

```bash
docker compose down -v
```

## Endpoint inicial

- `GET /healthz` -> `{"status":"ok"}`
