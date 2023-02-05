.DEFAULT_GOAL = help
SHELL := /bin/bash
.ONESHELL:
.SHELLFLAGS := -eu -o pipefail -c
.DELETE_ON_ERROR:
BUILD_OUTPUT = "./bin/lol"
MAKEFLAGS += --warn-undefined-variables
MAKEFLAGS += --no-builtin-rules


build: test ## build the go package
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -v -trimpath -buildmode=pie -ldflags '-s -w -buildid=none' -o $(BUILD_OUTPUT) nrvn.cc/go/lol/cmd

##@ Development
mod: ## verify and download go modules
	@true >go.sum
	@go mod tidy -v
	@go mod verify

mod-outdated: ## list outdated modules
	@printf "Direct dependencies\n"
	@go list -u -m -f '{{if (and (not .Indirect) .Update)}}{{.}}{{end}}' all
	@printf "Indirect dependencies\n"
	@go list -u -m -f '{{if (and (not .Indirect) .Update)}}{{.}}{{end}}' all

lint: ## lint source code
	@golangci-lint run

test: ## test the go package
	go test -failfast -race -cover -coverprofile cover.out nrvn.cc/go/lol/...

##@ Helpers
help: ## Display this help
	@awk 'BEGIN {FS = ":.*##"; printf "\n\033[1mUsage\033[0m\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)
