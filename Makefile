up:
	docker-compose up
test-cover:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out