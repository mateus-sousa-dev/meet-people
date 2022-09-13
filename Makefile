up:
	docker-compose up
bup:
	docker-compose up --build
test:
	go test ./...
generate-doc: ## Generate Swagger Api Documentation
	swag init --parseDependency --parseInternal --parseDepth 1
