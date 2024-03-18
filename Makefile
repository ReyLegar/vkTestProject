.PHONY: build
build:
		go build -v ./cmd/vkTestProject

.PHONY: test
test: 
		go test -v -timeout 30s ./...

.DEFAULT_GOAL := build