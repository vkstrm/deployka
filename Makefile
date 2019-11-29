
deps:
	go get ./...

build:
	go build ./...

install:
	go install ./cmd/deployka

test:
	go test ./internal/...
