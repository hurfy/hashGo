# Variables
MAIN_PKG := ./cmd/hashGo
LDFLAGS  := -ldflags "-s -w"

# Default mode
.DEFAULT_GOAL := build

build:
	go build $(LDFLAGS) $(MAIN_PKG)

run:
	go build $(MAIN_PKG)
	./hashGo.exe -p D:\Games\m2net -f sha256 -sd -o -s -op ./output.json

.PHONY: build run