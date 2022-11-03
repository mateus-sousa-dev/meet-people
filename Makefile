up:
	docker-compose down
	docker-compose up -d
build:
	docker-compose down
	docker-compose up --build -d
test:
	go test ./...
test-cover:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out
generate-doc: ## Generate Swagger Api Documentation
	cd cmd/api/ && swag init --parseDependency --parseInternal --parseDepth 1 -o ../../docs
