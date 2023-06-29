# Makefile

# Variables
DOCKER_IMAGE_NAME = '2fa-api'
DOCKER_CONTAINER_NAME = '2fa-api-local'

.PHONY: build
build:
	@go build -o ./bin/main ./cmd/main.go

.PHONY: run
run:
	@go run ./cmd/main.go

.PHONY: test
test:
	@go test ./...

.PHONY: test-coverage-html
test-coverage-html:
	@go test -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out -o coverage.html

.PHONY: docker-build
docker-build:
	docker build -t $(DOCKER_IMAGE_NAME) .

.PHONY: docker-run
docker-run:
	@docker run -p 8080:8080 --name $(DOCKER_CONTAINER_NAME) $(DOCKER_IMAGE_NAME)

.PHONY: docker-stop
docker-stop:
	@docker stop $(DOCKER_CONTAINER_NAME)
	@docker rm $(DOCKER_CONTAINER_NAME)

.PHONY: docker-build-run
docker-build-run:
	@make docker-build
	@docker-compose up -d

.PHONY: docker-clean
docker-clean:
	@docker rm $(DOCKER_CONTAINER_NAME)

.PHONY: docker-up
docker-up:
	@docker-compose up -d

.PHONY: docker-down
docker-down:
	@docker-compose down
