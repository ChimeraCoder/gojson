INSTALL_PATH?=/usr/local/bin

compile: gojson

gojson:
	go build gojson.go

clean:
	rm gojson

install: gojson
	[ ! -f ${INSTALL_PATH}/gojson ] && install gojson ${INSTALL_PATH}

.PHONY: compile clean install
