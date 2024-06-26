gen:
	@templ generate ./...
	@go generate ./...

test: gen
	@go test ./...

infra-start:
	docker-compose -f build/compose.yaml up --detach --remove-orphans

infra-stop:
	docker-compose -f build/compose.yaml down

migrate:
	GOOSE_MIGRATION_DIR=./storage/postgres/migrations/ goose postgres "host=127.0.0.1 user=postgres password=postgres dbname=anauction" up

run: gen
	@go run cmd/anauction/main.go
