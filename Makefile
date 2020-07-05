.PHONY: start build

NOW = $(shell date -u '+%Y%m%d%I%M%S')

all: start

build:
	@go build -ldflags "-w -s" -o ./cmd/phone main.go

start:
	go run main.go -f phone.txt