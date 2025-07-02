# Load environment variables from .env if the file exists
ifneq ("$(wildcard .env)","")
	include .env
	export
endif

# Run the main application
run:
	go run ./cmd/server/

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

# Generate Go code from Protobuf definitions
.PHONY: proto
proto-user:
	protoc --proto_path=api/proto/user/v1 \
	       --go_out=api/proto/user/v1/gen --go_opt=paths=source_relative \
	       --go-grpc_out=api/proto/user/v1/gen --go-grpc_opt=paths=source_relative \
	      	api/proto/user/v1/*.proto

