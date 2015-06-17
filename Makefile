build: format test
	go build -o build/gojson cmd/gojson/gojson.go

test:
	go test -v

format:
	gofmt -w -e -s -l *.go

clean:
	rm -rf build

.PHONY: test format clean
