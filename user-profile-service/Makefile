# Makefile

APP_NAME=main

build:
	go build -o $(APP_NAME) .

test:
	go test ./... -v

lint:
	@command -v golangci-lint >/dev/null 2>&1 || go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	golangci-lint run ./...

run:
	go run main.go

clean:
	rm -f $(APP_NAME)

fmt:
	go fmt ./...

deps:
	go mod tidy
