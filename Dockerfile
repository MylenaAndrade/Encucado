# Build stage
FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/api ./cmd/api

# Runtime stage
FROM alpine:3.20

WORKDIR /app
COPY --from=builder /bin/api /app/api

EXPOSE 8080

ENTRYPOINT ["/app/api"]
