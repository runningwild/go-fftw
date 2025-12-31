# Build the library
build:
    go build -v ./...

# Run all tests
test:
    go test -v -count=1 ./...

# Run benchmarks
bench:
    go test -bench=. -benchmem -run=^$ ./...

# Run linters
lint:
    golangci-lint run

# Run linters and fix issues
lint-fix:
    golangci-lint run --fix

# Format code using treefmt
fmt:
    treefmt . --allow-missing-formatter

# Check if code is formatted
fmt-check:
    treefmt --allow-missing-formatter --fail-on-change

# Generate coverage report
cover:
    go test -coverprofile=coverage.txt -covermode=atomic ./...
    go tool cover -html=coverage.txt -o coverage.html

# Clean build artifacts
clean:
    rm -f coverage.txt coverage.html

# Run all checks (test, lint, coverage)
check: test lint cover

# Default target
default: build
