# App configuration
APP_NAME = server

# Load environment variables from .env if the file exists
ifneq ("$(wildcard .env)","")
	include .env
	export
endif

# Run the main application
run:
	go run ./cmd/$(APP_NAME)/

# Create a new migration file
# Usage: make create-migration name=create_todos_table
create-migration:
	migrate create -ext sql -dir sql/migration -seq $(name)

# Apply all available up migrations
migrate-up:
	migrate -path sql/migration -database "$(PG_URL)" up

# Roll back the most recent migration
migrate-down:
	migrate -path sql/migration -database "$(PG_URL)" down 1

# Roll back and re-apply the last migration (useful during development)
migrate-redo:
	migrate -path sql/migration -database "$(PG_URL)" redo

# Show the current migration version
migrate-status:
	migrate -path sql/migration -database "$(PG_URL)" version

# Force set the migration version
# Usage: make migrate-force version=1
migrate-force:
	migrate -path sql/migration -database "$(PG_URL)" force $(version)

# Generate Go code from SQL using sqlc
sqlc-generate:
	sqlc generate
