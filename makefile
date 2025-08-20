OPENAPI_SPEC=spec/openapi.yaml
OPENAPI_GEN_DIR=internal/apigen

.PHONY: run build test lint fmt install-tools clean check-line-length

run:
	go run .

build:
	go build -o bin/evm-indexer-go .

test:
	go test -v ./...

.PHONY: gen-api
gen-api: $(OPENAPI_SPEC)
	@echo "Generating API code from OpenAPI specification..."
	@mkdir -p $(OPENAPI_GEN_DIR)
	oapi-codegen -generate types,echo,spec -package apigen $(OPENAPI_SPEC) > $(OPENAPI_GEN_DIR)/api.gen.go


abigen:
	./scripts/abigen.sh

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

# Check line length specifically
check-line-length:
	@echo "Checking line length (max 100 characters)..."
	@golangci-lint run --disable-all --enable=lll

clean:
	@rm -rf bin/