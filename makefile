.PHONY: run build test lint fmt install-tools clean

run:
	go run .

build:
	go build -o bin/evm-indexer-go .

test:
	go test -v ./...

# Install required tools
install-tools:
	@echo "Installing golangci-lint..."
	@curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.55.2
	@echo "Installing goimports..."
	@go install golang.org/x/tools/cmd/goimports@latest

# Format code
fmt:
	@echo "Formatting code..."
	@goimports -w .
	@gofmt -w .

# Run linter
lint:
	@echo "Running linter..."
	@golangci-lint run

# Fix linting issues automatically where possible
lint-fix:
	@echo "Fixing linting issues..."
	@golangci-lint run --fix

clean:
	@rm -rf bin/