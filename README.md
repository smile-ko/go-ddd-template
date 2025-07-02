# Go DDD Template

## 1. Project Overview

This project is a template for backend applications using Domain-Driven Design (DDD) architecture with Golang, Gin, Postgres, SQLC, Docker, and best practices for structure, configuration, logging, migration, and more.

- **Entry point:** `cmd/server/main.go`
- **Initialization & wiring:** `internal/bootstrap.go`
- **Configuration:** `config/config.go` (via environment variables, supports `.env`)
- **Database:** PostgreSQL, migrations with `migrate`, queries via SQLC
- **HTTP API:** Gin framework, RESTful, handlers in `internal/interfaces/http/v1/`
- **Repository pattern:** Clear separation of domain, repository, usecase, handler

---

## 2. Project Structure

```
go-ddd-template/
├── api/                  # OpenAPI/proto specs
├── cmd/
│   └── server/           # Entry point (main.go)
├── config/               # Application configuration
├── docs/                 # Architecture docs
├── internal/
│   ├── application/      # Usecase, service, DTO (application layer)
│   ├── domain/           # Entity, repository interface (pure domain)
│   ├── infrastructure/   # Repository implementation, DB, SQLC
│   ├── interfaces/       # HTTP handler, transport layer
│   └── bootstrap.go      # Run function to initialize the app
├── pkg/                  # Shared libraries (logger, httpserver, postgres)
├── sql/                  # Migration & SQL queries
├── Dockerfile
├── docker-compose.yml
├── Makefile
└── go.mod, go.sum
```

---

## 3. Layer Flow Diagram

```mermaid
flowchart TD
    A[main.go] --> B[bootstrap.Run]
    B --> C[Config Loader]
    B --> D[Logger Init]
    B --> E[Postgres Connect]
    B --> F[SQLC Queries]
    B --> G[Repository]
    B --> H[UseCase]
    B --> I[HTTP Handler]
    I --> J[HTTP Server]
    J --> K[Client]
    E --> F
    F --> G
    G --> H
    H --> I
```

**Explanation:**

- `main.go` calls `bootstrap.Run`, which initializes config, logger, DB, repository, usecase, handler, and runs the HTTP server.
- Layers are separated: infrastructure (DB, repo), application (usecase), domain (entity, repo interface), interfaces (HTTP handler).

---

## 4. Setup & Usage

### 4.1. Requirements

- Go 1.20+
- Docker, docker-compose
- PostgreSQL (recommended via docker-compose)

### 4.2. Configuration

- Create a `.env` file in the project root, for example:
  ```
  APP_NAME=server
  APP_VERSION=1.0.0
  ENV_NAME=dev
  HTTP_PORT=8001
  LOG_LEVEL=debug
  PG_URL=postgres://postgres:postgres@localhost:5432/go_ddd_template?sslmode=disable
  ...
  ```

### 4.3. Database Migration

- Create a new migration:
  ```
  make create-migration name=create_todos_table
  ```
- Apply migrations:
  ```
  make migrate-up
  ```
- Rollback:
  ```
  make migrate-down
  ```

### 4.4. Generate SQLC Code

- Generate Go code from SQL:
  ```
  make sqlc-generate
  ```

### 4.5. Run Locally

- Start Postgres with docker-compose:
  ```
  docker-compose up -d
  ```
- Run the app:
  ```
  make run
  ```
  or
  ```
  go run ./cmd/server/
  ```

### 4.6. Build & Run with Docker

- Build the image:
  ```
  docker build -t go-ddd-template .
  ```
- (You can edit docker-compose.yml to enable the app service)

---

## 5. Highlights

- **Config:** Uses env variables, supports `.env`, all config in clear structs.
- **Logging:** Zap, supports file, console, JSON logging.
- **Migration:** Managed by Makefile, easy rollback/redo.
- **Repository pattern:** Easy to extend, test, swap DB.
- **SQLC:** Safe SQL queries, auto-generates Go code.
- **HTTP server:** Gin, handler separation, easy API versioning.

---

For API examples, more detailed flow for each layer, or testing instructions, feel free to ask!
