MAKEFLAGS += --silent

ldflags := -X 'github.com/preetbiswas12/Kage/constant.BuiltAt=$(shell date -u)'
ldflags += -X 'github.com/preetbiswas12/Kage/constant.BuiltBy=$(shell whoami)'
ldflags += -X 'github.com/preetbiswas12/Kage/constant.Revision=$(shell git rev-parse --short HEAD)'
ldflags += -s
ldflags += -w

build_flags := -ldflags=${ldflags}

all: help

help:
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@echo "  build        Build the kage binary"
	@echo "  install      Install the kage binary"
	@echo "  uninstall    Uninstall the kage binary"
	@echo "  test         Run the tests"
	@echo "  gif          Generate usage gifs"
	@echo "  help         Show this help message"
	@echo ""

install:
	@go install "$(build_flags)"


build:
	@go build "$(build_flags)"

test:
	@go test ./...

uninstall:
	@rm -f $(shell which kage)

gif:
	@vhs assets/tui.tape
	@vhs assets/inline.tape
