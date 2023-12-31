SHELL := /bin/bash

VERSION=`git describe --exact-match --tags HEAD 2> /dev/null`
COMMIT=`git rev-parse HEAD 2> /dev/null`
DATE_BUILD=`date +%Y-%m-%d\_%H:%M`

GOBIN = $(GOPATH)/bin
STATICCHECK = $(GOBIN)/staticcheck

.PHONY: all lint test cover build clean help fmt

all: build

$(STATICCHECK):
	go install honnef.co/go/tools/cmd/staticcheck@latest

fmt: ## Make gofmt persistant
	gofmt -s -w -l .

lint: $(GOLINT) ## Start lint
	diff -u <(echo -n) <(gofmt -s -d .); [ $$? -eq 0 ]
	go vet ./...
	$(STATICCHECK) ./...

test: ## Run test
	go test -race -v -coverprofile=coverage.txt ./...

cover: test ## Display test coverage percent
	go tool cover -func coverage.txt

build: ## Build debug binary
	GOOS=linux GOARCH=arm GOARM=7 go build -v -ldflags "-X main.version=debug" -o "build/tic-linky" ./cmd/tic-linky/*.go

release: ## Build release binary
	GOOS=linux GOARCH=arm GOARM=7 go build -v -ldflags "-X main.version=${VERSION} -X main.commit=${COMMIT} -X main.buildDate=${DATE_BUILD}" -o "build/tic-linky" ./cmd/tic-linky/*.go

clean: ## Remove vendors and build
	rm -rf build
	rm -f coverage.txt

help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
