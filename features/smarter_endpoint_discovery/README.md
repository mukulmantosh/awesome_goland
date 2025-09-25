# Gorilla Mux + SQLite Basic Web App

This is a minimal web application written in Go using:
- Gorilla Mux for HTTP routing
- SQLite (pure Go driver) for persistence

It exposes a tiny API to manage a list of items and demonstrates a straightforward project structure.

## Features
- Healthcheck endpoint
- Create an item
- List items
- Auto-create the SQLite database and schema if missing
- Graceful shutdown

## Requirements
- Go 1.20+ (tested with Go 1.22). The `go.mod` may state a newer version but the code uses standard features.
- No CGO required (uses `modernc.org/sqlite` pure-Go driver)

## Project Layout
```
web_app/
  README.md
  go.mod
  main.go
  internal/
    db/
      db.go            # DB connection + schema migration
    handlers/
      handlers.go      # HTTP handlers (health, items)
    models/
      item.go          # Data model
```

## Running
From the web_app directory:

```bash
# Fetch dependencies
go mod tidy

# Run
go run ./...
```

By default the server listens on `:8080` and creates a SQLite database at `data/app.db`.

You can override defaults via environment variables:
- `PORT`: server port (default `8080`)
- `DB_PATH`: sqlite database file path (default `data/app.db`)

## API
- Health
  - `GET /health` -> `{ "status": "ok" }`

- List items
  - `GET /items` -> `[{ "id": 1, "name": "foo", "created_at": "2025-01-01T00:00:00Z" }]`

- Create item
  - `POST /items`
    - Content-Type: application/json
    - Body: `{ "name": "my item" }`
    - Response: `201 Created` with the created item JSON

## Examples
```bash
# Health
curl -s http://localhost:8080/health | jq

# Create
curl -s -X POST http://localhost:8080/items \
  -H 'Content-Type: application/json' \
  -d '{"name":"first"}' | jq

# List
curl -s http://localhost:8080/items | jq
```

## Notes
- The database schema is created on startup if it does not exist.
- Errors are returned with appropriate HTTP status codes and JSON messages.
- The code is intentionally minimal and well-commented for learning purposes.
