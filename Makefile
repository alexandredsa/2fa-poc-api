.PHONY: run

run:
	go run ./cmd/main.go

.PHONY: build test test-coverage test-coverage-html

build:
	go build -o app ./cmd/main.go

test:
	go test -v ./...

test-coverage:
	go test -cover ./...

test-coverage-html:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
	rm coverage.out
