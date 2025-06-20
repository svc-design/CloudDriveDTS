APP_NAME := cloudvault
MAIN_PACKAGE := ./cmd/cloudvault
GO ?= go
PROVIDER ?=
SYNC_FLAGS ?=
DAEMON_FLAGS ?=

.PHONY: all build run login sync daemon clean init test help

all: build

init:
	$(GO) mod tidy

build:
	$(GO) build -o $(APP_NAME) $(MAIN_PACKAGE)

run:
	$(GO) run $(MAIN_PACKAGE)

login:
	$(GO) run $(MAIN_PACKAGE) login $(PROVIDER)

sync:
	$(GO) run $(MAIN_PACKAGE) sync $(SYNC_FLAGS)

daemon:
	$(GO) run $(MAIN_PACKAGE) daemon $(DAEMON_FLAGS)

clean:
	rm -f $(APP_NAME)

test:
	$(GO) test ./...

help:
	@echo "☁️  CloudVault CLI Usage"
	@echo ""
	@echo "make build                 Build cloudvault binary"
	@echo "make run                   Run main entry"
	@echo "make login PROVIDER=name   Login to provider"
	@echo "make sync SYNC_FLAGS=...   Sync files"
	@echo "make daemon DAEMON_FLAGS=... Start daemon"
	@echo "make init                  Install dependencies (go mod tidy)"
	@echo "make test                  Run unit tests"
	@echo "make clean                 Remove build artifacts"
