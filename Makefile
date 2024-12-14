build: 
	@go build -o bin/cleaner-v0

run: build
	@./bin/cleaner-v0

test: 
	@go test -v ./...
