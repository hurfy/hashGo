.PHONY: build
build:
	go build ./cmd/hashGo
	./hashGo.exe

.PHONY: test
test:
	go build ./cmd/hashGo
	./hashGo.exe --path ./Makefile --format sha256

.DEFAULT_GOAL := build