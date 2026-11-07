.PHONY: build dev

build:
	go run ./cmd/site build

dev:
	go tool modd
