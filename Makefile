gen:
	@templ generate ./...
	@go generate ./...

test: gen
	@go test ./...

run: gen
	@go run cmd/anauction/main.go
