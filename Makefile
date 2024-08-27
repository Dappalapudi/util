.PHONY: help build test integration

.EXPORT_ALL_VARIABLES:
GOPRIVATE              = github.com/Dappalapudi/*
CGO_ENABLED           := 1
CGO_CFLAGS             = -g -O2 -Wno-return-local-addr

# Sets a default for m	ake
.DEFAULT_GOAL := help


help:; ## Output help
	@printf "%s\\n" \
		"The following targets are available:" \
		""
	@awk 'BEGIN {FS = ":.*?## "} /^[\/.%a-zA-Z_-]+:.*?## / {printf "  \033[36m%-15s\033[0m - %s\n", $$1, $$2}' ${MAKEFILE_LIST}

	@printf "%s\\n" "" "" \
		"Examples:" \
		"make build" \
		"    build all packages" 

build: ## Build all packages
	go build ./...

test: ## Test all packages
	go test ./...

integration: ## Integration testing
	go test --tags=integration ./...

