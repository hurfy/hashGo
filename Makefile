# Variables
MAIN_PKG := ./cmd/hashGo
LDFLAGS  := -ldflags "-s -w"

# Default mode
.DEFAULT_GOAL := run

build:
	go build $(LDFLAGS) $(MAIN_PKG)

run:
	go run $(MAIN_PKG) -f md5 -s -p D:\Games\m2net

test:
	go test -v $(MAIN_PKG)

.PHONY: build run test