# Colors
YELLOW := $(shell tput -Txterm setaf 3)
GREEN  := $(shell tput -Txterm setaf 2)
RESET  := $(shell tput -Txterm sgr0)

# Default variables
APP_NAME ?= chronolog
GO_FILES := $(shell find . -type f -name '*.go' -not -path "./vendor/*")

# Versioning
VERSION := $(shell git describe --tags --abbrev=0 2>/dev/null || echo "0.0.1")
BUILD_TIME := $(shell date -u +%Y-%m-%dT%H:%M:%SZ)
COMMIT := $(shell git rev-parse --short HEAD)

# Build flags
LDFLAGS = -X 'github.com/mvleandro/chronolog/internal.LibraryVersion=$(VERSION)' \
          -X 'github.com/mvleandro/chronolog/internal.LibraryCommit=$(COMMIT)' \
          -X 'github.com/mvleandro/chronolog/internal.LibraryBuildTime=$(BUILD_TIME)'

## ---------------------------------------------------------------------
## 🆘 Help
## ---------------------------------------------------------------------
help:
	@echo ""
	@echo "${YELLOW}Available commands:${RESET}"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "  ${GREEN}%-22s${RESET} %s\n", $$1, $$2}'
	@echo ""

## ---------------------------------------------------------------------
## 🚀 Build & Run
## ---------------------------------------------------------------------
build: ## Build the Go project
	@echo "${YELLOW}Building the module...${RESET}"
	go build -ldflags "$(LDFLAGS)" -o bin/$(APP_NAME) ./examples

run-example: ## Run the example application
	@echo "${YELLOW}Running example...${RESET}"
	go run ./examples/main.go

## ---------------------------------------------------------------------
## 🧪 Test & Quality
## ---------------------------------------------------------------------
test: ## Run unit tests
	@echo "${YELLOW}Running tests...${RESET}"
	go test ./... -v

coverage: ## Generate code coverage report
	@echo "${YELLOW}Generating coverage report...${RESET}"
	go test ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out -o coverage.html
	@echo "${GREEN}✔ Coverage report saved to coverage.html${RESET}"

lint: ## Run golangci-lint
	@echo "${YELLOW}Linting code...${RESET}"
	golangci-lint run

format: ## Format code (gofmt + goimports)
	@echo "${YELLOW}Formatting code...${RESET}"
	gofmt -w $(GO_FILES)
	goimports -w $(GO_FILES)

tidy: ## Tidy up go.mod
	@echo "${YELLOW}Tidying modules...${RESET}"
	go mod tidy

## ---------------------------------------------------------------------
## 📦 Release & Changelog
## ---------------------------------------------------------------------
changelog: ## Generate CHANGELOG.md using git-chglog
	@echo "${YELLOW}Generating changelog...${RESET}"
	git-chglog -o CHANGELOG.md

release: ## Build release using GoReleaser
	@echo "${YELLOW}Building release with GoReleaser...${RESET}"
	goreleaser release --clean --skip-validate

## ---------------------------------------------------------------------
## 🔧 Dev Tools
## ---------------------------------------------------------------------
install-tools: ## Install dev tools (lint, goimports, etc.)
	@echo "${YELLOW}Installing developer tools...${RESET}"
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install golang.org/x/tools/cmd/goimports@latest
	go install github.com/git-chglog/git-chglog/cmd/git-chglog@latest
	go install github.com/goreleaser/goreleaser@latest

## ---------------------------------------------------------------------
## 🧽 Clean
## ---------------------------------------------------------------------
clean: ## Remove build files
	@echo "${YELLOW}Cleaning project...${RESET}"
	rm -f coverage.out coverage.html
	rm -rf bin

.PHONY: help build run-example test coverage lint format tidy changelog release install-tools clean
