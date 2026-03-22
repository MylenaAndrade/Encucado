# Encucado Backend (Go)

Estrutura inicial de backend em Go.

## Requisitos

- Go 1.22+

Ou, se preferir, Docker + Docker Compose.

## Configuracao de ambiente

Copie `.env.example` para `.env` e ajuste os valores se necessario.

Variaveis usadas pela API:

- `PORT` (padrao: `8080`)
- `DATABASE_URL` (obrigatoria)

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

## Migrations iniciais

Arquivos SQL:

- `migrations/000001_create_users.up.sql`
- `migrations/000001_create_users.down.sql`

Aplicar migration inicial no banco do Docker:

```bash
docker compose exec -T db psql -U encucado -d encucado -f /migrations/000001_create_users.up.sql
```

Reverter migration inicial:

```bash
docker compose exec -T db psql -U encucado -d encucado -f /migrations/000001_create_users.down.sql
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
- `GET /readyz` -> valida conexao com PostgreSQL
- `GET /users?limit=50` -> lista usuarios
- `POST /users` -> cria usuario

Exemplo de listagem de usuarios:

```bash
curl "http://localhost:8080/users?limit=20"
```

Exemplo de criacao de usuario:

```bash
curl -X POST http://localhost:8080/users \
	-H "Content-Type: application/json" \
	-d '{"name":"Mylen","email":"mylen@example.com"}'
```
