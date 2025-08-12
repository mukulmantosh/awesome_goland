Basic REST API with SQLite

This is a minimal Go REST API using Gorilla Mux and SQLite (modernc.org/sqlite, pure Go). It exposes:

- GET /health -> {"status":"ok"}
- GET /items -> list items
- POST /items -> create item with JSON body {"name":"..."}

Local run:
- go run ./...
- Environment variables:
  - PORT (default 8080)
  - DB_PATH (default data/app.db)

Docker build & run:
- docker build -t basic-rest-sqlite .
- docker run -p 8080:8080 -e PORT=8080 -e DB_PATH=/data/app.db -v $(pwd)/data:/data basic-rest-sqlite
