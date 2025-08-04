BIN_NAME=huff

.PHONY: build
build:
	@go build -o bin/huff .

.PHONY: execute
execute: build
	@cd bin && ./huff