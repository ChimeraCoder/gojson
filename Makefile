build/gojson: format test
	mkdir -p build
	go build -o build/gojson ./gojson

test:
	go test -v -cover

format:
	gofmt -w -e -s -l *.go **/*.go

clean:
	rm -rf build

.PHONY: test format clean
