# Variables
MAIN_PKG := ./cmd/hashGo
LDFLAGS  := -ldflags "-s -w"

# Default mode
.DEFAULT_GOAL := build

.PHONY: build
build:
	go build $(LDFLAGS) $(MAIN_PKG)

.PHONY: run
run: build
	./hashGo.exe --path ./Makefile --format md5