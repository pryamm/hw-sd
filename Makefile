install:
	go mod tidy

lint:
	go fmt ./...

test: install
	go test -v ./...

build: install
	go build -v .

run: install
	go run main.go
