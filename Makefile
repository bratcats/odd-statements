all: tidy build test

tidy:
	go mod tidy

build:
	go build

test:
	go test ./...