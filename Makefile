build/gojson: format test
	mkdir -p build
	go build -o build/gojson ./gojson

test:
	go test -v

format:
	gofmt -w -e -s -l *.go **/*.go

.PHONY: test format
