build: format test _build
	go build -o _build/gojson ./gojson

test:
	go test

format:
	gofmt -w -e -s -l *.go **/*.go

_build:
	mkdir _build

.PHONY: build test format
