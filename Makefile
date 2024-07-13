GO ?= go
GOFMT ?= gofmt "-s"
GOFILES := $(shell find . -name "*.go")
CURRENT_DIR = $(shell pwd)
SERVER_DIR = cd server/
CLIENT_DIR = cd client/
DOCKER ?= docker

.PHONY: air
# Install air for live-reload.
air:
	@mkdir -p bin
	@GOBIN=$(CURRENT_DIR)/$(SERVER_DIR)/bin $(GO) install github.com/air-verse/air@latest

.PHONY: dev
# Run the application in watch mode.
dev:
	@$(CURRENT_DIR)/bin/air

.PHONY: build
# Build the client.
build:
	@$(CLIENT_DIR) && npm run build	

.PHONY: run
# Run the application.
run:
	@$(SERVER_DIR) && $(GO) run .

.PHONY: up
# Start the containers.
up:
	@$(DOCKER) compose up -d

.PHONY: down
# Stop the containers.
down:
	@$(DOCKER) compose down

.PHONY: enter
# Enter the database.
enter:
	@$(DOCKER) exec -it gin-postgres psql -d gin-postgres -U demystif -W

.PHONY: tidy
# Tidy the Go module.
tidy:
	@$(SERVER_DIR) && $(GO) mod tidy

.PHONY: fmt
# Format the Go files.
fmt:
	@$(SERVER_DIR) && $(GOFMT) -w $(GOFILES)

help:
	@echo ''
	@echo 'Usage:'
	@echo ' make [target]'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
	helpMessage = match(lastLine, /^# (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 2, RLENGTH); \
			printf " - \033[36m%-20s\033[0m %s\n", helpCommand, helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help
