all: test build lint

build:
	go build .
	go test . -run xxx

test:
	go test .
	go test -race .

lint:
	golangci-lint run .

tidy:
	go mod tidy

fmt:
	gofmt -s -w .