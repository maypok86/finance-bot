SHELL := /bin/bash

BIN := "./bin/bot"
SRC := "./cmd/bot"

GIT_HASH := $(shell git log --format="%h" -n 1)
LDFLAGS := -X main.release="develop" -X main.buildDate=$(shell date -u +%Y-%m-%dT%H:%M:%S) -X main.gitHash=$(GIT_HASH)

.PHONY: setup
setup: ## Install all the build and lint dependencies
	(which golangci-lint > /dev/null) || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(shell go env GOPATH)/bin v1.37.0
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
	go get golang.org/x/tools/cmd/goimports

.PHONY: build
build: ## Build a version
	go build -v -o $(BIN) -ldflags "$(LDFLAGS)" $(SRC)

.PHONY: run
run: build ## Build and run a version
	echo -n > develop.log
	source ./scripts/env.sh && $(BIN) -config configs/config.yml

.PHONY: version
version: build ## Build and view project version
	$(BIN) version

.PHONY: fmt
fmt: ## Run goimports on all go files
	find . -name '*.go' -not -wholename './vendor/*' | while read -r file; do goimports -w "$$file"; done

.PHONY: lint
lint: ## Run all the linters
	golangci-lint run ./...

.PHONY: test
test: ## Run all the tests
	echo -n > coverage.txt
	echo -n > develop.log
	echo 'mode: atomic' > coverage.txt && go test -covermode=atomic -coverprofile=coverage.txt -v -race -timeout=30s ./...

.PHONY: cover
cover: test ## Run all the tests and opens the coverage report
	go tool cover -html=coverage.txt

.PHONY: ci
ci: lint test ## Run all the tests and code checks

.PHONY: generate
generate: ## Generate all mocks
	go generate ./...

.PHONY: clean
clean: ## Remove temporary files
	go clean
	rm -rf bin/
	rm -rf coverage.txt
	rm -rf develop.log

.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.DEFAULT_GOAL:= help