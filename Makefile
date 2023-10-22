.PHONY=build

build:
	@go build -o bin/main main.go

run: build
	@./bin/main

test: build
	@go test -v ./...