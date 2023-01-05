build:
	@go build -o bin/go-commerce

run: build
	@./bin/go-commerce

test:
	@go test -v ./...
