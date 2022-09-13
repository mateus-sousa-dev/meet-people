up:
	docker-compose up
bup:
	docker-compose up --build
test:
	go test ./...
test-cover:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out
generate-doc: ## Generate Swagger Api Documentation
	swag init --parseDependency --parseInternal --parseDepth 1
