run: build
	@./bin/authn

build:
	@go build -o bin/authn cmd/api/main.go
