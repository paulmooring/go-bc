.PHONY: run test build fmt vet clean

run:
	go run .

test:
	go test ./...

build:
	go build -o bin/go-bc .

fmt:
	go fmt ./...

vet:
	go vet ./...

clean:
	rm -rf bin
