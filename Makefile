.DEFAULT_GOAL := all

.PHONY: all
all: lint test

.PHONY: test
test:
	@echo "Running tests..."
	@go test -v ./...

.PHONY: coverage
coverage:
	@echo "Running tests with coverage..."
	@go test -cover ./...


.PHONY: vet
vet:
	@echo "Running go vet..."
	@go vet ./...

.PHONY: fmt
fmt:
	@echo "Formatting code..."
	@go fmt ./...